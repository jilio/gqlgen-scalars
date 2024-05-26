# gqlgen-scalars

This repository contains custom scalar types for [gqlgen](https://gqlgen.com/), a Go library for building GraphQL servers.

## Scalar Types

- `Base64`: Represents base64-encoded data.
- `BigInt`: Represents a large integer using `math/big.Int`.

## Installation

To install the package, run:

```bash
go get github.com/jilio/gqlgen-scalars
```

## Usage

To use these scalar types in your gqlgen project:

1. Import the package in your Go code:

  ```go
   import "github.com/your-username/gqlgen-scalars/scalar"
   ```

2. Add the scalar types to your GraphQL schema:

  ```graphql
  scalar Base64
  scalar BigInt
  ```

3. Use the scalar types in your GraphQL resolvers:

  ```go
  type Resolver struct{}

  func (r *Resolver) Base64() scalar.Base64 {
    return scalar.Base64("SGVsbG8gV29ybGQ=")
  }

  func (r *Resolver) BigInt() scalar.BigInt {
    return scalar.BigInt("123456789012345678901234567890")
  }
  ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
