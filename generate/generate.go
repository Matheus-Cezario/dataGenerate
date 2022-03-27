package generate

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func Generate(jsonData map[string]string, quantity int) ([]map[string]interface{}, error) {
	var listData []map[string]interface{}
	for i := 0; i < quantity; i++ {
		var data map[string]interface{} = map[string]interface{}{}
		for key, value := range jsonData {
			values := strings.Split(value, ":")
			if len(values) == 0 {
				return nil, fmt.Errorf("error: Value len of key %s equals 0", key)
			}
			var gValue interface{}
			switch strings.ToLower(values[0]) {
			case Name:
				gValue = generateName(values...)
			case Number:
				gValue = generateNumber(values...)
			case Choise:
				gValue = generateChoise(values...)
			case ChoiseAny:
				gValue = generateChoiseAny(values...)
			default:
				gValue = "Not Found!"
			}
			data[key] = gValue

		}
		listData = append(listData, data)
	}

	return listData, nil
}

func generateName(parans ...string) string {
	lenParans := len(parans)
	var gender string
	if lenParans >= 2 {
		gender = parans[1]
	} else {
		gender = "A"
	}
	switch strings.ToUpper(gender) {
	case "M":
		randomIndex := rand.Intn(len(MascNames))
		return MascNames[randomIndex]
	case "F":
		randomIndex := rand.Intn(len(FemNames))
		return FemNames[randomIndex]
	default:
		listName := MascNames
		if rand.Float32() > 0.5 {
			listName = FemNames
		}
		randomIndex := rand.Intn(len(listName))
		return listName[randomIndex]
	}
}

func generateNumber(parans ...string) int {
	var (
		start int
		end   int
	)

	if len(parans) >= 2 {
		start, _ = strconv.Atoi(parans[1])
		if len(parans) >= 3 {
			end, _ = strconv.Atoi(parans[2])
		} else {
			end = 1000
		}
	} else {
		start = 0
		end = 100
	}
	if start > end {
		start, end = end, start
	}

	return start + int(rand.Float32()*float32(end-start))
}

func generateChoise(parans ...string) string {
	choises := parans[1:]
	if len(choises) == 0 {
		return "No data"
	}
	randomIndex := rand.Intn(len(choises))
	return strings.ReplaceAll(choises[randomIndex], "\r", "")
}

func generateChoiseAny(parans ...string) []string {
	if len(parans) <= 3 {
		return []string{"No data"}
	}
	var qtd int
	var err error
	if strings.ToLower(parans[1]) == "r" {
		qtd = rand.Intn(len(parans) - 2)
	} else {
		qtd, err = strconv.Atoi(parans[1])
		if err != nil {
			return []string{err.Error()}
		}
	}
	canRepete := strings.ToLower(parans[2])
	choises := parans[2:]
	dataReturn := []string{}
	for i := 0; i < qtd; i++ {
		randomIndex := rand.Intn(len(choises))
		data := strings.ReplaceAll(choises[randomIndex], "\r", "")
		dataReturn = append(dataReturn, data)
		if canRepete != "y" {
			choises = remove(choises, randomIndex)
		}
	}
	return dataReturn

}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
