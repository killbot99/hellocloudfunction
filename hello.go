package hello

import (
	"net/http"
	"fmt"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World! %s", time.Now())
}


