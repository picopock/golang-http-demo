package controller

import (
  "fmt"
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
)

func FormPost(c *gin.Context) {
  message := c.PostForm("message")
  nick := c.DefaultPostForm("nick", "anonymous")

  c.JSON(http.StatusOK, gin.H{
    "status":  "posted",
    "message": message,
    "nick":    nick,
  })
}

func Post(c *gin.Context) {
  id := c.Query("id")
  page := c.DefaultQuery("page", "0")
  name := c.PostForm("name")
  message := c.PostForm("message")

  fmt.Printf("id %s; page %s; name: %s; message: %s", id, page, name, message)
}

func PostMap(c *gin.Context) {
  ids := c.QueryMap("ids")
  names := c.PostFormMap("names")

  c.JSON(http.StatusOK, gin.H{
    "ids":   ids,
    "names": names,
  })
}

func Upload(c *gin.Context) {
  // single file
  file, _ := c.FormFile("file")
  log.Println(file.Filename)

  // c.SaveUploadedFile(file, dst)

  c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func UploadMulti(c *gin.Context) {
  // multiple files
  form, _ := c.MultipartForm()
  files := form.File["upload[]"]

  for _, file := range files {
    log.Println(file.Filename)

    // upload the file to specific dst
    // c.SaveUploadedFile(file, dst)
  }
  c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
