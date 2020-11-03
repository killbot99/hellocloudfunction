package hello

import (
	"context"
	"fmt"
	"time"
)

func hello(ctx context.Context, message Message) {
	fmt.Printf("Hello World! %s", time.Now())
}

func main() {
	var message Message
	var ctx context.Context
	hello(ctx, message)
}

type Message struct {
	Data []byte `json:"data"`
}
