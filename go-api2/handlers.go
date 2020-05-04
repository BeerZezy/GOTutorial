package main
import (
    "encoding/json"
    "fmt"
    "net/http"
)
// Our first handler
func getIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello world!")
}
// Our first json response handler
func getJson(w http.ResponseWriter, r *http.Request) {
    // set JSON response header
    w.Header().Set("Content-type", "application/json; charset=UTF-8;");
    json.NewEncoder(w).Encode(BasicResponse{
    0,
    "Hello JSON!",
 })
}