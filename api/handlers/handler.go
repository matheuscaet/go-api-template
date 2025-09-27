package handlers

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK GET"))
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK POST"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK PUT"))
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK DELETE"))
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
	http.ListenAndServe(":8080", nil)
}
