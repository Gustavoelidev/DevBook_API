package controllers

import "net/http"

// CriarUsuarios insere um usuario no banco de dados.

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario!"))
}

// BuscarUsuarios busca todos os usuarios no banco de dados.
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios!"))
}

// BuscarUsuario busca um usuario no banco de dados.
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando o usuario!"))
}

// AtualizarUsuarios atualiza um usuario no banco de dados.
func AtualizarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario!"))
}

// DeletarUsuarios deleta um usuario no banco de dados.
func DeletarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario!"))
}
