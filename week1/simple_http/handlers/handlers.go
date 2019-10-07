package handlers

import (
	"encoding/json"
	"net/http"

	db "github.com/thai2902/mono-go/week1/model"
)

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var students = []db.Student{
		db.Student{FirstName: "Bessie", LastName: "Von", ClassName: "Golang191001", AcademyName: "Golang backend", Email: "Bessie.Ferry6@gmail.com", Age: 25},
		db.Student{FirstName: "Marisol", LastName: "Kirlin", ClassName: "Golang191001", AcademyName: "Golang backend", Email: "Marisol_Pfannerstill@hotmail.com", Age: 38},
		db.Student{FirstName: "Florencio", LastName: "Jacobs", ClassName: "Golang191001", AcademyName: "Golang backend", Email: "Florencio.Auer90@hotmail.com", Age: 32},
		db.Student{FirstName: "Rodrick", LastName: "Hills", ClassName: "Golang191001", AcademyName: "Golang backend", Email: "Rodrick11@yahoo.com", Age: 35},
		db.Student{FirstName: "Timmy", LastName: "Stanton", ClassName: "Golang191001", AcademyName: "Golang backend", Email: "Timmy_Cormier57@gmail.com", Age: 26},
		db.Student{FirstName: "Candice", LastName: "Schmeler", ClassName: "Golang191001", AcademyName: "Golang backend", Email: "Candice71@hotmail.com", Age: 30},
		db.Student{FirstName: "Desiree", LastName: "Kihn", ClassName: "Golang191001", AcademyName: "Golang backend", Email: "Desiree_Thompson@yahoo.com", Age: 40},
	}
	jsonStudents, _ := json.Marshal(students)
	w.Header().Set("Content-Type", "apllication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonStudents))
}
