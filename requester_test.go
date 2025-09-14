package wigal

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestRequester(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	response, err := requester[interface{}](ctx, "https://jsonplaceholder.typicode.com/todos/1", http.MethodGet, nil,
		nil,
		nil)
	if err != nil {
		t.Errorf("expected to get a respose with status Ok but got err %v", err)
	}

	fmt.Printf("response: %v\n", response)
}
