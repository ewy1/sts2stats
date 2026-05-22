//go:build api

package api

import (
	"fmt"
	"reflect"
	"strings"
	"sts2stats/spool"
)

type Filter interface {
	WhereQuery() string
}

type RunFilter struct {
	Win       *bool
	Character *string
	Ascension *int
	Version   *string
}

func Query(f any) string {
	var parts []string
	for field, value := range reflect.ValueOf(f).Fields() {
		if !value.IsNil() {
			elem := value.Elem()

			switch elem.Kind() {
			case reflect.String:
				parts = append(parts, fmt.Sprintf("%v = '%v'", field.Name, elem.String()))
			case reflect.Int:
				parts = append(parts, fmt.Sprintf("%v = %v", field.Name, elem.Int()))
			case reflect.Bool:
				parts = append(parts, fmt.Sprintf("%v = %v", field.Name, elem.Bool()))
			default:
				spool.Warn("unsupported filter field kind: %v", elem.Kind())
			}
		}
	}
	if len(parts) == 0 {
		return ""
	}
	return fmt.Sprintf("WHERE %s", strings.Join(parts, " AND "))
}
