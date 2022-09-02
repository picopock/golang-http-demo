package router

import (
  "github.com/gin-gonic/gin"
  "github.com/picopico/go-demo/src/controller"
)

func applyUser(r *gin.Engine) {
  user := r.Group("/user")
  {
    user.GET("/:name", controller.UserName)

    user.GET("/:name/*action", controller.UserNameAction)

    user.POST("/:name/*action", controller.PostUserNameAction)

    user.GET("/groups", controller.UserGroup)
  }
}
