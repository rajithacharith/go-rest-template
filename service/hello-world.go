package service

import (
	"context"
)

type HelloWorldService interface {
	GetHelloWorld(ctx context.Context) string
}

func NewHelloWorldService() HelloWorldService {
	return &helloWorldService{}
}

type helloWorldService struct{}

func (h *helloWorldService) GetHelloWorld(ctx context.Context) string {
	return "Hello World"
}
