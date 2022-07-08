package main

type ComplinaceDomainLogic interface {
	check(request ComplianceCheckRequest) ComplianceCheckResponse
}

type ComplianceServiceImpl struct {
	repository ComplianceRepository
}

type SanctionedUser struct {
	name    string
	surname string
}

func (service ComplianceServiceImpl) check(request ComplianceCheckRequest) ComplianceCheckResponse {
	response := ComplianceCheckResponse{}
	problemsDetected := false
	response.TransferId = request.TransferId
	senders := service.repository.findByNameAndLastName(request.SenderData.Surname, request.SenderData.Name)
	receivers := service.repository.findByNameAndLastName(request.ReceiverData.Surname, request.ReceiverData.Name)

	if len(senders) > 0 {
		response.SenderStatus = append(response.SenderStatus, ComplianceProblem{ProblemId: "1", ProblemDescription: "User on sanctions list"})
		problemsDetected = true
	}
	if len(receivers) > 0 {
		response.ReceiverStatus = append(response.ReceiverStatus, ComplianceProblem{ProblemId: "1", ProblemDescription: "User on sanctions list"})
		problemsDetected = true
	}

	if problemsDetected {
		response.Status = "ALERT"
	} else {
		response.Status = "OK"
	}
	return response
}
