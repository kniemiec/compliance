package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetComplianceStatus(t *testing.T) {
	assert := assert.New(t)
	request := ComplianceCheckRequest{TransferId: "id"}
	complianceService := ComplianceServiceImpl{}
	response := complianceService.check(request)
	assert.Equal(request.TransferId, response.TransferId, "Unexpected TransferId")
	assert.Equal("OK", response.Status, "Unexpected address")
}
