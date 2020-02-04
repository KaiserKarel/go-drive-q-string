package qstring

import (
	"strings"
	"time"
)

const format = "2006-01-02T15:04:05"

// DateTimeBuilders are used for querying by datetimes.
//
// Prefer not using builders directly, instead using the constructors.
type DateTimeBuilder struct {
	builder *strings.Builder
}

func (b *DateTimeBuilder) EQ(t time.Time) *Builder {
	b.builder.WriteString(" = '")
	b.builder.WriteString(t.Format(format))
	b.builder.WriteString("'")
	return &Builder{builder: b.builder}
}

func (b *DateTimeBuilder) NE(t time.Time) *Builder {
	b.builder.WriteString(" != '")
	b.builder.WriteString(t.Format(format))
	b.builder.WriteString("'")
	return &Builder{builder: b.builder}
}

func (b *DateTimeBuilder) GT(t time.Time) *Builder {
	b.builder.WriteString(" > '")
	b.builder.WriteString(t.Format(format))
	b.builder.WriteString("'")
	return &Builder{builder: b.builder}
}

func (b *DateTimeBuilder) GTE(t time.Time) *Builder {
	b.builder.WriteString(" >= '")
	b.builder.WriteString(t.Format(format))
	b.builder.WriteString("'")
	return &Builder{builder: b.builder}
}

func (b *DateTimeBuilder) LT(t time.Time) *Builder {
	b.builder.WriteString(" < '")
	b.builder.WriteString(t.Format(format))
	b.builder.WriteString("'")
	return &Builder{builder: b.builder}
}

func (b *DateTimeBuilder) LTE(t time.Time) *Builder {
	b.builder.WriteString(" <= '")
	b.builder.WriteString(t.Format(format))
	b.builder.WriteString("'")
	return &Builder{builder: b.builder}
}
