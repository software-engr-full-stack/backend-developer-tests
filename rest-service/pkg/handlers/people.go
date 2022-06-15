package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

func People(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")

    // Notes: only GET method is allowed as per requirements.
    if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusNotFound)
        msg := jsonError(fmt.Errorf("%s", http.StatusText(http.StatusNotFound)))
        fmt.Fprintln(w, string(msg))
        return
    }

    // Notes: handle GET /people/:id
    wid := newWithIDHandler(req.URL.Path, "/people")
    if wid.idPresent {
        wid.handle(w)
        return
    }

    // Notes: handle GET /people?first_name=:first_name&last_name=:last_name
    wnph := newWithNamesParamsHandler(req)
    if wnph.paramsPresent {
        wnph.handle(w)
        return
    }

    // Notes: GET /people?phone_number=:phone_number
    wpph := newWithPhoneNumberParamHandler(req)
    if wpph.paramPresent {
        wpph.handle(w)
        return
    }

    // Notes: handle GET /people
    data, err := json.Marshal(models.AllPeople())
    if err != nil {
        msg := jsonError(fmt.Errorf("marshaling of all people failed"))
        fmt.Fprintln(w, string(msg))
        return
    }

    fmt.Fprintln(w, string(data))
}

func jsonError(err error) []byte {
    msg, err := json.Marshal(map[string]string{"error": err.Error()})
    if err != nil {
        panic(err)
    }
    return msg
}
