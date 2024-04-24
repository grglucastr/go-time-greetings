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

	if hour > 17 {
		return fmt.Sprintf("%v", result["goodEvening"])
	}

	if hour > 12 {
		return fmt.Sprintf("%v", result["goodAfternoon"])
	}

	return fmt.Sprintf("%v", result["goodMorning"])
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
