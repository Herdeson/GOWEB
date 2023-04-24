package routes

import (
	"GOWEB/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/inserir", controllers.Insere)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/update", controllers.Update)

}
