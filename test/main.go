package main

import (
	"mynginx/logic"
	"mynginx/pkg/ttviper"
	"net/http"
)

func main05() {
	handler := http.FileServer(http.Dir("./etc"))
	http.ListenAndServe("localhost:9090", handler)
}

func main() {
	conf := ttviper.ReadConfig("./etc", "config.yaml")
	c := logic.ParseNginx(conf)
	c.Run()
}

//func main03() {
//	conf := ttviper.ReadConfig()
//	val := conf.Viper.Get("http")
//	fmt.Println(val)
//	m, ok := val.([]any)
//	if !ok {
//		fmt.Println("!ok")
//	} else {
//		fmt.Println(m)
//	}
//	for _, a := range m {
//		m2, ok2 := a.(map[string]any)
//		if !ok2 {
//			fmt.Println("!ok2")
//		} else {
//			fmt.Println(m2)
//		}
//		fmt.Printf("listen_on: %v\n", m2["listen_on"])
//		m3, ok3 := m2["routers"].([]any)
//		if !ok3 {
//			fmt.Printf("!ok3")
//		} else {
//			for _, a2 := range m3 {
//				m4, ok4 := a2.(map[string]any)
//				if !ok4 {
//					fmt.Println("!ok4")
//				} else {
//					fmt.Printf("balancer: %v\n", m4["balancer"])
//					fmt.Printf("location: %v\n", m4["location"])
//					fmt.Printf("proxy: %v\n", m4["proxy"])
//					fmt.Printf("root: %v\n", m4["root"])
//					fmt.Printf("index: %v\n", m4["index"])
//				}
//			}
//		}
//
//	}
//}
//
//func main02() {
//	h := sha256.New()
//	h.Write([]byte("hello"))
//	s := string([]byte(fmt.Sprintf("%X", h.Sum([]byte(""))))[:10])
//
//	i32, err := strconv.ParseInt(s, 16, 64)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(i32)
//}
//
//func main01() {
//	go Gin()
//	time.Sleep(time.Second)
//	go http.ListenAndServe("localhost:8080", new(Handler))
//	time.Sleep(time.Hour)
//}
//
//type Handler struct {
//}
//
//func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	u, err := url.Parse("http://localhost:8081" + r.URL.Path)
//	if nil != err {
//		myLog.Log.Error(err)
//		return
//	}
//	proxy := httputil.ReverseProxy{
//		Director: func(r *http.Request) {
//			r.URL = u
//		},
//	}
//	proxy.ServeHTTP(w, r)
//}
//
//func Gin() {
//	r := gin.Default()
//	r.GET("/gin", func(c *gin.Context) {
//		c.String(200, "from gin")
//	})
//	r.Static("/text", "./etc")
//	r.Run(":8081")
//}
