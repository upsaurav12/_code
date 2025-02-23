package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Add(2,3) = %d , want 5", result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	if result != 2 {
		t.Errorf("Subtract(5,3) = %d, want 2", result)
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(4, 2)
	if err != nil || result != 2 {
		t.Errorf("Divide(4,2) = %d, want 2", result)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10,0) did not return an error; want error")
	}
}

func TestReverseString(t *testing.T) {
	result := ReverseString("hello")
	expected := "olleh"

	if result != expected {
		t.Errorf("ReverseString(\"hello\") = %s; want %s", result, expected)
	}
}
