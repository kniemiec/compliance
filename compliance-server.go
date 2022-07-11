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
	manager ComplianceManager
}

func (server ComplianceServerImpl) Run() {
	router := gin.Default()
	router.PUT("/checkCompliance", server.getComplianceStatus)
	router.POST("/users", server.createUser)
	router.GET("/users", server.retrieveAllUsers)
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

func (server ComplianceServerImpl) retrieveAllUsers(c *gin.Context) {
	response := server.manager.retrieveAllUsers()
	c.JSON(http.StatusOK, response)
}

func (server ComplianceServerImpl) createUser(c *gin.Context) {
	var user UserData
	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
	}
	response := server.manager.create(user)
	c.JSON(http.StatusOK, response)
}

func main() {
	repository := ComplianceRepositoryImpl{}
	service := ComplianceServiceImpl{repository: repository}
	dataManager := ComplianceManagerImpl{repository: repository}
	server := ComplianceServerImpl{service: service, manager: dataManager}

	repository.initializeRepository()
	server.Run()
}
