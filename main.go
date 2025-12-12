// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"log"
// )

// func helloWorld(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Hello, World!\n")
// }

// func main() {
// 	http.HandleFunc("/", helloWorld)
// 	fmt.Println("Starting server on localhost:8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id    		int
	name  		string
	inventory int
	price 		float64
}

func main() {
	db, err := sql.Open("sqlite3", "./practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("SELECT id, name, inventory, price FROM products")
	if err!= nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p Product

		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Println("Product ID: ", p.id, "Product Name: ", p.name, "Product Inventory: ", p.inventory, "Product Price: ", p.price)
	}

}