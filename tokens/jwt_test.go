package tokens

import "testing"

var sampleJWT = struct {
	encoded string
	decoded string
}{
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ",
	`Header:
{
 "alg": "HS256",
 "typ": "JWT"
}
Payload:
{
 "admin": true,
 "name": "John Doe",
 "sub": "1234567890"
}
Signature: TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ`,
}

func TestParseJwt(t *testing.T) {
	_, err := ParseJwt(sampleJWT.encoded)
	if err != nil {
		t.Errorf("Unexpected ParseJwt error: %s", err.Error())
	}
}

func TestString(t *testing.T) {
	jwt, _ := ParseJwt(sampleJWT.encoded)
	if jwt.String() != sampleJWT.decoded {
		t.Errorf("Error in string output: was expecting: ```\n%s```\nbut got: ```\n%s```", sampleJWT.decoded, jwt.String())
	}
}
