package qstring

import (
	"strconv"
	"strings"
)

// BoolBuilder is used for queries which may only be matched against true or false.
// See also EnumBuilder.
//
// Prefer not using builders directly, instead using the constructors.
type BoolBuilder struct {
	builder *strings.Builder
}

func (b *BoolBuilder) EQ(ok bool) {
	b.builder.WriteString(" = ")
	b.builder.WriteString(strconv.FormatBool(ok))
}

func (b *BoolBuilder) NE(ok bool) {
	b.builder.WriteString(" != ")
	b.builder.WriteString(strconv.FormatBool(ok))
}
