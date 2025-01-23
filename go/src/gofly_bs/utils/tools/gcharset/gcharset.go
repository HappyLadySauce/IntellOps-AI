package gcharset

import (
	"bytes"
	"context"
	"io"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"

	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"
	"gofly/utils/tools/intlog"
)

var (
	// Alias for charsets.
	charsetAlias = map[string]string{
		"HZGB2312": "HZ-GB-2312",
		"hzgb2312": "HZ-GB-2312",
		"GB2312":   "HZ-GB-2312",
		"gb2312":   "HZ-GB-2312",
	}
)

// Convert converts `src` charset encoding from `srcCharset` to `dstCharset`,
// and returns the converted string.
// It returns `src` as `dst` if it fails converting.
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	if dstCharset == srcCharset {
		return src, nil
	}
	dst = src
	// Converting `src` to UTF-8.
	if srcCharset != "UTF-8" {
		if e := getEncoding(srcCharset); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewDecoder()),
			)
			if err != nil {
				return "", gerror.Wrapf(err, `convert string "%s" to utf8 failed`, srcCharset)
			}
			src = string(tmp)
		} else {
			return dst, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported srcCharset "%s"`, srcCharset)
		}
	}
	// Do the converting from UTF-8 to `dstCharset`.
	if dstCharset != "UTF-8" {
		if e := getEncoding(dstCharset); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader([]byte(src)), e.NewEncoder()),
			)
			if err != nil {
				return "", gerror.Wrapf(err, `convert string from utf8 to "%s" failed`, dstCharset)
			}
			dst = string(tmp)
		} else {
			return dst, gerror.NewCodef(gcode.CodeInvalidParameter, `unsupported dstCharset "%s"`, dstCharset)
		}
	} else {
		dst = src
	}
	return dst, nil
}

// getEncoding returns the encoding.Encoding interface object for `charset`.
// It returns nil if `charset` is not supported.
func getEncoding(charset string) encoding.Encoding {
	if c, ok := charsetAlias[charset]; ok {
		charset = c
	}
	enc, err := ianaindex.MIB.Encoding(charset)
	if err != nil {
		intlog.Errorf(context.TODO(), `%+v`, err)
	}
	return enc
}
