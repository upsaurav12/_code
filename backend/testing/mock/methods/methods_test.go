package methods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllOperations(t *testing.T) {
	operations := []struct {
		a             int
		b             int
		operation     string
		expected      int
		expectedError string
	}{
		{1, 2, "add", 3, ""},
		{4, 5, "subtract", -1, ""},
		{3, 2, "multiply", 6, ""},
		{9, 3, "divide", 3, ""},
		{9, 0, "divide", -1, "Cannot divide by zero"},
	}

	for _, operation := range operations {
		var result int
		var err error

		switch operation.operation {
		case "add":
			result = Add(operation.a, operation.b)
		case "subtract":
			result = Subtract(operation.a, operation.b)
		case "multiply":
			result = Multiply(operation.a, operation.b)
		case "divide":
			result, err = Divide(operation.a, operation.b)
		default:
			t.Fatalf("Invalid operation: %s", operation.operation)
		}

		if operation.expectedError != "" {
			assert.Equal(t, operation.expectedError, err.Error())
		} else {
			assert.NoError(t, err)
		}

		assert.Equal(t, operation.expected, result)
	}
}

/*
func TestAdd(t *testing.T) {
	adds := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{4, 5, 9},
	}

	for _, add := range adds {
		result := Add(add.a, add.b)
		assert.Equal(t, add.expected, result)
	}
}

func TestSubtract(t *testing.T) {
	subtracts := []struct {
		a        int
		b        int
		expected int
	}{
		{3, 2, 1},
		{9, 5, 4},
	}

	for _, subtract := range subtracts {
		result := Subtract(subtract.a, subtract.b)
		assert.Equal(t, subtract.expected, result)
	}
}*/
