package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthUserKey is the cookie name for user credential in digest auth.
const AuthUserKey = "user"

// Accounts defines a key/value for username/password list of authorized logins.
type Accounts map[string]string

func (as Accounts) searchCredential(authValue, method, realm, url string, body []byte) (string, bool) {
	if authValue == "" {
		return "", false
	}

	d := readDigest(authValue)
	for username, password := range as {
		d.Username = username
		if d.Validate(password, method, realm, url, body) {
			return username, true
		}
	}

	return "", false
}

// DigestAuthForRealm returns a Digest HTTP Authorization middleware. It takes as arguments a map[string]string where
// the key is the username and the value is the password, as well as the name of the Realm.
// If the realm is empty, "gin@golang" will be used by default.
// (see https://datatracker.ietf.org/doc/html/rfc2617)
func DigestAuthForRealm(accounts Accounts, realm string) gin.HandlerFunc {
	if realm == "" {
		realm = "gin@golang"
	}

	nonce := make([]byte, 32)
	rand.Read(nonce)
	algorithm := "MD5"

	authenticate := fmt.Sprintf(
		"Digest realm=%s, nonce=%s, algorithm=%s, qop=%s",
		strconv.Quote(realm),
		strconv.Quote(hex.EncodeToString(nonce)),
		algorithm,
		strconv.Quote("auth"),
	)

	return func(c *gin.Context) {
		// Search user in the slice of allowed credentials
		user, found := accounts.searchCredential(
			c.GetHeader("Authorization"),
			c.Request.Method,
			realm,
			c.Request.URL.Path,
			nil,
		)
		if !found {
			// Credentials don't match, we return 401 and abort handlers chain.
			c.Header("WWW-Authenticate", authenticate)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// The user credentials was found, set user's id to key AuthUserKey in this context, the user's id can be read later using
		// c.MustGet(gin.AuthUserKey).
		c.Set(AuthUserKey, user)
	}
}

// DigestAuth returns a Digest HTTP Authorization middleware. It takes as argument a map[string]string where
// the key is the username and the value is the password.
func DigestAuth(accounts Accounts) gin.HandlerFunc {
	return DigestAuthForRealm(accounts, "")
}
