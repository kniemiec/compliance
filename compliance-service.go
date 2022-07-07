package main

type ComplinaceDomainLogic interface {
	check(request ComplianceCheckRequest) ComplianceCheckResponse
}

type ComplianceServiceImpl struct {
}

type SanctionedUser struct {
	lastname string
	surname  string
}

var sanctionedUsers = []SanctionedUser{
	{lastname: "Jan", surname: "Kowalski"},
	{lastname: "Andrzej", surname: "Nowak"},
}

func (service ComplianceServiceImpl) check(request ComplianceCheckRequest) ComplianceCheckResponse {
	response := ComplianceCheckResponse{}
	problemsDetected := false
	response.TransferId = request.TransferId
	for _, user := range sanctionedUsers {
		if user.lastname == request.SenderData.Lastname && user.surname == request.SenderData.Surname {
			response.SenderStatus = append(response.SenderStatus, ComplianceProblem{ProblemId: "1", ProblemDescription: "User on sanctions list"})
			problemsDetected = true
		}
		if user.lastname == request.ReceiverData.Lastname && user.surname == request.ReceiverData.Surname {
			response.ReceiverStatus = append(response.ReceiverStatus, ComplianceProblem{ProblemId: "1", ProblemDescription: "User on sanctions list"})
			problemsDetected = true
		}
	}
	if problemsDetected {
		response.Status = "ALERT"
	} else {
		response.Status = "OK"
	}
	return response
}

func test() int {
	return 5
}
