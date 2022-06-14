package handlers

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/software-engr-full-stack/backend-developer-tests/rest-service/pkg/models"
)

type withPhoneNumberParam struct {
    paramPresent bool
    phoneNumber string
}

func newWithPhoneNumberParamHandler(req *http.Request) withPhoneNumberParam {
    q := req.URL.Query()
    noParams := withPhoneNumberParam{paramPresent: false}
    if len(q) == 0 {
        return noParams
    }

    qpn := q["phone_number"]
    if qpn == nil {
        return noParams
    }

    return withPhoneNumberParam{
        paramPresent: true,
        phoneNumber: qpn[0],
    }
}

func (wpnp *withPhoneNumberParam) handle(w http.ResponseWriter) {
    people := models.FindPeopleByPhoneNumber(wpnp.phoneNumber)

    data, err := json.Marshal(people)
    if err != nil {
        msg := jsonError(fmt.Errorf("marshaling of people failed"))
        fmt.Fprintln(w, string(msg))
        return
    }

    fmt.Fprintln(w, string(data))
}
