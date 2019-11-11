package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {

	var err error

	p := log.Println // the alias for log.Println in order to simplify the code

	p("-------------------", time.Now().Format(http.TimeFormat), "A 'task 3' request is received :) -------------------")
	p("На сервер обратился клиент с адреса", r.RemoteAddr)


	// Разбираем query
	say := r.URL.Query().Get("say")
	if say == "" {
		say = "Say something by the fallowing way: url?say=[yourphrase]"
		_, err = fmt.Fprintf(w, say)
		if err != nil {
			p("Ошибка работы веб-сервера!")
			fmt.Fprintf(w, "Ошибка работы веб-сервера!")
		}
	} else {
		_, err = fmt.Fprintf(w, "Hello, %s\n", say)
		if err != nil {
			p("Ошибка работы веб-сервера!")
			fmt.Fprintf(w, "Ошибка работы веб-сервера!")
		}
	}
	p("Обработка запроса успешно завершена")
}

func main() {
	p := log.Println // the alias for log.Println in order to simplify the code

	// Обслуживание корня сайта и всё что выше
	http.HandleFunc("/", hello)
	myServer := http.Server{ // настройки моего сервера
		ReadHeaderTimeout: 20 * time.Second, // защита от SLOWLORIS
		ReadTimeout:       1 * time.Minute,  // защита от SLOWLORIS
		WriteTimeout:      2 * time.Minute,  // защита от SLOWLORIS
		Addr:              ":8181",
	}

	// Запуск сервера
	p("Server is listening...")
	myServer.ListenAndServe()
}
