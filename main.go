package main

import (
	"fmt"
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

		slice_movimentosValidos = append(slice_movimentosValidos, positionValidar)
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

func visitarTodoTabuleiro(slice2Dint_tabuleiro [][]int8, int_linhaInicial uint8, int_colunaInicial uint8) [][]int8 {

	//Visitar posicao inicial
	slice2Dint_tabuleiro[int_linhaInicial][int_colunaInicial] = 1

	uint8_posicaoLinhaAtual := int_linhaInicial
	uint8_posicaoColunaAtual := int_colunaInicial

	slice_movimentosValidos := retornaMovimentosValidos(slice2Dint_tabuleiro, uint8_posicaoLinhaAtual, uint8_posicaoColunaAtual)

	return slice2Dint_tabuleiro

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

	slice2Dint_tabuleiro = visitarTodoTabuleiro(slice2Dint_tabuleiro, int_linhaInicial, int_colunaInicial)

	fmt.Println(slice2Dint_tabuleiro[int_linhaInicial][int_colunaInicial])

}
