package routers

import (
	_ "applyUpLoadFile/docs"
	"applyUpLoadFile/routers/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"strings"
)

func InitRouter(r *gin.Engine, g []string) {

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	for i, _ := range g {
		matchRouter(g[i], r)
	}
}

func matchRouter(tag string, r *gin.Engine) {
	path := fmt.Sprintf("%s/%s", strings.ToLower(tag), strings.ToLower(""))
	switch tag {
	case "web":
		r.Static("/static", "./static")
		web.LoadWebRouter(r.Group(path))
		break
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		//c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		//c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length,access-control-allow-origin, access-control-allow-headers")
		//c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
