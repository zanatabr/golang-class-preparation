package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// (*) Como não há uma referência direta ao driver no arquivo .go atual,
//     é feito um “import implícito” com o uso do caractere
// 	   underline ( _ ) :: ( _ "github.com/go-sql-driver/mysql" ).
//
// 	   Quem vai efetivamente usar esse import é o pacote “database/sql”,
//     quando a instrução sql.Open("mysql", stringConexao) for executada.
//
// (*) “defer” é semelhante ao “finally” do Java

func main() {

	// formato da String de Conexão: "usuario:senha@/banco"
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Dentro do sql.Open")
		log.Fatal(erro)
		// só vai cair aqui se ocorrer um erro inesperado na conexão
		// mas não um erro de autenticação, por exemplo.
		// (*) Fazer um teste alterando para uma senha inválida na
		//     string de conexão.
	}
	defer db.Close()

	// db.Ping para verificar se a conexão está realmente aberta
	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do db.Ping")
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
