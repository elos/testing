package modeltest

import (
	"testing"

	"github.com/elos/data"
	"github.com/elos/testing/expect"
	"gopkg.in/mgo.v2/bson"
)

/*
	TestAnonReadAccess ensures that an anonymous access can not
	read a model
*/
func AnonReadAccess(s data.Store, m data.Model, t *testing.T) {
	if err := s.Save(m); err != nil {
		t.Fatalf("Error while saving model: %s", err)
	}

	access := data.NewAnonAccess(s)

	m, err := access.Unmarshal(m.Kind(), data.AttrMap{
		"id": m.ID().(bson.ObjectId).Hex(),
	})

	if err != nil {
		t.Errorf("Error while unmarshalling user: %s", err)
	}

	expect.AccessDenial("Reading User Anonymously", access.PopulateByID(m), t)
}
