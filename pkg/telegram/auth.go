package telegram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const AuthToken = "7035388245:AAGDQDjRZa3apH-X7R1XEgy0tlonP6YbTmQ"
const AuthExpire = 60 * 30

func Validate(sp string, token string, options ValidateOptions) error {
	searchParams, err := url.ParseQuery(sp)
	if err != nil {
		return err
	}

	authDate := time.Time{}
	var hash string
	var pairs []string

	for key, values := range searchParams {
		value := values[0] // assuming single-value params
		if key == "hash" {
			hash = value
			continue
		} else if key == "auth_date" {
			authDateInt, err := strconv.Atoi(value)
			if err != nil {
				return errors.New("auth date should present integer")
			}
			authDate = time.Unix(int64(authDateInt), 0)
		}
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
	}

	if hash == "" || authDate.IsZero() {
		return errors.New("data is invalid")
	}

	if options.ExpiresIn > 0 {
		if time.Now().After(authDate.Add(time.Duration(options.ExpiresIn) * time.Second)) {
			return errors.New("session has expired")
		}
	}

	sort.Strings(pairs)
	message := strings.Join(pairs, "\n")

	computedHash := computeHMACSHA256(message, token)
	if computedHash != hash {
		return errors.New("session is invalid")
	}
	return nil
}

func computeHMACSHA256(message, token string) string {
	key := computeHMACSHA256ForWebAppData(token)
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func computeHMACSHA256ForWebAppData(token string) string {
	mac := hmac.New(sha256.New, []byte("WebAppData"))
	mac.Write([]byte(token))
	return string(mac.Sum(nil))
}
