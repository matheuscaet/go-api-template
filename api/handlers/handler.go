package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matheuscaet/go-api-template/business"
	task "github.com/matheuscaet/go-api-template/business/types"
	"github.com/matheuscaet/go-api-template/internal/config"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := business.NewTaskService().GetTasks(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	task, err = business.NewTaskService().CreateTask(r.Context(), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task task.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	task, err = business.NewTaskService().UpdateTask(r.Context(), task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	err := business.NewTaskService().DeleteTask(r.Context(), r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted"})
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetTasks(w, r)
	case "POST":
		CreateTask(w, r)
	case "PUT":
		UpdateTask(w, r)
	case "DELETE":
		DeleteTask(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func StartServer() {
	http.HandleFunc("/health", HealthCheck)
	http.HandleFunc("/tasks", HandleTasks)
	http.ListenAndServe(":"+config.Port, nil)
}
