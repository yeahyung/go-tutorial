package router

import (
	"github.com/gin-gonic/gin"
	"go-tutorial/gin-tutorial/handler"
)

type ServerRouter struct {
	routerGroup *gin.RouterGroup
	handler     *handler.ServerHandler
}

func NewServerRouter(routerGroup *gin.RouterGroup, handler *handler.ServerHandler) *ServerRouter {
	router := &ServerRouter{
		routerGroup: routerGroup,
		handler:     handler,
	}

	router.SetServerRouter()
	return router
}

func (a *ServerRouter) SetServerRouter() {
	v1 := a.routerGroup.Group("/api/v1/getPerson")
	{
		v1.GET("/:person-id", a.handler.GetPersonWithPathHandler)
	}
}
