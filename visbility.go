package qstring

type VisibilityLevel string

const (
	AnyoneCanFind VisibilityLevel = "anyoneCanFind"
	AnyoneWithLink VisibilityLevel = "anyoneWithLink"
	DomainCanFind VisibilityLevel = "domainCanFind"
	DomainWithLink VisibilityLevel = "domainWithLink"
	Limited VisibilityLevel = "limited"
)

// VisibilityBuilders are used to match the visibility enum field.
//
// Prefer not using builders directly, instead using the constructors.
type VisibilityBuilder struct {
	builder *EnumBuilder
}

func (b *VisibilityBuilder) EQ(level VisibilityLevel) *Builder {
	return b.builder.EQ(string(level))
}

func (b *VisibilityBuilder) NE(level VisibilityLevel) *Builder {
	return b.builder.NE(string(level))
}