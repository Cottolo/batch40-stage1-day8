package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	// route := mux.NewRouter()
	// route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// }).Methods("GET")

	route := mux.NewRouter()

	// path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	// routing
	route.HandleFunc("/home", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/detail/{id}", detailProject).Methods("GET")
	route.HandleFunc("/project", formAddProject).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")
	route.HandleFunc("/delete-project/{id}", deleteProject).Methods("GET")
	route.HandleFunc("/edite-project/{id}", editeProject).Methods("GET")

	fmt.Println("server running at localhost:5000")
	http.ListenAndServe("localhost:5000", route)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/index.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	response := map[string]interface{}{

		"Projects": dataProject,
	}

	tmpl.Execute(w, response)
}

func formAddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/project.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	tmpl.Execute(w, nil)
}

// Type Data
type Project struct {
	ProjectName        string
	StartDate          string
	EndDate            string
	Duration           float64
	ProjectDescription string
	NodeJs             string
	NextJs             string
	ReactJs            string
	TypeScript         string
}

// ARRAY
var dataProject = []Project{}

func addProject(w http.ResponseWriter, r *http.Request) {
	error := r.ParseForm()
	if error != nil {
		log.Fatal(error)
	}

	var projectName = r.PostForm.Get("project-name")
	var startDate = r.PostForm.Get("start-date")
	var endDate = r.PostForm.Get("end-date")
	var projectDescription = r.PostForm.Get("project-description")
	var nodeJs = r.PostForm.Get("node-js")
	var nextJs = r.PostForm.Get("next-js")
	var reactJs = r.PostForm.Get("react-js")
	var typeScript = r.PostForm.Get("typescript")
	var layout = "2006-01-02"
	var start, _ = time.Parse(layout, startDate)
	var end, _ = time.Parse(layout, endDate)
	var duration = math.Round(end.Sub(start).Hours() / 24 / 30)

	// fmt.Println("Project Name :" + r.PostForm.Get("project-name"))
	// fmt.Println("Start Date :" + r.PostForm.Get("start-date"))
	// fmt.Println("End Date : " + r.PostForm.Get("end-date"))
	// fmt.Println("Project Description :" + r.PostForm.Get("project-description"))
	// fmt.Println("Technology : " + r.PostForm.Get("node-js"))
	// fmt.Println("Technology : " + r.PostForm.Get("next-js"))
	// fmt.Println("Technology : " + r.PostForm.Get("react-js"))
	// fmt.Println("Technology : " + r.PostForm.Get("typescript"))
	// fmt.Println(start)
	// fmt.Println(end)
	// fmt.Println(duration)

	//OBJECT
	var newProject = Project{
		ProjectName:        projectName,
		ProjectDescription: projectDescription,
		StartDate:          startDate,
		EndDate:            endDate,
		NextJs:             nextJs,
		ReactJs:            reactJs,
		NodeJs:             nodeJs,
		TypeScript:         typeScript,
		Duration:           duration,
	}

	//PUSH
	dataProject = append(dataProject, newProject)

	// fmt.Println(dataProject)

	//HALAMAN SETELAH POST
	http.Redirect(w, r, "/home", http.StatusMovedPermanently)

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/contact.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	tmpl.Execute(w, nil)
}

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/detail.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	var ProjectDetail = Project{}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// fmt.Println(id)

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				ProjectName:        data.ProjectName,
				ProjectDescription: data.ProjectDescription,
				StartDate:          data.StartDate,
				EndDate:            data.EndDate,
				NextJs:             data.NextJs,
				ReactJs:            data.ReactJs,
				NodeJs:             data.NodeJs,
				TypeScript:         data.TypeScript,
				Duration:           data.Duration,
			}
		}
	}

	// OBJECT
	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	// fmt.Println(data)

	tmpl.Execute(w, data)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// fmt.Println(id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	http.Redirect(w, r, "/home", http.StatusFound)

}

func editeProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("views/project.html")

	if error != nil {
		w.Write([]byte("not found 404"))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	fmt.Println(id)

	tmpl.Execute(w, nil)
}
