package src

import (
	"course/src/controller"
	"course/src/middleware"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net/http"
)

func Run(db *gorm.DB, rdb *redis.Client) {

	r := mux.NewRouter()

	userController := controller.InitUserController(db, rdb)
	loginController := controller.InitLoginController(db, rdb)
	registerController := controller.InitRegisterController(db, rdb)

	loginRouter := r.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", loginController.Login).Methods("Post")

	registerRoute := r.PathPrefix("/register").Subrouter()
	registerRoute.HandleFunc("", registerController.Register).Methods("Post")

	adminRouter := r.PathPrefix("/admin").Subrouter()

	adminRouter.Use(middleware.InitAdminMiddleware(userController).AdminMiddleware)
	adminRouter.HandleFunc("/user", userController.CreateUser).Methods("POST")
	adminRouter.HandleFunc("/user", userController.GetAllUsers).Methods("GET")
	adminRouter.HandleFunc("/user/{userId}", userController.GetUser).Methods("GET")
	adminRouter.HandleFunc("/user/{userId}", userController.DeleteUser).Methods("DELETE")
	adminRouter.HandleFunc("/user/{userId}/assign/admin", userController.AssignAdmin).Methods("POST")

	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.Use(middleware.InitUserMiddleware(userController).UserMiddleware)
	userRouter.HandleFunc("/profile", userController.Profile).Methods("GET")

	http.Handle("/", r)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
