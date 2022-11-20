package httpx

import (
	"net/http"

	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/google/wire"
)

type impl struct {
}

func NewImpl() httpx.Client {
	return &impl{}
}

func (i *impl) Do(req *http.Request) (resp *http.Response, err error) {
	client := http.DefaultClient

	return client.Do(req)
}

// ProviderSet is a provider set for httpx
var ProviderSet = wire.NewSet(NewImpl)
