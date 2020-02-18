package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	//exibeNomes()
	exibeIntroducao()

	for {
		exibeMenu()

		//Go não utiliza os parenteses no IF
		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo do programa")
		// } else {
		// 	fmt.Println("Não conheço este comando!")
		// }

		// ValorBoleano := false
		// if ValorBoleano {
		// 	fmt.Println("Este IF funciona!")
		// }

		// idade := 32
		// expressaoBoleana := idade > 18
		// if expressaoBoleana {
		// 	fmt.Println("É maior de idade!")
		// }

		//Mais sucinto que utilizar IF
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
			break //Não é obrigatório, logo o Go não vai "reclamar"
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		// case 4: // Várias linhas dentro do case
		// 	fmt.Println("Exemplo 4")
		// 	fmt.Println("Outra linha")
		// 	fmt.Println("Última linha")
		default:
			fmt.Println("Não conheço este comando!")
			os.Exit(-1)
		}
	}

	//É case sensitive. Precisa tratar como lowercase ou uppercase
	// fmt.Println("Escolha um prato")
	// fmt.Println("B - Batata frita")
	// fmt.Println("P - Pizza")
	// fmt.Println("H - Hambúrger")

	// var prato string
	// fmt.Scan(&prato)

	// switch prato {
	// case "B":
	// 	fmt.Println("Você escolheu Batata frita!")
	// case "P":
	// 	fmt.Println("Você escolheu Pizza!")
	// case "H":
	// 	fmt.Println("Você escolheu hambúrguer!")
	// default:
	// 	fmt.Println("Prato não existe!")

	// }

}

func exibeIntroducao() {
	nome := "Fabio Rubim"
	versao := 1.1
	fmt.Println("Olá, Sr. ", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Inciar monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	//Função que captura o que usuário digitar
	var comandoLido int
	//fmt.Scanf("%d", &comandoLido) //Não é utilizado inferência, você informa com o %d o tipo que será guardado no ponteiro &comandoLido, neste caso um int
	//Scanf("%d", &comandoLido) não precisaria saber o tipo, já que a variável comando já foi declarada como int
	//Como já se sabe o tipo utilize somente o scan(&comandoLido). Não precisa colocar o %d
	fmt.Scan(&comandoLido) //Se colocar algo que não seá um INT, nada será "pego", logo o valor de comando será 0.
	fmt.Println("O ponteiro da minha variável comando é", &comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	//utilizando slices
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	// fmt.Println("Sites", sites)
	// //Go só possui FOR como estrutura de repetição
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println("Site ", sites[i])
	// }

	for i := 0; i < monitoramentos; i++ {
		//Esse tipo de FOR retorna dois valores, um indice e o próprio elementos
		for indice, site := range sites {
			fmt.Println("Na posicao ", indice, "do slice tem o site ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)//5 segundos
		fmt.Println("")
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas! Status code:", resp.StatusCode)
	}
}

//Utilizando arrays
// func iniciarMonitoramento() {
// 	fmt.Println("Monitorando...")
// 	var sites [4]string //Todo array no Go tem um tamanho fixo [4] por exemplo. Normalmente não se trabalha com array, e sim com slice. Slice é um abstração, funciona "em com a" de um array
// 	sites[0] = "https://random-status-code.herokuapp.com/"
// 	sites[1] = "https://www.alura.com.br"
// 	sites[2] = "https://www.caelum.com.br"
// 	//sites[3] = "https://www.google.com.br"
// 	fmt.Println("Site 3", sites[3]) //Será impressa uma string vazia, já que não foi atribuído nenhum valor. Isso ocorre pois o valor padrão para o tipo string é vazio.
// 	fmt.Println("Sites TypeOf", reflect.TypeOf(sites))
// 	//sites[4] = "https://www.santander.com.br"
// 	fmt.Println("Sites", sites)
// 	//site := "https://www.alura.com.br"
// 	site := "https://random-status-code.herokuapp.com/"
// 	resp, _ := http.Get(site)
// 	//fmt.Println(resp)

// 	if resp.StatusCode == 200 {
// 		fmt.Println("Site:", site, "foi carregado com sucesso!")
// 	} else {
// 		fmt.Println("Site:", site, "está com problemas! Status code:", resp.StatusCode)
// 	}
// }

// func exibeNomes() {
// 	//TUdo é um aray em Go, os slices são arrays em Go. O que diferecia é que os slices abstraem algumas funcionalidade, como adcionar um novo item.
// 	nomes := []string{"Fabio", "Marina", "Renato", "Valéria"} //Slices não necessitam declaram o tamanho, ele infere a partir da quantidade de itens passados.
// 	fmt.Println(nomes)
// 	fmt.Println("Nomes TypeOf", reflect.TypeOf(nomes))
// 	fmt.Println("O meu slice tem", len(nomes), "itens")                 //Quantidade de items
// 	fmt.Println("O meu slice tem a capacidade de", cap(nomes), "itens") //Capacidade disponível

// 	//O append dobra o tamanho do slice de acordo com a quantidade de itens de acordo com o tamanho do slice anterior
// 	//Se o tamnho dos itens adicionados superarem o dobro do tamanho, então o novo slice terá o tamanho da quantidade total de itens
// 	nomes = append(nomes, "Tomaz", "Zure", "Euclydes", "Lito", "Cláudia", "Regina", "Rodrigo")

// 	fmt.Println(nomes)
// 	fmt.Println("Nomes TypeOf", reflect.TypeOf(nomes))
// 	fmt.Println("O meu slice tem", len(nomes), "itens")
// 	fmt.Println("O meu slice tem a capacidade de", cap(nomes), "itens")
// }

/*
func main() {
	estados := devolveEstadosDoSudeste()
	fmt.Println(estados)

func devolveEstadosDoSudeste() [4]string {
	var estados [4]string
	estados[0] = "RJ"
	estados[1] = "SP"
	estados[2] = "MG"
	estados[3] = "ES"
	return estados
}
*/
