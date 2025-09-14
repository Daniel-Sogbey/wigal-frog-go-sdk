package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRequester(t *testing.T) {
	response, err := requester[interface{}]("https://jsonplaceholder.typicode.com/todos/1", http.MethodGet, nil, nil,
		nil)
	if err != nil {
		t.Errorf("expected to get a respose with status Ok but got err %v", err)
	}

	fmt.Printf("response: %v\n", response)
}
