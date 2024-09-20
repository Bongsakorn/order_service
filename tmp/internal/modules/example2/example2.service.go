package example2

import "order_service/internal/modules/example"

type ExampleService2Interface interface {
	GetExample() (string, error)
}

type ExampleService2 struct {
	ExampleService example.ExampleServiceInterface
}

func NewService(example.ExampleServiceInterface) ExampleService2Interface {
	return &ExampleService2{}
}

func (service *ExampleService2) GetExample() (string, error) {
	return "Hello, World! 2", error(nil)
}
