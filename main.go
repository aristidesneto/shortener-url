package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type DataUrl struct {
  Url   string `json:"url"`
}

type MemoryDatabase struct {
  Id string `json:"id"`
  Url string `json:"url"`
}

var newUrls []MemoryDatabase
var baseUrl = os.Getenv("BASE_URL")

func shortUrl(c *gin.Context) {
  const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  var url DataUrl

  if baseUrl == "" {
    log.Fatalln("Missing variable BASE_URL")
  }

  if err := c.BindJSON(&url); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  var seedRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
  b := make([]byte, 5)
  for i := range b {
    b[i] = charset[seedRand.Intn(len(charset))]
  }

  dbMongo := MemoryDatabase{
    Id: string(b),
    Url: url.Url,
  }

  newUrls = append(newUrls, dbMongo)

  c.JSON(http.StatusOK, gin.H{
    "url": baseUrl + "/" + string(b),
  })
}

func getShortUrl(c *gin.Context)  {
  url := c.Param("url")

  for _, item := range newUrls {
    if item.Id == url {
      c.Redirect(http.StatusMovedPermanently, item.Url)
      return
    }
  }

  c.JSON(http.StatusBadRequest, gin.H{
    "error": "invalid url: " + baseUrl + "/" + url,
  })
}

func main()  {  
	r := gin.Default()

  r.POST("/short-url", shortUrl)
  r.GET("/:url", getShortUrl)
  
  r.Run()
}