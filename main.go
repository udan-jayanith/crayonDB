// crayonDB is a concurrency-safe NoSQL database that enforces writing data to a database in a structured way.
//
// It stores data as JSON files, making it easier to send over a network.
package crayonDB

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// CrayonDB Database
type CrayonDB struct {
	DatabasePath string
}

// OpenedDB gives a map of Opened CrayonDB where key is path. And value is CrayonDB struct.
var OpenedDB map[string]*CrayonDB

// Open() open's the database. If it doesn't exists Open will create a database and return CrayonDB, nil.
func Open(path, name string) (CrayonDB, error) {
	dbPath := filepath.Join(path, name)
	crayonDB := CrayonDB{
		DatabasePath: dbPath,
	}

	if OpenedDB == nil {
		OpenedDB = make(map[string]*CrayonDB)
	}
	OpenedDB[path] = &crayonDB

	err := os.MkdirAll(dbPath, os.ModePerm)
	return *(OpenedDB[path]), err
}

// IsPathExists() returns true if the path exists. else false.
func (cdb *CrayonDB) IsPathExists(path string) bool {
	folderPath := filepath.Join(cdb.DatabasePath, path)
	info, err := os.Stat(folderPath)

	return (err == nil && info.IsDir())
}

// IsDocExists() returns true if the doc exists. else false.
func (cdb *CrayonDB) IsDocExists(path, doc string) bool {
	folderPath := filepath.Join(cdb.DatabasePath, path, doc+".json")
	info, err := os.Stat(folderPath)

	return err == nil && !info.IsDir()
}

func IsDocNameValid(doc string) {

}

// UpdatePath() creates a directory named path, along with any necessary parents, and returns nil, or else returns an error.
//
// If a path named directory exists UpdatePath() does nothing and return nil.
func (cdb *CrayonDB) UpdatePath(path string) error {
	folderPath := filepath.Join(cdb.DatabasePath, path)
	err := os.MkdirAll(folderPath, os.ModePerm)
	return err
}

// UpdateDoc() replaces the doc with the given struct.
// If path doesn't exists UpdateDoc() returns err = PathDoesNotExists
// If doc doesn't exists UpdateDoc() will create the doc with given struct.
func (cdb *CrayonDB) UpdateDoc(path, doc string, docStruct any) error {
	if !cdb.IsPathExists(path) {
		return fmt.Errorf(PathDoesNotExists)
	}

	docPath := filepath.Join(cdb.DatabasePath, path, doc+".json")
	jsonByte, err := json.Marshal(docStruct)
	if err != nil {
		return err
	}
	err = os.WriteFile(docPath, jsonByte, os.ModePerm)
	return err
}

// GetDocAsJson() returns doc as it is saved.
func (cdb *CrayonDB) GetDocAsBytes(path, doc string) ([]byte, error) {
	docPath := filepath.Join(cdb.DatabasePath, path, doc+".json")
	DocBytes, err := os.ReadFile(docPath)
	if err != nil {
		return DocBytes, err
	}

	return DocBytes, err
}

// WriteDocAsJson() writes to a writer.
func (cdb *CrayonDB) WriteDocAsJson(w io.Writer, path, doc string) {
	//json.NewEncoder(w)
}

// GetDoc() takes a address of a struct and change it with the doc.
func (cdb *CrayonDB) GetDoc(path, doc string, _ any) {}

// GetFieldFromDoc() returns the field value from given path and doc.
func GetFieldFromDoc() {}

// GetItemsInPath() returns a map of folders and docs in the path as key as the name of the folder or doc and value as item type.
func (cdb *CrayonDB) GetItemsInPath(path string) {}

//Delete Options
