package main

import (
  "github.com/gin-gonic/gin"
  "github.com/picopico/go-demo/src/middleware"
  "github.com/picopico/go-demo/src/router"
)

func main() {
  // create a router without any middleware by default
  // r := gin.New()
  // Default with the logger and Recovery middleware already attached
  r := gin.Default()
  // register middleware
  middleware.RegisterMiddleware(r)
  // register router
  router.RegisterRouter(r)
  // start server
  r.Run("127.0.0.1:8080")
}
