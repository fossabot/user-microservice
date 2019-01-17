package entity

import (
	"time"

	"github.com/guregu/dynamo"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/thomaspoignant/user-microservice/db"
)

type User struct {
	// id of the User (format UUID)
	ID        string `dynamo:"id,hash" json:"id" example:"8da8adc3-0ae9-47b2-884c-ee41e691ff57"`
	FirstName string `dynamo:"first_name" json:"first_name" example:"John"`
	LastName  string `dynamo:"last_name" json:"last_name" example:"Doe"`
	// creation date of the entry (format example 2019-01-17T21:03:08.373394+01:00)
	CreatedAt time.Time `dynamo:"created_at" json:"created_at" example:"2019-01-17T21:03:08.373394+01:00"`
	// last update date of the entry (format example 2019-01-17T21:03:08.373394+01:00)
	UpdatedAt time.Time `dynamo:"updated_at" json:"updated_at" example:"2019-01-17T21:03:08.373394+01:00"`
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
