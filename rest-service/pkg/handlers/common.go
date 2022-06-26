package handlers

import (
    "testing"
    "io"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/lib/libtest"
)

func onlyGET(t *testing.T, path string) {
    req := httptest.NewRequest(http.MethodPost, path, nil)
    w := httptest.NewRecorder()
    People(w, req)
    res := w.Result()
    defer res.Body.Close()

    ptb := libtest.PathTitleBuilder{Path: fmt.Sprintf("%s %s", http.MethodPost, path)}
    title := ptb.Build("response code")
    if actual, expected := res.StatusCode, 404; actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    data, err := io.ReadAll(res.Body)
    title = ptb.Build("read response body error")
    if actual, expected := err, error(nil); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    var eresp map[string]string
    title = ptb.Build("unmarshal response body error")
    if actual, expected := json.Unmarshal(data, &eresp), error(nil); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    title = ptb.Build("error details")
    if actual, expected := eresp["error"], http.StatusText(http.StatusNotFound); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }
}
