package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

func People(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    data, err := json.Marshal(models.AllPeople())
    if err != nil {
        marshalErrorMsg, err := json.Marshal(map[string]string{"error": "marshaling of all people failed"})
        if err != nil {
            panic(err)
        }

        fmt.Fprintln(w, string(marshalErrorMsg))
        return
    }

    fmt.Fprintln(w, string(data))
}
