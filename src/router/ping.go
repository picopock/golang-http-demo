package router

import (
  "github.com/gin-gonic/gin"
  "github.com/picopico/go-demo/src/controller"
)

func applyPing(r *gin.Engine) {
  r.GET("/ping", controller.Ping)
}
