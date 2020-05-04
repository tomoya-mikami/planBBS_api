package Plan

type Service struct {
	repository RepositoryInterface
}

func (s Service) Add() error {
	events := []Event{
		{Title: "空港", Description: "羽田空港にむかいます", URL: "https://tokyo-haneda.com/index.html", Date: 1588491615665},
		{Title: "ねこ", Description: "ねこです", URL: "", Date: 1588491620000},
	}
	plan := Plan{Title: "旅行プラン", Description: "赤坂の旅行プランです", Events: &events}

	return s.repository.Add(plan)
}

func (s Service) findAll() ([]Plan, error) {
	return s.repository.findAll()
}
