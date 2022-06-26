package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/lib/libtest"
    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

func TestPeopleNames(t *testing.T) {
    // TODO: test case using path /peopleTHISHOULD404
    type input struct {
        firstName string
        lastName string
    }
    type testType struct {
        input
        expected []*models.Person
    }
    expectedPeople := libtest.ExpectedPeople
    tests := []testType{
        testType{
            input: input{firstName: "John", lastName: "Doe"},
            expected: []*models.Person{expectedPeople[0], expectedPeople[3]},
        },

        testType{
            input: input{firstName: "Jenny", lastName: "Smith"},
            expected: []*models.Person{expectedPeople[4]},
        },

        testType{
            input: input{firstName: "NONE-EXISTENT", lastName: "USER"},
        },
    }

    for _, test := range tests {
        req := httptest.NewRequest(http.MethodGet, "/people/", nil)
        q := req.URL.Query()
        q.Add("first_name", test.input.firstName)
        q.Add("last_name", test.input.lastName)
        req.URL.RawQuery = q.Encode()

        path := req.URL.String()
        onlyGET(t, path)

        w := httptest.NewRecorder()
        People(w, req)
        res := w.Result()
        defer res.Body.Close() //nolint:gocritic,deferInLoop

        data := libtest.TestResponseMeta(
            t,
            libtest.TestHTTPResponseType{
                Path: path,
                Response: res,
                ExpectedStatusCode: 200,
                ExpectedHeader: []string{"application/json; charset=utf-8"},
            },
        )

        libtest.TestResponseData(t, path, data, libtest.BuildPeopleMap(test.expected))
    }
}
