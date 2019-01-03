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
	}
	setupTable(uniqDbName)
	return uniqDbName
}

func getUserService(uniqDbName string, t *testing.T) UserService {
	if !userServiceTestInit {
		var err error
		currentService, err := newUserService(uniqDbName)
		assert.Nil(t, err)

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

// newUserService with an empty table name
func TestEmptyTableName(t *testing.T) {
	_, err := newUserService("")
	assert.NotNil(t, err)
}

// insert a user and test that object is updated
func TestInsertUser(t *testing.T) {
	uniqDbName := createTableBeforeTest()
	service := getUserService(uniqDbName, t)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	err := service.saveUser(&user)
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
	service := getUserService(uniqDbName, t)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	err := service.saveUser(&user)
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

func TestInsertAndDeleteUser(t *testing.T) {
	uniqDbName := createTableBeforeTest()
	service := getUserService(uniqDbName, t)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}
	err := service.saveUser(&user)
	assert.Nil(t, err)

	numberOfUser, err := db.DynamoDbClient.Table(uniqDbName).Get("id", user.ID).Count()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), numberOfUser)

	err = service.deleteUser(&user)
	assert.Nil(t, err)

	numberOfUserAfterDelete, err := db.DynamoDbClient.Table(uniqDbName).Get("id", user.ID).Count()
	assert.Nil(t, err)
	assert.Equal(t, int64(0), numberOfUserAfterDelete)
}

func TestUpdateUser(t *testing.T) {
	uniqDbName := createTableBeforeTest()
	service := getUserService(uniqDbName, t)

	user := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	err := service.saveUser(&user)
	assert.Nil(t, err)

	value, err := db.DynamoDbClient.Table(uniqDbName).Get("id", user.ID).Count()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), value)

	result := user
	err = service.getUser(&user)
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

	updatedUser := user
	updatedUser.FirstName = "John2"
	updatedUser.LastName = "Doe2"

	log.Infof("updatedUser : %s", updatedUser)
	err = service.saveUser(&updatedUser)
	assert.Nil(t, err)
	gotUpdatedCreatedAt, _ := updatedUser.CreatedAt.MarshalJSON()
	assert.Equal(t, expectedCreatedAt, gotUpdatedCreatedAt)
	gotUpdatedUpdatedAt, _ := updatedUser.UpdatedAt.MarshalJSON()
	assert.NotEqual(t, expectedUpdatedAt, gotUpdatedUpdatedAt)
	assert.NotEqual(t, user.FirstName, updatedUser.FirstName)
	assert.NotEqual(t, user.LastName, updatedUser.LastName)
	assert.Equal(t, user.ID, result.ID)

	//we verify that we only have one entry
	value, err = db.DynamoDbClient.Table(uniqDbName).Get("id", updatedUser.ID).Count()
	assert.Nil(t, err)
	assert.Equal(t, int64(1), value)

}
