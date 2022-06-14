package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

type withParamsHandler struct {
    paramsPresent bool
    firstName string
    lastName string
}

func newWithParamsHandler(req *http.Request) withParamsHandler {
    q := req.URL.Query()
    noParams := withParamsHandler{paramsPresent: false}
    if len(q) == 0 {
        return noParams
    }

    qfn := q["first_name"]
    qln := q["last_name"]
    if qfn == nil || qln == nil {
        return noParams
    }

    return withParamsHandler{
        paramsPresent: true,
        firstName: qfn[0],
        lastName: qln[0],
    }
}

func (wph *withParamsHandler) handle(w http.ResponseWriter) {
    people := models.FindPeopleByName(wph.firstName, wph.lastName)

    data, err := json.Marshal(people)
    if err != nil {
        msg := jsonError(fmt.Errorf("marshaling of people failed"))
        fmt.Fprintln(w, string(msg))
        return
    }

    fmt.Fprintln(w, string(data))
}
