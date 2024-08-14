package Results

type GetCurrentRunningTimerResult struct {
	Status          string `json:"status"`
	Duration        int    `json:"duration"`
	User            User   `json:"user"`
	StartedAt       string `json:"startedAt"`
	UserDate        string `json:"userDate"`
	Task            Task   `json:"task"`
	CurrentTaskTime struct {
		User         int           `json:"user"`
		History      []TaskHistory `json:"history"`
		LockReasons  []interface{} `json:"lockReasons"`
		IsLocked     bool          `json:"isLocked"`
		ManualTime   int           `json:"manualTime"`
		ID           int           `json:"id"`
		Date         string        `json:"date"`
		Time         int           `json:"time"`
		TimerTime    int           `json:"timerTime"`
		PastDateTime int           `json:"pastDateTime"`
		Task         Task          `json:"task"`
		CreatedAt    string        `json:"createdAt"`
	} `json:"currentTaskTime"`
	Today int `json:"today"`
}

type User struct {
	AvatarUrl      string      `json:"avatarUrl"`
	AvatarUrlLarge string      `json:"avatarUrlLarge"`
	ID             int         `json:"id"`
	Email          string      `json:"email"`
	Name           string      `json:"name"`
	Headline       string      `json:"headline"`
	Capacity       int         `json:"capacity"`
	Cost           int         `json:"cost"`
	CostHistory    interface{} `json:"costHistory"`
}

type Task struct {
	Iteration  string            `json:"iteration"`
	Projects   []string          `json:"projects"`
	Attributes map[string]string `json:"attributes"`
	Completed  bool              `json:"completed"`
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Name       string            `json:"name"`
	URL        string            `json:"url"`
	Status     string            `json:"status"`
	Labels     []string          `json:"labels"`
	CreatedAt  string            `json:"createdAt"`
	Time       struct {
		Total     int            `json:"total"`
		Users     map[string]int `json:"users"`
		TimerTime int            `json:"timerTime"`
	} `json:"time"`
	Assignees []struct {
		AccountID   string `json:"accountId"`
		AccountName string `json:"accountName"`
		UserID      int    `json:"userId"`
	} `json:"assignees"`
}

type TaskHistory struct {
	ID           int    `json:"id"`
	Time         int    `json:"time"`
	PreviousTime int    `json:"previousTime"`
	Action       string `json:"action"`
	Source       string `json:"source"`
	CreatedAt    string `json:"createdAt"`
	CreatedBy    int    `json:"createdBy"`
}