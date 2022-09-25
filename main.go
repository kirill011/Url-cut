package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	connStr     = "user=postgres password=4526 dbname=Url-cut sslmode=disable"
)

var mapMas = make(map[string]string)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage)
	router.HandleFunc("/{key}", shortPage)
	http.ListenAndServe(":8080", router)
}

type Result struct {
	Link   string
	Short  string
	Status string
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("templates/index.html")
	result := Result{}

	if r.Method == "POST" {

		if !testUrl(r.FormValue("s")) {

			result.Status = "Неправильный формат!"
			result.Link = ""

		} else {

			result.Link = r.FormValue("s")
			result.Short = shorter()

			if os.Args[len(os.Args)-1] == "-d" {

				db, err := sql.Open("postgres", connStr)

				if err != nil {

					panic(err)

				}

				defer db.Close()
				db.Exec("insert into url (link, short) values ($1, $2)", result.Link, result.Short)
			} else {

				mapMas[result.Short] = result.Link

			}

			result.Status = "Сокращение было выполнено успешно"
		}
	}

	templ.Execute(w, result)
}

func shortPage(w http.ResponseWriter, r *http.Request) {

	var link string
	vars := mux.Vars(r)
	if os.Args[len(os.Args)-1] == "-d" {

		db, err := sql.Open("postgres", connStr)

		if err != nil {

			panic(err)

		}

		defer db.Close()

		rows := db.QueryRow("select link from url where short=$1 limit 1", vars["key"])
		rows.Scan(&link)
	} else {

		link = mapMas[vars["key"]]

	}

	fmt.Fprintf(w, "<script>location='%s';</script>", link)

}

func shorter() string {

	b := make([]byte, 32)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func testUrl(rawUrl string) bool {

	matched, _ := regexp.MatchString(`http(s)?://(((\w+)\.(\w+)(.\w+)?)|(\w+):(\d+))(/)?(\S+)?`, rawUrl)

	if !matched {
		return false
	}

	_, err := url.ParseRequestURI(rawUrl)

	if err != nil {

		return false

	}
	u, err := url.Parse(rawUrl)

	if err != nil || u.Host == "" {
		return false
	}
	return true
}
