package nginx

import (
	"github.com/gin-gonic/gin"
	"mynginx/logic/router"
)

type Config struct {
	Servers []*Server
}

type Server struct {
	ListenOn string
	Routers  []router.Router
	Engine   *gin.Engine
}

func (s *Server) Srv() {
	for _, r := range s.Routers {
		if r.IsStatic {
			s.Engine.Static(r.Location, r.Root)
		} else {

		}
	}

}
