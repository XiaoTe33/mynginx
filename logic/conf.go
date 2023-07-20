package logic

import (
	"fmt"
	"log"
	"mynginx/pkg/myLog"
	"mynginx/pkg/ttviper"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Nginx struct {
	Servers []*Server
}

func (nginx *Nginx) Run() {
	for _, server := range nginx.Servers {
		err := http.ListenAndServe(":"+server.ListenOn, server)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
	time.Sleep(time.Hour)
}

type Server struct {
	ListenOn string
	Routers  []Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	myLog.Log.Infof("request in")
	myLog.Log.Infof("path:%v", r.URL.Path)
	for _, rt := range s.Routers {
		myLog.Log.Infof("location: %v, path: %v", rt.Location, r.URL.Path)
		if strings.HasPrefix(r.URL.Path, rt.Location) {
			if rt.Root != "" {
				rt.ServeFS(w, r)
			} else {
				rt.ServeHTTP(w, r, strings.TrimPrefix(r.URL.Path, rt.Location))
			}

		}
	}
}

// ParseNginx 把viper读取到的配置解析到结构体中
func ParseNginx(config *ttviper.Config) *Nginx {
	myLog.Log.Info("parsing config to struct...")
	c := new(Nginx)
	v := config.Viper

	httpC := v.Get("http")
	m, ok := httpC.([]any)
	if !ok {
		myLog.Log.Errorf("http field is not []any type")
		return nil
	}

	for _, a := range m {
		server := new(Server)
		serverC, ok2 := a.(map[string]any)
		if !ok2 {
			myLog.Log.Errorf("http elements should be map[string]any type")
			return nil
		}

		server.ListenOn = fmt.Sprintf("%v", serverC["listen_on"])

		routersC, ok3 := serverC["routers"].([]any)
		if !ok3 {
			myLog.Log.Errorf("routers should be a []any type")
			return nil
		}
		var rts []Router
		for _, routerCI := range routersC {
			rt := Router{}
			routerC, ok4 := routerCI.(map[string]any)
			if !ok4 {
				myLog.Log.Errorf("routers' elements should be map[string]any type")
				return nil
			}
			rt.BalanceMethod = fmt.Sprintf("%v", routerC["balancer"])
			rt.Location = fmt.Sprintf("%v", routerC["location"])
			rt.Root = fmt.Sprintf("%v", routerC["root"])
			rt.Index = fmt.Sprintf("%v", routerC["index"])

			proxysC, ok5 := routerC["proxy"].([]any)
			if rt.Root == "" {
				if !ok5 {
					myLog.Log.Errorf("proxys shuld be []any type")
					return nil
				}
				for _, proxyCI := range proxysC {
					proxy := Proxy{}
					proxyC, ok6 := proxyCI.(map[string]any)
					if !ok6 {
						myLog.Log.Errorf("proxy shuld be map[string]any type")
						return nil
					}
					proxy.IP = fmt.Sprintf("%v", proxyC["ip"])
					i, err := strconv.Atoi(fmt.Sprintf("%v", proxyC["weight"]))
					if err != nil {
						myLog.Log.Errorf("weight should be a integer")
						return nil
					}
					proxy.Weight = i
					rt.Proxy = append(rt.Proxy, proxy)
				}
			}
			rt.Balancer = NewBalancer(rt)
			rts = append(rts, rt)
		}
		server.Routers = rts
		c.Servers = append(c.Servers, server)
	}
	myLog.Log.Info("parse config to struct successfully")
	return c
}
