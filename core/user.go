package core

import (
	"time"

	"github.com/guregu/dynamo"
	uuid "github.com/satori/go.uuid"
	"github.com/thomaspoignant/user-microservice/db"
)

type User struct {
	ID        string    `dynamo:"id,hash"`
	FirstName string    `dynamo:"first_name"`
	LastName  string    `dynamo:"last_name"`
	CreatedAt time.Time `dynamo:"created_at"`
	UpdatedAt time.Time `dynamo:"updated_at"`
}

type UserService struct {
	UserTable *dynamo.Table
}

func newUserService(tableName string) (*UserService, error) {
	table, err := db.GetDynamodbTable(tableName)
	if err != nil {
		return nil, err
	}
	return &UserService{
		UserTable: table,
	}, nil
}

func (service *UserService) createUser(user *User) error {
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	user.ID = uuid.NewV4().String()
	return service.UserTable.Put(user).Run()
}

func (service *UserService) updateUser(user *User) error {
	now := time.Now()
	user.UpdatedAt = now
	return service.UserTable.Update(user.ID, user).Run()
}

func (service *UserService) getUser(user *User) error {
	return service.UserTable.Get("id", user.ID).One(user)
}
