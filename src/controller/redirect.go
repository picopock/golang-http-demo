package controller

import (
  "io"
  "net/http"

  "github.com/gin-gonic/gin"
)

func RedirectChecker(req *http.Request, via []*http.Request) error {
  return http.ErrUseLastResponse
}

func RedirectReq1(c *gin.Context) {
  client := &http.Client{
    CheckRedirect: RedirectChecker,
  }
  // req1 反向代理到 req2
  resp, err := client.Get("http://127.0.0.1:8080/redirect/req2")
  if err != nil {
    c.JSON(200, gin.H{
      "message": "ohhhh, an error occured for request req2",
    })
    return
  }
  body, erro := io.ReadAll(resp.Body)
  if erro != nil {
    c.JSON(200, gin.H{
      "message": "ohhhh, an error occured for read resp body",
    })
    return
  }
  // 检测到 req2 302 重定向的动作，获取响应 location, redirect req1 请求到 location
  if resp.StatusCode == http.StatusFound {
    location := resp.Header.Get("Location")
    c.Redirect(http.StatusFound, location)
    return
  }

  c.JSON(200, string(body))
  // if err != nil && resp.StatusCode != http.StatusFound {
  //   location := resp.Header.Get("Location")
  //   c.Redirect(302, location)
  // }
  defer resp.Body.Close()
}

// 中间层重定向 req3
func RedirectReq2(c *gin.Context) {
  c.Redirect(302, "http://127.0.0.1:8080/redirect/req3")
}

func RedirectReq3(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "redirect success!",
  })
}
