package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

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

    wid := newWithIDHandler(req.URL.Path, "/people")
    if wid.idPresent {
        wid.handle(w)
        return
    }

    wph := newWithParamsHandler(req)
    if wph.paramsPresent {
        wph.handle(w)
        return
    }

    wpph := newWithPhoneNumberParamHandler(req)
    if wpph.paramPresent {
        wpph.handle(w)
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

func jsonError(err error) []byte {
    msg, err := json.Marshal(map[string]string{"error": err.Error()})
    if err != nil {
        panic(err)
    }
    return msg
}
