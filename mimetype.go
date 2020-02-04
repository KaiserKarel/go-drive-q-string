package qstring

type MIMEType string

const (
	Audio        MIMEType = "application/vnd.google-apps.audio"
	Document     MIMEType = "application/vnd.google-apps.document"
	Drawing      MIMEType = "application/vnd.google-apps.drawing"
	File         MIMEType = "application/vnd.google-apps.file"
	Folder       MIMEType = "application/vnd.google-apps.folder"
	Form         MIMEType = "application/vnd.google-apps.form"
	FusionTable  MIMEType = "application/vnd.google-apps.fusiontable"
	Map          MIMEType = "application/vnd.google-apps.map"
	Photo        MIMEType = "application/vnd.google-apps.photo"
	Presentation MIMEType = "application/vnd.google-apps.presentation"
	Script       MIMEType = "application/vnd.google-apps.script"
	Site         MIMEType = "application/vnd.google-apps.site"
	SpreadSheet  MIMEType = "application/vnd.google-apps.spreadsheet"
	Unknown      MIMEType = "application/vnd.google-apps.unknown"
	Video        MIMEType = "application/vnd.google-apps.video"
	DriveSDK     MIMEType = "application/vnd.google-apps.drive-sdk	"
)

// MimeTypeBuilder are used to construct queries for matching against mimeTypes.
// As opposed to regular enums, mimeTypes also support the contains operator.
//
// Prefer not using builders directly, instead using the constructors.
type MimeTypeBuilder struct {
	builder *EnumBuilder
}

func (b *MimeTypeBuilder) EQ(name MIMEType) *Builder {
	return b.builder.EQ(string(name))
}

func (b *MimeTypeBuilder) NE(name MIMEType) *Builder {
	return b.builder.NE(string(name))
}

func (b *MimeTypeBuilder) Contains(param MIMEType) *Builder {
	b.builder.builder.WriteString(" contains '")
	b.builder.builder.WriteString(string(param))
	b.builder.builder.WriteString("'")
	return &Builder{builder: b.builder.builder}
}
