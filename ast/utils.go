package ast

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

func appendSQLFragment(buff *bytes.Buffer, fragment string) {
	if fragment == "" {
		return
	}
	if buff.Len() > 0 {
		prev := buff.String()
		if !hasBoundarySpace(prev, fragment) {
			buff.WriteByte(' ')
		}
	}
	buff.WriteString(fragment)
}

func hasBoundarySpace(prev, next string) bool {
	return unicode.IsSpace(rune(prev[len(prev)-1])) ||
		unicode.IsSpace(rune(next[0]))
}

func replaceWhere (data string) string {
	lowerData := strings.ToLower(strings.TrimSpace(data))
	// find where keyword.
	hasWhereKeyword := false
	if strings.HasPrefix(lowerData,"where") {
		trimData := strings.TrimPrefix(lowerData, "where")
		if len(trimData) > 0 && unicode.IsSpace(rune(trimData[0])){
			hasWhereKeyword = true
		}
	}
	if hasWhereKeyword {
		return fmt.Sprintf("AND%s", strings.TrimSpace(data)[5:]/* string slice, skip 5 for `where` */)
	}
	return data
}
