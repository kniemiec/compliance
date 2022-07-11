package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSuccessfull(t *testing.T) {
	assert := assert.New(t)
	user := UserData{
		UserId:  "id",
		Surname: "Kowalski",
		Name:    "Andrzej",
	}
	repository := &MockRepository{}
	complinaceManager := ComplianceManagerImpl{repository: repository}

	response := complinaceManager.create(user)
	assert.Equal(int32(0), response.status, "Unexpected Error when inserting data")
	assert.Equal("", response.description, "Unexpected description when nil expected")
}

func TestFindAllSuccessfull(t *testing.T) {
	assert := assert.New(t)
	repository := &MockRepository{}
	complinaceManager := ComplianceManagerImpl{repository: repository}

	response := complinaceManager.retrieveAllUsers()
	assert.Equal(1, len(response), "Unexpected number of elements in response")
}
