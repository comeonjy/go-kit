package search

import (
	"bytes"
	"github.com/google/martian/log"
	"os"
	"regexp"
	"strconv"
	"unsafe"
)

// Str2Bytes convert string to byte slice. Warning: it's unsafe!!!
func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// ReverseStringSlice reverses StringSlice
func ReverseStringSlice(s []string) []string {
	// make a copy of s
	l := len(s)
	t := make([]string, l)
	for i := 0; i < l; i++ {
		t[i] = s[i]
	}

	// reverse
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	return t
}

// ReverseStringSliceInplace reverses StringSlice
func ReverseStringSliceInplace(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// EscapeSymbols escape custom symbols
func EscapeSymbols(s, symbols string) string {
	m := make(map[rune]struct{})
	for _, c := range symbols {
		m[c] = struct{}{}
	}
	var buf bytes.Buffer
	var ok bool
	for _, c := range s {
		if _, ok = m[c]; ok {
			buf.WriteByte('\\')
		}
		buf.WriteRune(c)
	}
	return buf.String()
}

// UnEscaper returns a function for unescaping string
func UnEscaper() func(s string) string {
	var re = regexp.MustCompile(`\\([abfnrtv'"?])`)
	var m = map[string]string{
		`\a`: "\a",
		`\b`: "\b",
		`\f`: "\f",
		`\n`: "\n",
		`\r`: "\r",
		`\t`: "\t",
		`\v`: "\v",
		`\\`: "\\",
		`\'`: "'",
		`\"`: "\"",
		`\?`: "?",
	}
	var mapping = func(key string) string {
		if v, ok := m[key]; ok {
			return v
		}
		return key
	}

	return func(s string) string {
		return re.ReplaceAllStringFunc(s, mapping)
	}
}

func IdentifyHashType(hash string) string {
	length := len(hash)
	typeStr := ""

	switch length {
	case 32: // MD5
		typeStr = "MD5"
	case 40: // SHA-1
		typeStr = "SHA-1"
	case 56: // SHA-224
		typeStr = "SHA-224"
	case 64: // SHA-256
		typeStr = "SHA-256"
	case 96: // SHA-384
		typeStr = "SHA-384"
	case 128: // SHA-512
		typeStr = "SHA-512"
	default:
		return ""
	}

	for i := 0; i < length; i++ {
		c := hash[i]
		if !(('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')) {
			return ""
		}
	}

	return typeStr
}

func GetEnvInteger[T int | int8 | int32 | int64](key string, defaultVal T) T {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Errorf("GetEnvInteger occur err: %s %s %d %v", key, value, defaultVal, err)
		return defaultVal
	}
	return T(i)
}

func ArrayEqual[T comparable](a1, a2 []T) bool {
	if len(a1) != len(a2) {
		return false
	}
	m := make(map[T]int)
	for _, v := range a1 {
		m[v]++
	}
	for _, v := range a2 {
		if _, ok := m[v]; ok {
			m[v]--
		}
		if m[v] == 0 {
			delete(m, v)
		}
	}
	return len(m) == 0
}
