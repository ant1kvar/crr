package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"crr/internal/data"

	rb "github.com/randomtoy/radiobrowser-go"
)

// Radio-browser base URL (one of the public mirrors).
//
// NOTE: for this mirror, HTTPS can intermittently stall while streaming larger
// responses (e.g. limit=20), causing ctx deadline exceeded. Plain HTTP has been
// observed to respond consistently fast in the same environment.
var baseURL = "http://radio.telekost.ru"

func GetStations(ctx context.Context, country, tag string) ([]data.Station, error) {
	// Safety net: library respects ctx, but our UI often uses Background().
	// Keep a reasonable default deadline to avoid hanging forever.
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 90*time.Second)
		defer cancel()
	}

	country = strings.TrimSpace(country)
	tag = strings.TrimSpace(tag)
	if country == "" && tag == "" {
		return nil, fmt.Errorf("country and tag are empty")
	}

	rbStations, err := rb.StationSearch(ctx, baseURL, rb.StationSearchOptions{
		CountryCode: country,
		Tag:         strings.ToLower(tag),
		Limit:       20,
		Order:       rb.StationOrderClickCount,
		Reverse:     true,
		HideBroken:  true,
	})
	if err != nil {
		return nil, err
	}

	out := make([]data.Station, 0, len(rbStations))
	for _, s := range rbStations {
		link := strings.TrimSpace(s.URLResolved)
		if link == "" {
			link = strings.TrimSpace(s.URL)
		}
		if strings.TrimSpace(s.Name) == "" || link == "" {
			continue
		}
		out = append(out, data.Station{
			Name: s.Name,
			Link: link,
		})
	}
	return out, nil
}
