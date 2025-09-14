package wigal

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func requester[T any](ctx context.Context, urlString, method string, queryParams url.Values, headers http.Header,
	body any) (*T, error) {
	var t T

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return &t, err
	}

	reqUrl, err := url.Parse(urlString)
	if err != nil {
		return &t, err
	}

	q := reqUrl.Query()

	for k, v := range queryParams {
		q.Set(k, strings.Join(v, ","))
	}

	reqUrl.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, method, reqUrl.String(), bytes.NewReader(jsonBytes))
	if err != nil {
		return &t, err
	}

	for k, v := range headers {
		req.Header[k] = v
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return &t, err
	}

	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return &t, err
	}

	err = json.Unmarshal(respBytes, &t)
	if err != nil {
		return &t, err
	}

	return &t, nil
}
