package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseByID).Methods("GET")
	r.HandleFunc("/createCourse", createCourse).Methods("POST")
	r.HandleFunc("/updateCourse/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/deleteCourse/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8020", r))

}


type Course struct {
	CourseId 		string   `json:"courseId"`
	CourseName 		string	 `json:"courseName"`
	CoursePrice  	int		 `json:"coursePrice"`
	Author     		*Author  `json:"author"`
}

type Author struct {
FullName string   `json:"fullName"`
Website  string   `json:"website"`
}

var courses = []Course{
	{
		CourseId: "4", 
		CourseName: "Full Stack",
		CoursePrice : 2000,
		Author : &Author{
			FullName: "john-Doe",
			Website: "john-doe.com",
		},
	},
		{
		CourseId: "12", 
		CourseName: "MERN",
		CoursePrice : 1500,
		Author : &Author{
			FullName: "Rohit",
			Website: "rohit.com",
		},
	},

}


func getAllCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

json.NewEncoder(w).Encode(courses)

}


func getCourseByID(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id := params["id"]

for _, course := range courses {
	if course.CourseId == id {
		json.NewEncoder(w).Encode(course)
		return
	} 
}


}
func createCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var newCourse Course

	err:= json.NewDecoder(r.Body).Decode(&newCourse)
	if err!=nil {
		// w.Write([]byte("Please send Data in Proper Format"))
		json.NewEncoder(w).Encode("Please send Data in Proper Format")
		return
	}

	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(courses)
}

func updateCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

params := mux.Vars(r)

id := params["id"]

for index, course :=  range courses {
	if course.CourseId == id {
		courses = append(courses[:index], courses[index+1:]...)
		break
	}
}

var updatedCourse Course 

json.NewDecoder(r.Body).Decode(&updatedCourse)

courses = append(courses, updatedCourse)

json.NewEncoder(w).Encode(courses)

}
func deleteCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

params := mux.Vars(r)

id := params["id"]

for index, course := range courses {
	if course.CourseId == id {
		courses = append(courses[:index], courses[index+1:]...)
		json.NewEncoder(w).Encode("Course deleted successfully.")
		return
	}
}


}






