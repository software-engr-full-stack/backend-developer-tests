package handlers

import (
    "net/http"
    "encoding/json"
    "strings"
    "fmt"

    "github.com/satori/go.uuid"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

func People(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusNotFound)
        msg := jsonError(fmt.Errorf("%s", http.StatusText(http.StatusNotFound)))
        fmt.Fprintln(w, string(msg))
        return
    }

    id := extractID(req.URL.Path, "/people")
    if id != "" {
        person, err := models.FindPersonByID(uuid.Must(uuid.FromString(id)))
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
        return
    }

    data, err := json.Marshal(models.AllPeople())
    if err != nil {
        msg := jsonError(fmt.Errorf("marshaling of all people failed"))
        fmt.Fprintln(w, string(msg))
        return
    }

    fmt.Fprintln(w, string(data))
}

func extractID(path, prefix string) string {
    id := strings.TrimPrefix(path, prefix)
    id = strings.TrimLeft(id, "/")
    return id
}

func jsonError(err error) []byte {
    msg, err := json.Marshal(map[string]string{"error": err.Error()})
    if err != nil {
        panic(err)
    }
    return msg
}
