package qstring

import (
	"strings"
)

// StringBuilders are used for matching against arbitrary string variables.
//
// Prefer not using builders directly, instead using the constructors.
type StringBuilder struct {
	builder *strings.Builder
}

func (b *StringBuilder) Equals(name string) *Builder {
	b.builder.WriteString(" = ")
	b.builder.WriteString("'")
	b.builder.WriteString(escape(name))
	b.builder.WriteString("'")
	return &Builder{builder:b.builder}
}

func (b *StringBuilder) NotEquals(name string) *Builder {
	b.builder.WriteString(" != ")
	b.builder.WriteString("'")
	b.builder.WriteString(escape(name))
	b.builder.WriteString("'")
	return &Builder{builder:b.builder}
}

func (b *StringBuilder) Contains(param string) *Builder {
	b.builder.WriteString(" contains '")
	b.builder.WriteString(escape(param))
	b.builder.WriteString("'")
	return &Builder{builder:b.builder}
}