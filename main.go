package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"
)

type Jogador struct {
	Nome  string `json:"nome"`
	Nivel int    `json:"nivel"`
}

func carregarJogadoresJSON(caminho string) ([]Jogador, error) {
	file, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var jogadores []Jogador
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&jogadores); err != nil {
		return nil, err
	}

	return jogadores, nil
}

func salvarTimesEmArquivo(time1, time2 []Jogador, nivelTime1, nivelTime2 int, formato string) error {
	timestamp := time.Now().Format("20060102_150405")
	var caminho string

	if formato == "md" {
		caminho = fmt.Sprintf("times/times_%s.md", timestamp)
		file, err := os.Create(caminho)
		if err != nil {
			return err
		}
		defer file.Close()

		fmt.Fprintf(file, "# *Atenção aos times!*\n\n\n")

		fmt.Fprintf(file, "*Time Branco(%d)*:\n", len(time1))
		for _, jogador := range time1 {
			fmt.Fprintf(file, "- %s\n", jogador.Nome)
		}

		fmt.Fprintf(file, "\n*Time Preto (%d)*:\n", len(time2))
		for _, jogador := range time2 {
			fmt.Fprintf(file, "- %s\n", jogador.Nome)
		}

	} else {
		caminho = fmt.Sprintf("times/times_%s.json", timestamp)
		file, err := os.Create(caminho)
		if err != nil {
			return err
		}
		defer file.Close()

		dados := map[string]interface{}{
			"time1": map[string]interface{}{
				"jogadores":       time1,
				"total_jogadores": len(time1),
				"nivel_total":     nivelTime1,
			},
			"time2": map[string]interface{}{
				"jogadores":       time2,
				"total_jogadores": len(time2),
				"nivel_total":     nivelTime2,
			},
		}

		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(dados); err != nil {
			return err
		}
	}

	fmt.Printf("Times salvos em '%s'\n", caminho)
	return nil
}

func balancearTimes(jogadores []Jogador) ([]Jogador, []Jogador, int, int) {
	sort.Slice(jogadores, func(i, j int) bool {
		return jogadores[i].Nivel > jogadores[j].Nivel
	})

	var time1, time2 []Jogador
	var nivelTime1, nivelTime2 int

	for _, jogador := range jogadores {
		if (nivelTime1 < nivelTime2) || (nivelTime1 == nivelTime2 && len(time1) <= len(time2)) {
			time1 = append(time1, jogador)
			nivelTime1 += jogador.Nivel
		} else {
			time2 = append(time2, jogador)
			nivelTime2 += jogador.Nivel
		}
	}

	return time1, time2, nivelTime1, nivelTime2
}

func main() {
	arquivo := flag.String("arquivo", "exemplo_jogadores.json", "Caminho do arquivo JSON com os jogadores")
	formato := flag.String("formato", "json", "Formato de saída (txt ou json)")
	flag.Parse()

	jogadores, err := carregarJogadoresJSON(*arquivo)
	if err != nil {
		fmt.Println("Erro ao carregar jogadores:", err)
		return
	}

	time1, time2, nivelTime1, nivelTime2 := balancearTimes(jogadores)

	err = salvarTimesEmArquivo(time1, time2, nivelTime1, nivelTime2, *formato)
	if err != nil {
		fmt.Println("Erro ao salvar times:", err)
	}
}
