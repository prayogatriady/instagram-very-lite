package collection

import "time"

type Feed struct {
	Caption   string    `json:"caption" bson:"caption"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
