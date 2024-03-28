package agg

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/blackhorseya/sion/entity/domain/rental/model"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

//go:embed asset.tmpl
var f embed.FS

// Asset is an aggregate root.
type Asset struct {
	*model.Car

	Distance      float64 `json:"distance"`
	GoogleMapsURL string  `json:"google_maps_url"`
}

// FlexContainer returns the asset as a flex message.
func (x *Asset) FlexContainer() (messaging_api.FlexContainerInterface, error) {
	x.GoogleMapsURL = x.Location.GetGoogleMaps()

	tmpl, err := template.New("asset.tmpl").ParseFS(f, "asset.tmpl")
	if err != nil {
		return nil, err
	}

	var layout bytes.Buffer
	err = tmpl.Execute(&layout, x)
	if err != nil {
		return nil, err
	}

	return messaging_api.UnmarshalFlexContainer(layout.Bytes())
}

// Assets is a collection of Asset.
type Assets []*Asset

// FlexContainer returns the assets as a flex message.
func (x Assets) FlexContainer() (messaging_api.FlexContainerInterface, error) {
	var containers []messaging_api.FlexBubble
	for _, asset := range x {
		container, err := asset.FlexContainer()
		if err != nil {
			return nil, err
		}

		containers = append(containers, container.(messaging_api.FlexBubble))
	}

	return &messaging_api.FlexCarousel{
		Contents: containers,
	}, nil
}
