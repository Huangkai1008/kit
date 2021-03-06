package jwtauth

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	kitmsg "github.com/Huangkai1008/kit/pkg/message"
)

// Options is a struct for specifying configuration options for the JWT.
type Options struct {
	JwtSecretKey           string
	JwtAccessTokenExpires  time.Duration
	JwtRefreshTokenExpires time.Duration
	JwtIssuer              string
	JwtSubject             string
	JwtAudience            []string
}

// NewOptions returns a new instance of the options with the given parameters.
func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)

	if err = v.UnmarshalKey("jwt", o); err != nil {
		return nil, errors.Wrap(err, kitmsg.LoadConfigError)
	}
	return o, err
}
