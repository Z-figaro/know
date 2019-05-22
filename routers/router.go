/**
 * @Author: figaro
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2019-05-22 16:55
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"know/pkg/setting"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"test",
		})
	})

	return r
}
