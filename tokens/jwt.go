package tokens

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JWT struct {
	Header    map[string]interface{}
	Payload   map[string]interface{}
	Signature string
	original  string
}

func ParseJwt(in string) (*JWT, error) {
	chunks := strings.Split(in, ".")
	if len(chunks) != 3 {
		return nil, fmt.Errorf("invalid JWT token, execting 3 dot separated b64 strings but got %d", len(chunks))
	}

	dec, err := base64.StdEncoding.DecodeString(chunks[0])
	if err != nil {
		return nil, fmt.Errorf("invalid jwt Header: %s", err.Error())
	}

	h := make(map[string]interface{})
	err = json.Unmarshal(dec, &h)
	if err != nil {
		return nil, fmt.Errorf("invalid jwt Header: %s", err.Error())
	}

	dec, err = base64.StdEncoding.DecodeString(chunks[1])
	if err != nil {
		return nil, fmt.Errorf("invalid jwt Payload: %s", err.Error())
	}
	p := make(map[string]interface{})
	err = json.Unmarshal(dec, &p)
	if err != nil {
		return nil, fmt.Errorf("invalid jwt Payload: %s", err.Error())
	}
	return &JWT{
		Header:    h,
		Payload:   p,
		Signature: chunks[2],
		original:  in,
	}, nil
}

func (j *JWT) String() string {
	h, err := json.MarshalIndent(j.Header, "", " ")
	if err != nil {
		panic(err)
	}
	p, err := json.MarshalIndent(j.Payload, "", " ")
	return "Header:\n" +
		string(h) +
		"\nPayload:\n" +
		string(p) +
		"\nSignature: " + j.Signature
}

func (j *JWT) Valid() (bool, error) {
	panic("not implemented")
}

func (j *JWT) Expiry() time.Time {
	panic("not implemented")
}

func (j *JWT) Emission() time.Time {
	panic("not implemented")
}

func (j *JWT) Dump() string {
	buf, err := json.MarshalIndent(j, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func (j *JWT) Parse(in string) error {
	tmp := &JWT{}
	err := json.Unmarshal([]byte(in), tmp)
	if err != nil {
		return err
	}
	j.Header = tmp.Header
	j.Payload = tmp.Payload
	j.String = tmp.Signature
}

func (j *JWT) Token() string {
	panic("not implemented")
}
