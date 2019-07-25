package output

import (
	"encoding/json"
	"log"

	"github.com/spf13/afero"
)

func writeJSON(file afero.File, outputData Data) {
	bytes, err := json.MarshalIndent(outputData, "", "\t")
	if err != nil {
		log.Fatalln(err)
	}
	file.WriteString(string(bytes))
}
