package types

import (
	"reflect"
	"testing"
)

func TestAutogeneration(t *testing.T) {
	manifest := new(Manifest)

	item := reflect.ValueOf(manifest.Organizations.Organization)
	res := reflect.ValueOf(manifest.Resources)

	if item.FieldByName("Item") == (reflect.Value{}) {
		t.Errorf("expected manifest to have an Item field")
	}

	if res.FieldByName("Resource") == (reflect.Value{}) {
		t.Errorf("expected manifest to have a []Resource field")
	}

	// todo extract metadata field?
}
