package scalar

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// Base64 is a custom GraphQL scalar type to represent base64-encoded data.
// It wraps a byte slice to provide base64 encoding and decoding functionality.
type Base64 []byte

// MarshalGQL implements the graphql.Marshaler interface found in gqlgen,
// allowing the type to be marshaled by gqlgen and sent over the wire.
// This will encode the byte slice as a base64 string.
func (b Base64) MarshalGQL(w io.Writer) {
	encoded := base64.StdEncoding.EncodeToString(b)
	io.WriteString(w, `"`+encoded+`"`) // JSON strings must be quoted
}

// UnmarshalGQL implements the graphql.Unmarshaler interface found in gqlgen,
// allowing the type to be received by a graphql client and unmarshaled.
// The input is expected to be a base64-encoded string, which will be decoded
// into the byte slice.
func (b *Base64) UnmarshalGQL(v interface{}) error {
	// Expect that the incoming value is a string.
	str, ok := v.(string)
	if !ok {
		return errors.New("Base64 must be a base64-encoded string")
	}

	// Decode the string into bytes.
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		// Use fmt.Errorf for error wrapping instead of errors.Wrap.
		return fmt.Errorf("failed to decode Base64 string: %w", err)
	}

	*b = decoded
	return nil
}
