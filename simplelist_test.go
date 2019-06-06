package simplelist

import (
	"testing"
)

const (
	Key1 = "key1"
	Key2 = "key2"
)

func TestBasic(t *testing.T) {
	l := New()
	if l.Exists(Key1) {
		t.Error("Did not expected Key1 to exist")
	}

	l.Add(Key1)

	if !l.Exists(Key1) {
		t.Error("Expected Key1 to exist")
	}

	l.Remove(Key1)

	if l.Exists(Key1) {
		t.Error("Did not expected Key1 to exist")
	}

	l.Set(Key1, true)

	if !l.Exists(Key1) {
		t.Error("Expected Key1 to exist")
	}

	l.Set(Key1, false)

	if l.Exists(Key1) {
		t.Error("Did not expected Key1 to exist")
	}

}
