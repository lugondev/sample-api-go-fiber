package handler

import (
	"encoding/hex"
	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	Limit     int `json:"limit"`
	QueryPage int `json:"-"`
	Offset    int `json:"-"`
	Page      int `json:"page"`
}

func GetPagination(c *fiber.Ctx) Pagination {
	pagination := Pagination{}
	// get page from query
	page := c.QueryInt("page")
	if page <= 0 {
		page = 1
	}
	pagination.Page = page
	pagination.QueryPage = page - 1
	// get limit from query: default 10, max 100
	limit := c.QueryInt("limit")
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	pagination.Limit = limit
	pagination.Offset = (page - 1) * limit

	return pagination
}

func IsHash(hash string) bool {
	if len(hash) != 64 {
		return false
	}
	source, err := hex.DecodeString(hash)
	if err != nil || len(source) != 32 {
		return false
	}
	return true
}
