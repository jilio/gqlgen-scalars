package scalar

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Address represents an Ethereum address,
// which is usable as a GraphQL scalar type.
type Address common.Address

// MarshalGQL implements the graphql.Marshaler interface found in gqlgen,
// allowing the type to be marshaled by gqlgen and sent over the wire.
// This will convert the Address to a lowercase hexadecimal string with "0x" prefix.
func (a Address) MarshalGQL(w io.Writer) {
	addr := common.Address(a)
	io.WriteString(w, fmt.Sprintf(`"%s"`, strings.ToLower(addr.Hex())))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface found in gqlgen,
// allowing the type to be received by a graphql client and unmarshaled.
func (a *Address) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("Ethereum address must be a string")
	}

	// Remove "0x" prefix if present
	str = strings.TrimPrefix(str, "0x")

	// Check if the string is a valid Ethereum address
	if !common.IsHexAddress("0x" + str) {
		return fmt.Errorf("invalid Ethereum address: %s", str)
	}

	addr := common.HexToAddress(str)
	*a = Address(addr)
	return nil
}

// Hex returns the hexadecimal string representation of the address.
func (a Address) Hex() string {
	return common.Address(a).Hex()
}

// String implements the fmt.Stringer interface.
func (a Address) String() string {
	return a.Hex()
}

// Bytes returns the address as a byte slice.
func (a Address) Bytes() []byte {
	return common.Address(a).Bytes()
}
