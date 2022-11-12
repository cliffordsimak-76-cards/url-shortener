package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/config"
	"github.com/cliffordsimak-76-cards/url-shortener/internal/app/utils"
	"github.com/labstack/echo/v4"
)

var ErrShortCookieValue = errors.New("error cookie value is too short")

// cookie middleware.
func Cookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		cookie, _ := req.Cookie(config.UserCookieName)
		if cookie != nil {
			err := validateCookie(cookie.Value)
			if err != nil {
				if errors.Is(err, ErrShortCookieValue) {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			}
			return next(c)
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
	if len(data) == 0 {
		return ErrShortCookieValue
	}
	if err != nil {
		return fmt.Errorf("cant decode cookie value")
	}
	h := hmac.New(sha256.New, []byte(config.SecretKey))
	n, err := h.Write(data[:config.UserIDLen])
	if err != nil {
		return err
	}
	if n != config.UserIDLen {
		return fmt.Errorf("wrong number of bytes written")
	}
	sign := h.Sum(nil)
	if !hmac.Equal(sign, data[config.UserIDLen:]) {
		return fmt.Errorf("wrong cookie value")
	}
	return nil
}

func generateCookie() (*http.Cookie, error) {
	userID, err := generateUserID()
	if err != nil {
		return nil, err
	}
	sign := utils.SignHMAC256(userID, []byte(config.SecretKey))
	cookieValue := bytes.Join([][]byte{userID, sign}, []byte(""))
	return &http.Cookie{
		Name:  config.UserCookieName,
		Value: hex.EncodeToString(cookieValue),
	}, nil
}

func generateUserID() ([]byte, error) {
	userID, err := utils.GenerateRandom(config.UserIDLen)
	if err != nil {
		return nil, fmt.Errorf("error generate userID")
	}
	return userID, nil
}
