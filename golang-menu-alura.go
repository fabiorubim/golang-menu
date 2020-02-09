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

}
