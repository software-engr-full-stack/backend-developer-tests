package libtest

import (
    "testing"
    "io"
    "net/http"
    "reflect"
    "fmt"
    "encoding/json"

    "github.com/satori/go.uuid"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

var ExpectedPeople = []*models.Person{
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

var ExpectedPeopleMap = BuildPeopleMap(ExpectedPeople)

func BuildPeopleMap(people []*models.Person) map[string]*models.Person{
    table := map[string]*models.Person{}

    for _, person := range people {
        key := buildUniquePersonKey(person.LastName, person.FirstName, person.PhoneNumber)
        table[key] = person
    }

    return table
}

func buildUniquePersonKey(ln, fn, pn string) string {
    return ln + " " + fn + " " + pn
}

type PathTitleBuilder struct {
    Path string
}

func (ptb *PathTitleBuilder) Build(title string) string {
    return fmt.Sprintf("%s, %s", ptb.Path, title)
}

type TestHTTPResponseType struct {
    Path string
    Response *http.Response
    ExpectedStatusCode int
    ExpectedHeader []string
    IsJSON bool
}

func TestResponseMeta(t *testing.T, test TestHTTPResponseType) []byte {
    ptb := PathTitleBuilder{Path: test.Path}

    title := ptb.Build("response code")
    if actual, expected := test.Response.StatusCode, test.ExpectedStatusCode; actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    title = ptb.Build("content type")
    actualHeader := test.Response.Header["Content-Type"]
    if isDeepEqual := reflect.DeepEqual(actualHeader, test.ExpectedHeader); !isDeepEqual {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actualHeader, test.ExpectedHeader)
    }

    data, err := io.ReadAll(test.Response.Body)
    if !test.IsJSON {
        return data
    }

    title = ptb.Build("read response body error")
    if actual, expected := err, error(nil); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }
    return data
}

func TestResponseData(t *testing.T, path string, data []byte, expectedPeopleMap map[string]*models.Person) {
    var actualPeople []*models.Person
    ptb := PathTitleBuilder{Path: path}
    title := ptb.Build("unmarshal response body error")
    if actual, expected := json.Unmarshal(data, &actualPeople), error(nil); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    title = ptb.Build("people count")
    if actual, expected := len(actualPeople), len(expectedPeopleMap); actual != expected {
        t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
    }

    for _, actualPerson := range actualPeople {
        key := buildUniquePersonKey(actualPerson.LastName, actualPerson.FirstName, actualPerson.PhoneNumber)
        title = ptb.Build(fmt.Sprintf("presence of person %#v", key))
        expectedPerson, ok := expectedPeopleMap[key]
        if actual, expected := ok, true; actual != expected {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actual, expected)
        }
        title = ptb.Build("person details")
        if isDeepEqual := reflect.DeepEqual(actualPerson, expectedPerson); !isDeepEqual {
            t.Fatalf("%s: actual not equal to expected, %#v != %#v", title, actualPerson, expectedPerson)
        }
    }
}
