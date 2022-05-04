package main


type ComplinaceDomainLogic interface {
	check(request ComplianceCheckRequest) ComplianceCheckResponse
}

type ComplianceServiceImpl struct {
}

type SanctionedUser struct { 
	lastname string
	surname string
}

var sanctionedUsers = []SanctionedUser{
	{ lastname: "Jan", surname: "Kowalski"},
	{ lastname: "Andrzej", surname: "Nowak"},
}


func (service ComplianceServiceImpl) check(request ComplianceCheckRequest) ComplianceCheckResponse {
	response := ComplianceCheckResponse{ }
	for _, user := range sanctionedUsers {
		if user.lastname == request.SenderData.Lastname && user.surname == request.SenderData.Surname {
			response.SenderStatus = append(response.SenderStatus, ComplianceProblem { ProblemId: "1", ProblemDescription : "User on sanctions list"})
		}
		if user.lastname == request.ReceiverData.Lastname && user.surname == request.ReceiverData.Surname {
			response.ReceiverStatus = append(response.ReceiverStatus, ComplianceProblem { ProblemId: "1", ProblemDescription : "User on sanctions list"})
		}
	}
	return response;
}