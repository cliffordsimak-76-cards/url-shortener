package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

var secretKey = []byte("secret key")
var userIDLen = 8

func Cookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		cookie, _ := req.Cookie("userID")

		if cookie != nil {
			err := validateCookie(cookie.Value)
			if err == nil {
				return next(c)
			}
		}

		newCookie, err := generateCookie()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		req.AddCookie(newCookie)
		c.SetCookie(newCookie)
		return next(c)
	}
}

func validateCookie(cookieValue string) error {
	data, err := hex.DecodeString(cookieValue)
	if err != nil {
		return fmt.Errorf("cant decode cookie value")
	}
	h := hmac.New(sha256.New, secretKey)
	h.Write(data[:userIDLen])
	sign := h.Sum(nil)
	if !hmac.Equal(sign, data[userIDLen:]) {
		return fmt.Errorf("wrong cookie value")
	}
	return nil
}

func generateCookie() (*http.Cookie, error) {
	userID, err := generateUserID()
	if err != nil {
		return nil, err
	}
	sign := utils.SignHMAC256(userID, secretKey)
	value := bytes.Join([][]byte{
		userID, sign,
	}, []byte(""))
	return &http.Cookie{
		Name:  "userID",
		Value: hex.EncodeToString(value),
	}, nil
}

func generateUserID() ([]byte, error) {
	userID, err := utils.GenerateRandom(userIDLen)
	if err != nil {
		return nil, fmt.Errorf("error generate userID")
	}
	return userID, nil
}
