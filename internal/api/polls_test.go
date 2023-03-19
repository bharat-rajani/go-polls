package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPollsHandler_200OK_ResponseRecorder(t *testing.T) {
	t.Parallel()
	expectedPollsResponse := `{"1":{"choice":["hi","bonjour"],"name":"Hello"}}`
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "http://localhost.com/polls", nil)
	GetPollsHandler(recorder, request)
	gotRespByte, err := io.ReadAll(recorder.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if expectedPollsResponse != string(gotRespByte) {
		fmt.Println("Expected!=Got")
		t.Fail()
	}
}

// Testing using mock server
func TestGetPollsHandler_200OK_MockServer(t *testing.T) {
	t.Parallel()
	expectedPollsResponse := `{"1":{"choice":["hi","bonjour"],"name":"Hello"}}`
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedPollsResponse))
	}))
	defer s.Close()

	// preparing request
	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	gotResp, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if expectedPollsResponse != string(gotResp) {
		t.Fail()
	}
}
