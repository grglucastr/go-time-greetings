package gotimegreetings

import (
	"encoding/json"
	"os"
	"time"
)

func Greet(lang string) string {

	t := time.Now()

	content := loadFile(lang)

	var result map[string]string
	json.Unmarshal(content, &result)

	if t.Hour() > 17 {
		return result["goodEvening"]
	}

	if t.Hour() > 12 {
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
		raw, err = os.ReadFile(filename)
	case "pt_br":
		filename = "strings-pt_br.json"
		raw, err = os.ReadFile(filename)
	}

	if err != nil {
		panic("Error load " + filename)
	}

	return raw
}
