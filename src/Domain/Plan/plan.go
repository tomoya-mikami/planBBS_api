package Plan

type Event struct {
	Title       string `firestore:"title"`
	Description string `firestore:"description"`
	URL         string `firestore:"url"`
	Date        int64  `firestore:"date"`
}

type Plan struct {
	Title       string `firestore:"title"`
	Description string `firestore:"description"`
	Events      *[]Event
}
