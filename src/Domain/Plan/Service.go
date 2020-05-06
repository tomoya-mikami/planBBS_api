package Plan

type ServiceInterface interface {
	Add(plan Plan) error
	FindAll() ([]Plan, error)
	Find(id string) (Plan, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) ServiceInterface {
	service := new(Service)
	service.repository = repository
	return service
}

func (s Service) Add(plan Plan) error {
	return s.repository.Add(&plan)
}

func (s Service) FindAll() ([]Plan, error) {
	return s.repository.FindAll()
}

func (s Service) Find(id string) (Plan, error) {
	return s.repository.Find(id)
}
