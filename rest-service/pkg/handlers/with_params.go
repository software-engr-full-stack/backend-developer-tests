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
    if len(q) == 0 {
        return withParamsHandler{paramsPresent: false}
    }

    return withParamsHandler{
        paramsPresent: true,
        firstName: q["first_name"][0],
        lastName: q["last_name"][0],
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
