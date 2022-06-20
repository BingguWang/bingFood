package prod

import (
    "bingFood/middleware/prod"
    "github.com/gin-gonic/gin"
)
func ProdRouter(r *gin.Engine) {
    group := r.Group("/shop")
    {
        CommonGroup := group.Group("/prod")
        //CommonGroup.Use(middleware.JWTAuthMiddleware()) // 这应该是商家用户的校验
        {
            CommonGroup.POST("/list", prod.GetProdList())
        }
    }
}
