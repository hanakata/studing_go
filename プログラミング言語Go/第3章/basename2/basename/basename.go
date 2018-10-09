//basenameはディレクトリ要素と接尾辞を取り除く
package basename

import "strings"

func Basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
