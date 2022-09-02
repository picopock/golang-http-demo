package controller

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func UserName(c *gin.Context) {
  name := c.Param("name")
  c.String(http.StatusOK, "Hello %s", name)
}

func UserNameAction(c *gin.Context) {
  name := c.Param("name")
  action := c.Param("action")
  message := name + " is " + action
  c.String(http.StatusOK, message)
}

func PostUserNameAction(c *gin.Context) {
  b := c.FullPath() == "/user/:name/*action"
  c.String(http.StatusOK, "is true ? %t", b)
}

func UserGroup(c *gin.Context) {
  c.String(http.StatusOK, "The available groups are [...]")
}
