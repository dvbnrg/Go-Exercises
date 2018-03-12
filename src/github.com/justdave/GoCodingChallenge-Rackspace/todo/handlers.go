package todo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

// Create will allow a user to create a new todo
// The supported body is {"title": "", "status": ""}
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//Open DB connection and catch errors
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Create TODO
	var todo CreateTodo

	json.NewDecoder(r.Body).Decode(&todo)

	//Log bad Requests
	if todo.Status == "" || todo.Title == "" {
		http.Error(w, "Todo request is missing status or title", http.StatusBadRequest)
	}

	//Catch invalid status's
	invalidStatus := true
	for _, status := range allowedStatuses {
		if todo.Status == status {
			invalidStatus = false
			break
		}
	}

	if !invalidStatus {
		http.Error(w, "The provided status is not supported", http.StatusBadRequest)
	}

	//Insert statement
	insertStmt := fmt.Sprintf(`INSERT INTO todo (title, status) VALUES ('%s', '%s') RETURNING id`, todo.Title, todo.Status)

	//create new ID for new TODO
	var todoID int

	// Insert and get back newly created todo ID
	if err := db.QueryRow(insertStmt).Scan(&todoID); err != nil {
		fmt.Printf("Failed to save to db: %s", err.Error())
	}

	//sysout newly created ID
	fmt.Printf("Todo Created -- ID: %d\n", todoID)

	//create new todo object
	newTodo := Todo{}
	db.QueryRow("SELECT id, title, status FROM todo WHERE id=$1", todoID).Scan(&newTodo.ID, &newTodo.Title, &newTodo.Status)
	defer db.Close()

	//marshall todo into json
	jsonResp, _ := json.Marshal(newTodo)

	//ship json into fprint
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, string(jsonResp))
}

// List will provide a list of all current to-dos
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	//get environment variables
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//open DB connection
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	//create todo array for all listings in the db
	todoList := []Todo{}

	//query db then close db connection
	rows, err := db.Query("SELECT id, title, status FROM todo")
	defer rows.Close()

	//parse rows returned from query into new todo object then append into list
	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Failed to build todo list")
		}

		todoList = append(todoList, todo)
	}

	//format to and return Json
	jsonResp, _ := json.Marshal(Todos{TodoList: todoList})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	//ship json
	fmt.Fprintf(w, string(jsonResp))
}

//Update returns the new state of updated property
func Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//get environment variables
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	//open DB connection
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	//Get id from incoming URl with gin

	//create TODO
	todo := Todo{}

	//Decode incoming Json
	json.NewDecoder(r.Body).Decode(&todo)

	//Update statement then close DB connection && Check if ID exists else toss out bad ID's
	db.QueryRow("UPDATE todo SET title=$2, status=$3 WHERE id=$1", todo.ID)
	db.QueryRow("SELECT id, title, status FROM todo WHERE id=$1", todo.ID).Scan(&todo.ID, &todo.Title, &todo.Status)
	defer db.Close()

	//format to and return Json
	jsonResp, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	//ship json
	fmt.Fprintf(w, string(jsonResp))
}
