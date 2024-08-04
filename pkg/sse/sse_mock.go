package sse

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type SSEMock struct {
}

func NewSSEMock() *SSEMock {
	return &SSEMock{}
}

func (s *SSEMock) Publish(event string, data []byte) {
	fmt.Println("publish")
}

func (s *SSEMock) Serve(ctx *gin.Context) {
	fmt.Println("serve")
}
