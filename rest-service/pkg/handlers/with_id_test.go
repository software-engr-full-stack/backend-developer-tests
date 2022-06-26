package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "reflect"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/lib/libtest"
    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

func TestPeopleID(t *testing.T) {
    type expectedType struct {
        person *models.Person
        statusCode int
        error
    }
    type testType struct {
        path string
        expected expectedType
    }

    expectedPerson := libtest.ExpectedPeople[2]
    nonExistentID := "d135b79c-ef02-4b1f-81c7-d8c25d423c55"
    tests := []testType{
        testType{
            path: "/people" + "/" + expectedPerson.ID.String(),
            expected: expectedType{person: expectedPerson, statusCode: 200},
        },

        testType{
            path: "/people" + "/" + nonExistentID,
            expected: expectedType{statusCode: 404, error: fmt.Errorf("user ID %s not found", nonExistentID)},
        },
    }

    for _, test := range tests {
        onlyGET(t, test.path)

        req := httptest.NewRequest(http.MethodGet, test.path, nil)
        w := httptest.NewRecorder()
        People(w, req)
        res := w.Result()
        defer res.Body.Close() //nolint:gocritic,deferInLoop

        data := libtest.TestResponseMeta(
            t,
            libtest.TestHTTPResponseType{
                Path: test.path,
                Response: res,
                ExpectedStatusCode: test.expected.statusCode,
                ExpectedHeader: []string{"application/json; charset=utf-8"},
            },
        )

        ptb := libtest.PathTitleBuilder{Path: test.path}

        if test.expected.error != nil {
            var eresp map[string]string
            title := ptb.Build("unmarshal response body error")
            if actual, expected := json.Unmarshal(data, &eresp), error(nil); actual != expected {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
            }
            title = ptb.Build("error details")
            if actual, expected := eresp["error"], test.expected.error.Error(); actual != expected {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
            }
            continue
        }

        var actualPerson *models.Person
        title := ptb.Build("unmarshal response body error")
        if actual, expected := json.Unmarshal(data, &actualPerson), error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }

        title = ptb.Build("person details")
        if isDeepEqual := reflect.DeepEqual(actualPerson, expectedPerson); !isDeepEqual {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actualPerson, expectedPerson)
        }
    }
}
