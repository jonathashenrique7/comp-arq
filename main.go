package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
)

func Secret(user, realm string) string {
	if user == "Informe o nome do usuário" {
		return "$Aqui você informa o hash md5" // Gerar criptografia MD5 Salt => https://unix4lyfe.org/crypt/
	} // Quando você logar, você utilizará a senha normal => Lembre-se que a criptografia acima é somente uma proteção
	return ""
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <diretório> <porta lógica>")
		os.Exit(1)
	}

	httpDir := os.Args[1] // [0] => o arquivo a ser compartilhado
	porta := os.Args[2]   // [1] e [2] => São argumentos que serão informados pelo usuário

	authenticator := auth.NewBasicAuthenticator("meuserver.com", Secret)
	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir(httpDir)).ServeHTTP(w, &r.Request)
	})) // Associa a url

	fmt.Printf("Servidor local na porta => %s", porta)
	log.Fatal(http.ListenAndServe(":"+porta, nil)) // Fazendo a referência do diretório para o "http"
}

// 1 ) Importante => Repare que na função "Dir()" eu informei um diretório de nome "arquivo-server"

// 2 ) Nesta pasta irei lançar todos os arquivos que serão disponibilizados no servidor

// 3 ) => Você irá digitar : go run main.go /nome/caminho/arquivo número_porta_lógica
