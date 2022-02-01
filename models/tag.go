package models

import (
	"time"
)

type Tag struct {
	Identifier string    `json:"id"`
	Name       string    `json:"name"`
	CreatedOn  time.Time `json:"created_on"`
}
