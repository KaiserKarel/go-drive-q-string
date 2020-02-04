# qstring

[![GoDoc](https://godoc.org/godoc.org/github.com/KaiserKarel/qstring?status.svg)](https://godoc.org/github.com/KaiserKarel/qstring)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/kaiserkarel/qstring)](https://goreportcard.com/report/github.com/kaiserkarel/qstring)

Typed query string builder for Google Drive's file API.

Not all query possibilities are currently supported, although it is trivial to extend. Drop a PR and I will merge it.

## Example 

```go 
func main() {
    ...
    // modifiedTime > '2012-06-04T12:00:00' and (mimeType contains 'image/' or mimeType contains 'video/')
    query := qstring.ModifiedTime().
                GT(time.Unix(1338811200, 0).UTC()).
            And().Sub(
                MimeType().
                Contains("image/").
                Or().MimeType().
                Contains("video/")).String()    

}
```

## Should I use this?

Depends. Drive's query string grammar is very simple, doing some simple string formatting solves 
most usecases. Query builders offer more type safety, but add a layer of indirection.
