package contexts

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// fetch long 3rd API with context timeout
// set value with context

type contextKey string

const messageKey contextKey = "message"

func FetchWithContext(ctx context.Context, msg string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	ctx = context.WithValue(ctx, messageKey, fmt.Sprintf("succeed-%s", msg))
	defer cancel()

	ch := make(chan string)

	fetch3rdParty := func(ctx context.Context, ch chan<- string) {
		time.Sleep(time.Millisecond * 300)
		val := ctx.Value(messageKey)
		message, _ := val.(string)
		ch <- message
	}

	go fetch3rdParty(ctx, ch)

	select {
	case <-ctx.Done():
		return "", errors.New("got timeout when fetch 3rd party")
	case resp := <-ch:
		return resp, nil
	}
}
