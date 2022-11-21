package adapter

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type requester struct {
	headerRequestParams map[string]string
	requestBody         io.Reader
	responseTarget      any
	basicAuth           *basicAuth
}

type basicAuth struct {
	username string
	password string
}

type requesterOption func(r *requester)

func withRequestHeaderParams(headerRequestParams map[string]string) requesterOption {
	return func(r *requester) {
		r.headerRequestParams = headerRequestParams
	}
}

func withRequestBody(body io.Reader) requesterOption {
	return func(r *requester) {
		r.requestBody = body
	}
}

func withResponseTarget(responseTarget any) requesterOption {
	return func(r *requester) {
		r.responseTarget = responseTarget
	}
}

func withBasicAuth(username, password string) requesterOption {
	return func(r *requester) {
		r.basicAuth = &basicAuth{
			username: username,
			password: password,
		}
	}
}

// doRequestWithDecode отправляет запрос с требуемыми настройками.
func doRequestWithDecode(
	ctx context.Context,
	method, url string,
	client *http.Client,
	opts ...requesterOption,
) error {
	r := &requester{}
	for i := range opts {
		opts[i](r)
	}

	request, err := http.NewRequestWithContext(ctx, method, url, r.requestBody)
	if err != nil {
		return errors.Wrap(err, "creating request")
	}

	if r.basicAuth != nil {
		request.SetBasicAuth(r.basicAuth.username, r.basicAuth.password)
	}

	for key, value := range r.headerRequestParams {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "doing request")
	}

	if response.StatusCode != http.StatusOK {
		b, bodyReadError := io.ReadAll(response.Body)
		if bodyReadError != nil {
			return errors.Wrapf(bodyReadError, "reading response body on error, status_code: %d", response.StatusCode)
		}
		err = errors.New("request unsuccessful")
		return errors.Wrapf(err, "status_code: %d, message: %s", response.StatusCode, string(b))
	}

	if r.responseTarget != nil {
		if err = json.NewDecoder(response.Body).Decode(r.responseTarget); err != nil {
			return errors.Wrap(err, "decoding to response target")
		}
	}

	if err = response.Body.Close(); err != nil {
		return errors.Wrap(err, "closing response body")
	}
	return nil
}
