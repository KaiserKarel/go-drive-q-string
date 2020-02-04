package qstring

import (
	"strings"
)

// CollectionBuilders are used for queries which match against predefined
// collections.
//
// Prefer not using builders directly, instead using the constructors.
type CollectionBuilder struct {
	builder *strings.Builder
}

// Parents finalized the expression with the parents collection.
func (b *CollectionBuilder) Parents() *Builder {
	b.builder.WriteString(" parents")
	return &Builder{builder:b.builder}
}

// Writers finalized the expression with the writers collection.
func (b *CollectionBuilder) Writers() *Builder {
	b.builder.WriteString(" writers")
	return &Builder{builder:b.builder}
}