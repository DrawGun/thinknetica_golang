package main

import "thinknetica_golang/task_17/pkg/api"

// Service сервис авторизации
type Service struct {
	api *api.API
}

func main() {
	service := new()
	service.api.Run()
}

func new() *Service {
	s := Service{
		api: api.New(api.APIPort),
	}

	return &s
}
