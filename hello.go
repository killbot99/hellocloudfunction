package hello

import (
	"http"
	"fmt"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World! %s", time.Now())
}

func main() {
	var w http.ResponseWriter
	var r http.Request
	Hello(r, *w)
}

