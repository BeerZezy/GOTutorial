package main
import (
    "log"
	"net/http"
)

// init main func
func main() {
    router := routes.NewRouter()
    log.Fatal(
        // start on port 3000 by default
        http.ListenAndServe(":3000", router),
    )
}