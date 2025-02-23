package check

import (
	"fmt"
)

// Struct
type MyService struct{}

// ProcessStatus handles different statuses and returns errors using fmt.Errorf
func (s *MyService) ProcessStatus(status string) (string, error) {
	switch status {
	case "ok":
		return "Success", nil
	case "not_found":
		return "", fmt.Errorf("error: resource not found")
	case "timeout":
		return "", fmt.Errorf("error: request timeout")
	default:
		return "", fmt.Errorf("error: unknown status %q", status)
	}
}
