package logic

import (
	"mynginx/pkg/myLog"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Router struct {
	Location      string
	Root          string
	Index         string
	BalanceMethod string
	Balancer      Balancer
	Proxy         []Proxy
}

type Proxy struct {
	IP     string
	Weight int
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request, pathAfter string) {
	ip := router.Proxy[router.Balancer.Index()].IP
	myLog.Log.Infof("request redirect: %v----->%v", r.URL, ip+pathAfter)
	u, err := url.Parse(ip + pathAfter)
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

func (router Router) ServeFS(w http.ResponseWriter, r *http.Request) {
	myLog.Log.Warn(router.Root)
	http.StripPrefix(router.Location, http.FileServer(http.Dir(router.Root))).ServeHTTP(w, r)
}
