package qstring

import (
	"strings"
)

// Enum builders are used to construct queries which may only be matched against
// a set of values. Enums differ from booleans since enums need to be
// encapsulated in single quotes, and the enum string needs to be escaped.
//
// Prefer not using builders directly, instead using the constructors.
type EnumBuilder struct {
	builder *strings.Builder
}

func (b *EnumBuilder) EQ(variant string) *Builder {
	b.builder.WriteString(" = ")
	b.builder.WriteString("'")
	b.builder.WriteString(escape(variant))
	b.builder.WriteString("'")
	return &Builder{builder:b.builder}
}

func (b *EnumBuilder) NE(variant string) *Builder {
	b.builder.WriteString(" != ")
	b.builder.WriteString("'")
	b.builder.WriteString(escape(variant))
	b.builder.WriteString("'")
	return &Builder{builder:b.builder}
}