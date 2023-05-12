package router

import (
  "github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
  applyPing(r)
  applyUser(r)
  applyForm(r)
  applyWelcome(r)
  applyRedirect(r)
}
