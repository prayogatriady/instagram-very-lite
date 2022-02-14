package collection

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FullName  string             `json:"full_name" bson:"full_name" validate:"required,min=2,max=500"`
	Password  string             `json:"password" bson:"password" validate:"required,min=6"`
	Email     string             `json:"email" bson:"email" validate:"email,required"`
	Phone     string             `json:"phone" bson:"phone" validate:"required"`
	UserType  string             `json:"user_type" bson:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Feeds     []Feed             `json:"feeds" bson:"feeds"`
}
