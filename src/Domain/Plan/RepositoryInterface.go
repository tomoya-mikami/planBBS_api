package Plan

type RepositoryInterface interface {
	Add(plan Plan) error
	findAll() ([]Plan, error)
}
