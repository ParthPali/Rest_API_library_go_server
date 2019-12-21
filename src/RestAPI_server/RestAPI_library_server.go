package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type book struct {
	BookID          string `json:"BookID"`
	BookTitle       string `json:"Title"`
	IsPopular       bool   `json:"IsPopular"`
	IsDemanded      bool   `json:"IsDemanded"`
	CheckedOut      bool   `json:"CheckedOut"`
	User 	        string `json:"User"`
}

var books []book

func main() {
    books = []book{
	book{BookID: "1",BookTitle: "Harry Potter1",IsPopular: true,IsDemanded: false,CheckedOut:false,User: ""},
	book{BookID: "2",BookTitle: "Harry Potter2",IsPopular: false,IsDemanded: true,CheckedOut:false,User: ""},
	book{BookID: "3",BookTitle: "Harry Potter3",IsPopular: false,IsDemanded: false,CheckedOut:true,User: "psquare"},
	book{BookID: "4",BookTitle: "Harry Potter4",IsPopular: false,IsDemanded: false,CheckedOut:true,User: "hsquare"},
    book{BookID: "5",BookTitle: "Harry Potter5",IsPopular: false,IsDemanded: false,CheckedOut:true,User: "psquare"},
    book{BookID: "6",BookTitle: "Harry Potter6",IsPopular: false,IsDemanded: false,CheckedOut:true,User: "nsquare"},
    }

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/book/{BookID}", returnbook)
    router.HandleFunc("/status-issued/", returnIssued)
    router.HandleFunc("/status-available/", returnAvailable)
    router.HandleFunc("/demanded/", returnDemandedBook)
    router.HandleFunc("/popular/", returnpopularbook)
    err := http.ListenAndServe(":7777", router)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func returnbook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["BookID"]

    for _, book := range books {
        if book.BookID == key {
            json.NewEncoder(w).Encode(book)
        }
    }
}

func returnpopularbook(w http.ResponseWriter, r *http.Request) {
    for _, book := range books {
        if book. IsPopular == true {
            json.NewEncoder(w).Encode(book)
        }
    }
}

func returnAvailable(w http.ResponseWriter, r *http.Request) {
    for _, book := range books {
        if book. CheckedOut == false {
            json.NewEncoder(w).Encode(book)
        }
    }
}

func returnIssued(w http.ResponseWriter, r *http.Request) {
    for _, book := range books {
        if book. CheckedOut == true {
            json.NewEncoder(w).Encode(book)
        }
    }
}

func returnDemandedBook(w http.ResponseWriter, r *http.Request) {
    for _, book := range books {
        if book. IsDemanded == true {
            json.NewEncoder(w).Encode(book)
        }
    }
}


