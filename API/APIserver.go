package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allarticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "hello world"},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func getUsers() []*User {
	dbconf := "test:test13579@tcp(localhost:3306)/test"

	db, err := sql.Open("mysql", dbconf)

	if err != nil {
		log.Print(err.Error())
	}
	err = db.Ping()

	if err != nil {
		fmt.Println("データベース接続失敗")
	} else {
		fmt.Println("データベース接続成功")
	}
	// 接続が終了したらクローズする
	defer db.Close()

	results, err := db.Query("select * from product")
	if err != nil {
		panic(err)
	}

	var users []*User
	for results.Next() {
		var u User
		err = results.Scan(&u.ID, &u.Name, &u.Price)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Name: '%s',price: %d\n", u.ID, u.Name, u.Price)
		users = append(users, &u)
	}
	return users
}

func userPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	users := getUsers()
	fmt.Println("sql")
	json.NewEncoder(w).Encode(users)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/users", userPage).Methods("GET")
	myRouter.HandleFunc("/articles", allarticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequests()
}
