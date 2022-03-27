package main

import (
	"dataGenerate/generate"
	"encoding/json"
	"fmt"
	"syscall/js"
)

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			fmt.Printf("Invalid no of arguments passed")
			return nil
		}
		inputJSON := args[0].String()
		qtd := args[1].Int()
		inputFormat := make(map[string]string)
		err := json.Unmarshal([]byte(inputJSON), &inputFormat)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return nil
		}
		pretty, err := generate.Generate(inputFormat, qtd)
		if err != nil {
			fmt.Printf("unable togenerate data %s\n", err)
			return nil
		}
		dataJson, _ := json.MarshalIndent(pretty, "", " ")
		return string(dataJson)
	})
	return jsonFunc
}

func main() {
	js.Global().Set("GOGenerate", jsonWrapper())
	for {
		<-make(chan bool)
	}
}
