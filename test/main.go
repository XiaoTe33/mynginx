package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"mynginx/pkg/myLog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

func main() {
	h := sha256.New()
	h.Write([]byte("hello"))
	s := string([]byte(fmt.Sprintf("%X", h.Sum([]byte(""))))[:10])

	i32, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i32)
}

func main01() {
	go Gin()
	time.Sleep(time.Second)
	go http.ListenAndServe("localhost:8080", new(Handler))
	time.Sleep(time.Hour)
}

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("http://localhost:8081" + r.URL.Path)
	if nil != err {
		myLog.Log.Error(err)
		return
	}
	proxy := httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL = u
		},
	}
	proxy.ServeHTTP(w, r)
}

func Gin() {
	r := gin.Default()
	r.GET("/gin", func(c *gin.Context) {
		c.String(200, "from gin")
	})
	r.Static("/text", "./etc")
	r.Run(":8081")
}
