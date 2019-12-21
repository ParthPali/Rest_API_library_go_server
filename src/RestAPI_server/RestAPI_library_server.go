package main

import (
    "fmt"
    "log"
    "net/http"
    //"encoding/json"
    "github.com/gorilla/mux"
    "sort"
    "strconv"
)

type Library struct{
    BookID          string `json:"BookID"`
    BookTitle       string `json:"Title"`
    User            string `json:"User"`
    IsReturned      bool   `json:"IsReturned"`
}

var LibraryHistory []Library

func main() {
    LibraryHistory = []Library{
        Library{BookID: "1",BookTitle: "Harry Potter1",User: "John",IsReturned: true},
        Library{BookID: "4",BookTitle: "Harry Potter4",User: "Jack",IsReturned: true},
        Library{BookID: "2",BookTitle: "Harry Potter2",User: "John",IsReturned: true},
        Library{BookID: "7",BookTitle: "Harry Potter7",User: "Alex" ,IsReturned: true},
        Library{BookID: "1",BookTitle: "Harry Potter1",User: "Noah",IsReturned: true},
        Library{BookID: "6",BookTitle: "Harry Potter6",User: "Luke",IsReturned: true},
        Library{BookID: "1",BookTitle: "Harry Potter1",User: "Alex",IsReturned: true},
        Library{BookID: "5",BookTitle: "Harry Potter5",User: "Jack",IsReturned: false},
        Library{BookID: "3",BookTitle: "Harry Potter3",User: "John",IsReturned: false},
        Library{BookID: "7",BookTitle: "Harry Potter7",User: "Luke",IsReturned: true},
        Library{BookID: "4",BookTitle: "Harry Potter4",User: "Jack",IsReturned: true},
        Library{BookID: "7",BookTitle: "Harry Potter7",User: "Jack",IsReturned: false},
        Library{BookID: "1",BookTitle: "Harry Potter1",User: "Ryan",IsReturned: false},
    }
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/book/{BookID}", returnbook)
    router.HandleFunc("/mostIssued/", returnMostIssued)
    router.HandleFunc("/availableBooks/", returnAvailable)
    router.HandleFunc("/issuedBooks",returnIssued)
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
    var isCheckedOut bool = false
    var bookname string

    for _, Library := range LibraryHistory {
        if Library.BookID == key && Library.IsReturned== false{
            fmt.Fprintf(w, "The book "+Library.BookTitle+" is currently issued to "+Library.User)
            isCheckedOut = true
        }
        if Library.BookID == key {
            bookname = Library.BookTitle
        }
    }

    if isCheckedOut == false && bookname != ""{
        fmt.Fprintf(w, "The book "+bookname+" is currently available")
    }
    if bookname == ""{
        fmt.Fprintf(w, "Book not found")
    }
}

func returnIssued(w http.ResponseWriter, r *http.Request) {
    //vars := mux.Vars(r)
    //key := vars["BookID"]
    //var isCheckedOut bool = false
    //var bookname string
    
    for _, Library := range LibraryHistory {
        if Library.IsReturned== false{
            fmt.Fprintf(w, "The book "+Library.BookTitle+" is currently issued to "+Library.User + "\n")
            //isCheckedOut = true
        }
    }
}

func checkavailable(b string) bool{
    var isavailable bool = false
    var isCheckedOut bool = false
    var bookname string
    for _, Library := range LibraryHistory {
        if Library.BookID == b && Library.IsReturned== false{
            //fmt.Fprintf(w, "The book "+Library.BookTitle+" is currently issued to "+Library.User)
            isCheckedOut = true
            isavailable = false
            return isavailable
        }
        if Library.BookID == b {
            bookname = Library.BookTitle
        }
    }
    if isCheckedOut == false && bookname != ""{
        //fmt.Fprintf(w, "The book "+bookname+" is currently available")
        isavailable = true
    }
    
    return isavailable
}

func returnAvailable(w http.ResponseWriter, r *http.Request) {
    //vars := mux.Vars(r)
    
    for i := 0; i < 7; i++ {
        
        if checkavailable(strconv.Itoa(i)){
            fmt.Fprintf(w, "The bookID# %d is currently available. \n",i)
        }
    }
}

func returnpopularbook(w http.ResponseWriter, r *http.Request) {
    var i [7]int
    var name [7]string 
    for _, Library := range LibraryHistory {
        if Library.BookID == "1" { 
            i[0]++
            name[0] = Library.BookTitle
        }
        if Library.BookID == "2" {
            i[1]++
            name[1] = Library.BookTitle
        }
        if Library.BookID == "3" {
            i[2]++
            name[2] = Library.BookTitle
        }
        if Library.BookID == "4" { 
            i[3]++
            name[3] = Library.BookTitle
        }
        if Library.BookID == "5" { 
            i[4]++
            name[4] = Library.BookTitle
        }
        if Library.BookID == "6" { 
            i[5]++
            name[5] = Library.BookTitle
        }
        if Library.BookID == "7" { 
            i[6]++
            name[6] = Library.BookTitle
        }
    }
    popularity := []struct {
        bookname string
        count int
    }{
        {name[0], i[0]},
        {name[1], i[1]},
        {name[2], i[2]},
        {name[3], i[3]},
        {name[4], i[4]},
        {name[5], i[5]},
        {name[6], i[6]},
    }
    sort.SliceStable(popularity, func(j, k int) bool{
        return popularity[j].count > popularity[k].count
    })
    fmt.Fprintf(w,"Top Trending books are as follows\n")
    x := 0
    for _, allbooks := range popularity {
        fmt.Fprintf(w," %v has been issued %d times \n",allbooks.bookname,allbooks.count)
        x++
    }
}

func returnMostIssued(w http.ResponseWriter, r *http.Request) {
    var i [7]int
    var name [7]string 
    for _, Library := range LibraryHistory {
        if Library.BookID == "1" { 
            i[0]++
            name[0] = Library.BookTitle
        }
        if Library.BookID == "2" {
            i[1]++
            name[1] = Library.BookTitle
        }
        if Library.BookID == "3" {
            i[2]++
            name[2] = Library.BookTitle
        }
        if Library.BookID == "4" { 
            i[3]++
            name[3] = Library.BookTitle
        }
        if Library.BookID == "5" { 
            i[4]++
            name[4] = Library.BookTitle
        }
        if Library.BookID == "6" { 
            i[5]++
            name[5] = Library.BookTitle
        }
        if Library.BookID == "7" { 
            i[6]++
            name[6] = Library.BookTitle
        }
    }
    popularity := []struct {
        bookname string
        count int
    }{
        {name[0], i[0]},
        {name[1], i[1]},
        {name[2], i[2]},
        {name[3], i[3]},
        {name[4], i[4]},
        {name[5], i[5]},
        {name[6], i[6]},
    }
    sort.SliceStable(popularity, func(j, k int) bool{
        return popularity[j].count > popularity[k].count
    })
    fmt.Fprintf(w,"Most issued book is ")
    fmt.Fprintf(w," %v which has been issued %d times.",popularity[0].bookname,popularity[0].count)
}

