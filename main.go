package main

// go-modulo-consumidor depende de go-modulo-biblioteca
import (
	"fmt"

	analisador "github.com/nbkmundotech/go-modulo-biblioteca"
)

func main() {
	var queryString = "?busca=termo&ordem=desc"
	var queryMap map[string]string = analisador.Analisar(queryString)
	fmt.Println(queryMap)
}
