package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

const AuthTokenCookieName = "Authorization"
const AuthTokenHeaderName = "Authorization"
const TokenExpiration = time.Hour * 72

func GetAuthTokenCookieOptions(token string) fiber.Cookie {
	return fiber.Cookie{
		Name:     AuthTokenCookieName,
		Value:    token,
		Expires:  time.Now().Add(TokenExpiration),
		SameSite: "strict",
	}
}
