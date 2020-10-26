package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

//var tmpl = template.Must(template.New("MyTemplate").ParseFiles("static/tmpl.html"))

// Set a Decoder instance as a package global, because it caches
// meta-data about structs, and an instance can be shared safely.
var (
	decoder = schema.NewDecoder()
	postID = 3
)


func getPostId() int{
	postID++
	return postID
}

type Post struct {
	Id        int
	Title     string `schema:"title,title2example"`
	Text      string `schema:"text"`
	Author    string `schema:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var posts = map[int]*Post{
	1: &(Post{
		Id:        1,
		Title:     "some text1" ,
		Text:      "some text1" ,
		Author:    "some author1" ,
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	2: &(Post{
		Id:        2,
		Title:     "some text2",
		Text:      "some text2",
		Author:    "some author2",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
	3: &(Post{
		Id:        3,
		Title:     "some text3",
		Text:      "some text3",
		Author:    "some author3",
		CreatedAt: time.Now().Add(-time.Hour),
	}),
}

func main() {
	fmt.Println("as")
	router := mux.NewRouter()
	//Шаблон со списком всех постов / короткие без Text
	router.HandleFunc("/", listPostHandler).Methods("GET")

	//Exapmple! delete
	router.HandleFunc("/list", viewList).Methods("GET")

	//Шаблон с текстовыми полями для задания  Title Text Author
	router.HandleFunc("/create", createPostHandlerGet).Methods("GET")
	router.HandleFunc("/create", createPostHandlerPost).Methods("POST")

	//Шаблон со страницей одного поста / полгого с отображением Text
	router.HandleFunc("/{id}", getPostHandlerID).Methods("GET")

	//Шаблон с текстовыми полями для обновления Title Text Author
	//router.HandleFunc("/{id}", updatePostHandleID).Methods("PUT")
	//r.HandleFunc("/articles", ArticlesHandler)
	//http.Handle("/", r)

	port := ":9096"
	fmt.Printf("Start server : port = %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// TaskList - список задач
type TaskList struct {
	Name        string
	Description string
	List        []Task
}

// Task - задача и ее статус
type Task struct {
	ID       string
	Text     string
	Complete bool
}

var simpleList = TaskList{
	Name:        "Название листа",
	Description: "Описание листа с задачами",
	List: []Task{
		Task{"first", "Первая задача", false},
		Task{"second", "Вторая задача", false},
		Task{"thrid", "Третья задача", true},
	},
}

func viewList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("Ttt").ParseFiles("./static/tmpl.html"))

	if err := tmpl.ExecuteTemplate(w, "list", simpleList); err != nil {
		log.Println(err)
	}
}

//GET
//2. роут и шаблон для просмотра конкретного поста в блоге.
func getPostHandlerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIDRaw, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	postID, err := strconv.Atoi(postIDRaw)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	post, ok := posts[postID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.New("tmpl_1page").ParseFiles("./static/tmpl_1page.html"))

	if err := tmpl.ExecuteTemplate(w, "list", simpleList); err != nil {
		log.Println(err)
	}

	if err := tmpl.ExecuteTemplate(w, "T", post); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

}

//Get
//3. Создайте роут и шаблон для создания материала
func createPostHandlerGet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("tmpl88").ParseFiles("./static/tmpl_create.html"))
	if err := tmpl.ExecuteTemplate(w, "CreatePosts", struct{ Success bool }{true}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

//POST
//3. Создайте роут и шаблон для создания материала
func createPostHandlerPost(w http.ResponseWriter, r *http.Request) {
	err:=r.ParseForm()
	if err != nil {
		log.Println(err)// Handle error
    }

	newID:=getPostId()
	post:= Post{ 
		Id:       newID,
	}

    // r.PostForm is a map of our POST form values
    err = decoder.Decode(&post, r.PostForm)
    if err != nil {
        log.Println(err)// Handle error
    }
	fmt.Println(post)
	posts[newID] = &post
	

	fmt.Fprintln(w, r.Form)
}

//PUT
//3. Создайте роут и шаблон для редактированияматериала.
func updatePostHandleID(w http.ResponseWriter, r *http.Request) {

}

//GET
//1. роут и шаблон для отображения всех постов в блоге.
func listPostHandler(w http.ResponseWriter, r *http.Request) {

	for key, value := range posts {
		fmt.Println("Key:", key, "Value:", value)
	}

	tmpl := template.Must(template.New("tmpl_1page").ParseFiles("./static/tmpl_allPosts.html"))

	if err := tmpl.ExecuteTemplate(w, "AllPosts", posts); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
