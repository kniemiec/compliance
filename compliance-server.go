package main

import (	
	"net/http"

    "github.com/gin-gonic/gin"
)

type ComplianceCheckRequest struct {
	transferId string `json:"transferId"`
	// senderData UserData 'json:"senderData"'
	// receiverData UserData 'json:"receiverData"'
}

type ComplianceCheckResponse struct {
	transferId string `json:"transferId"`
	// senderData UserData 'json:"senderData"'
	// receiverData UserData 'json:"receiverData"'
}

var status = ComplianceCheckRequest {
	transferId: "1",}

func getComplianceStatus(c * gin.Context){
	c.IndentedJSON(http.StatusOK, status)
}


func main() {
	router := gin.Default()
	router.GET("complianceStatus", getComplianceStatus)

	router.Run("localhost:8888")

}
