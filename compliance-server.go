package main

import (
	"fmt"
	"net/http"
    "github.com/gin-gonic/gin"
)


func getComplianceStatus(c *gin.Context){
	var request ComplianceCheckRequest
    if err := c.BindJSON(&request); err != nil {
        c.String(http.StatusBadRequest, "Invalid request")
    }	

	var complianceService ComplinaceDomainLogic = ComplianceServiceImpl { }
	response := complianceService.check(request)
	fmt.Println(response)	
	c.JSON(http.StatusOK, response)
}


func main() {
	router := gin.Default()
	router.PUT("/checkCompliance", getComplianceStatus)
	router.Run(":8092")
}
