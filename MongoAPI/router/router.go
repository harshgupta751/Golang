package router 

import (
	"github.com/gorilla/mux"
	controller "MongoAPI/controller"

)

func Router() *mux.Router{

router:= mux.NewRouter()

router.HandleFunc("/v1/addmovie", controller.InsertOneMovie).Methods("POST")
router.HandleFunc("/v1/getmovies", controller.GetAllMovies).Methods("GET")
router.HandleFunc("/v1/markwatched/{id}", controller.MarkWatched).Methods("PUT")
router.HandleFunc("/v1/deletemovie/{id}", controller.DeleteOneMovie).Methods("DELETE")
router.HandleFunc("/v1/deleteallmovies", controller.DeleteAllMovies).Methods("DELETE")



return router

}

