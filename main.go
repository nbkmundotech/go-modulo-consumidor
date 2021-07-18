package main

// go-modulo-consumidor depende de go-modulo-biblioteca
import (
	"fmt"

	"github.com/nbkmundotech/go-modulo-biblioteca/analisador"
	"github.com/nbkmundotech/go-modulo-biblioteca/consolidar"
)

func main() {
	var queryString = "?busca=termo&ordem=desc"
	var queryMap map[string]string = analisador.Analisar(queryString)
	fmt.Println(queryMap)

	var queryStringRevertida string = consolidar.MapaParaString(queryMap)
	fmt.Println(queryStringRevertida)
}
