package main

import (
    "testing"
    "net/http"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/lib/libtest"
)

func TestMain(t *testing.T) {
    cases := []struct{
        Path string
        ExpectedStatusCode int
        ExpectedHeader []string
        IsJSON bool
        ExpectedText string
    }{
        { "/people", 200, []string{"application/json; charset=utf-8"}, true, "" },
        { "/people/", 200, []string{"application/json; charset=utf-8"}, true, "" },
        { "/peopleTHISHOULD404", 404, []string{"text/plain; charset=utf-8"}, false, "404 page not found\n" },
    }

    for _, cs := range cases {
        path := cs.Path
        ptb := libtest.PathTitleBuilder{Path: path}
        res, err := http.Get("http://localhost:8000" + path) //nolint:golangcilint,noctx
        title := ptb.Build("http.Get error return value")
        if actual, expected := err, error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
        defer res.Body.Close() //nolint:gocritic,deferInLoop

        data := libtest.TestResponseMeta(
            t,
            libtest.TestHTTPResponseType{
                Path: path,
                Response: res,
                ExpectedStatusCode: cs.ExpectedStatusCode,
                ExpectedHeader: cs.ExpectedHeader,
            },
        )

        if !cs.IsJSON {
            title = ptb.Build("non-JSON string response data")
            if actual, expected := string(data), cs.ExpectedText; actual != expected {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
            }
            continue
        }

        libtest.TestResponseData(t, path, data, libtest.ExpectedPeopleMap)
    }
}
