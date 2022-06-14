package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

type withNamesParamsHandler struct {
    paramsPresent bool
    firstName string
    lastName string
}

func newWithNamesParamsHandler(req *http.Request) withNamesParamsHandler {
    q := req.URL.Query()
    noParams := withNamesParamsHandler{paramsPresent: false}
    if len(q) == 0 {
        return noParams
    }

    qfn := q["first_name"]
    qln := q["last_name"]
    if qfn == nil || qln == nil {
        return noParams
    }

    return withNamesParamsHandler{
        paramsPresent: true,
        firstName: qfn[0],
        lastName: qln[0],
    }
}

func (wnph *withNamesParamsHandler) handle(w http.ResponseWriter) {
    people := models.FindPeopleByName(wnph.firstName, wnph.lastName)

    data, err := json.Marshal(people)
    if err != nil {
        msg := jsonError(fmt.Errorf("marshaling of people failed"))
        fmt.Fprintln(w, string(msg))
        return
    }

    fmt.Fprintln(w, string(data))
}
