// for gbk is not supported in sift.
package main

import (
	"bufio"
	"io"
	"regexp"
	"unicode"

	"github.com/axgle/mahonia"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// just for gbk-system environment.
func HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

// for only utf8 is supported in golang's regexp.
func CheckToUtf8(data string) string {
	switch options.CodingInput {
	case "gbk":
		dec := mahonia.NewDecoder("GBK")
		data = dec.ConvertString(data)
	}
	switch options.Coding {
	case "gbk":
		enc := mahonia.NewEncoder("UTF-8")
		data = enc.ConvertString(data)
	}
	return data
}

func CheckToGBK(data string) string {
	switch options.CodingInput {
	case "gbk":
		dec := mahonia.NewDecoder("UTF-8")
		data = dec.ConvertString(data)
		enc := mahonia.NewEncoder("GBK")
		data = enc.ConvertString(data)
	}
	return data
}

// transform io.reader to utf8 reader.
func Utf8ReaderAny(reader io.Reader) io.Reader {
	bytes, err := bufio.NewReader(reader).Peek(1024)
	if err == nil || err.Error() == "EOF" {
		e, _, _ := charset.DetermineEncoding(bytes, "")
		if e == charmap.Windows1252 {
			e = simplifiedchinese.GBK
		}
		reader = transform.NewReader(reader, e.NewDecoder())
	}
	return reader
}

// transform gbk reader to utf8 reader.
func Utf8ReaderGbk(reader io.Reader) io.Reader {
	reader = transform.NewReader(reader, simplifiedchinese.GBK.NewDecoder())
	return reader
}

func Utf8Reader(reader io.Reader) io.Reader {
	switch options.Coding {
	case "gbk":
		return Utf8ReaderGbk(reader)
	}
	return reader
}
