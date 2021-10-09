package insta

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"os"
)

type Person struct {
	Name    string  `json:"name"`
	ID   string  `json:"id"`
	Password  string  `json:"password"`
	Email string
}

type Post struct {
	ID string 
	Caption  string
	time Time 
}


func newuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	for _, item := range Person {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func newPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Post Person
	_ = json.NewDecoder(r.Body).Decode(&Post)
	Person.ID = strconv.Itoa(rand.Intn(100000000))
	Post = append(Post, Person)
	json.NewEncoder(w).Encode(Person)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Post {
		if item.ID == params["id"] {
			Post = append(Post[:index], Post[index+1:]...)
			var post Person
			_ = json.NewDecoder(r.Body).Decode(&Person)
			Person.ID = params["id"]
			post = append(post, Person)
			json.NewEncoder(w).Encode(Person)
			return
		}
	}
}

func getAllPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Post {
		if item.ID == params["id"] {
			Post = append(Post[:index], Post[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Post)
}

func main() {
	r := mux.NewRouter()

	
	Person = append(person, Person{Name: "Yash", ID: "438227", Email: "yash786@gmail.com", Post: &Post{ID: "8954", Caption: "Happy Navratri"}})
	Person = append(person, Person{Name: "Rahul", ID: "454555", Email: "rahulofficial@gmail.com", Post: &Post{ID: "3794", Caption: "Good Morning"}})

	// Route handles & endpoints
	r.HandleFunc("/users", newUser).Methods("GET")
	r.HandleFunc("/users/{id}", getuser).Methods("GET")
	r.HandleFunc("/posts", newpost).Methods("POST")
	r.HandleFunc("/post/{id}", getpost).Methods("GET")
	r.HandleFunc("/post/users/{id}", getAllPost).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

