package middleware

import "testing"

func TestDigest_Validate(t *testing.T) {
	authValue := `Digest username="Mufasa",
		realm="testrealm@host.com",
		nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
		uri="/dir/index.html",
		qop=auth,
		nc=00000001,
		cnonce="0a4f113b",
		response="6629fae49393a05397450978507c4ef1",
		opaque="5ccc069c403ebaf9f0171e9517f40e41"`

	d := readDigest(authValue)
	d.Username = "Mufasa"

	if !d.Validate("Circle Of Life", "GET", "testrealm@host.com", "/dir/index.html", nil) {
		t.Error("validation failed")
	}
}

func TestDigest_Validate2(t *testing.T) {
	authValue := `Digest username="Mufasa",
		realm="testrealm@host.com",
		nonce="dcd98b7102dd2f0e8b11d0f600bfb0c093",
		uri="/dir/index.html",
		qop=auth-int,
		algorithm=MD5-sess,
		nc=00000001,
		cnonce="0a4f113b",
		response="4bc52afdeedaf89d44503b5b50ab7652",
		opaque="5ccc069c403ebaf9f0171e9517f40e41"`

	d := readDigest(authValue)
	d.Username = "Mufasa"

	if !d.Validate("Circle Of Life", "GET", "testrealm@host.com", "/dir/index.html", nil) {
		t.Error("validation failed")
	}
}
