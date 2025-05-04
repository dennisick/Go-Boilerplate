package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	AccessTokenExpiry  = 15 * time.Minute
	RefreshTokenExpiry = 7 * 24 * time.Hour
)

type UserToken struct {
	ID        uuid.UUID
	SubjectID uuid.UUID
	ExpiresAt time.Time
	CreatedAt time.Time
}

type UserTokenDTO struct {
	JTI string `json:"jti"`
	Sub string `json:"sub"`
	Exp string `json:"exp"`
}

func (s *UserToken) ToJWT(secretKey []byte) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"jti": s.ID.String(),
			"sub": s.SubjectID.String(),
			"exp": s.ExpiresAt.Unix(),
			"iat": s.CreatedAt.Unix(),
		},
	)

	return token.SignedString(secretKey)
}
