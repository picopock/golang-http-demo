package router

import (
  "github.com/gin-gonic/gin"
  "github.com/picopico/go-demo/src/controller"
)

func applyForm(r *gin.Engine) {
  form := r.Group("/form")
  {
    form.POST("/form_post", controller.FormPost)

    form.POST("/post", controller.Post)

    form.POST("/post_map", controller.PostMap)

    r.MaxMultipartMemory = 8 << 20 // 8 MiB
    form.POST("/upload", controller.Upload)

    form.POST("/upload_multi", controller.UploadMulti)
  }
}
