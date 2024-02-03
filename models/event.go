package models

import (
	"fmt"

	"book-exchange.com/rest/db"
)

type Event struct {
	EventID           int64    
	NameOfBook        string    `binding:"required"`
	Description 	  string    `binding:"required"`
	//Request    Request    `binding:"required"`
	//DateTime    time.Time `binding:"required"`
	UserID      	  int64 
}
var events = []Event{}
func (e Event) Save() error{
	fmt.Printf("book:%s",e.NameOfBook)
	query := `
	INSERT INTO events(nameofbook, description, user_id)
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.NameOfBook, e.Description, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.EventID = id
	return err
} 

func GetEvents(id int64) ([]Event, error){
	//get all post where id is not equal to id as user is getting feed so we won't show user his feed
	for index, value := range events {
        fmt.Printf("Index: %d, Value: %s\n", index, value.NameOfBook)
    }
	return events, nil
}

func GetAllEvents() ([]Event, error){
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.EventID, &event.NameOfBook, &event.Description, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}