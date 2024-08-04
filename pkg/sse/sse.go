package sse

import (
	"github.com/gin-gonic/gin"
	"github.com/r3labs/sse/v2"
	"time"
)

type Events interface {
	Publish(event string, data []byte)
}

type Handler interface {
	Serve(ctx *gin.Context)
}

type Server struct {
	server *sse.Server
}

func NewServer() *Server {
	server := sse.New()
	server.AutoStream = true
	server.EventTTL = time.Second * 30
	return &Server{server: server}
}

func (s *Server) Publish(event string, data []byte) {
	s.server.CreateStream(event)
	s.server.Publish(event, &sse.Event{

		Data: data,
	})
}

func (s *Server) Serve(ctx *gin.Context) {
	s.server.ServeHTTP(ctx.Writer, ctx.Request)
}
