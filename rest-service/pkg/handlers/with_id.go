package handlers

import (
    "net/http"
    "encoding/json"
    "strings"
    "fmt"

    "github.com/satori/go.uuid"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

type withIDHandler struct {
    idPresent bool
    id string
}

func newWithIDHandler(path, prefix string) (withIDHandler) {
    id := strings.TrimPrefix(path, prefix)

    // Notes: because strings.TrimPrefix will return the the path value if a
    // prefix match isn't found.
    if id == "" || id == path {
        return withIDHandler{idPresent: false}
    }

    return withIDHandler{idPresent: true, id: id}
}

func (widh *withIDHandler) handle(w http.ResponseWriter) {
    person, err := models.FindPersonByID(uuid.Must(uuid.FromString(widh.id)))
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        msg := jsonError(err)
        fmt.Fprintln(w, string(msg))
        return
    }

    data, err := json.Marshal(person)
    if err != nil {
        msg := jsonError(fmt.Errorf("marshaling of person failed"))
        fmt.Fprintln(w, string(msg))
        return
    }

    fmt.Fprintln(w, string(data))
}
