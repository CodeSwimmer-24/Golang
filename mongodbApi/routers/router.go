package routers

import (
	"net/http"

	"github.com/CodeSwimmer-24/mongodbApi/controllers"
)

// RegisterRoutes registers all the routes for the API
func RegisterRoutes() {
	http.HandleFunc("/api/postData", controllers.HandlePostData) // POST /api/postData
}
