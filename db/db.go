package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *sql.DB
var DB2 *mongo.Client


func InitDB2(){
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err!=nil{
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}
func InitDB(){
	DB2, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	fmt.Println(DB2)
	InitDB2()
	// fmt.Println(client)
	//coll := client.Database("db").Collection("books")
	//uid := "Peeyuwsh"
	//id := utils.UniqueId(uid)
	//fmt.Printf("id:%s \t again:%s\n",id,utils.UniqueId(uid))
	//doc := models.Event{NameOfBook : id, Description: " A test book1"}
	//doc := Book{Title: "Atonement", Author: "Ian McEwan"}
	//result, err := coll.InsertOne(context.TODO(), doc)
	//fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	//filter := bson.D{{"nameofbook", "740286226"}}
	// cursor, err := coll.Find(context.TODO(), bson.D{})
	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// defer cursor.Close(context.TODO())

	// // Decode results into a slice of Tea
	// var results []models.Event
	// if err = cursor.All(context.TODO(), &results); err != nil {
	// 	fmt.Print(err.Error())
	// }

	// // Print JSON representation of each result
	// for _, result := range results {
	// 	res, _ := bson.MarshalExtJSON(result, false, false)
	// 	fmt.Println(string(res))
	// }
}

func createTables() {
	createUserTable :=`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)
	if err !=nil{
		panic("Could not create users table.")
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		eventid INTEGER PRIMARY KEY AUTOINCREMENT,
		nameofbook TEXT NOT NULL,
		description TEXT NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`

	_, err = DB.Exec(createEventsTable)

	if err!=nil{
		fmt.Println("ERROR:",err)
		panic("Could not create events table.")
	}
}