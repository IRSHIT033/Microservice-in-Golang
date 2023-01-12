package tokenutil

import (
	"fmt"
	"strconv"
	"time"

	"github.com/IRSHIT033/E-comm-GO-/server/User_service/domain_user"
	"github.com/golang-jwt/jwt/v4"
)

func CreateAccessToken(user *domain_user.User, secret string, expiry int) (accessToken string, err error) {
	exp := jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry)))
	claims := &domain_user.JwtCustomClaims{
		Name: user.Name,
		ID:   strconv.FormatUint(uint64(user.ID), 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: exp,
		},
	}
	//: jwt.RegisteredClaims{ExpiresAt: exp},
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(user *domain_user.User, secret string, expiry int) (refreshToken string, err error) {
	claimsrefresh := &domain_user.JwtCustomRefreshClaims{
		ID: strconv.FormatUint(uint64(user.ID), 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsrefresh)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method : %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["id"].(string), nil

}
