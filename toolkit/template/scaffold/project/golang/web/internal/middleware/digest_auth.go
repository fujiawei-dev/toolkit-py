package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"net/textproto"
	"strings"
	"unsafe"
)

// https://datatracker.ietf.org/doc/html/rfc2617

type Digest struct {
	Username  string
	Realm     string
	Nonce     string
	Uri       string
	Algorithm string
	Cnonce    string
	Nc        string
	Qop       string
	Response  string
}

func (d *Digest) Validate(password, method, realm, uri string, body []byte) bool {
	if d == nil || d.Realm != realm || !strings.HasPrefix(d.Uri, strings.TrimSuffix(uri, "/")) {
		return false
	}

	var (
		ha1      [md5.Size]byte
		ha2      [md5.Size]byte
		response [md5.Size]byte
	)

	ha1Str := d.Username + ":" + d.Realm + ":" + password
	if d.Algorithm == "MD5-sess" {
		ha1 = md5.Sum(StringToBytes(ha1Str))
		ha1Str = hex.EncodeToString(ha1[:]) + ":" + d.Nonce + ":" + d.Cnonce
	}
	ha1 = md5.Sum(StringToBytes(ha1Str))

	// fmt.Printf("HA1 = %x\n", ha1)

	ha2Str := method + ":" + d.Uri
	if d.Qop == "auth-int" {
		ha2 = md5.Sum(body)
		ha2Str = method + ":" + d.Uri + ":" + BytesToString(ha2[:])
	}
	ha2 = md5.Sum(StringToBytes(ha2Str))

	// fmt.Printf("HA2 = %x\n", ha2)

	var responseStr string

	if d.Qop == "auth" || d.Qop == "auth-int" {
		responseStr = hex.EncodeToString(ha1[:]) + ":" + d.Nonce + ":" + d.Nc + ":" + d.Cnonce + ":" + d.Qop + ":" + hex.EncodeToString(ha2[:])
	} else {
		responseStr = hex.EncodeToString(ha1[:]) + ":" + d.Nonce + ":" + hex.EncodeToString(ha2[:])
	}

	// fmt.Printf("responseStr = %s\n", responseStr)
	response = md5.Sum(StringToBytes(responseStr))
	// fmt.Printf("L = %s\n", d.Response)
	// fmt.Printf("R = %x\n", response)

	return hex.EncodeToString(response[:]) == d.Response
}

// StringToBytes converts string to byte slice without a memory allocation.
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func assertBool(guard bool, text string) {
	if !guard {
		panic(text)
	}
}

func readDigest(authValue string) *Digest {
	if len(authValue) == 0 {
		return nil
	}

	sp := strings.Index(authValue, " ")
	if sp < strings.Index(authValue, ",") {
		authValue = authValue[sp+1:]
	}

	pairs := make(map[string]string, strings.Count(authValue, ";"))
	authValue = textproto.TrimString(authValue)

	var part string
	for len(authValue) > 0 { // continue since we have rest
		part, authValue, _ = strings.Cut(authValue, ",")
		if part = textproto.TrimString(part); part != "" {
			key, value, _ := strings.Cut(part, "=")
			if val, ok := parseDigestValue(value, true); ok {
				pairs[strings.ToLower(key)] = val
			}
		}
	}

	return &Digest{
		Username:  pairs["username"],
		Realm:     pairs["realm"],
		Nonce:     pairs["nonce"],
		Uri:       pairs["uri"],
		Algorithm: pairs["algorithm"],
		Cnonce:    pairs["cnonce"],
		Nc:        pairs["nc"],
		Qop:       pairs["qop"],
		Response:  pairs["response"],
	}
}

func validDigestValueByte(b byte) bool {
	return 0x20 <= b && b < 0x7f && b != '"' && b != ';' && b != '\\'
}

func parseDigestValue(raw string, allowDoubleQuote bool) (string, bool) {
	// Strip the quotes, if present.
	if allowDoubleQuote && len(raw) > 1 && raw[0] == '"' && raw[len(raw)-1] == '"' {
		raw = raw[1 : len(raw)-1]
	}
	for i := 0; i < len(raw); i++ {
		if !validDigestValueByte(raw[i]) {
			return "", false
		}
	}
	return raw, true
}
