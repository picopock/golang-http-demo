package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
  firstname := c.DefaultQuery("firstname", "Guest")
  lastname := c.Query("lastname")

  c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
