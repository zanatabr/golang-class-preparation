package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Nick  string `json:"nick`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

// CriarUsuario: insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	// Lê o corpo da requisição
	// o pacote "ioutil" é usado para ler o conteúdo (body)
	// da requisição que é um JSON.
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return // Não retorna nada (void), mas interrompe a execução
	}

	// Trata o conteúdo do request (valor enviado no Body)
	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	// Conecta ao BD
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!"))
		return
	}
	defer db.Close()

	// Faz a preparação do Statement (PreparedStatement)
	statement, erro := db.Prepare("insert into usuarios (nome, email, nick, senha) values (?, ?, ?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	// Executa a instrução no BD
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email,
		usuario.Nick, usuario.Senha)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	// Obtém o último ID inserido
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	// STATUS CODE
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))

}

// BuscarUsuarios: traz todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	// Conecta ao BD
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	// Executa a Query
	linhas, erro := db.Query("select id, nome, nick, email from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer linhas.Close()

	// Iterar sobre as linhas retornadas no BD
	// e preencher o slice de usuarios
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		// .Scan
		// Importante: Tem que ser na ordem em que as colunas aparecem
		//             durante a consulta (ordem em que elas foram criadas
		//             no BD ou organizadas na cláusula SQL)
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick,
			&usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	// Status Code
	w.WriteHeader(http.StatusOK)

	// Converte para a saída usando o Encode
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

// BuscarUsuario: traz um usuário específico salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	// Lê todos os parâmetro que estão sendo passado na rota da requisição
	parametros := mux.Vars(r)

	// Trata o parâmetro ID
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	// Conecta ao BD
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}

	linha, erro := db.Query("select id, nome, nick, email from usuarios where id=?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuário!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick,
			&usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}
}

// AtualizarUsuario: altera os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {

	// Lê todos os parâmetro que estão sendo passado na rota da requisição
	parametros := mux.Vars(r)

	// Trata o parâmetro ID
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	// Conecta ao BD
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	// Faz a preparação do Statement (PreparedStatement)
	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}

	// Status Code
	w.WriteHeader(http.StatusNoContent)

}

// DeletarUsuario remove um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	// Lê todos os parâmetro que estão sendo passado na rota da requisição
	parametros := mux.Vars(r)

	// Trata o parâmetro ID
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	// Conecta ao BD
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	// Faz a preparação do Statement (PreparedStatement)
	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuário!"))
		return
	}

	// Status Code
	w.WriteHeader(http.StatusNoContent)
}
