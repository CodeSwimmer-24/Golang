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
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PATCH")

	// Starting the server on port 4000
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers

// homePage handles requests to the home route, responding with a welcome message.
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Gherdekho.com</h1>")) // Send a simple HTML response
}

// createCourse handles POST requests to create a new course entry.
// It decodes JSON data, checks for validity, generates a unique ID, and adds the course to the database.
func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create new course")
	w.Header().Set("Content-Type", "application/json")

	// Parse form data from request
	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode("Error parsing form data")
		return
	}

	// Get form data values
	courseName := r.FormValue("courseName")
	coursePriceStr := r.FormValue("coursePrice")
	authorName := r.FormValue("authorName")
	authorWebsite := r.FormValue("authorWebsite")

	// Convert coursePrice to integer
	coursePrice, err := strconv.Atoi(coursePriceStr)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid course price")
		return
	}

	// Create a new course struct
	course := Course{
		CourseName:  courseName,
		CoursePrice: coursePrice,
		Author: &Author{
			FullName: authorName,
			Website:  authorWebsite,
		},
	}

	// Check if the course is empty
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data in form")
		return
	}

	// Check for duplicate course name
	for _, value := range courses {
		if course.CourseName == value.CourseName {
			json.NewEncoder(w).Encode("Course name already exists")
			return
		}
	}

	// Generate a unique ID for the course
	id, _ := rand.Int(rand.Reader, big.NewInt(100))
	course.CourseID = strconv.Itoa(int(id.Int64()))

	// Append the new course to the courses slice
	courses = append(courses, course)

	// Send the created course as a response
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
