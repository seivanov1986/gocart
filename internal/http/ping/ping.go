package ping

import (
	"net/http"

	"bitbucket.org/ivanovse/gocart/external/helpers"
)

func (u *ping) Ping(w http.ResponseWriter, r *http.Request) {
	helpers.HttpResponse(w, nil, http.StatusOK)
}
