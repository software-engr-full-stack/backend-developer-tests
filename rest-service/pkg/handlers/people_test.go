package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/lib/libtest"
)

func TestPeople(t *testing.T) {
    path := "/people"

    onlyGET(t, path)

    req := httptest.NewRequest(http.MethodGet, path, nil)
    w := httptest.NewRecorder()
    People(w, req)
    res := w.Result()
    defer res.Body.Close()

    data := libtest.TestResponseMeta(
        t,
        libtest.TestHTTPResponseType{
            Path: path,
            Response: res,
            ExpectedStatusCode: 200,
            ExpectedHeader: []string{"application/json; charset=utf-8"},
        },
    )

    libtest.TestResponseData(t, path, data, libtest.ExpectedPeopleMap)

    path = "/peopleTHISHOULD404"
    onlyGET(t, path)

    req = httptest.NewRequest(http.MethodGet, path, nil)
    w = httptest.NewRecorder()
    People(w, req)
    res = w.Result()
    defer res.Body.Close()

    ptb := libtest.PathTitleBuilder{Path: path}
    title := ptb.Build("response code")
    if actual, expected := res.StatusCode, 404; actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }
}
