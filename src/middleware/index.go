package middleware

import "github.com/gin-gonic/gin"

func RegisterMiddleware(r *gin.Engine) {
  // Global middleware
  // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release
  // By default gin.DefaultWrite =  os.Stdout
  r.Use(gin.Logger())
  // Recovery middleware recovers from any panics and write a 500 if there was one.
  r.Use(gin.Recovery())

}
