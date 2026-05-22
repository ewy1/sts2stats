//go:build api

package api

import "testing"

func ptr[T any](obj T) *T {
	return &obj
}

func TestFilter(t *testing.T) {
	f := RunFilter{
		Win:       ptr(true),
		Character: ptr("guy"),
		Ascension: ptr(3),
		Version:   ptr("0.333.2"),
	}
	Query(f)
}
