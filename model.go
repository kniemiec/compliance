package main

type ComplianceCheckRequest struct {
	TransferId   string   `json:"transferId"`
	SenderData   UserData `json:"senderData"`
	ReceiverData UserData `json:"receiverData"`
}

type ComplianceCheckResponse struct {
	TransferId     string              `json:"transferId"`
	Status         string              `json:"status"`
	SenderStatus   []ComplianceProblem `json:"senderStatus"`
	ReceiverStatus []ComplianceProblem `json:"receiverStatus"`
}

type UserData struct {
	UserId  string      `json:"userId"`
	Surname string      `json:"surname"`
	Name    string      `json:"name"`
	Address AddressData `json:"address"`
}

type AddressData struct {
	Country    string `json:"country"`
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	Street     string `json:"street"`
}

type ComplianceProblem struct {
	ProblemId          string `json:"problemId"`
	ProblemDescription string `json:"problemDescription"`
}
