package core

import (
	"time"

	"github.com/guregu/dynamo"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
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

func (service *UserService) saveUser(user *User) error {
	now := time.Now()
	if user.ID == "" {
		user.CreatedAt = now
		user.ID = uuid.NewV4().String()
	}
	user.UpdatedAt = now
	return service.UserTable.Put(user).Run()
}

func (service *UserService) getUser(user *User) error {
	log.WithField("user", user).Info("Param for update")
	return service.UserTable.Get("id", user.ID).One(user)
}

func (service *UserService) deleteUser(user *User) error {
	return service.UserTable.Delete("id", user.ID).Run()
}
