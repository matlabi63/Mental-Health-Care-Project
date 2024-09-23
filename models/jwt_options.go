// models.go or jwt_options.go (within models package)
package models
type JWTOptions struct {
    SecretKey string
    Issuer    string
    Expiry    int64
    ExpiresDuration int
}
