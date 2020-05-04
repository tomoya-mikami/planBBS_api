package Plan

type RepositoryInterface interface {
	Add(plan Plan) error
	FindAll() ([]Plan, error)
}
