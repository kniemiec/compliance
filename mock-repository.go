package main

type MockRepository struct {
}

func (repository *MockRepository) findByNameAndLastName(lastName string, name string) []SanctionedUser {
	var result []SanctionedUser
	user := SanctionedUser{surname: "Nowak", name: "Andrzej"}
	if lastName == user.surname && name == user.name {
		result = append(result, user)
	}
	return result
}

func (repository MockRepository) initializeRepository() {}

func (repository MockRepository) insert(user UserData) OperationResult {
	return OperationResult{
		status: 0,
	}
}

func (repository MockRepository) findAll() []UserData {
	return []UserData{{UserId: "1", Surname: "Nowak", Name: "Andrzej"}}
}
