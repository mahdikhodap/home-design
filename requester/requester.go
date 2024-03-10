package requester

import (
	"fmt"
	"net/http"
)

type Requester struct {
	Client *http.Client
}

type RequestEntity struct {
	Endpoint string
}

func NewRequester() *Requester {
	return &Requester{
		Client: &http.Client{},
	}
}

// TODO add necessary config like time out and etc...
func (r *Requester) Load() *Requester {
	return r
}

func (r *Requester) Get(entity RequestEntity) (*http.Response, error) {
	res, err := r.Client.Get(entity.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request to %s: %v", entity.Endpoint, err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for %s", res.StatusCode, entity.Endpoint)
	}

	return res, nil
}
