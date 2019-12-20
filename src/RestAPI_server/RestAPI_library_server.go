package main

import (
	"fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", libraryHome)
    http.HandleFunc("/status/", status)
    http.HandleFunc("/demanded/", demanded)
    http.HandleFunc("/popular/", popular)
    http.HandleFunc("/checkedOutUser/", checkedOutUser)
    err := http.ListenAndServe(":7777", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func libraryHome(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "This is the library Homepage.\n")
	fmt.Fprintf(writer, "To see the status of a book type /status/\n")
	fmt.Fprintf(writer, "To see the most issued book type /demanded/\n")
	fmt.Fprintf(writer, "To see the top trending book type /popular/\n")
	fmt.Fprintf(writer, "To see the user of a book type /checkedOutUser/")
}

func status(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "This book is currently issued.")
}

func demanded(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Harry Potter is the most issued book in our branch.")
}

func popular(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Percy Jackson is the top trending book.")
}

func checkedOutUser(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Psquare has currently checked out 'To Kill a Mockingbird'.")
}
