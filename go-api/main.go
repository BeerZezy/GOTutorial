package main
import (
   "fmt"
   "net/http"
   "encoding/json"
)

func main() {
	handleRequest()
}

func handleRequest() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/getAddress", getAddressBookAll)
	http.ListenAndServe(":8080", nil)
}

type addressBook struct {
    Firstname string
    Lastname  string
    Code      int
    Phone     string
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func getAddressBookAll(w http.ResponseWriter, r *http.Request) {
	addBook := addressBook{
		Firstname: "Phanudet",
		Lastname:  "Khawilai",
		Code:      1995,
		Phone:     "0623625698",
	}
	json.NewEncoder(w).Encode(addBook)
}