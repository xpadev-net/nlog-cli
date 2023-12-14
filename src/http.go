package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	mainpb "main/grpc/proto/main"
	"net/http"
)

func getConnection() {
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = hellopb.NewGreetingServiceClient(conn)
}

func endTask(taskId int, exitCode int) error {
	if taskId < 0 {
		return nil
	}
	body := EndTaskBody{
		TaskId:   taskId,
		ExitCode: exitCode,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = HttpPost(fmt.Sprintf("%s/api/logger/endTask", config.Endpoint), string(jsonData))
	if err != nil {
		return err
	}
	return nil
}

func appendLog(taskId int, logType string, message string) (int, error) {
	if taskId < 0 {
		return -1, nil
	}
	body := AppendLogBody{
		TaskId:  taskId,
		Type:    logType,
		Message: message,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return -1, err
	}
	res, err := HttpPost(fmt.Sprintf("%s/api/logger/appendLog", config.Endpoint), string(jsonData))
	if err != nil {
		return -1, err
	}
	var resObj AppendLogResponse
	err = json.Unmarshal([]byte(res), &resObj)
	if err != nil {
		return -1, err
	}
	return resObj.LogId, nil
}

func createTask(itemId int, issuer string, workDir string, command string, pid int) (int, error) {
	body := CreateTaskBody{
		ItemId:    itemId,
		Issuer:    issuer,
		WorkDir:   workDir,
		Command:   command,
		ProcessId: pid,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return -1, err
	}
	res, err := HttpPost(fmt.Sprintf("%s/api/logger/createTask", config.Endpoint), string(jsonData))
	if err != nil {
		return -1, err
	}
	var resObj CreateTaskResponse
	err = json.Unmarshal([]byte(res), &resObj)
	if err != nil {
		return -1, err
	}
	return resObj.TaskId, nil
}

func HttpPost(url string, json string) (string, error) {

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(json)),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return string(b), nil
}
