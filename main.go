package main

import (
	"fmt"
	"sort"
)

type Jogador struct {
	Nome  string
	Nivel int
}

type Time struct {
	Jogadores []Jogador
}

var jogadores []Jogador = []Jogador{
	{Nome: "Thiago", Nivel: 4},
	{Nome: "Paulo", Nivel: 5},
	{Nome: "Marcel", Nivel: 3},
	{Nome: "Thales", Nivel: 5},
	{Nome: "Brunão", Nivel: 1},
	{Nome: "Juliano", Nivel: 3},
	{Nome: "Claudinho", Nivel: 4},
	{Nome: "Serginho", Nivel: 1},
	{Nome: "Waldeka", Nivel: 2},
	{Nome: "Higor", Nivel: 2},
}

func OrdenarTime(jogadores []Jogador) []Jogador {

	sort.SliceStable(jogadores, func(i, j int) bool {
		return jogadores[i].Nivel > jogadores[j].Nivel
	})

	return jogadores
}

func SepararTimes(jogadores []Jogador) (timeA, timeB Time) {
	jgOrdenados := OrdenarTime(jogadores)

	a := Time{}
	b := Time{}

	for i, jogador := range jgOrdenados {
		if i%2 == 0 {
			a.Jogadores = append(a.Jogadores, jogador)
		} else {
			b.Jogadores = append(b.Jogadores, jogador)
		}
	}

	return a, b
}

func main() {
	fmt.Println("Hello Separador de Times")
}
