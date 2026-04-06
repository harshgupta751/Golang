package controller

import (
	models "MongoAPI/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString="mongodb+srv://harshachieve300:Harhgpt7@cluster0.8uvrt2i.mongodb.net/"

const dbName="netflix"
const collectionName="watchlist"

var collection *mongo.Collection

func init(){

	clientOptions := options.Client().ApplyURI(connectionString)

client, err:=	mongo.Connect(clientOptions)

if err!=nil {
	log.Fatal(err)
}

fmt.Println("MongoDB connection established")

collection = client.Database(dbName).Collection(collectionName)

fmt.Println("Collection instance is Ready")




}

func insertOneMovie(movie models.NetFlix){

	result, err := collection.InsertOne(context.Background(), movie)

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in DB with id :", result.InsertedID)

}

func getAllMovies() []models.NetFlix {
 cursor , err := collection.Find(context.Background(), bson.M{})

 if err!=nil {
	log.Fatal(err)
 }

 var movies []models.NetFlix
 for cursor.Next(context.Background()) {
	var movie models.NetFlix
	err := cursor.Decode(&movie)
	if err!=nil {
		log.Fatal(err)
	}
	movies = append(movies, movie)

 }

 defer cursor.Close(context.Background())
 fmt.Println(movies)

 return movies

}


func updateOneMovie(movieID string){
id, _:= bson.ObjectIDFromHex(movieID)

filter := bson.M{"_id": id}
update := bson.M{
	"$set" : bson.M{
			"watched" : true,
	},
}

result, err := collection.UpdateOne(context.Background(), filter, update)

if err!=nil {
	log.Fatal(err)
}

fmt.Println("Modified count is : ", result.ModifiedCount)

}

func deleteOneMovie(movieId string){

	id, _:=bson.ObjectIDFromHex(movieId)

filter := bson.M{
	"_id" : id,
}

	result, err:= collection.DeleteOne(context.Background(), filter)

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count is :", result.DeletedCount)


}


func deleteAllMovies(){

result, err:= collection.DeleteMany(context.Background(), bson.M{})

if err != nil {
	log.Fatal(err)
}

fmt.Println("Deleted Count is :", result.DeletedCount)


}



func InsertOneMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/json")
w.Header().Set("Access-Control-Allow-Methods", "POST")

var movie models.NetFlix
json.NewDecoder(r.Body).Decode(&movie)

insertOneMovie(movie)

json.NewEncoder(w).Encode("Movie inserted successfully")

}

func GetAllMovies(w http.ResponseWriter, r *http.Request){
w.Header().Set("Access-Control-Allow-Methods", "GET")

movies := getAllMovies()

json.NewEncoder(w).Encode(movies)

}


func UpdateOneMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/x-www-form-urlencode")
w.Header().Set("Access-Control-Allow-Methods", "PUT")

params:= mux.Vars(r)

id:= params["id"]

updateOneMovie(id)

json.NewEncoder(w).Encode("Movie watch status updated successfully")

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type", "application/x-www-form-urlencode")
w.Header().Set("Access-Control-Allow-Methods", "DELETE")

params:= mux.Vars(r)

id:= params["id"]

deleteOneMovie(id)

json.NewEncoder(w).Encode("Movie deleted successfully")


}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
w.Header().Set("Access-Control-Allow-Methods", "DELETE")

deleteAllMovies()

json.NewEncoder(w).Encode("All movies deleted successfully")

}




