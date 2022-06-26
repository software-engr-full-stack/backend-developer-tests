package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/lib/libtest"
    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

func TestPeoplePhoneNumber(t *testing.T) {
    // TODO: test case using path /peopleTHISHOULD404
    type input struct {
        phoneNumber string
    }
    type testType struct {
        input
        expected []*models.Person
    }
    expectedPeople := libtest.ExpectedPeople
    tests := []testType{
        testType{
            input: input{phoneNumber: "+44 7700 900077"},
            expected: []*models.Person{expectedPeople[2], expectedPeople[4]},
        },

        testType{
            input: input{phoneNumber: "+1 (800) 555-1313"},
            expected: []*models.Person{expectedPeople[1]},
        },

        testType{
            input: input{phoneNumber: "123 456 7890"},
        },
    }

    for _, test := range tests {
        req := httptest.NewRequest(http.MethodGet, "/people/", nil)
        q := req.URL.Query()
        q.Add("phone_number", test.input.phoneNumber)
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
