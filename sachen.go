package sachen

import (
	"unicode"
	"unicode/utf8"
  )


func FirstStringUpper(val string) bool {
  ch, _ := utf8.DecodeRuneInString(val)
  return unicode.IsUpper(ch)
}
