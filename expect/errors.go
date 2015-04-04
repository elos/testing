package expect

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/models"
)

/*
 is a helper that ensures err
	is data.ErrAccessDenial and prints a failure message
	if not
*/
func AccessDenial(property string, err error, t *testing.T) {
	if err != data.ErrAccessDenial {
		t.Errorf("Expected access denial on %s, got %s", property, err)
	}
}

func EmptyLinkError(property string, err error, t *testing.T) {
	_, ok := err.(*data.EmptyLinkError)
	if !ok {
		t.Errorf("Expected empty link on %s, got %s", property, err)
	}
}

func EmptyRelationship(property string, err error, t *testing.T) {
	if err != models.ErrEmptyRelationship {
		t.Errorf("Expected empty relationship on %s, got %s", property, err)
	}
}

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
