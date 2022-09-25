package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	alpabet = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // Строка из которой shorter выбирает символы
	connStr = "user=postgres password=4526 dbname=Url-cut sslmode=disable"     // Строка для подключения к БД
)

var mapMas = make(map[string]string) // Карта для соответствия short(ключ) и link(значение)

func main() {

	rand.Seed(time.Now().UTC().UnixNano()) // Устанавливаем сид для рандомайзера (без него выдаёт одинаковые значения)
	router := mux.NewRouter()              // инициализация роутера
	router.HandleFunc("/", indexPage)      // отслеживаем переход по localhost:8080/
	router.HandleFunc("/{key}", shortPage) //  отслеживаем переход по localhost:8080/{key}
	http.ListenAndServe(":8080", router)   // слушаем порт и по нему отображаем
}

type Result struct {
	Link   string
	Short  string
	Status string
}

func indexPage(w http.ResponseWriter, r *http.Request) { //функция забирает со страницы link и помещает её в БД вместе с short
	templ, _ := template.ParseFiles("templates/index.html")
	result := Result{}

	if r.Method == "POST" {

		if !testUrl(r.FormValue("s")) { // Если link имеет неправильный формат

			result.Status = "Неправильный формат!"
			result.Link = ""

		} else {

			result.Link = r.FormValue("s") // передаём link в экземпляр структуры Result
			result.Short = shorter()       // Помещаем результат от вызова shorter() в экземпляр структуры Result

			if os.Args[len(os.Args)-1] == "-d" { // Если последний аргумент командной строки совпадает с "-d"

				db, err := sql.Open("postgres", connStr) //Открываем соединение с БД

				if err != nil {

					panic(err)

				}

				defer db.Close()                                                                    // После завершения функции закрываем соединение с БД
				db.Exec("insert into url (link, short) values ($1, $2)", result.Link, result.Short) // Помещаем в БД значения link и short
			} else {

				mapMas[result.Short] = result.Link // ставим в соответствие ключ short и значение link

			}

			result.Status = "Сокращение было выполнено успешно" //Обновляем статус
		}
	}

	templ.Execute(w, result) //Передаём на сервер переменную result
}

func shortPage(w http.ResponseWriter, r *http.Request) {

	var link string
	vars := mux.Vars(r)                  // Получаем параметры пути запроса
	if os.Args[len(os.Args)-1] == "-d" { // Если последний аргумент командной строки совпадает с "-d"

		db, err := sql.Open("postgres", connStr) // Открываем соединение с БД

		if err != nil {

			panic(err)

		}

		defer db.Close() // После завершения функции закрываем соединение с БД

		rows := db.QueryRow("select link from url where short=$1 limit 1", vars["key"]) // Берём из базы данных link, которой соответствует vars["key"]
		rows.Scan(&link)
	} else {

		link = mapMas[vars["key"]] // Берём из карты link, которой соответствует vars["key"]

	}

	fmt.Fprintf(w, "<script>location='%s';</script>", link) //Вставляем link в поисковый запрос

}

func shorter() string { // Берёт из строки alphabet 32 псевдослучайных элемента и выдаёт их одной строкой

	b := make([]byte, 32)

	for i := range b {
		b[i] = alpabet[rand.Intn(len(alpabet))]
	}
	return string(b)
}

func testUrl(rawUrl string) bool { //Проверяет формат ссылки

	matched, _ := regexp.MatchString(`http(s)?:\/\/(((\w+)\.(\w+)(.\w+)?)\/|(\w+):(\d+))(\/)?(\S+)?`, rawUrl)

	if !matched {
		return false
	}

	return true
}
