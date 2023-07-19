package router

import "mynginx/logic/balancer"

type Router struct {
	IsStatic      bool
	Location      string
	Root          string
	Index         string
	BalanceMethod string
	Balancer      balancer.Balancer
	Proxy         []Proxy
}

type Proxy struct {
	IP     string
	Weight int
}
