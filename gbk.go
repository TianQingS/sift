// for gbk is not supported in sift.
package main

import (
	"bufio"
	"io"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// for gbk is not supported in golang's regexp.
func ToUtf8(data string) string {
	enc := mahonia.NewEncoder("UTF-8")
	data = enc.ConvertString(data)
	return data
}

// transform io.reader to utf8 reader.
func Utf8ReaderAny(reader io.Reader) io.Reader {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err == nil {
		e, _, _ := charset.DetermineEncoding(bytes, "")
		reader = transform.NewReader(reader, e.NewDecoder())
	}
	return reader
}

// transform gbk reader to utf8 reader.
func Utf8Reader(reader io.Reader) io.Reader {
	reader = transform.NewReader(reader, simplifiedchinese.GB18030.NewDecoder())
	return reader
}
