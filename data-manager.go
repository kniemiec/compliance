package main

type OperationResult struct {
	status      int32
	description string
}

type ComplianceManager interface {
	create(user UserData) OperationResult
	retrieveAllUsers() []UserData
	// update(user UserData) OperationResult
	// delete(userId String) OperationResult
}

type ComplianceManagerImpl struct {
	repository ComplianceRepository
}

func (manager ComplianceManagerImpl) create(user UserData) OperationResult {
	return manager.repository.insert(user)
}

func (manager ComplianceManagerImpl) retrieveAllUsers() []UserData {
	return manager.repository.findAll()
}
