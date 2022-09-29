package src

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

// Read a file and return content
func ReadFile(name string) []byte {
	b, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

// Convert []PathUrl to map[path]url
func buildMap(paths []PathUrl) map[string]string {
	ret := map[string]string{}

	for _, value := range paths {
		ret[value.Path] = value.Url
	}

	return ret
}

type PathUrl struct {
	Path string
	Url  string
}

func ParseYaml(data []byte) (map[string]string, error) {
	var m []PathUrl

	err := yaml.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}

	return buildMap(m), nil
}

// Struct for containing db connection and read statement
type DB struct {
	db   *sql.DB
	read *sql.Stmt
}

// Opens a connection to the database
func OpenDB(name string) (*DB, error) {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		return nil, err
	}

	// Prepare read statement
	read, err := db.Prepare("SELECT path, url FROM urls")
	if err != nil {
		return nil, err
	}

	return &DB{
		db:   db,
		read: read,
	}, nil
}

// Read paths & urls from DB, return values as map[path]url
func ReadDb(name string) (map[string]string, error) {
	conn, err := OpenDB(name)
	if err != nil {
		return nil, err
	}
	defer conn.db.Close()

	rows, err := conn.read.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paths []PathUrl
	for rows.Next() {

		var path PathUrl
		err := rows.Scan(&path.Path, &path.Url)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return buildMap(paths), nil

}
