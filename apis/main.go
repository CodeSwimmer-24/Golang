package main

import (
	"crypto/rand"   // For generating random data
	"encoding/json" // For encoding/decoding JSON
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" // For routing
)

// Models for courses

// Course struct represents a course entity with fields for the course name, ID, price, and author details.
type Course struct {
	CourseName  string  `json:"courseName"`  // Field for course name with JSON key "courseName"
	CourseID    string  `json:"courseId"`    // Field for course ID with JSON key "courseId"
	CoursePrice int     `json:"coursePrice"` // Field for course price with JSON key "coursePrice"
	Author      *Author `json:"author"`      // Pointer to an Author struct
}

// Author struct represents an author entity with fields for full name and website.
type Author struct {
	FullName string `json:"fullName"` // Field for author's full name
	Website  string `json:"website"`  // Field for author's website (fixed JSON tag)
}

// Fake DB
// courses is a slice that acts as a mock database for storing Course entries.
var courses []Course

// Middleware/helper

// IsEmpty checks if the given Course struct is empty based on CourseName and CoursePrice.
func (c *Course) IsEmpty() bool {
	return c.CourseName == "" && c.CoursePrice == 0
}

func main() {
	// Initializing a new router
	r := mux.NewRouter()

	// seeding

	courses = append(courses, Course{
		CourseID:    "2F",
		CourseName:  "GoLang",
		CoursePrice: 200,
		Author: &Author{
			FullName: "Fahad Mahmood",
			Website:  "www.golang.com",
		},
	},
	)
	courses = append(courses, Course{
		CourseID:    "29",
		CourseName:  "NodeJs",
		CoursePrice: 120,
		Author: &Author{
			FullName: "Shail Mahmood",
			Website:  "www.nodejs.com",
		},
	},
	)

	// Defining routes and linking them to handler functions
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PATCH")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")

	// Starting the server on port 4000
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers

// homePage handles requests to the home route, responding with a welcome message.
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Gherdekho.com</h1>")) // Send a simple HTML response
}

// getAllCourses handles requests to retrieve all courses.
// It sets the Content-Type header to JSON and encodes the courses slice as JSON.
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) // Encode and send the courses slice as JSON
}

// getOneCourse retrieves a course by its ID from the courses slice.
// It fetches the ID from the request, searches in the slice, and responds accordingly.
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Retrieve the ID from the URL path parameters using mux
	params := mux.Vars(r)
	id := params["id"]

	// Loop through the courses slice to find the course with the matching ID
	for _, course := range courses {
		if course.CourseID == id {
			json.NewEncoder(w).Encode(course) // Encode and send the found course as JSON
			return                            // Exit function after finding and sending the course
		}
	}
	// Send a message if the course is not found
	json.NewEncoder(w).Encode("No data found")
}

// createCourse handles POST requests to create a new course entry.
// It decodes JSON data, checks for validity, generates a unique ID, and adds the course to the database.
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create new course")
	w.Header().Set("Content-Type", "application/json")

	// Check if the request body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data") // Send a response if no data was sent
		return
	}

	var course Course
	// Decode the JSON request body into the course struct
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		json.NewEncoder(w).Encode("Error decoding JSON") // Handle decoding errors
		return
	}

	// Check if the course struct is empty using the IsEmpty method
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in JSON") // Respond if the data is empty
		return
	}

	// Check the duplicate data

	for _, value := range courses {
		if course.CourseName == value.CourseName {
			json.NewEncoder(w).Encode("Cource name already exist")
			return
		}
	}

	// Generate a unique ID for the course and convert it to a string
	id, _ := rand.Int(rand.Reader, big.NewInt(100))
	course.CourseID = strconv.Itoa(int(id.Int64()))

	// Append the new course to the courses slice (mock database)
	courses = append(courses, course)

	// Send a success response with the added course data
	json.NewEncoder(w).Encode(course)
}

// updateCourse handles PATCH requests to update an existing course.
// It finds the course by ID, deletes the old version, and appends the updated one.
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course")
	w.Header().Set("Content-Type", "application/json")

	// Retrieve the ID from the URL path parameters using mux
	params := mux.Vars(r)
	id := params["id"]

	// Loop through the courses slice to find the course with the matching ID
	for index, course := range courses {
		if course.CourseID == id {
			// Remove the existing course
			courses = append(courses[:index], courses[index+1:]...)

			// Decode the updated course data from the request body
			if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
				json.NewEncoder(w).Encode("Error decoding JSON")
				return
			}

			// Set the ID of the new course to match the old one
			course.CourseID = id
			courses = append(courses, course) // Add the updated course
			json.NewEncoder(w).Encode(course) // Send the updated course back as JSON
			return
		}
	}
	// Send a message if the course is not found
	json.NewEncoder(w).Encode("Course not found")
}

// deleteCourse handles DELETE requests to remove a course by ID.
// It finds and removes the course from the courses slice.
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course")
	w.Header().Set("Content-Type", "application/json")

	// Retrieve the ID from the URL path parameters using mux
	params := mux.Vars(r)
	id := params["id"]

	// Loop through the courses slice to find the course with the matching ID
	for index, course := range courses {
		if course.CourseID == id {
			// Remove the course from the slice
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted successfully") // Send a success message
			return
		}
	}
	// Send a message if the course is not found
	json.NewEncoder(w).Encode("Course not found")
}

// deleteAllCourses handles DELETE requests to remove all courses.
func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")
	w.Header().Set("Content-Type", "application/json")

	// Clear the courses slice
	courses = []Course{}

	// Send a message confirming deletion of all courses
	json.NewEncoder(w).Encode("All courses deleted successfully")
}
