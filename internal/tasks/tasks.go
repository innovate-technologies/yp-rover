package tasks

// Task defines a dispatched task
type Task struct {
	Unit     string            `json:"unit"`
	Function string            `json:"function"`
	Args     map[string]string `json:"args"`
}
