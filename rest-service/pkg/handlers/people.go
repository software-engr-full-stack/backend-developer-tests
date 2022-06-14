package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

func People(w http.ResponseWriter, req *http.Request) {
    data, err := json.Marshal(models.AllPeople())
    if err != nil {
        panic(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    fmt.Fprintln(w, string(data))
}
