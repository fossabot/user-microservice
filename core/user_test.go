package core

import (
	"testing"

	"github.com/rs/xid"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/thomaspoignant/user-microservice/db"
)

var userServiceTest UserService
var userServiceTestInit bool
var uniqDbName string

func createTableBeforeTest() string {
	if uniqDbName == "" {
		uniqDbName = xid.New().String()
		log.Infof("Database name for Test : %s", uniqDbName)
		setupTable(uniqDbName)
	}
	return uniqDbName
}

func getUserService(uniqDbName string, t *testing.T) UserService {
	if !userServiceTestInit {
		var err error
		currentService, err := newUserService(uniqDbName)
		if err != nil {
			assert.Nil(t, err)
		}
		userServiceTest = *currentService
		userServiceTestInit = true
	}
	return userServiceTest
}

func setupTable(uniqDbName string) {
	db.DynamoDbClient.Table(uniqDbName).DeleteTable()
	//creation of the dynamo table
	db.DynamoDbClient.CreateTable(uniqDbName, User{}).Run()
}

// insert a user and test that object is updated
func TestInsertUser(t *testing.T) {
	uniqDbName := createTableBeforeTest()
	setupTable(uniqDbName)
	service := getUserService(uniqDbName, t)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	err := service.createUser(&user)
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.UpdatedAt)
	assert.NotNil(t, user.CreatedAt)

	value, err := db.DynamoDbClient.Table(uniqDbName).Get("id", user.ID).Count()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), value)

}

//insert a user and try to read what is in database
func TestInsertAndReadUser(t *testing.T) {
	uniqDbName := createTableBeforeTest()
	setupTable(uniqDbName)
	service := getUserService(uniqDbName, t)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	err := service.createUser(&user)
	assert.Nil(t, err)

	value, err := db.DynamoDbClient.Table(uniqDbName).Get("id", user.ID).Count()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), value)

	result := User{
		ID: user.ID,
	}

	err = service.getUser(&result)
	assert.Nil(t, err)

	expectedCreatedAt, _ := user.CreatedAt.MarshalJSON()
	gotCreatedAt, _ := result.CreatedAt.MarshalJSON()
	assert.Equal(t, expectedCreatedAt, gotCreatedAt)

	expectedUpdatedAt, _ := user.UpdatedAt.MarshalJSON()
	gotUpdatedAt, _ := result.UpdatedAt.MarshalJSON()
	assert.Equal(t, expectedUpdatedAt, gotUpdatedAt)

	assert.Equal(t, user.FirstName, result.FirstName)
	assert.Equal(t, user.LastName, result.LastName)
	assert.Equal(t, user.ID, result.ID)
}
