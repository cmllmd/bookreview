package models

import "github.com/cmllmd/bookreview/db"

type Comentario struct {
	Id         int
	Nome       string
	Comentario string
	Nota       int
}

func BuscaTodosOsComentarios() []Comentario {
	db := db.ConectaComBD()

	selectDeTodosOsComentarios, err := db.Query("select * from comentarios")
	if err != nil {
		panic(err.Error())
	}
	c := Comentario{}
	comentarios := []Comentario{}

	for selectDeTodosOsComentarios.Next() {
		var id, nota int
		var nome, comentario string

		err = selectDeTodosOsComentarios.Scan(&id, &nome, &comentario, &nota)
		if err != nil {
			panic(err.Error())
		}
		c.Id = id
		c.Nome = nome
		c.Comentario = comentario
		c.Nota = nota

		comentarios = append(comentarios, c)
	}
	defer db.Close()
	return comentarios
}

func CriaNovoComentario(nome, comentario string, nota int) {
	db := db.ConectaComBD()

	insereDadosnoBd, err := db.Prepare("insert into comentarios(nome, comentario, nota) values ($1, $2,$3)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosnoBd.Exec(nome, comentario, nota)
	defer db.Close()
}

func DeletaComentario(id string) {
	db := db.ConectaComBD()

	deletarComentario, _ := db.Prepare("delete from comentarios where id=$1")

	deletarComentario.Exec(id)
	defer db.Close()
}

func EditaComentario(id string) Comentario {
	db := db.ConectaComBD()
	comentarioDoBanco, err := db.Query("select * from comentarios where id=$1", id)

	comentarioParaAtualizar := Comentario{}

	for comentarioDoBanco.Next() {
		var id, nota int
		var nome, comentario string

		err = comentarioDoBanco.Scan(&id, &nome, &comentario, &nota)
		if err != nil {
			panic(err.Error())
		}
		comentarioParaAtualizar.Id = id
		comentarioParaAtualizar.Nome = nome
		comentarioParaAtualizar.Comentario = comentario
		comentarioParaAtualizar.Nota = nota
	}
	defer db.Close()
	return comentarioParaAtualizar
}

func AtualizaComentario(id int, nome, comentario string, nota int) {
	db := db.ConectaComBD()

	AtualizaComentario, err := db.Prepare("update comentarios set nome=$1, comentario=$2, nota=$3, where id=$4")
	if err != nil {
		panic(err.Error())
	}

	AtualizaComentario.Exec(nome, comentario, nota, id)
	defer db.Close()
}
