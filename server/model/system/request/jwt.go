package request

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          uint
	UserName    string
	NickName    string
	AuthorityId string
}
