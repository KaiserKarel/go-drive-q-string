package qstring

import (
	"strings"
)

var r = strings.NewReplacer("'", `\'`)

func escape(s string) string {
	return r.Replace(s)
}
