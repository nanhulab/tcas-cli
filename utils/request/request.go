/*
 * @Author: jffan
 * @Date: 2024-07-29 14:56:25
 * @LastEditTime: 2024-07-29 15:23:49
 * @LastEditors: jffan
 * @FilePath: \tcas-cli\utils\request\request.go
 * @Description: Copyright © 2024 <jffan@nanhulab.ac.cn>
 */
package request

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Request is a struct to hold request parameters
type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Query   map[string]string
	Body    interface{}
}

// NewRequest creates a new Request instance
func NewRequest(method, url string) *Request {
	return &Request{
		Method:  method,
		URL:     url,
		Headers: make(map[string]string),
		Query:   make(map[string]string),
	}
}

// SetHeader sets a header for the request
func (r *Request) SetHeader(key, value string) *Request {
	r.Headers[key] = value
	return r
}

// SetQuery sets a query parameter for the request
func (r *Request) SetQuery(key, value string) *Request {
	r.Query[key] = value
	return r
}

// SetBody sets the body for the request
func (r *Request) SetBody(body interface{}) *Request {
	r.Body = body
	return r
}

// Do sends the request and returns the response
func (r *Request) Do() (*http.Response, error) {
	// Create the URL with query parameters
	u, err := url.Parse(r.URL)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	for k, v := range r.Query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	// Create the request body
	var body io.Reader
	if r.Body != nil {
		switch b := r.Body.(type) {
		case string:
			body = strings.NewReader(b)
		case []byte:
			body = bytes.NewReader(b)
		default:
			jsonBody, err := json.Marshal(r.Body)
			if err != nil {
				return nil, err
			}
			body = bytes.NewReader(jsonBody)
		}
	}

	// Create the HTTP request
	req, err := http.NewRequest(r.Method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DoWithBody sends the request and returns the response with body as an interface{}
func (r *Request) DoWithBody(v interface{}) (*http.Response, error) {
	resp, err := r.Do()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return nil, err
	}

	return resp, nil
}

// Get is a shortcut for creating a GET request
func Get(url string) *Request {
	return NewRequest(http.MethodGet, url)
}

// Post is a shortcut for creating a POST request
func Post(url string) *Request {
	return NewRequest(http.MethodPost, url)
}

// Put is a shortcut for creating a PUT request
func Put(url string) *Request {
	return NewRequest(http.MethodPut, url)
}

// Delete is a shortcut for creating a DELETE request
func Delete(url string) *Request {
	return NewRequest(http.MethodDelete, url)
}
