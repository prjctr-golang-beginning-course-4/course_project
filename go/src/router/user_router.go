package router

import (
	"github.com/gorilla/mux"
)

func SetRoutes(r *mux.Router) {
	//apiRoute := r.PathPrefix("/api").Subrouter()
	//apiRoute.HandleFunc("/register", controller.RegisterUser).Methods("POST")
	//apiRoute.HandleFunc("/login", controller.LoginUser).Methods("POST")
	//
	//adminRoute := apiRoute.PathPrefix("/admin").Subrouter()
	//adminRoute.Use(middleware.AdminMiddleware)
	//adminRoute.HandleFunc("/user", controller.GetAllUser).Methods("GET")
	//adminRoute.HandleFunc("/user/{id}", controller.GetAllUser).Methods("GET")
	//adminRoute.HandleFunc("/user/{id}/profile", controller.GetUserProfile).Methods("GET")
	//
	//userRoute := apiRoute.PathPrefix("/user").Subrouter()
	//userRoute.Use(middleware.UserMiddleware)
	//userRoute.HandleFunc("/profile", controller.GetAllUser).Methods("GET")
	//userRoute.HandleFunc("/profile", controller.GetAllUser).Methods("POST")
	//userRoute.HandleFunc("/reset_password", controller.GetAllUser).Methods("GET")
}
