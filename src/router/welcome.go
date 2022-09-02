package router

import (
  "github.com/gin-gonic/gin"
  "github.com/picopico/go-demo/src/controller"
)

func applyWelcome(r *gin.Engine) {
  r.GET("/welcome", controller.Welcome)
}
