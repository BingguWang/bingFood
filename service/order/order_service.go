package order

import (
    "bingFood/entity/basket"
    "bingFood/entity/order"
    "bingFood/entity/order/req"
    "bingFood/entity/order/resp"
    "bingFood/entity/prod"
    "bingFood/global"
    "bingFood/utils"
    "encoding/json"
    "fmt"
    "github.com/bwmarrin/snowflake"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/copier"
    "github.com/nsqio/go-nsq"
    "github.com/pkg/errors"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
    "log"
    "strconv"
    "time"
)

func SettleOrder(ctx *gin.Context, param req.SettleOrderReq) (res interface{}, err error) {
    log.Printf("request args are : %v", utils.ToJsonString(param))

    db := global.MYSQL_DB
    var basketList []basket.Basket

    if err = db.Where("basket_id IN ? ", param.BasketIds).Preload("Sku").Find(&basketList).Error; err != nil {
        return
    }

    fmt.Println(utils.ToJsonString(basketList))

    var (
        oriPriceTotal   int // 原价总和
        packingFeeTotal int // 打包费
        priceTotal      int // 现价总和
        finalTotal      int // 最后需支付金额
        discountTotal   int // 总共优惠的金额
        deliverFeeTotal int // 配送费
        redPacket       int // 红包
        itemList        []order.OrderItem
        prodNums        int // 总商品个数
        shopId          uint64
        prodName        string // 商品名，用分号连接
    )
    for _, v := range basketList {
        sku := v.Sku
        item := order.OrderItem{
            UserId:     v.UserId,
            ShopId:     v.ShopId,
            ProdId:     0,
            ProdName:   sku.ProdName,
            ProdNums:   v.ProdNums,
            Pic:        sku.Pic,
            Price:      sku.Price,
            ProdAmount: sku.Price * v.ProdNums,
            OriPrice:   sku.OriPrice,
            SkuId:      sku.SkuId,
            SkuName:    sku.SkuName,
            PropId:     sku.ProdId,
            PropName:   sku.ProdName + sku.SkuName,
        }
        oriPriceTotal += sku.OriPrice
        priceTotal += sku.Price
        packingFeeTotal += sku.PackingFee

        prodNums += v.ProdNums
        itemList = append(itemList, item)
        shopId = v.ShopId
        prodName += sku.ProdName + sku.SkuName + ";"
    }

    // TODO 配送费应该从配送系统计算得到，这里只是用个数值替一下
    deliverFeeTotal = 5 * 100 // 假设是固定的配送费

    discountTotal = (oriPriceTotal - priceTotal) + redPacket
    finalTotal = packingFeeTotal + priceTotal + deliverFeeTotal - discountTotal

    claims, _ := ctx.Get("claims")
    userClaims := claims.(*utils.UserClaims)
    fmt.Println(itemList)

    orderRes := resp.SettleOrderResp{
        ShopId:         shopId,
        UserMobile:     userClaims.UserMobile,
        ProdNums:       prodNums,
        PackingAmount:  packingFeeTotal,
        DeliverAmount:  deliverFeeTotal,
        ProdAmount:     priceTotal,
        DiscountAmount: discountTotal,
        FinalTotal:     finalTotal,
        OrderItems:     itemList,
        ProdName:       prodName,
    }

    fmt.Println(utils.ToJsonString(orderRes))

    // 返回的结算内容存到redis里,后面的提交订单时不需要前端再传过来了,提交订单的时候删掉
    cli := global.GVA_REDIS
    key := "settledOrder_" + strconv.FormatUint(shopId, 10) + "_" + userClaims.UserMobile // TODO 规范,常数写到其他地方去
    _, err = cli.Set(ctx, key, utils.ToJsonString(orderRes), 10*time.Minute).Result()     // 停留在结算页面没操作超过10分钟结算就作废
    if err != nil {
        return
    }

    return orderRes, nil
}

func ConfirmOrder(ctx *gin.Context, param req.ConfirmOrderReq) (err error) {
    log.Printf("request args are : %v", utils.ToJsonString(param))

    // 去redis取出结算好的订单信息
    claims, _ := ctx.Get("claims")
    userClaims := claims.(*utils.UserClaims)
    cli := global.GVA_REDIS
    key := "settledOrder_" + strconv.FormatUint(param.ShopId, 10) + "_" + userClaims.UserMobile
    var data string
    data, err = cli.Get(ctx, key).Result()
    if err != nil {
        if data == "" {
            err = errors.New("表单已过期，请重新结算")
        }
        return
    }
    var settledOrder resp.SettleOrderResp
    _ = json.Unmarshal([]byte(data), &settledOrder)

    var od order.Order
    _ = copier.CopyWithOption(&od, &settledOrder, copier.Option{
        IgnoreEmpty: true,
        DeepCopy:    true,
    })
    od.ReceiveAddr = param.ReceiveAddr

    // 生成orderNumber
    var node *snowflake.Node
    node, err = snowflake.NewNode(1) // 新建一个节点号为1的node
    if err != nil {
        return
    }

    number := node.Generate()
    od.OrderNumber = strconv.FormatInt(number.Int64(), 10)
    for i := 0; i < len(od.OrderItems); i++ {
        od.OrderItems[i].OrderNumber = od.OrderNumber
    }
    fmt.Println(utils.ToJsonString(od.OrderItems))
    od.Remark = param.Remarks

    // 插入order及order_item,更新库存
    if err = InsertOrder(od); err != nil {
        log.Printf("InsertOrder failed : %v", err.Error())
        return
    }

    // redis删除之前保存的结算信息
    if _, err = cli.Del(ctx, key).Result(); err != nil {
        return
    }

    // 把订单号存入到MQ里
    if err = PubOrderNumberToMQ(od.OrderNumber); err != nil {
        return
    }

    return nil
}

func InsertOrder(od order.Order) error {
    log.Printf("order is : %v", utils.ToJsonString(od))

    db := global.MYSQL_DB
    if err := db.Transaction(func(tx *gorm.DB) error {
        if err := tx.Omit(clause.Associations).Create(&od.OrderItems).Error; err != nil {
            log.Printf("insert orderItem failed : %v", err.Error())
            return err
        }
        if err := tx.Omit(clause.Associations).Create(&od).Error; err != nil {
            log.Printf("insert order failed : %v", err.Error())
            return err
        }

        // 更新库存
        for _, item := range od.OrderItems {
            if err := tx.Model(&prod.Sku{}).Where("sku_id = ? AND stock - ? >=0 ", item.SkuId, item.ProdNums).Update("stock", gorm.Expr("stock - ?", item.ProdNums)).Error; err != nil {
                log.Printf("update sku stock failed : %v", err)
                return err
            }
        }

        return nil
    }); err != nil {
        return err
    }
    return nil
}

func PubOrderNumberToMQ(orderNumber string) (err error) {
    log.Printf("未支付的订单号存入MQ : %v", orderNumber)

    config := global.NSQ_CONFIG
    // TODO yaml写配置
    producer, err := nsq.NewProducer("127.0.0.1:4150", config)
    if err != nil {
        return
    }

    // 5分钟未支付的订单就消费(视为订单取消)掉
    if err = producer.DeferredPublish("unPayOrder", 5*time.Second, []byte(orderNumber)); err != nil {
        panic(err)
        return
    }

    fmt.Println(time.Now().String())
    return nil
}
