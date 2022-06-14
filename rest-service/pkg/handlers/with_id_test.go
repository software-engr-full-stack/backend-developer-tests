package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "reflect"
    "fmt"

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

    expectedPerson := expectedPeople[2]
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

        data := runCommonTests(t, test.path, res, test.expected.statusCode)

        tb := titleBuilder{path: test.path}

        if test.expected.error != nil {
            var eresp map[string]string
            title := tb.build("unmarshal response body error")
            if actual, expected := json.Unmarshal(data, &eresp), error(nil); actual != expected {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
            }
            title = tb.build("error details")
            if actual, expected := eresp["error"], test.expected.error.Error(); actual != expected {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
            }
            continue
        }

        var actualPerson *models.Person
        title := tb.build("unmarshal response body error")
        if actual, expected := json.Unmarshal(data, &actualPerson), error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }

        title = tb.build("person details")
        if isDeepEqual := reflect.DeepEqual(actualPerson, expectedPerson); !isDeepEqual {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actualPerson, expectedPerson)
        }
    }
}
