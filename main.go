package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"website/database"
)

// used by the listen and serve function to serve the webpage
func viewHandler(w http.ResponseWriter, r *http.Request) {
	var page *template.Template
	var homepage = "body.html"
	pageName, err := db.DocSearch(r.URL.Path[len("/"):])
	if err != nil {
		page, _ = intoTemplate(get404())
		page.Execute(w, "")
		return
	}

	switch fileSuffix := pageName[strings.LastIndex(pageName, ".")+1:]; fileSuffix {
	case "html":
		data, err := getContent(pageName, "blog_posts")
		if err != nil {
			data = get404()
		}
		page, err = intoTemplate(data)
		if err != nil {
			page, _ = intoTemplate(get404())
		}
		page.Execute(w, "")

	case "":
		//this handles the homepage
		data, err := getContent(homepage, "blog_posts")
		if err != nil {
			data = get404()
		}
		page, err = intoTemplate(data)
		if err != nil {
			page, _ = intoTemplate(get404())
		}
		page.Execute(w, "")

	case "css":
		data, err := getContent(pageName, "css")
		if err != nil {
			data = nil
		}
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Write(data)

	case "png":
		data, err := getContent(pageName, "images")
		if err != nil {
			data = nil
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(data)

	case "svg":
		data, err := getContent(pageName, "images")
		if err != nil {
			data = nil
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		w.Write(data)
	default:
		println("unsupported file type: " + fileSuffix)
	}

}

// maybe this is not needed as it is just wrapping a function call
func getContent(name string, table string) ([]byte, error) {
	return db.SingleResult(name, table)
}

// takes in data, gets the header template and merges them into a single
// tempate
func intoTemplate(data []byte) (*template.Template, error) {
	var strBase string
	var strHeader string

	base, err := getContent("home.html", "blog_posts")
	if err != nil {
		strBase = ""
	} else {
		strBase = string(base)
	}

	header, err := getContent("head.html", "blog_posts")
	if err != nil {
		strHeader = ""
	} else {
		strHeader = string(header)
	}

	FullTemplate, err := template.New("final").Parse(strBase + strHeader + string(data))
	return FullTemplate, err
}

// returns the data for the 404 content
func get404() []byte {
	data, _ := getContent("404.html", "blog_posts")
	return data
}


// global database variable
var db *database.Database

func main() {
	// chage this to change the db
	selectedDatabase := "postgres"
	var err error
	db, err = database.Connect(selectedDatabase)

	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to connect to", selectedDatabase, "database")
		return
	}

	http.HandleFunc("/", viewHandler)
	//change the ListenAndServeTLS to add https requirment later
	http.ListenAndServe("localhost:8080", nil)
}

/*Certificate is saved at: /usr/local/etc/letsencrypt/live/wieds.ca/fullchain.pem
Key is saved at:         /usr/local/etc/letsencrypt/live/wieds.ca/privkey.pem
This certificate expires on 2024-10-14.
*/
