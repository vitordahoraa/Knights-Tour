package main

import (
	"fmt"
	"sort"
)

var validKnightsMove = [][]int8{
	{1, 2},
	{1, -2},
	{2, 1},
	{2, -1},
	{-1, 2},
	{-1, -2},
	{-2, 1},
	{-2, -1},
}

// Validar se o tabueleiro esta todo feito
func validaTabuleiro(slice2Dint_tabuleiro [][]int8) bool {
	haPosicoesNaoVisitadas := false

	for i := range slice2Dint_tabuleiro {
		for j := range slice2Dint_tabuleiro[i] {
			if !haPosicoesNaoVisitadas {
				haPosicoesNaoVisitadas = slice2Dint_tabuleiro[i][j] == -1
			}
		}
	}

	return !haPosicoesNaoVisitadas
}

func retornaMovimentosValidos(slice2Dint_tabuleiro [][]int8, int_posicaoLinhaAtual uint8, int_posicaoColunaAtual uint8) [][]uint8 {

	slice_movimentosValidos := make([][]uint8, 0)

	//Iterar posicoes validas
	for i := range validKnightsMove {

		//Salva as posicoes a ser validada
		linhaIteration := (int8(int_posicaoLinhaAtual) + validKnightsMove[i][0])
		colunaIteration := (int8(int_posicaoColunaAtual) + validKnightsMove[i][1])
		//ignorar posicoes fora do tabuleiro
		if (linhaIteration < 0) || (linhaIteration >= int8(len(slice2Dint_tabuleiro))) || (colunaIteration >= int8(len(slice2Dint_tabuleiro))) || (colunaIteration < 0) {
			continue
		}
		//Adicionar posicao a validar na var positionValidar
		positionValidar := make([]uint8, 2)
		positionValidar[0] = uint8(linhaIteration)
		positionValidar[1] = uint8(colunaIteration)

		if slice2Dint_tabuleiro[positionValidar[0]][positionValidar[1]] == -1 {
			slice_movimentosValidos = append(slice_movimentosValidos, positionValidar)
		}
	}
	return slice_movimentosValidos
}

func initVars(int_tamanhoTabuleiro *uint8, int_linhaInicial *uint8, int_colunaInicial *uint8) {

	//Solicitando tamanho do tabuleiro, apenas tamanhos validos
	for tabuleiroInvalido := true; tabuleiroInvalido; tabuleiroInvalido = !(*int_tamanhoTabuleiro%2 == 0 && *int_tamanhoTabuleiro != 1 && *int_tamanhoTabuleiro != 2 && *int_tamanhoTabuleiro != 4) {

		fmt.Print("Digite o tamanho da tabuleiro: ")
		fmt.Scanln(int_tamanhoTabuleiro)

	}
	//Solicitando linha inicial do cavalo
	fmt.Print("Digite o linha inicial do cavalo (index 0): ")
	fmt.Scanln(int_linhaInicial)

	//Solicitando coluna do cavalo
	fmt.Print("Digite o coluna inicial do cavalo (index 0): ")
	fmt.Scanln(int_colunaInicial)
}

func criarTabuleiro(int_tamanhoTabuleiro uint8) [][]int8 {
	// Criar Array2D
	slice2Dint_tabuleiro := make([][]int8, int_tamanhoTabuleiro)

	//Para cada posicao em i
	for i := range slice2Dint_tabuleiro {

		//Criar array 1D
		slice2Dint_tabuleiro[i] = make([]int8, int_tamanhoTabuleiro)

		//Inicializar valor em -1
		for j := range slice2Dint_tabuleiro[i] {
			slice2Dint_tabuleiro[i][j] = -1
		}
	}

	return slice2Dint_tabuleiro
}

// Funcao utilizada pra retornar o numero de posicoes validas aparece na slice de movimentosValidos
func numeroDeInstanciasDeMovimentosValidos(slice_movimentosValidos []int8) map[int8]int8 {
	dict := make(map[int8]int8)
	for _, num := range slice_movimentosValidos {
		dict[num] = dict[num] + 1
	}
	return dict
}

func visitarPosicaoTabuleiro(slice2Dint_tabuleiro [][]int8, uint8_posicaoLinhaAtual uint8, uint8_posicaoColunaAtual uint8, int8_numeroVisita int8) bool {

	//Verifica movimentos validos dessa posicao
	slice_movimentosValidos := retornaMovimentosValidos(slice2Dint_tabuleiro, uint8_posicaoLinhaAtual, uint8_posicaoColunaAtual)

	//Visitar posicao
	slice2Dint_tabuleiro[uint8_posicaoLinhaAtual][uint8_posicaoColunaAtual] = int8_numeroVisita

	printTabuleiroAndWaitUserInput(slice2Dint_tabuleiro)

	//Caso essa posicao nao seja valida, retornar true caso o tabuleiro esteja completo, e false caso contrario
	if len(slice_movimentosValidos) == 0 {
		return validaTabuleiro(slice2Dint_tabuleiro)
	}

	//Incrementa a visita
	int8_numeroVisita++

	slice_proxMovimentos := make([]int8, len(slice_movimentosValidos))

	for posicao := range slice_movimentosValidos {
		// Colocar a posicao valida de todas as posicoes
		slice_proxMovimentos[posicao] = int8(len(retornaMovimentosValidos(slice2Dint_tabuleiro, slice_movimentosValidos[posicao][0], slice_movimentosValidos[posicao][1])))
		//slice2Dint_tabuleiro[slice_movimentosValidos[posicao][0]][slice_movimentosValidos[posicao][1]] = int8(len(retornaMovimentosValidos(slice2Dint_tabuleiro, slice_movimentosValidos[posicao][0], slice_movimentosValidos[posicao][1])))
	}

	//slice_numVezes := numeroDeInstanciasDeMovimentosValidos(slice_proxMovimentos)

	// Sorteia do menor ao maior
	sort.SliceStable(slice_proxMovimentos, func(i, j int) bool { return slice_proxMovimentos[i] < slice_proxMovimentos[j] })

	//Criar variavel para iterar sobre todas as casas que tem a mesma casas validas
	//contadorNumVezes := 0
	//Para cada movimento proximo, do menor para o maior
	for proxMov := range slice_proxMovimentos {
		//Inicia contador de vezes dessa iteracao
		//contadorNumVezesIteration := 0
		for posicao := range slice_movimentosValidos {
			casasValidas := int8(len((retornaMovimentosValidos(slice2Dint_tabuleiro, slice_movimentosValidos[posicao][0], slice_movimentosValidos[posicao][1]))))
			//Caso o numero de casas validas seja o menor iterado
			if casasValidas == slice_proxMovimentos[proxMov] {
				// Se retornou false
				fmt.Println(int8_numeroVisita)
				if !visitarPosicaoTabuleiro(slice2Dint_tabuleiro, slice_movimentosValidos[posicao][0], slice_movimentosValidos[posicao][1], int8_numeroVisita) {
					continue
				}

				break
			}
		}
		// Colocar a posicao valida de todas as posicoes
		//slice2Dint_tabuleiro[slice_movimentosValidos[posicao][0]][slice_movimentosValidos[posicao][1]] = int8(slice_proxMovimentos[posicao])
		//
	}

	return true
}

func visitarTodoTabuleiro(slice2Dint_tabuleiro [][]int8, uint8_posicaoInicialLinha uint8, uint8_posicaoInicialColuna uint8) {
	//Visitar primeira posicao
	visitarPosicaoTabuleiro(slice2Dint_tabuleiro, uint8_posicaoInicialLinha, uint8_posicaoInicialColuna, 1)
}

// Printar o tabuleiro
func printTabuleiroAndWaitUserInput(slice2Dint_tabuleiro [][]int8) {

	fmt.Print("IMPRIMINDO TABULEIRO\n\n")

	for i := range slice2Dint_tabuleiro {
		for j := range slice2Dint_tabuleiro[i] {
			fmt.Printf("%d  ", slice2Dint_tabuleiro[i][j])
		}
		fmt.Print("\n")
	}

	//Wait user input
	//fmt.Print("Aperte uma tecla para continuar!\n")
	//fmt.Scanln()
}

func main() {
	//Iniciando as vars
	var int_tamanhoTabuleiro uint8
	var int_linhaInicial uint8
	var int_colunaInicial uint8

	initVars(&int_tamanhoTabuleiro, &int_linhaInicial, &int_colunaInicial)

	// Criar tabuleiro
	// -1 = nao visitado, qualquer outro valor eh a ordem de visitacao
	slice2Dint_tabuleiro := criarTabuleiro(int_tamanhoTabuleiro)

	visitarTodoTabuleiro(slice2Dint_tabuleiro, int_linhaInicial, int_colunaInicial)

}
