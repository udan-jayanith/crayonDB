package crayonDB_test

import (
	"crayonDB"
	"testing"
)

func TestOpen(t *testing.T) {
	_, err := crayonDB.Open(crayonDB.CurrentDir, "Database")
	if err != nil {
		t.Log(err)
		return
	}
}

func TestUpdatePath(t *testing.T) {
	crayondb, err := crayonDB.Open(crayonDB.CurrentDir, "Database")
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
	crayondb, err := crayonDB.Open(crayonDB.CurrentDir, "Database")
	if err != nil {
		t.Log(err)
		return
	}

	if !crayondb.IsPathExists("Users") {
		t.Log("Unexpected value")
	}
}

func TestIsDocExists(t *testing.T) {
	crayondb, err := crayonDB.Open(crayonDB.CurrentDir, "Database")
	if err != nil {
		t.Log(err)
		return
	}

	if crayondb.IsDocExists("Users", "userDoc") {
		t.Log("Unexpected value")
	}
}

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func TestUpdateDoc(t *testing.T) {
	user := User{
		FirstName: "Udan",
		LastName:  "Jayakody",
		Age:       16,
	}

	crayondb, err := crayonDB.Open(crayonDB.CurrentDir, "Database")
	if err != nil {
		t.Log(err)
		return
	}

	err = crayondb.UpdateDoc("Users", user.FirstName, &user)
	if err != nil {
		t.Log(err)
		return
	}

	err = crayondb.UpdateDoc("user", user.FirstName, &user)
	if err == nil {
		t.Log("Unexpected behavior.")
	}
}

func TestGetDocAsBytes(t *testing.T) {
	crayondb, err := crayonDB.Open(crayonDB.CurrentDir, "Database")
	if err != nil {
		t.Log(err)
		return
	}

	_, err = crayondb.GetDocAsBytes("Users", "Udan")
	if err != nil {
		t.Log("Unexpected behavior.")
		return
	}

	m, err := crayondb.GetDocAsBytes("Users", "uhdan")
	if err == nil {
		t.Log(err)
		t.Log(m)
	}
}
