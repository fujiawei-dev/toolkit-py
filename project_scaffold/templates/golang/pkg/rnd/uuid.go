{{GOLANG_HEADER}}

package {{GOLANG_PACKAGE}}

import (
	"strings"

	"github.com/google/uuid"
)

// UUID returns a standard, random UUID as string.
func UUID() string {
	return uuid.New().String()
}

// IsUUID tests if the string looks like a standard UUID.
func IsUUID(s string) bool {
	return len(s) == 36 && IsHex(s)
}

// SanitizeUUID normalizes UUIDs found in XMP or Exif metadata.
func SanitizeUUID(s string) string {
	if s == "" {
		return ""
	}

	s = strings.Replace(strings.TrimSpace(s), "\"", "", -1)

	if start := strings.LastIndex(s, ":"); start != -1 {
		s = s[start+1:]
	}

	if !IsUUID(s) {
		return ""
	}

	return strings.ToLower(s)
}
