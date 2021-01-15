package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
)


var tasks []Task


type Task struct {
	Description string	
	Done bool
}


func main() {
	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)
	http.ListenAndServe(":8000",nil)
}


func list(rw http.ResponseWriter, _ *http.Request) {
	data, err := ioutil.ReadFile("Task.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(data, &tasks)
	
	fmt.Println()
	
	for i, tsk := range tasks {
		//Je suis resté bloqué sur la condition
		if tsk.Done ==false {
			fmt.Printf("ID:%v, task: %v\n",i,tsk.Description)
		}
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}


func add(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		http.Error(res,"can'treadbody", http.StatusBadRequest,)
		return
	}
	
	d:= string(body)
	fmt.Println(d)

	tasks = append(tasks,Task{d, false})
	fmt.Println(tasks)
}


func done(rw http.ResponseWriter, _ *http.Request) {
	
}


