package qstring

import (
	"testing"
	"time"
)

func TestExamples(t *testing.T) {
	var tests = []struct {
		want string
		got  string
	}{
		{
			"name = 'hello'",
			Name().Equals("hello").String(),
		},
		{
			"name contains 'hello' and name contains 'goodbye'",
			Name().Contains("hello").And().Name().Contains("goodbye").String(),
		},
		{
			"not name contains 'hello'",
			Not().Name().Contains("hello").String(),
		},
		{
			"mimeType = 'application/vnd.google-apps.folder'",
			MimeType().EQ("application/vnd.google-apps.folder").String(),
		},
		{
			"mimeType != 'application/vnd.google-apps.folder'",
			MimeType().NE("application/vnd.google-apps.folder").String(),
		},
		{
			"fullText contains 'important' and trashed = true",
			FullText().Contains("important").And().Trashed(true).String(),
		},
		{
			"fullText contains 'hello'",
			FullText().Contains("hello").String(),
		},
		{
			"not fullText contains 'hello'",
			Not().FullText().Contains("hello").String(),
		},
		{
			`fullText contains 'hello world'`,
			FullText().Contains("hello world").String(),
		},
		{
			"'appDataFolder' in parents",
			Q("'appDataFolder'").In().Parents().String(),
		},
		{
			"'test@example.org' in writers",
			Q("'test@example.org'").In().Writers().String(),
		},
		{
			"modifiedTime > '2012-06-04T12:00:00'",
			ModifiedTime().GT(time.Unix(1338811200, 0).UTC()).String(),
		},
		{
			"sharedWithMe and name contains 'hello'",
			SharedWithMe().And().Name().Contains("hello").String(),
		},
		{
			"visibility = 'limited'",
			Visibility().EQ(Limited).String(),
		},
		{
			"modifiedTime > '2012-06-04T12:00:00' and (mimeType contains 'image/' or mimeType contains 'video/')",
			ModifiedTime().GT(time.Unix(1338811200, 0).UTC()).And().Sub(MimeType().Contains("image/").Or().MimeType().Contains("video/")).String(),
		},
		{
				ViewedByMe().EQ(time.Unix(1338811200, 0).UTC()).String(),
				"viewedByMe = '2012-06-04T12:00:00'",
		},
	}

	for _, test := range tests {
		if test.got != test.want {
			t.Errorf("got %s, want %s", test.got, test.want)
		}
	}
}
