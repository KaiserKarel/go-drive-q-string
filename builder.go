package qstring

import (
	"fmt"
	"strconv"
	"strings"
)

// Builder continues writing to previous statements, meaning new builders should
// be initialized with an initial value. Prefer using the named constructors.
type Builder struct {
	builder *strings.Builder
}

// String calls the underlying strings.Builder String method.
func (b *Builder) String() string {
	return b.builder.String()
}

// Q returns a Builder preloaded with obj.
func Q(obj string) *Builder {
	b := &Builder{&strings.Builder{}}
	b.builder.WriteString(obj)
	return b
}

// ModifiedTime is used to query by the date of the last modification of the file.
func ModifiedTime() *DateTimeBuilder {
	b := DateTimeBuilder(*Q("modifiedTime"))
	return &b
}

// ViewedByMeTime is used to query by the last time the file was viewed by the user.
func ViewedByMeTime() *DateTimeBuilder {
	b := DateTimeBuilder(*Q("viewedByMeTime"))
	return &b
}

// CreatedTime is used to query by the time at which the file was created.
func CreatedTime() *DateTimeBuilder {
	b := DateTimeBuilder(*Q("createdTime"))
	return &b
}

// ModifiedByMeTime is used to query by the last time the file was modified by anyone.
func ModifiedByMeTime() *DateTimeBuilder {
	b := DateTimeBuilder(*Q("modifiedByMeTime"))
	return &b
}

// SharedWithMeTime is used to query by the time at which the file was shared with the user, if applicable.
func SharedWithMeTime() *DateTimeBuilder {
	b := DateTimeBuilder(*Q("sharedWithMeTime"))
	return &b
}

// Visibility is used to query by visibility level.
func Visibility() *VisibilityBuilder {
	builder := EnumBuilder(*Q("visibility"))
	return &VisibilityBuilder{&builder}
}

// MimeType is used to query by MIME type.
func MimeType() *MimeTypeBuilder {
	builder := EnumBuilder(*Q("mimeType"))
	return &MimeTypeBuilder{&builder}
}

// SharedWithMe is used to query for files that are in the user's "Shared with me" collection.
func SharedWithMe() *Builder {
	return Q("sharedWithMe")
}

// Name is used to query by the name of the file. This is not necessarily unique
// within a folder. Note that for immutable items such as the top level folders of
// shared drives, My Drive root folder, and Application Data folder the name is
// constant.
func Name() *StringBuilder {
	b := StringBuilder(*Q("name"))
	return &b
}

// FullText is used to query full text of the file including name,
// description, content, and indexable text. Whether the 'title', 'description',
// or 'indexableText' properties, or the content of the file matches.
func FullText() *FullTextBuilder {
	b := FullTextBuilder(*Q("fullText"))
	return &b
}

// ViewedByMe is used for queries by whether the file has been viewed by this
// user.
func ViewedByMe() *DateTimeBuilder {
	b := DateTimeBuilder(*Q("viewedByMe"))
	return &b
}

// Properties are key value pairs, which can be set by all apps.
func Properties() *MapBuilder {
	b := MapBuilder(*Q("properties"))
	return &b
}

// AppProperties are key value private to the requesting app.
func AppProperties() *MapBuilder {
	b := MapBuilder(*Q("properties"))
	return &b
}

// Trashed is used for queries where the file is trashed.
func Trashed(ok bool) *Builder {
	b := Q("trashed = ")
	b.builder.WriteString(strconv.FormatBool(ok))
	return b
}

// Starred is used for queries where the file is starred.
func Starred(ok bool) *Builder {
	b := Q("starred = ")
	if ok {
		b.builder.WriteString("true")
	} else {
		b.builder.WriteString("false")

	}
	return b
}

// Not negates the following expression.
func Not() *Builder {
	b := &Builder{&strings.Builder{}}
	b.builder.WriteString("not")
	return b
}

// Sub takes a nested subquery, which will be enclosed by parenthesis.
func (b *Builder) Sub(query fmt.Stringer) *Builder {
	b.builder.WriteString(" (")
	b.builder.WriteString(query.String())
	b.builder.WriteString(")")
	return b
}

// See Sub
func (b *Builder) SubS(query string) *Builder {
	b.builder.WriteString(" (")
	b.builder.WriteString(query)
	b.builder.WriteString(")")
	return b
}

// And is the logical conjunction operator of two expressions.
func (b *Builder) And() *Builder {
	b.builder.WriteString(" and")
	return b
}

// Visibility is used to query by visibility level.
func (b *Builder) Visibility() *VisibilityBuilder {
	b.builder.WriteString(" visibility")
	builder := EnumBuilder(*b)
	return &VisibilityBuilder{builder: &builder}
}

// SharedWithMe is used to query for files that are in the user's "Shared with me" collection.
func (b *Builder) SharedWithMe() *BoolBuilder {
	b.builder.WriteString(" sharedWithMe")
	return &BoolBuilder{builder: b.builder}
}

// In is used for query by inclusion within one of the predefined collections.
func (b *Builder) In() *CollectionBuilder {
	b.builder.WriteString(" in")
	return &CollectionBuilder{builder: b.builder}
}

// Or is the logical disjunction operator
func (b *Builder) Or() *Builder {
	b.builder.WriteString(" or")
	return b
}

func (b *Builder) Name() *StringBuilder {
	b.builder.WriteString(" name")
	a := StringBuilder(*b)
	return &a
}

func (b *Builder) ModifiedTime() *DateTimeBuilder {
	b.builder.WriteString(" modifiedTime")
	a := DateTimeBuilder(*b)
	return &a
}

func (b *Builder) ViewedByMeTime() *DateTimeBuilder {
	b.builder.WriteString(" viewedByMeTime")
	a := DateTimeBuilder(*b)
	return &a
}

func (b *Builder) CreatedTime() *DateTimeBuilder {
	b.builder.WriteString(" createdTime")
	a := DateTimeBuilder(*b)
	return &a
}

func (b *Builder) ModifiedByMeTime() *DateTimeBuilder {
	b.builder.WriteString(" modifiedByMeTime")
	a := DateTimeBuilder(*b)
	return &a
}

func (b *Builder) SharedWithMeTime() *DateTimeBuilder {
	b.builder.WriteString(" sharedWithMeTime")
	a := DateTimeBuilder(*b)
	return &a
}

func (b *Builder) FullText() *FullTextBuilder {
	b.builder.WriteString(" fullText")
	a := FullTextBuilder(*b)
	return &a
}

func (b *Builder) MimeType() *MimeTypeBuilder {
	b.builder.WriteString(" mimeType")
	a := EnumBuilder(*b)
	return &MimeTypeBuilder{builder: &a}
}

func (b *Builder) Properties() *MapBuilder {
	b.builder.WriteString("properties")
	return &MapBuilder{builder: b.builder}
}

func (b *Builder) AppProperties() *MapBuilder {
	b.builder.WriteString("appProperties")
	return &MapBuilder{builder: b.builder}
}

func (b *Builder) Trashed(ok bool) *Builder {
	b.builder.WriteString(" trashed = ")
	if ok {
		b.builder.WriteString("true")
	} else {
		b.builder.WriteString("false")

	}
	return b
}

func (b *Builder) Starred(ok bool) *Builder {
	b.builder.WriteString(" starred = ")
	if ok {
		b.builder.WriteString("true")
	} else {
		b.builder.WriteString("false")

	}
	return b
}
