package crayonDB_test

import (
	"crayonDB"
	"testing"
)

func TestOpen(t *testing.T) {
	_, err := crayonDB.Open("", "Database")
	if err != nil {
		t.Log(err)
		return
	}
}

func TestUpdatePath(t *testing.T) {
	crayondb, err := crayonDB.Open("", "Database")
	if err != nil {
		t.Log(err)
		return
	}

	err = crayondb.UpdatePath("Users")
	if err != nil {
		t.Log(err)
	}
}

func TestIsPathExists(t *testing.T) {
	crayondb, err := crayonDB.Open("", "Database")
	if err != nil {
		t.Log(err)
		return
	}

	if !crayondb.IsPathExists("Users") {
		t.Log("Unexpected value")
	}
}

func TestIsDocExists(t *testing.T) {
	crayondb, err := crayonDB.Open("", "Database")
	if err != nil {
		t.Log(err)
		return
	}

	if crayondb.IsDocExists("Users", "userDoc") {
		t.Log("Unexpected value")
	}
}
