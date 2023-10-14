package main

type EndTaskBody struct {
	TaskId   int `json:"taskId"`
	ExitCode int `json:"exitCode"`
}

type AppendLogBody struct {
	TaskId  int    `json:"taskId"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type AppendLogResponse struct {
	LogId int `json:"logId"`
}

type CreateTaskBody struct {
	ItemId    int    `json:"itemId"`
	Issuer    string `json:"issuer"`
	WorkDir   string `json:"workDir"`
	Command   string `json:"command"`
	ProcessId int    `json:"processId"`
}

type CreateTaskResponse struct {
	TaskId int `json:"taskId"`
}
