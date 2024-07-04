package telegram_test

import (
	"encoding/json"
	"fmt"
	"sample/api/pkg/telegram"
	"testing"
)

func TestParseQueryTelegramAuth(t *testing.T) {
	rawQuery := "query_id=AAEnANY9AgAAACcA1j3AVDux&user=%7B%22id%22%3A5332402215%2C%22first_name%22%3A%22Narker%22%2C%22last_name%22%3A%22Zero%22%2C%22username%22%3A%22znarker%22%2C%22language_code%22%3A%22en%22%2C%22allows_write_to_pm%22%3Atrue%7D&auth_date=1716894823&hash=780f09ce5f9b57e3d81c2381b788a74bb5dae0efb78ad08202a6219b9ca5313f"
	authData, err := telegram.ParseQueryString(rawQuery)
	if err != nil {
		t.Fatalf("Error parsing query: %v", err)
	}

	// Output the result
	result, _ := json.MarshalIndent(authData, "", "  ")
	fmt.Println(string(result))
}
