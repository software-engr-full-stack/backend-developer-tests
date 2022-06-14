package handlers

import (
    "testing"
    "io"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "reflect"
    "fmt"

    "github.com/satori/go.uuid"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

var expectedPeople = []*models.Person{
    {
        ID:          uuid.Must(uuid.FromString("81eb745b-3aae-400b-959f-748fcafafd81")),
        FirstName:   "John",
        LastName:    "Doe",
        PhoneNumber: "+1 (800) 555-1212",
    },
    {
        ID:          uuid.Must(uuid.FromString("5b81b629-9026-450d-8e46-da4f8c7bd513")),
        FirstName:   "Jane",
        LastName:    "Doe",
        PhoneNumber: "+1 (800) 555-1313",
    },
    {
        ID:          uuid.Must(uuid.FromString("df12ce76-767b-4bf0-bccb-816745df9e70")),
        FirstName:   "Brian",
        LastName:    "Smith",
        PhoneNumber: "+44 7700 900077",
    },
    // This is another person with the name John Doe
    {
        ID:          uuid.Must(uuid.FromString("135af595-aa86-4bb5-a8f7-df17e6148e63")),
        FirstName:   "John",
        LastName:    "Doe",
        PhoneNumber: "+1 (800) 555-1414",
    },
    // This is another person with the phone number +44 7700 900077
    {
        ID:          uuid.Must(uuid.FromString("000ebe58-b659-422b-ab48-a0d0d40bd8f9")),
        FirstName:   "Jenny",
        LastName:    "Smith",
        PhoneNumber: "+44 7700 900077",
    },
}

var expectedPeopleMap = func() map[string]*models.Person{
    table := map[string]*models.Person{}

    for _, person := range expectedPeople {
        key := buildKey(person.LastName, person.FirstName, person.PhoneNumber)
        table[key] = person
    }

    return table
}()

func TestPeople(t *testing.T) {
    path := "/people"
    req := httptest.NewRequest(http.MethodGet, path, nil)
    w := httptest.NewRecorder()
    People(w, req)
    res := w.Result()
    defer res.Body.Close()

    buildTitle := func(title string) string {
        return fmt.Sprintf("%s, %s", path, title)
    }

    title := buildTitle("response code")
    if actual, expected := res.StatusCode, 200; actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    data, err := io.ReadAll(res.Body)
    title = buildTitle("read response body error")
    if actual, expected := err, error(nil); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    var actualPeople []*models.Person
    title = buildTitle("unmarshal response body error")
    if actual, expected := json.Unmarshal(data, &actualPeople), error(nil); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    title = buildTitle("people count")
    if actual, expected := len(actualPeople), len(expectedPeople); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    for _, actualPerson := range actualPeople {
        key := buildKey(actualPerson.LastName, actualPerson.FirstName, actualPerson.PhoneNumber)
        title = fmt.Sprintf("presence of person %#v", key)
        expectedPerson, ok := expectedPeopleMap[key]
        if actual, expected := ok, true; actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
        isDeepEqual := reflect.DeepEqual(actualPerson, expectedPerson)
        if actual, expected := isDeepEqual, true; actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
    }
}

func buildKey(ln, fn, pn string) string {
    return ln + " " + fn + " " + pn
}
