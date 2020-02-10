package main

import "fmt"

func main() {
	nome := "Fabio Rubim"
	versao := 1.1
	fmt.Println("Olá, Sr. ", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Inciar monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("0- Sair do programa")

	//Função que captura o que usuário digitar
	var comando int
	//fmt.Scanf("%d", &comando) //Não é utilizado inferência, você informa com o %d o tipo que será guardado no ponteiro &comando, neste caso um int
	//Scanf("%d", &comando) não precisaria saber o tipo, já que a variável comando já foi declarada como int
	//Como já se sabe o tipo utilize somente o scan(&comando). Não precisa colocar o %d
	fmt.Scan(&comando) //Se colocar algo que não seá um INT, nada será "pego", logo o valor de comando será 0.
	fmt.Println("O ponteiro da minha variável comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)

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
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
		break //Não é obrigatório, logo o Go não vai "reclamar"
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
	// case 4: // Várias linhas dentro do case
	// 	fmt.Println("Exemplo 4")
	// 	fmt.Println("Outra linha")
	// 	fmt.Println("Última linha")
	default:
		fmt.Println("Não conheço este comando!")
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
