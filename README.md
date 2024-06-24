# gqlgen-scalars

This repository contains custom scalar types for [gqlgen](https://gqlgen.com/), a Go library for building GraphQL servers.

## Scalar Types

- `Base64`: Represents base64-encoded data.
- `BigInt`: Represents a large integer using `math/big.Int`.
- `Decimal`: Represents a decimal number with arbitrary precision using `github.com/shopspring/decimal`.

## Installation

To install the package, run:

```bash
go get github.com/jilio/gqlgen-scalars
```

## Usage

To use these scalar types in your gqlgen project:

1. Import the package in your Go code:

   ```go
   import "github.com/jilio/gqlgen-scalars/scalar"
   ```

2. Add the scalar types to your GraphQL schema:

   ```graphql
   scalar Base64
   scalar BigInt
   scalar Decimal
   ```

3. Configure gqlgen to use these custom scalar types. In your `gqlgen.yml` file, add:

   ```yaml
   models:
     Base64:
       model: github.com/jilio/gqlgen-scalars/scalar.Base64
     BigInt:
       model: github.com/jilio/gqlgen-scalars/scalar.BigInt
     Decimal:
       model: github.com/jilio/gqlgen-scalars/scalar.Decimal
   ```

4. Use the scalar types in your GraphQL resolvers:

   ```go
   type Resolver struct{}

   func (r *Resolver) Base64() scalar.Base64 {
     return scalar.Base64("SGVsbG8gV29ybGQ=")
   }

   func (r *Resolver) BigInt() scalar.BigInt {
     return scalar.BigInt{big.NewInt(123456789012345678901234567890)}
   }

   func (r *Resolver) Decimal() scalar.Decimal {
     return scalar.NewDecimal(decimal.NewFromFloat(123.45))
   }
   ```

## Examples

Here are some examples of how to use these scalar types in your GraphQL queries and mutations:

```graphql
query {
  getBase64Data
  getLargeNumber
  getPreciseDecimal
}

mutation {
  uploadBase64Data(data: "SGVsbG8gV29ybGQ=")
  setLargeNumber(number: "123456789012345678901234567890")
  setPreciseDecimal(value: "123.45000000000000000001")
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
