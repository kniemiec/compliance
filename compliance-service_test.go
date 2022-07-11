package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComplianceStatusWhenSanctionedUsersFound(t *testing.T) {
	assert := assert.New(t)
	request := ComplianceCheckRequest{
		TransferId:   "id",
		SenderData:   UserData{Surname: "Nowak", Name: "Andrzej"},
		ReceiverData: UserData{Surname: "Kowalski", Name: "Alojzy"},
	}
	repository := &MockRepository{}
	complianceService := ComplianceServiceImpl{repository: repository}

	response := complianceService.check(request)
	assert.Equal(request.TransferId, response.TransferId, "Unexpected TransferId")
	assert.Equal("ALERT", response.Status, "Unexpected Status")
}

func TestGetComplianceStatusWhenSanctionedUsersNotFound(t *testing.T) {
	assert := assert.New(t)
	request := ComplianceCheckRequest{
		TransferId:   "id",
		SenderData:   UserData{Surname: "Kowalski", Name: "Andrzej"},
		ReceiverData: UserData{Surname: "Kowalski", Name: "Alojzy"},
	}
	repository := &MockRepository{}
	complianceService := ComplianceServiceImpl{repository: repository}

	response := complianceService.check(request)
	assert.Equal(request.TransferId, response.TransferId, "Unexpected TransferId")
	assert.Equal("OK", response.Status, "Unexpected Status")
}
