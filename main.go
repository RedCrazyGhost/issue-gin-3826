package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

)

func Method(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _method, ok := c.GetQuery("_method"); ok && strings.ToUpper(_method) != c.Request.Method {
			c.Request.Method = strings.ToUpper(_method)
			r.HandleContext(c)
			c.Abort()
		} else if _header := c.GetHeader("X-HTTP-Method-Override"); _header != "" && strings.ToUpper(_header) != c.Request.Method {
			c.Request.Method = strings.ToUpper(_header)
			r.HandleContext(c)
			c.Abort()
		}else{
			c.Next()
		}
	}
}


func main() {
	r := gin.Default()
	r.Use(Method(r))
	
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"GET"})
	})
	
	r.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"POST"})
	})
	r.PUT("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"PUT"})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}