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

func TestPeople(t *testing.T) {
    path := "/people"

    onlyGET(t, path)

    req := httptest.NewRequest(http.MethodGet, path, nil)
    w := httptest.NewRecorder()
    People(w, req)
    res := w.Result()
    defer res.Body.Close()

    data := runCommonTests(t, path, res, 200)

    tb := titleBuilder{path: path}

    var actualPeople []*models.Person
    title := tb.build("unmarshal response body error")
    if actual, expected := json.Unmarshal(data, &actualPeople), error(nil); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    title = tb.build("people count")
    if actual, expected := len(actualPeople), len(expectedPeople); actual != expected {
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

    path = "/peopleTHISHOULD404"
    onlyGET(t, path)

    req = httptest.NewRequest(http.MethodGet, path, nil)
    w = httptest.NewRecorder()
    People(w, req)
    res = w.Result()
    defer res.Body.Close()

    tb = titleBuilder{path: path}
    title = tb.build("response code")
    if actual, expected := res.StatusCode, 404; actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }
}
