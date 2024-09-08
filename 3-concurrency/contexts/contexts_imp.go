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

	fetch3rdParty := func(ctx context.Context, ch chan<- string, timeInMs int64) {
		time.Sleep(time.Millisecond * time.Duration(timeInMs))
		val := ctx.Value(messageKey)
		message, _ := val.(string)
		ch <- message
	}

	go fetch3rdParty(ctx, ch, 300)

	select {
	case <-ctx.Done():
		return "", errors.New("got timeout when fetch 3rd party")
	case resp := <-ch:
		return resp, nil
	}
}

func callFirstApi(ctx context.Context) {
	time.Sleep(time.Millisecond * 200)
	fmt.Println("from first api")
}
func callSecondApi(ctx context.Context) {
	time.Sleep(time.Millisecond * 250)
	fmt.Println("from second api")
}
func callThirdApi(ctx context.Context) {
	time.Sleep(time.Millisecond * 400)
	fmt.Println("from third api")
}

func FetchSomeParties(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*300)
	defer cancel()

	go callFirstApi(ctx)
	go callSecondApi(ctx)
	go callThirdApi(ctx)

	fmt.Println("I am from called func")
	time.Sleep(300 * time.Millisecond)
}
