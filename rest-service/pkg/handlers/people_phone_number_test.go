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

func TestPeoplePhoneNumber(t *testing.T) {
    type input struct {
        phoneNumber string
    }
    type testType struct {
        input
        expected []*models.Person
    }
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

        data := runCommonTests(t, path, res, 200)

        tb := titleBuilder{path: path}

        var actualPeople []*models.Person
        title := tb.build("unmarshal response body error")
        if actual, expected := json.Unmarshal(data, &actualPeople), error(nil); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }

        title = tb.build("people count")
        if actual, expected := len(actualPeople), len(test.expected); actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }

        for _, actualPerson := range actualPeople {
            key := buildKey(actualPerson.LastName, actualPerson.FirstName, actualPerson.PhoneNumber)
            title = tb.build(fmt.Sprintf("presence of person %#v", key))
            expectedPerson, ok := expectedPeopleMap[key]
            if actual, expected := ok, true; actual != expected {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
            }
            title = tb.build("person details")
            if isDeepEqual := reflect.DeepEqual(actualPerson, expectedPerson); !isDeepEqual {
                t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actualPerson, expectedPerson)
            }
        }
    }
}
