package common

type middleware struct{}

func New() *middleware {
	return &middleware{}
}
