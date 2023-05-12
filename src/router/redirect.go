package router

import (
  "github.com/gin-gonic/gin"
  "github.com/picopico/go-demo/src/controller"
)

func applyRedirect(r *gin.Engine) {
  redirect := r.Group("/redirect")
  {
    redirect.GET("/req1", controller.RedirectReq1)
    redirect.GET("/req2", controller.RedirectReq2)
    redirect.GET("/req3", controller.RedirectReq3)
  }
}
