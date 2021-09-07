package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/cmllmd/bookreview/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsComentarios := models.BuscaTodosOsComentarios()
	temp.ExecuteTemplate(w, "Index", todosOsComentarios)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		comentario := r.FormValue("comentario")
		nota := r.FormValue("nota")

		notaConvertida, _ := strconv.Atoi(nota)

		models.CriaNovoComentario(nome, comentario, notaConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoComentario := r.URL.Query().Get("id")
	models.DeletaComentario(idDoComentario)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoComentario := r.URL.Query().Get("id")
	comentario := models.EditaComentario(idDoComentario)
	temp.ExecuteTemplate(w, "Edit", comentario)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		comentario := r.FormValue("comentario")
		nota := r.FormValue("nota")

		idConvertido, _ := strconv.Atoi(id)

		notaConvertida, _ := strconv.Atoi(nota)

		models.AtualizaComentario(idConvertido, nome, comentario, notaConvertida)
	}
	http.Redirect(w, r, "/", 301)
}
