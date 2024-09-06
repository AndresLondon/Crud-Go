package services_user

import (
	conection "main/db"
	templates_sistem "main/resources"
	"net/http"
)

type Users struct {
	Id     int
	Nombre string
	Correo string
}

func Index(w http.ResponseWriter, r *http.Request) {
	registros, err := conection.DbGlobalUltra2.Query(`select * from users order by id`)
	if err != nil {
		panic(err.Error())
	}

	user := Users{}
	arregloUser := []Users{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Nombre = nombre
		user.Correo = correo

		arregloUser = append(arregloUser, user)
	}

	err = templates_sistem.Plantillas.ExecuteTemplate(w, "index_user", arregloUser)
	if err != nil {
		panic(err.Error())
	}

}

func Create(w http.ResponseWriter, r *http.Request) {
	templates_sistem.Plantillas.ExecuteTemplate(w, "new_user", nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name_user")
		email := r.FormValue("email_user")

		insertRegistros, err := conection.DbGlobalUltra2.Prepare(`INSERT INTO users (name, email) VALUES ($1, $2)`)

		if err != nil {
			panic(err.Error())
		}

		defer insertRegistros.Close()

		_, err = insertRegistros.Exec(name, email)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id_user := r.URL.Query().Get("id")

	delet_user, err := conection.DbGlobalUltra2.Prepare(`DELETE FROM users WHERE id = $1`)

	if err != nil {
		panic(err.Error())
	}

	defer delet_user.Close()

	_, err = delet_user.Exec(id_user)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id_user := r.URL.Query().Get("id")

	registros, err := conection.DbGlobalUltra2.Query(`select * from users where id = $1`, id_user)

	if err != nil {
		panic(err.Error())
	}

	user := Users{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		user.Id = id
		user.Nombre = nombre
		user.Correo = correo
	}

	templates_sistem.Plantillas.ExecuteTemplate(w, "edit_user", user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id_user")
		name := r.FormValue("name_user")
		email := r.FormValue("email_user")

		updateUser, err := conection.DbGlobalUltra2.Prepare(`UPDATE users SET name = $1, email = $2 where id = $3`)

		if err != nil {
			panic(err.Error())
		}

		defer updateUser.Close()

		_, err = updateUser.Exec(name, email, id)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/", 301)
	}
}
