package scalar

import (
	"errors"
	"fmt"
	"io"

	"github.com/shopspring/decimal"
)

// Decimal represents a decimal number using shopspring/decimal.Decimal,
// which is usable as a GraphQL scalar type.
type Decimal struct {
	decimal.Decimal
}

// MarshalGQL implements the graphql.Marshaler interface found in gqlgen,
// allowing the type to be marshaled by gqlgen and sent over the wire.
// This will convert the Decimal to a JSON number as a string.
func (d Decimal) MarshalGQL(w io.Writer) {
	io.WriteString(w, `"`+d.String()+`"`) // Wrap in quotes to ensure it's treated as a string
}

// UnmarshalGQL implements the graphql.Unmarshaler interface found in gqlgen,
// allowing the type to be received by a graphql client and unmarshaled.
func (d *Decimal) UnmarshalGQL(v interface{}) error {
	switch value := v.(type) {
	case string:
		dec, err := decimal.NewFromString(value)
		if err != nil {
			return fmt.Errorf("invalid decimal value: %w", err)
		}
		d.Decimal = dec
	case float64:
		d.Decimal = decimal.NewFromFloat(value)
	case int:
		d.Decimal = decimal.NewFromInt(int64(value))
	case int64:
		d.Decimal = decimal.NewFromInt(value)
	default:
		return errors.New("invalid type for Decimal")
	}
	return nil
}

// NewDecimal creates a new Decimal from a decimal.Decimal
func NewDecimal(d decimal.Decimal) Decimal {
	return Decimal{Decimal: d}
}
