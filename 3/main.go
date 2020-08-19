package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    lib "./lib"
)

type Palindrome struct {
    N1 int `json:"n1"`
    N2 int `json:"n2"`
    SUM int `json:"sum"`
    PAL []int `json:"pal"`
}

func Pal(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var Pal Palindrome
    if r.Method == "GET" {
        Pal = Palindrome{
            N1: 99,
            N2: 100,
            SUM: len(lib.SearchPal(99,100)),
            PAL: lib.SearchPal(99,100),
        }
        data, _ := json.Marshal(Pal)
        w.Write(data)
        return
    }else if r.Method == "POST" {
        getN1 := r.PostFormValue("n1")
        n1, _ := strconv.Atoi(getN1)
        getN2 := r.PostFormValue("n2")
        n2, _ := strconv.Atoi(getN2)
        Pal = Palindrome{
            N1: n1,
            N2: n2,
            SUM: len(lib.SearchPal(n1,n2)),
            PAL: lib.SearchPal(n1,n2),
        }
 
        data, _ := json.Marshal(Pal)
        w.Write(data)
        return
    }
 
    http.Error(w, "error", http.StatusNotFound)
    return
}

func handleRequests() {
    http.HandleFunc("/", Pal)
    fmt.Println("server running...")
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    handleRequests()
}