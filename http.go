package main

import (
	"net/http"
	"io"
	"time"
	tea "charm.land/bubbletea/v2"
)
const url = "https://dog.ceo/api/breeds/image/random"

func sendQuery(client *http.Client, target string) tea.Cmd {
return func() tea.Msg {
	 req, err := http.NewRequest("GET", url, nil)
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
		start := time.Now()
resp, err := client.Do(req)
		elapsed := time.Since(start).Milliseconds()
	if err != nil {
		logToFile("error.log", err.Error())
		return errMsg{err: err}
	}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		logToFile("response.log", string(body))
	return respMsg{
		resp: resp,
			body: string(body),
			ms: elapsed,
	}
}
}
type QueryMethod int
const (
MethodGet QueryMethod = iota	
MethodHead
MethodPost
MethodPut
MethodPatch
MethodDelete
MethodConnect
MethodOptions
MethodTrace
)
var methodName = map[QueryMethod]string{
	MethodGet:     "GET",
	MethodHead:    "HEAD",
	MethodPost:    "POST",
	MethodPut:     "PUT",
	MethodPatch:   "PATCH",
	MethodDelete:  "DELETE",
	MethodConnect: "CONNECT",
	MethodOptions: "OPTIONS",
	MethodTrace:   "TRACE",
}
func (qm QueryMethod) String() string{
	if name, ok := methodName[qm]; ok {
		return name
	}
	return "UNKNOWN"
}

type respMsg struct {
	resp *http.Response
	body string
	ms int64
}
type errMsg struct {
	err error
}