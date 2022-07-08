package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ComplianceServer interface {
	Run()
}

type ComplianceServerImpl struct {
	service ComplinaceDomainLogic
}

func (server ComplianceServerImpl) Run() {
	router := gin.Default()
	router.PUT("/checkCompliance", server.getComplianceStatus)
	router.Run(":8092")
}

func (server ComplianceServerImpl) getComplianceStatus(c *gin.Context) {
	var request ComplianceCheckRequest
	if err := c.BindJSON(&request); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
	}

	response := server.service.check(request)
	fmt.Println(response)
	c.JSON(http.StatusOK, response)
}

func main() {
	repository := ComplianceRepositoryImpl{}
	service := ComplianceServiceImpl{repository: repository}
	server := ComplianceServerImpl{service: service}

	repository.initializeRepository()
	server.Run()
}
