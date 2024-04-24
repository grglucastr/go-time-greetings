package gotimegreetings

import (
	"embed"
	"encoding/json"
)

//go:embed strings.json
//go:embed strings-pt_br.json
//go:embed strings-err.json
var f embed.FS

func Greet(lang string, hour int) string {

	content := loadFile(lang)

	var result map[string]string
	json.Unmarshal(content, &result)

	if hour > 17 {
		return result["goodEvening"]
	}

	if hour > 12 {
		return result["goodAfternoon"]
	}

	return result["goodMorning"]
}

func loadFile(lang string) []byte {

	filename := "strings.json"
	var raw []byte
	var err error

	switch lang {
	case "en":
		raw, err = f.ReadFile(filename)
	case "pt_br":
		filename = "strings-pt_br.json"
		raw, err = f.ReadFile(filename)
	default:
		filename = "strings-" + lang + ".json"
		panic("Error load " + filename)
	}

	if err != nil {
		panic(err)
	}

	return raw
}
