//go:build api

package api

import (
	"sts2stats/stats"
	"sts2stats/storage"
)

func CardChoices() (any, error) {
	return storage.Entities[*stats.CardChoice]("SELECT * FROM CardChoice")
}
