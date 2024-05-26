package scalar

import (
	"errors"
	"fmt"
	"io"
	"math/big"
)

// BigInt represents a large integer using math/big.Int,
// which is usable as a GraphQL scalar type.
type BigInt struct {
	*big.Int
}

// MarshalGQL implements the graphql.Marshaler interface found in gqlgen,
// allowing the type to be marshaled by gqlgen and sent over the wire.
// This will convert the big.Int to a JSON number as a string.
func (b BigInt) MarshalGQL(w io.Writer) {
	io.WriteString(w, b.String()) // big.Int's String method gives the decimal representation
}

// UnmarshalGQL implements the graphql.Unmarshaler interface found in gqlgen,
// allowing the type to be received by a graphql client and unmarshaled.
func (b *BigInt) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("BigInt must be a string formatted as a large integer")
	}

	b.Int = new(big.Int)
	if _, valid := b.SetString(str, 10); !valid {
		return fmt.Errorf("BigInt could not be parsed from the provided string: %s", str)
	}

	return nil
}
