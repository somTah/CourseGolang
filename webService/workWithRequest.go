package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Course struct {
	Id         int     `json: "id"`
	Name       string  `json: "name"`
	Price      float64 `json: "price"`
	Instructor string  `json: "instructor"`
}

var courseList []Course

func init() {
	CourseJSON := `[
		{
			"id": 1,
			"name": "Python",
			"price": 2590,
			"instructor": "BorntoDev"
		},
		{
			"id": 2,
			"name": "JavaScript",
			"price": 1000,
			"instructor": "BorntoDev"
		},
		{
			"id": 3,
			"name": "SQL",
			"price": 3000,
			"instructor": "BorntoDev"
		}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &courseList)
	if err != nil {
		log.Fatal(err)
	}
}

func courseHandler(w http.ResponseWriter, r *http.Request) { //hand
	courseJSON, err := json.Marshal(courseList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	case http.MethodPost:
		var newCourse Course
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newCourse.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.Id = getNextId()
		courseList = append(courseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func getNextId() int {
	highestId := -1
	for _, course := range courseList {
		if highestId < course.Id {
			highestId = course.Id
		}
	}
	return highestId + 1
}
func main() {
	http.HandleFunc("/course", coursesHandler)
	http.ListenAndServe(":5000", nil)
}
