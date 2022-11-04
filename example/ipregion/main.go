package main

import (
	_ "database/sql"
	_ "fmt"

	_ "github.com/mattn/go-sqlite3"

	"coinsky_go_project/example/ipregion/ip"
)

func main() {
	/*db, err := sql.Open("sqlite3", "ip.db")
	if err != nil {
		panic(err)
	}
	fmt.Println(db)*/

	//ip.Pull()
	ip.Create()

}
