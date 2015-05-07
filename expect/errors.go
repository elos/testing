package expect

import (
	"testing"
)

func NoError(op string, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Error while %s: %s", op, err)
	}
}

func Error(ee error, op string, err error, t *testing.T) {
	if err != ee {
		t.Fatalf("Expected %s while %s, but got %s", ee, op, err)
	}
}
