package simple

type SimpleRepository struct {
}

type SimpleService struct {
	*SimpleRepository // dependency
}

// PROVIDER

// used to create SimpleRepository
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

// used to create SimpleService, using parameter dependency
func NewSimpleService(simpleRepository *SimpleRepository) *SimpleService {
	return &SimpleService{SimpleRepository: simpleRepository}
}
