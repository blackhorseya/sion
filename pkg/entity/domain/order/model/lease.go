package model

import (
	"encoding/json"

	"github.com/blackhorseya/irent/pkg/timex"
)

func (x *Lease) MarshalJSON() ([]byte, error) {
	type Alias Lease

	type output struct {
		*Alias
		StartAt    string `json:"start_at"`
		EndAt      string `json:"end_at"`
		LastPickAt string `json:"last_pick_at"`
	}

	object := &output{
		Alias:      (*Alias)(x),
		StartAt:    x.StartAt.AsTime().UTC().Format(timex.RFC3339Mill),
		EndAt:      x.EndAt.AsTime().UTC().Format(timex.RFC3339Mill),
		LastPickAt: x.LastPickAt.AsTime().UTC().Format(timex.RFC3339Mill),
	}

	return json.Marshal(object)
}
