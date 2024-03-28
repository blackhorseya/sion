package irent

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/blackhorseya/sion/entity/domain/rental/agg"
	"github.com/blackhorseya/sion/entity/domain/rental/repo"
	"github.com/blackhorseya/sion/pkg/configx"
	"github.com/blackhorseya/sion/pkg/contextx"
)

type impl struct {
	endpoint string
	version  string
}

// NewAssetRepo is a function to create a new asset repository.
func NewAssetRepo() (repo.IAssetRepo, error) {
	return &impl{
		endpoint: configx.C.IRent.HTTP.URL,
		version:  configx.C.IRent.Version,
	}, nil
}

func (i *impl) FetchAvailableCars(ctx contextx.Contextx) ([]*agg.Asset, error) {
	uri, err := url.ParseRequestURI(i.endpoint + "/api/AnyRent")
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(map[string]interface{}{
		"Radius":    0,
		"Latitude":  0,
		"Longitude": 0,
		"ShowAll":   1,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var got anyRentResponse
	err = json.NewDecoder(resp.Body).Decode(&got)
	if err != nil {
		return nil, err
	}

	if got.ErrorMessage != "Success" {
		return nil, errors.New(got.ErrorMessage)
	}

	return got.ToAggregate(), nil
}
