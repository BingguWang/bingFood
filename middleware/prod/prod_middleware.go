package prod

import (
	"bingFood/common/response"
	"bingFood/entity/prod"
	"bingFood/entity/prod/req"
	"bingFood/global"
	"bingFood/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 获取商品列表
func GetProdList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var param req.ListProdReq
		_ = ctx.ShouldBindJSON(&param)

		fmt.Println(utils.ToJsonString(&param))
		db := global.MYSQL_DB

		var prodList []prod.Prod
		limit := param.PageInfo.PageSize
		offset := param.PageInfo.PageSize * (param.PageInfo.Page - 1)
		var total int64
		db.Where("shop_id = ?", param.ShopId).Count(&total)

		fmt.Println(utils.ToJsonString(param.ShopId))

		if err := db.Limit(limit).Offset(offset).
			Preload("Skus").Preload("Properties").
			Where(&prod.Prod{ShopId: param.ShopId}). // 这样当没传入shopId时零值就不会作为条件了
			Find(&prodList).Error; err != nil {
			response.FailWithMessage(fmt.Sprintf("获取列表失败, err:%v", err.Error()), ctx)
			ctx.Abort()
			return
		}
		response.OkWithDetailed(response.PageResult{
			List:     prodList,
			Total:    total,
			Page:     param.PageInfo.Page,
			PageSize: param.PageInfo.PageSize,
		}, "获取列表成功", ctx)
		ctx.Abort()
		return
	}
}
