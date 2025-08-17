package task

type Task struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (t *Task) MarkAsDone() {
	t.Done = true
}
