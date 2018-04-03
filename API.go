package main

import (
    "fmt"
    "log"
    "net/http"
    //"encoding/json"
    "github.com/gorilla/mux"
    "os"
    "io/ioutil"
)
type Test struct{
    Name string
}
func main() {
    
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", GetBeaconInfo)
    log.Fatal(http.ListenAndServe(":8080", router))
   
}


func GetBeaconInfo(w http.ResponseWriter, r *http.Request) {    
    
    file, err := os.Create("BeaconInfo.txt")
    if err != nil {
	   log.Fatal(err)
    }

    body := ioutil.readAll(r.Body)
    

    //json.NewDecoder(r.Body).Decode(&t)
    //vars := mux.Vars(r)
    //fmt.Fprintf(w, "OK")
    
    fmt.Println(string(body))
    defer r.Body.Close()
    defer file.Close()
}