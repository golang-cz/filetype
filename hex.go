package main

import "strings"

func key(first []byte, second ...[]byte) string {
	var b strings.Builder
	b.WriteString(EncodeToString(first))

	// for _, str := range second {
	// 	b.WriteString(string(str))
	// }

	return b.String()
}

const hextable = "0123456789ABCDEF"

func EncodeToString(src []byte) string {
	dst := make([]byte, len(src)*3)
	encodeToHex(dst, src)
	return string(dst)
}

func encodeToHex(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		dst[j+2] = ' '
		j += 3
	}
	return len(src) * 3
}
