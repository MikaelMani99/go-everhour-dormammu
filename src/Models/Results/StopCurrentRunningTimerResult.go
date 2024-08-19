package Results

type StopCurrentRunningTimerResult struct {
	Status   string   `json:"status"`
	TaskTime TaskTime `json:"taskTime"`
}

type TaskTime struct {
	User         int       `json:"user"`
	History      []History `json:"history"`
	LockReasons  []string  `json:"lockReasons"`
	IsLocked     bool      `json:"isLocked"`
	ManualTime   int       `json:"manualTime"`
	ID           int       `json:"id"`
	Date         string    `json:"date"`
	Time         int       `json:"time"`
	TimerTime    int       `json:"timerTime"`
	PastDateTime int       `json:"pastDateTime"`
	Task         Task      `json:"task"`
	CreatedAt    string    `json:"createdAt"`
}

type History struct {
	ID           int    `json:"id"`
	Time         int    `json:"time"`
	PreviousTime int    `json:"previousTime"`
	Action       string `json:"action"`
	Source       string `json:"source"`
	CreatedAt    string `json:"createdAt"`
	CreatedBy    int    `json:"createdBy"`
}

type Time struct {
	Total     int            `json:"total"`
	Users     map[string]int `json:"users"`
	TimerTime int            `json:"timerTime"`
}

type Assignee struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	UserID      int    `json:"userId"`
}
