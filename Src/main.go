package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/MikaelMani99/go-everhour-dormammu/Models/Requests"
	"github.com/MikaelMani99/go-everhour-dormammu/Models/Results"
	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

const CurrentRunningTimerUrl = "https://api.everhour.com/timers/current"
const STartTimerUrl = "https://api.everhour.com/timers"

func main() {
	fmt.Println("Starting the application...")
	godotenv.Load(".env")
	var apiKey = os.Getenv("EVERHOUR_API_KEY")

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	// add a job to the scheduler
	job, err := s.NewJob(
		gocron.DurationJob(
			time.Hour,
		),
		gocron.NewTask(
			func(apiKey string) {
				refreshClockedTask(apiKey)
			},
			apiKey,
		),
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// each job has a unique id
	fmt.Println(job.ID())

	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Duration(time.Hour * 10)):
	}

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		panic(err)
	}
}

func refreshClockedTask(apiKey string) error {
	fmt.Println("Refreshing the task...")
	var currentTask, err = fetchCurrentRunningTask(apiKey)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Stopping the task...")
	err = stopCurrentTask(apiKey)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Starting the task...")
	err = startCurrentTask(apiKey, currentTask)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func fetchCurrentRunningTask(apiKey string) (Results.GetCurrentRunningTimerResult, error) {
	var responseData Results.GetCurrentRunningTimerResult

	request, err := http.NewRequest(
		"GET",
		CurrentRunningTimerUrl,
		nil)

	if err != nil {
		return responseData, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return responseData, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return responseData, err
	}

	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return responseData, err
	}

	if responseData.Status != "active" {
		return responseData, errors.New("no active task")
	}

	return responseData, nil
}

func stopCurrentTask(apiKey string) error {
	request, err := http.NewRequest(
		"DELETE",
		CurrentRunningTimerUrl,
		nil)

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	var responseData Results.GetCurrentRunningTimerResult

	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return err
	}

	return nil
}

func startCurrentTask(apiKey string, currentTask Results.GetCurrentRunningTimerResult) error {

	var startTimerRequest = Requests.StartTimerRequest{
		Task:     currentTask.Task.ID,
		UserDate: currentTask.UserDate,
		Comment:  "Resuming task",
	}

	startTimerJson, err := json.Marshal(startTimerRequest)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		"POST",
		STartTimerUrl,
		bytes.NewBuffer(startTimerJson))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	var responseData Results.GetCurrentRunningTimerResult

	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return err
	}

	if responseData.Status != "active" {
		return errors.New("task not started")
	}

	return nil
}
