package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	conn *pgxpool.Pool
	name string
}

// creates a connection pool . The connection is closed when the program stops
func Connect(db string) (*Database, error) {
	if db == "postgres" {
		pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, errors.New(("Database Connection Error"))
		}

		return &Database{pool, "postgres"}, nil
	}

	if db == "file" {
		return &Database{nil, "file"}, nil
	}
	return nil, errors.New("invalid database type")
}

// this will return the sanitized file name ie take in the url request
// and return the file name that matches from DB
func (db *Database) DocSearch(url string) (string, error) {
	//for Postgres
	if db.name == "postgres" {
		var page string = ""
		var a string
		var err error
		var sql string
		fmt.Println("url: " + url)
		switch fileSuffix := url[strings.LastIndex(url, ".")+1:]; fileSuffix {
		case "html":
			sql = "select name from blog_posts;"
		case "css":
			sql = "select name from css;"
		case "png":
			sql = "select name from images;"
		case "svg":
			sql = "select name from images;"
		case "jpg":
			sql = "select name from images;"
		}
		fmt.Println("searching with running: " + sql + "for " + url)
		rows, err := db.conn.Query(context.Background(), sql)
		if err != nil {
			return "", errors.New("404")
		}
		pgx.ForEachRow(rows, []any{&a}, func() error {
			if a == url {
				page = a
			}
			return nil

		})

		if page == "" {
			err = errors.New("404")
		}
		return page, err
	}

	//for files
	if db.name == "file" {
		var file string = ""
		var pageError error = errors.New("404")
		drFileSlice, _ := os.ReadDir("pages/")
		for i := range drFileSlice {
			if url == drFileSlice[i].Name() {
				file = drFileSlice[i].Name()
				pageError = nil
				break
			}
		}

		return file, pageError
	}

	return "", errors.New("connection lost when searching post")
}

// returns the content of a entry be it a file or database entry
func (db *Database) SingleResult(name string, table string) ([]byte, error) {
	var err error
	var data []byte
	//fmt.Println(name)
	if db.name == "postgres" {
		var post *string
		var field string
		image := false
		switch fileSuffix := name[strings.LastIndex(name, ".")+1:]; fileSuffix {
		case "html":
			field = "post"

		case "css":
			field = "doc"

		case "png":
			field = "image_path"
			image = true

		case "svg":
			field = "image_path"
			image = true
		}
		fmt.Println("Looking for a specific document with: select " + field + " from " + table + " where name =" + name)
		err = db.conn.QueryRow(context.Background(), "select "+field+" from "+table+" where name =$1", name).Scan(&post)

		if err != nil {
			fmt.Println("query errored for content")
		} else {

			data = []byte(*post)
			//fmt.Println(fileName + " : " +*post)
		}

		if image {
			data, err = os.ReadFile(string(data))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	if db.name == "file" {
		data, err = os.ReadFile("pages/" + name)
	}
	return data, err
}


