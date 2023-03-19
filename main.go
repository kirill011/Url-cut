package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

const (
	alpabet = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ" // Строка из которой shorter выбирает символы
	connStr = "user=postgres password=4526 dbname=Url-cut sslmode=disable"     // Строка для подключения к БД
)

var dFlag bool

var mapMas = make(map[string]string) // Карта для соответствия short(ключ) и link(значение)

func main() {
	dFlagptr := flag.Bool("d", false, "В зависимости от флага -d данные будут храниться либо в памяти(если нет флага -d), либо в PostgreSQl(если флаг присутствует)")
	flag.Parse()
	dFlag = *dFlagptr
	router := mux.NewRouter()              // инициализация роутера
	router.HandleFunc("/", indexPage)      // отслеживаем переход по localhost:8080/
	router.HandleFunc("/{key}", shortPage) //  отслеживаем переход по localhost:8080/{key}
	http.ListenAndServe(":8080", router)   // слушаем порт и по нему отображаем
}

// Структура, которая хранит:
//
// Link - ссылка которую нужно сократить.
//
// Short - её сокращённый код.
//
// DeleteTime - Время, когда эту ссылку необходимо удалить из PostgeSQL.
//
// Status - Статус обработки ссылки.
type Result struct {
	Link       string
	Short      string
	DeleteTime time.Time
	Status     string
}

// функция забирает со страницы link и помещает её в БД вместе с short
func indexPage(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("templates/index.html")
	result := Result{}
	result.Link = r.FormValue("s") // передаём link в экземпляр структуры Result

	if r.Method == "POST" {

		if !testUrl(r.FormValue("s")) { // Если link имеет неправильный формат

			result.Status = "Ссылка имеет неправильный формат"
			result.Link = ""
		} else {

			result.Short = shorter() // Помещаем результат от вызова shorter() в экземпляр структуры Result

			days, err := strconv.Atoi(r.FormValue("d"))              //Берём с формы количество дней до удаления
			if err != nil || days < 0 || days%1 != 0 || days > 365 { // Проверяем

				result.Status = "Количество дней до удаления должно быть целым неотрицательным числом" + string(rune(8804)) + " 365!"
				result.Link = ""
				fmt.Println(result.Link, result.Status)
				templ.Execute(w, result) //Передаём на сервер переменную result
				return
			}
			result.DeleteTime = time.Now().AddDate(0, 0, days) //Добавляем к текущей дате количество дней до удаления
			fmt.Println(result.DeleteTime)
			if dFlag { // Если присутствует флаг -d

				db, err := sql.Open("postgres", connStr) //Открываем соединение с БД

				if err != nil {

					panic(err)

				}

				defer db.Close()                                                                                               // После завершения функции закрываем соединение с БД
				db.Exec("insert into url (link, short, dt) values ($1, $2, $3)", result.Link, result.Short, result.DeleteTime) // Помещаем в БД значения link, short и deleteTime
			} else {

				mapMas[result.Short] = result.Link // ставим в соответствие ключ short и значение link

			}

			result.Status = "Сокращение выполнено" //Обновляем статус
		}
	}
	fmt.Println(result.Link, result.Status)
	templ.Execute(w, result) //Передаём на сервер переменную result
}

// Обработчик перехода по сокращённой ссылке
func shortPage(w http.ResponseWriter, r *http.Request) {

	var link string
	vars := mux.Vars(r) // Получаем параметры пути запроса
	if dFlag {          // Если присутствует флаг -d

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

// Новая функция для составления коротких ссылок
func shorter() string {
	if dFlag {
		var maxShort string
		var maxLen int
		db, err := sql.Open("postgres", connStr) // Открываем соединение с БД

		if err != nil {

			panic(err)

		}

		defer db.Close()                                                                                   // После завершения функции закрываем соединение с БД
		rowsMax := db.QueryRow("SELECT length(short) FROM public.url order by length(short) desc limit 1") // Берём из базы данных максимальную длинну short
		rowsMax.Scan(&maxLen)                                                                              // Записываем в maxLen

		rows := db.QueryRow("select short from url where length(short) = $1 order by short desc limit 1", maxLen) // Берём из базы данных short с максимальным значением
		rows.Scan(&maxShort)                                                                                      // Записываем в maxShort

		for i := len(maxShort) - 1; i >= 0; i-- {
			if maxShort[i] != 'Z' {
				return maxShort[:i] + string(alpabet[strings.Index(alpabet, string(maxShort[i]))+1]) + maxShort[i+1:]
			}
		}
		return string(alpabet[0]) + strings.Repeat(string(alpabet[0]), maxLen)
	} else {
		maxLen := 0
		keys := make([]string, 0, len(mapMas))
		for key, _ := range mapMas {
			if len(key) > maxLen {
				maxLen = len(key)
			}
			keys = append(keys, key)
		}
		maxShort := maxShort(keys)
		for i := len(maxShort) - 1; i >= 0; i-- {
			if maxShort[i] != 'Z' {
				return maxShort[:i] + string(alpabet[strings.Index(alpabet, string(maxShort[i]))+1]) + maxShort[i+1:]
			}
		}
		return string(alpabet[0]) + strings.Repeat(string(alpabet[0]), maxLen)
	}
}

// Функция определяет максимальное значение short при хранении локально
func maxShort(s []string) string {
	maxSort := "0"
	for _, val := range s {
		if len(val) > len(maxSort) {
			maxSort = val
			continue
		} else if len(val) < len(maxSort) {
			continue
		} else {
			for i, v := range val {
				if strings.Index(alpabet, string(v)) > strings.Index(alpabet, string(maxSort[i])) {
					maxSort = val
				}
			}
		}
	}
	return maxSort
}

// Проверяет введённый URL
func testUrl(rawUrl string) bool {

	matched, _ := regexp.MatchString(`[-a-zA-Z0-9@:%_\+.~#?&\/=]{2,256}\.[a-z]{2,4}\b(\/[-a-zA-Z0-9@:%_\+.~#?&\/=]*)?`, rawUrl)

	return matched
}
