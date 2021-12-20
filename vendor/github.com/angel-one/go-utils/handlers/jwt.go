package handlers

import (
	"github.com/angel-one/go-utils/log"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

// JWTHandler is the handler for auth with JWT
type JWTHandler interface {
	Middleware() gin.HandlerFunc
	Login() gin.HandlerFunc
}

// JWTOptions is the set of configurations for jwt handling
type JWTOptions struct {
	Realm             string        // The protected space, pass in your App Name here
	SigningAlgorithm  string        // The signing algorithms to use (HS256, HS384, HS512, RS256, RS384, RS512)
	Timeout           time.Duration // The expiration period of a token
	IdentityKey       string        // The key which acts as identity in claims. Will be set as is in your context
	PrivatePassPhrase string        // Passphrase to decrypt the Private key file, if required
	PrivateKeyBytes   []byte        // Private Key used to asymmetrically sign the JWT while creation
	PublicKeyBytes    []byte        // Public Key used to asymmetrically sign the JWT while validation
	SymmetricKeyBytes []byte        // Secret key which will be used to symmetrically sign the JWT

	UnauthorizedFunc func(ctx *gin.Context, code int, message string)
	Authenticator    func(ctx *gin.Context) (interface{}, error)
	PayloadFunc      func(data interface{}) map[string]interface{}
}

func (o *JWTOptions) getIdentityHandler() func(*gin.Context) interface{} {
	return func(ctx *gin.Context) interface{} {
		defer func() {
			if r := recover(); r != nil {
				log.Error(ctx).Stack().Msg("error getting claims")
			}
		}()
		c := jwt.ExtractClaims(ctx)
		if c == nil {
			return nil
		}
		return c[o.IdentityKey]
	}
}

func (o *JWTOptions) getPayloadFunction() func(interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		return o.PayloadFunc(data)
	}
}

// jwtHandler is the implementation for handling jwt
type jwtHandler struct {
	j *jwt.GinJWTMiddleware
}

// Middleware is used to get the middleware handler
func (j *jwtHandler) Middleware() gin.HandlerFunc {
	return j.j.MiddlewareFunc()
}

// Login is used to get the login handler
func (j *jwtHandler) Login() gin.HandlerFunc {
	return j.j.LoginHandler
}

// NewJWTHandler is used to get the new jwt handler
func NewJWTHandler(options JWTOptions) (JWTHandler, error) {
	m, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:                options.Realm,
		SigningAlgorithm:     options.SigningAlgorithm,
		PrivateKeyPassphrase: options.PrivatePassPhrase,
		PrivKeyBytes:         options.PrivateKeyBytes,
		PubKeyBytes:          options.PublicKeyBytes,
		Key:                  options.SymmetricKeyBytes,
		Timeout:              options.Timeout,
		IdentityKey:          options.IdentityKey,
		IdentityHandler:      options.getIdentityHandler(),
		Authorizator:         getAuthorizer(),
		Unauthorized:         options.UnauthorizedFunc,
		TokenLookup:          "header:Authorization",
		TokenHeadName:        "Bearer",
		TimeFunc:             time.Now,
		Authenticator:        options.Authenticator,
		PayloadFunc:          options.getPayloadFunction(),
	})
	if err != nil {
		return nil, err
	}
	return &jwtHandler{j: m}, nil
}

func getAuthorizer() func(interface{}, *gin.Context) bool {
	return func(data interface{}, ctx *gin.Context) bool {
		return true
	}
}
