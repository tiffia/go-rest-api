module http.example.com

go 1.18

require http.example.com/backend v0.0.0

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/mattn/go-sqlite3 v1.14.32 // indirect
)

replace http.example.com/backend => ../backend
