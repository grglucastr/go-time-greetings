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
	var err error
	switch lang {
	case "en":
		raw, err = os.ReadFile(filename)
	case "pt_br":
		filename = "strings-pt_br.json"
		raw, err = os.ReadFile(filename)
	default:
		filename = "strings-" + lang + ".json"
		panic("Error load " + filename)
	}

	if err != nil {
		fmt.Println("Deu erro vice...")
		panic(err)
	}

	return raw
}
