package repository_test

import (
	"log"
	"testing"

	"testjun/pkg/repository"
)




func Test_UserService(t *testing.T) {
	//Arrange
	session, err := repository.New("mongodb://127.0.0.1:27017")

	if err != nil {
		log.Fatalf("Unable to connect to mongo session: %s", err)
	}
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}


	testUsername := "integration_test_RRRr"
	testPassword := "integration_test_password"

	_, err = session.Create(testUsername, testPassword)

	//Assert
	if err != nil{
		log.Fatalf("Unable to create user: %s", err)
	}
	result, _, _ := session.Find(testUsername)
	if result.GUID != testUsername {
		log.Fatalf("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, result.GUID)
	}
}