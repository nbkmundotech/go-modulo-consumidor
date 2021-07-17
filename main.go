package main

// go-modulo-consumidor depende de go-modulo-biblioteca
import (
	"fmt"

	"github.com/nbkmundotech/go-modulo-biblioteca/analisador"
)

func main() {
	var queryString = "?busca=termo&ordem=desc"
	var queryMap map[string]string = analisador.Analisar(queryString)
	fmt.Println(queryMap)
}
