package repository

import (
	"context"
	"log"
	"test-mongodb/model/collection"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindUser(ctx context.Context, email string) (collection.User, error)
	FindUsers(ctx context.Context) []collection.User
	CreateUser(ctx context.Context, user collection.User) error
	DeleteUser(ctx context.Context, email string) error
	UpdateUser(ctx context.Context, email string, user collection.User) error
	CountByEmail(ctx context.Context, email string) (int, error)
	CountByEmailPass(ctx context.Context, email string, password string) (int, error)

	CreateFeed(ctx context.Context, email string, feed collection.Feed) error
}

type UserRepositoryImpl struct {
	Collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserRepositoryImpl{
		Collection: collection,
	}
}

func (r *UserRepositoryImpl) FindUser(ctx context.Context, email string) (collection.User, error) {
	var user collection.User
	err := r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	return user, err
}

func (r *UserRepositoryImpl) FindUsers(ctx context.Context) []collection.User {
	var users []collection.User
	cursor, err := r.Collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err)
	}

	return users
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user collection.User) error {
	_, err := r.Collection.InsertOne(ctx, user)

	return err
}

func (r *UserRepositoryImpl) DeleteUser(ctx context.Context, email string) error {
	_, err := r.Collection.DeleteOne(ctx, bson.M{"email": email})

	return err
}

func (r *UserRepositoryImpl) UpdateUser(ctx context.Context, email string, user collection.User) error {

	update := bson.D{{"$set",
		bson.M{
			"full_name":  user.FullName,
			"password":   user.Password,
			"phone":      user.Phone,
			"user_type":  user.UserType,
			"updated_at": user.UpdatedAt,
		}}}

	_, err := r.Collection.UpdateOne(ctx, bson.M{"email": email}, update)

	return err
}

func (r *UserRepositoryImpl) CountByEmail(ctx context.Context, email string) (int, error) {
	count, err := r.Collection.CountDocuments(ctx, bson.M{"email": email})

	return int(count), err
}

func (r *UserRepositoryImpl) CountByEmailPass(ctx context.Context, email string, password string) (int, error) {
	count, err := r.Collection.CountDocuments(ctx, bson.M{"email": email, "password": password})

	return int(count), err
}

func (r *UserRepositoryImpl) CreateFeed(ctx context.Context, email string, feed collection.Feed) error {
	update := bson.D{{"$push",
		bson.M{
			"feeds": bson.M{
				"caption":    feed.Caption,
				"created_at": feed.CreatedAt,
			},
		}}}

	_, err := r.Collection.UpdateOne(ctx, bson.M{"email": email}, update)

	return err
}
