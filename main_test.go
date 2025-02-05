package main

import (
	"reflect"
	"testing"
)

func TestOrdenarTime(t *testing.T) {
	want := []Jogador{
		{Nome: "Paulo", Nivel: 5},
		{Nome: "Thales", Nivel: 5},
		{Nome: "Thiago", Nivel: 4},
		{Nome: "Claudinho", Nivel: 4},
		{Nome: "Marcel", Nivel: 3},
		{Nome: "Juliano", Nivel: 3},
		{Nome: "Waldeka", Nivel: 2},
		{Nome: "Higor", Nivel: 2},
		{Nome: "Brunão", Nivel: 1},
		{Nome: "Serginho", Nivel: 1},
	}

	got := OrdenarTime(jogadores)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestSepararTimes(t *testing.T) {
	wantTime1 := Time{
		[]Jogador{
			{Nome: "Paulo", Nivel: 5},
			{Nome: "Thiago", Nivel: 4},
			{Nome: "Marcel", Nivel: 3},
			{Nome: "Waldeka", Nivel: 2},
			{Nome: "Brunão", Nivel: 1},
		},
	}
	wantTime2 := Time{
		[]Jogador{
			{Nome: "Thales", Nivel: 5},
			{Nome: "Claudinho", Nivel: 4},
			{Nome: "Juliano", Nivel: 3},
			{Nome: "Higor", Nivel: 2},
			{Nome: "Serginho", Nivel: 1},
		},
	}
	gotTime1, gotTime2 := SepararTimes(jogadores)

	if !reflect.DeepEqual(wantTime1, gotTime1) {
		t.Errorf("want time 1 %v got time1 %v", wantTime1, gotTime1)
	}

	if !reflect.DeepEqual(wantTime2, gotTime2) {
		t.Errorf("want time 1 %v got time1 %v", wantTime2, gotTime2)
	}

}
