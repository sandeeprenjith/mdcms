package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
	"fmt"
)

type content struct {
	Title string 
	Content string 
}

func (c content) Date() string {
	t :=  time.Now().Local().Format("2006-01-02")
	return t
}

func CreateNewContent(title string) {
	db, _ := sql.Open("mysql", "root:tester@tcp(127.0.0.1:3306)/mdcms")
	defer db.Close()

	c := content{Title: title}

	query1 := "delete from content where title = \"" + c.Title + "\";" 
	_, _ = db.Exec(query1)
}

func main() {
	CreateNewContent("testing")
}


