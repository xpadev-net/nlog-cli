package main

import (
	"context"
	"fmt"
	pb "github.com/xpadev-net/nlog-cli/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func getConnection() pb.LoggingServiceClient {
	address := config.GrpcEndpoint
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return nil
	}
	return pb.NewLoggingServiceClient(conn)
}

func endTask(conn pb.LoggingServiceClient, taskId int, exitCode int) error {
	if taskId < 0 {
		return nil
	}
	body := &pb.EndTaskRequest{
		TaskId:   int64(taskId),
		ExitCode: int32(exitCode),
	}
	_, err := conn.EndTask(context.Background(), body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func appendLog(conn pb.LoggingServiceClient, taskId int, logType pb.Log_LogType, message string) (int, error) {
	if taskId < 0 {
		return -1, nil
	}
	body := &pb.AppendLogRequest{
		Log: &pb.Log{
			TaskId:  int64(taskId),
			Type:    logType,
			Message: message,
		},
	}
	res, err := conn.AppendLog(context.Background(), body)
	if err != nil {
		return -1, err
	}
	return int(res.LogId), nil
}

func createTask(conn pb.LoggingServiceClient, itemId int, issuer string, workDir string, command string, pid int) (int, error) {
	body := pb.CreateTaskRequest{
		ItemId:    int64(itemId),
		Issuer:    issuer,
		WorkDir:   workDir,
		Command:   command,
		ProcessId: int64(pid),
	}
	res, err := conn.CreateTask(context.Background(), &body)
	if err != nil {
		return -1, err
	}
	return int(res.TaskId), nil
}

func ping(conn pb.LoggingServiceClient) {
	_, err := conn.Ping(context.Background(), &pb.PingRequest{})
	if err != nil {
		log.Fatal(err)
	}
}
