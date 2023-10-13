package main

import (
    "errors"
    "math"
    "testing"
)

func TestSquareGetValue(t *testing.T) {
    // Test a valid square
    s := square{5}
    area, err := s.getValue()
    if err != nil {
        t.Errorf("Expected no error, got: %v", err)
    }
    if area != 25 {
        t.Errorf("Expected area of 25, got: %.2f", area)
    }

    // Test a square with a negative side
    s = square{-5}
    area, err = s.getValue()
    expectedError := errors.New("The Value Can't be negative")
    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error: %v, got: %v", expectedError, err)
    }
}

func TestCircleGetValue(t *testing.T) {
    // Test a valid circle
    c := circle{3}
    area, err := c.getValue()
    if err != nil {
        t.Errorf("Expected no error, got: %v", err)
    }
    expectedArea := math.Pi * 9
    if area != expectedArea {
        t.Errorf("Expected area of %.2f, got: %.2f", expectedArea, area)
    }

    // Test a circle with a negative radius
    c = circle{-3}
    area, err = c.getValue()
    expectedError := errors.New("The Value cant be negative")
    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error: %v, got: %v", expectedError, err)
    }
}
