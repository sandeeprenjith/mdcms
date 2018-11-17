package main

import (
	_ "github.com/go-sql-driver/mysql" 
	"database/sql"
	"crypto/sha256"
	"encoding/base64"
	"github.com/russross/blackfriday"
	"net/http"
	"html/template"
)
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
type user {
	User string
	Password string
}

type content {
	Title string
	Date string
	Content string
}

func (c content) Date() string {
        t :=  time.Now().Local().Format("2006-01-02")
        return t
}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////
//Error handler
func handle_err(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//Create password hash
func passwordhash(password string) string {
        passbyte := []byte(password)
        passhash := sha256.New()
        passhash.Write(passbyte)
        phencoded := base64.URLEncoding.EncodeToString(passhash.Sum(nil))
        return phencoded
}

//Create tables
func CreateTables() {
        db, err := sql.Open("mysql", "root:tester@tcp(127.0.0.1:3306)/mdcms")
	handle_err(err)
        defer db.Close()

        c := content{"example", time.Now().Local().Format("2006-01-25"), "asoaijsdojaosjdoajdoiajsodjosaij" }
        _, err = db.Exec("create table content ( title VARCHAR(50), date DATE, content MEDIUMTEXT);" )
	handle_err(err)

}

//Creating new content
func CreateNewContent(title string, newcontent string) {
        db, err := sql.Open("mysql", "root:tester@tcp(127.0.0.1:3306)/mdcms")
	handle_err(err)
        defer db.Close()

        c := content{title, newcontent}

        query1 := "insert into content(title, date, content) values (\"" + c.Title + "\", \"" + c.Date() + "\", \"" + c.Content + "\" );"
        _, err = db.Exec( query1 )
	handle_err(err)
}

//Editing Content


func EditContent(title string, newcontent string) {
        db, err := sql.Open("mysql", "root:tester@tcp(127.0.0.1:3306)/mdcms")
	handle_err(err)
        defer db.Close()

        c := content{title, newcontent}

        query1 := "update content set date = \"" + c.Date() + "\", content = \"" + c.Content + "\" where title = \"" + c.Title +"\";"
        _, err = db.Exec(query1)
	handle_err(err)
}

//Deleting content
func DeleteContent(title string) {
        db, _ := sql.Open("mysql", "root:tester@tcp(127.0.0.1:3306)/mdcms")
        defer db.Close()

        c := content{Title: title}

        query1 := "delete from content where title = \"" + c.Title + "\";"
        _, _ = db.Exec(query1)
}

// Parsing markdown
func mdparse(markdown string) string {
        input := []byte(markdown)
        output := string(blackfriday.MarkdownCommon(input))
        return output
}


///////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main(){
}
