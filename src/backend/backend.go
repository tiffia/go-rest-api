// package backend

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, World!")
// }

// func Run(addr string) {
// 	http.HandleFunc("/", helloWorld)
// 	fmt.Println("Starting server on: ", addr)
// 	log.Fatal(http.ListenAndServe(addr, nil))
// }

package backend

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB   *sql.DB
	Port string
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func (a *App) Initialize() {
	var err error
	a.DB, err = sql.Open("sqlite3", "../../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (a *App) Run() {
	http.HandleFunc("/", helloWorld)
	log.Printf("Starting server on %s", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))

}
