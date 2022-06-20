package shop

import (
    "bingFood/common/response"
    "bingFood/entity/shop"
    "bingFood/entity/shop/req"
    "bingFood/entity/shop/resp"
    "bingFood/global"
    "bingFood/utils"
    "fmt"
    "github.com/gin-gonic/gin"
)

// 获取商家列表
func GetShopList() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var shopListReq req.ShopReq
        _ = ctx.ShouldBindJSON(&shopListReq)
        fmt.Println(shopListReq)
        shopCond := shopListReq.ShopCond

        limit := shopListReq.PageInfo.PageSize
        offset := shopListReq.PageInfo.PageSize * (shopListReq.PageInfo.Page - 1)
        fmt.Println(utils.ToJsonString(shopCond))

        db := global.MYSQL_DB
        var shopList []shop.Shop
        var count int64
        db.Where(&shopCond).Count(&count)
        if err := db.Limit(limit).
            Offset(offset).
            Where(&shopCond).
            Find(&shopList).Error; err != nil {
            response.FailWithMessage(fmt.Sprintf("获取列表失败,err:%v", err), ctx)
            ctx.Abort()
            return
        }
        response.OkWithDetailed(response.PageResult{
            List:     shopList,
            Total:    count,
            Page:     shopListReq.PageInfo.Page,
            PageSize: shopListReq.PageInfo.PageSize,
        }, "获取列表成功", ctx)
        ctx.Abort()
        return
    }
}

// 获取商家详情
func GetShopDetail() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        var shopListReq req.ShopReq
        _ = ctx.ShouldBindJSON(&shopListReq)
        fmt.Println(shopListReq)
        shopCond := shopListReq.ShopCond

        limit := shopListReq.PageInfo.PageSize
        offset := shopListReq.PageInfo.PageSize * (shopListReq.PageInfo.Page - 1)
        fmt.Println(utils.ToJsonString(shopCond))

        db := global.MYSQL_DB
        var shopDetailResp resp.ShopDetailResp

        if err := db.Limit(limit).
            Offset(offset).
            Where(&shopCond).
            Take(&shopDetailResp.Shop).Error; err != nil {
            response.FailWithMessage(fmt.Sprintf("获取列表失败,err:%v", err), ctx)
            ctx.Abort()
            return
        }

        if err := db.Limit(limit).
            Offset(offset).
            Where(&shopListReq.ShopCond).
            Take(&shopDetailResp.ProdList).Error; err != nil {
            response.FailWithMessage(fmt.Sprintf("获取列表失败,err:%v", err), ctx)
            ctx.Abort()
            return
        }

        response.OkWithDetailed(shopDetailResp, "获取列表成功", ctx)
        ctx.Abort()
        return
    }
}
