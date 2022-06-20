package basket

import (
    "bingFood/common/response"
    "bingFood/entity/basket/req"
    bs "bingFood/service/basket"
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
)

func GetBasketDetail() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var basketReq req.BasketReq
        _ = ctx.ShouldBindJSON(&basketReq)

        pageSize := basketReq.PageInfo.PageSize
        page := basketReq.PageInfo.Page

        res, total, err := bs.GetBasketDetails(basketReq)
        if err != nil {
            response.FailWithMessage(fmt.Sprintf("获取购物车列表失败, err:%v", err.Error()), ctx)
            ctx.Abort()
            return
        }
        response.OkWithDetailed(response.PageResult{
            List:     res,
            Total:    total,
            Page:     page,
            PageSize: pageSize,
        }, "获取列表成功", ctx)
        ctx.Abort()
        return
    }
}

func AddProdToBasket() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var reqParam req.AddProdToBasketReq
        _ = ctx.ShouldBindJSON(&reqParam)

        if reqParam.ProdNums <= 0 {
            _ = bs.DeleteBasketRecord(reqParam)
            response.OkWithMessage("商品已移出购物车", ctx)
            ctx.Abort()
            return
        }

        if err := bs.InsertBasketRecord(reqParam); err != nil {
            log.Printf("insert basket failed : %v", err.Error())
            response.FailWithMessage(err.Error(), ctx)
            ctx.Abort()
            return
        }
        response.OkWithMessage("添加商品到购物车成功", ctx)
        ctx.Abort()
        return
    }
}
