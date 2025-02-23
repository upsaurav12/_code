package mock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockGreeter struct{}

func (m MockGreeter) Greet(name string) string {
	return "Mocked greeting for " + name
}

type Service struct {
	greeter Greeter
}

func (s *Service) SayHello(name string) string {
	return s.greeter.Greet(name)
}

func TestSayHello(t *testing.T) {
	// Create a service that uses the MockGreeter
	service := &Service{greeter: MockGreeter{}}

	// Test that SayHello uses the mocked Greet method
	result := SayHello(service.greeter, "Bob")

	fmt.Println(result)

	// Assert that the result matches the mocked behavior
	assert.Equal(t, "Mocked greeting for Bob", result)
}
