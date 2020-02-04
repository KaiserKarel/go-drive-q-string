package qstring

import (
	"strings"
)

// FullTextBuilder are used for contain queries against ill-defined collections (fullText is defined as the union of the file.description, file.title and file.contents (if indexable).
//
// Prefer not using builders directly, instead using the constructors.
type FullTextBuilder struct {
	builder *strings.Builder
}

// Contains creates a query text contained within the indexable text. Single quotes
// are escaped.
//
// See also: https://developers.google.com/drive/api/v3/ref-search-terms#fn1
func (b *FullTextBuilder) Contains(param string) *Builder {
	b.builder.WriteString(" contains '")
	b.builder.WriteString(escape(param))
	b.builder.WriteString("'")
	return &Builder{builder: b.builder}
}
