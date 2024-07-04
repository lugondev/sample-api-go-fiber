package telegram

import (
	"encoding/json"
	"net/url"
)

func ParseQueryString(rawQuery string) (*TeleAuthData, error) {
	parsedQuery, _ := url.ParseQuery(rawQuery)

	// Initialize the TeleAuthData structure from parsed values
	authData := &TeleAuthData{
		QueryID:  parsedQuery.Get("query_id"),
		AuthDate: parsedQuery.Get("auth_date"),
		Hash:     parsedQuery.Get("hash"),
	}

	// Decode the `user` field from URL-encoded JSON
	userJson := parsedQuery.Get("user")
	userJsonDecoded, _ := url.QueryUnescape(userJson)

	// Unmarshal the user JSON into the `User` struct
	user := UserTeleAuth{}
	err := json.Unmarshal([]byte(userJsonDecoded), &user)
	if err != nil {
		return nil, err
	}

	// Assign the user to the authData
	authData.User = user
	return authData, nil
}
