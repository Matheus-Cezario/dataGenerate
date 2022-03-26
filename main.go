package main

import (
	"dataGenerate/generate"
	"encoding/json"
	"fmt"
)

func main() {

	data, _ := generate.Generate(map[string]string{
		"nome":"name:m",
		"id":"number:10:25",
		"cidade":"choise:Ouro Branco:Ouro Preto:Belo Horizonte",
		"produtos":"choiseAny:r:Bala:Sorvete:Ovo:Arroz:Batata:Chocolate",
		},5)
	dataJson,_ := json.MarshalIndent(data,""," ")
	fmt.Println(string(dataJson))
}