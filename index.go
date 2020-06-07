package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

const (
	USERNAME = "test"
	PASSWORD = "test"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "blog"
)

func user(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "用户信息")
}

func getUser(w http.ResponseWriter, req *http.Request) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK,
		SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select id,name from user ")
	if err != nil {
		panic(err)
	}
	arr := []map[string]interface{}{}
	for rows.Next() {
		var m = make(map[string]interface{})
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
			return
		}

		m["id"] = id
		m["name"] = name
		arr = append(arr, m)
	}

	jsonString, _ := json.Marshal(arr)
	//fmt.Println(err)
	fmt.Fprintf(w, string(jsonString))

}

func main() {
	http.HandleFunc("/user", user)
	http.HandleFunc("/userIfo", getUser)
	http.ListenAndServe(":8091", nil)
}
