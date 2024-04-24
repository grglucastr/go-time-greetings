package gotimegreetings

import (
	"encoding/json"
	"fmt"
	"os"
)

func Greet(lang string, hour int) string {

	content := loadFile(lang)

	var result map[string]string
	json.Unmarshal(content, &result)

	fmt.Println("Eita caroço")

	a := fmt.Sprintf("kiding me: %v", result)
	fmt.Println(a)

	if hour > 17 {
		fmt.Println(result["goodEvening"])
		return result["goodEvening"]
	}

	if hour > 12 {
		fmt.Println(result["goodAfternoon"])
		return result["goodAfternoon"]
	}

	fmt.Println(result["goodAfternoon"])
	return result["goodMorning"]
}

func loadFile(lang string) []byte {

	filename := "strings.json"
	var raw []byte

	switch lang {
	case "en":
		raw, _ = os.ReadFile(filename)
	case "pt_br":
		filename = "strings-pt_br.json"
		raw, _ = os.ReadFile(filename)
	default:
		filename = "strings-" + lang + ".json"
		panic("Error load " + filename)
	}

	return raw
}
