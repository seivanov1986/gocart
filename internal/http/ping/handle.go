package ping

import (
	"net/http"
)

type Ping interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type ping struct {
}

func New() *ping {
	return &ping{}
}
