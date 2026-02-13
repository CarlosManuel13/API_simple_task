package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
)

type Task struct {
	ID string
	Title string
	Completed bool
}

var tasks = []Task{
	{ID: "1", Title: "Run the server", Completed: true},
	{ID: "2", Title: "Make Post function", Completed: true},

}


func manageTasks(w http.ResponseWriter, r *http.Request){
	
	switch r.Method {

		case http.MethodGet: 
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(tasks)
		
		
		case http.MethodPost: 
			newTask := Task{}
			err := json.NewDecoder(r.Body).Decode(&newTask)
		
			if err != nil {
				//We have one error
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			
			}
			
			tasks = append(tasks, newTask)
			
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newTask)
		
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method %s not allowed", r.Method)
	}
}


func main () {

	http.HandleFunc("/tasks", manageTasks)
	fmt.Println("Starting Server...")
	log.Fatal(http.ListenAndServe(":8000", nil))


}
