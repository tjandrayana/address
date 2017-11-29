package address

import (
	"encoding/csv"
	"log"
	"os"
)

func (m *Module) LoadStopWords() {
	generalPath := "files/stopwords/general.csv"
	provincePath := "files/stopwords/export-ws_province.csv"
	cityPath := "files/stopwords/export-ws_city.csv"
	districtPath := "files/stopwords/export-ws_district.csv"

	okGeneral := ReadStopAddress(generalPath, "general")
	if !okGeneral {
		log.Println("failed to read general stopwords")
	}

	okProvince := ReadStopAddress(provincePath, "province")
	if !okProvince {
		log.Println("failed to read province stopwords")
	}

	okCity := ReadStopAddress(cityPath, "city")
	if !okCity {
		log.Println("failed to read city stopwords")
	}

	okDistrict := ReadStopAddress(districtPath, "district")
	if !okDistrict {
		log.Println("failed to read district stopwords")
	}
}

func ReadStopAddress(path, info string) bool {
	err := setStopAddress(path, info)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func setStopAddress(path, info string) error {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Println("[Error] Set Stopwords from ", path)
		return err
	}

	csvFile := csv.NewReader(f)
	content, err := csvFile.ReadAll()
	if err != nil {
		log.Println(err)
	}

	csvData := content

	// Create a new reader.
	for i := range csvData {
		if info == "province" {
			province[csvData[i][0]] = csvData[i][1]
		} else if info == "city" {
			city[csvData[i][0]] = csvData[i][1]
		} else if info == "district" {
			district[csvData[i][0]] = csvData[i][1]
		} else {
			for j := range csvData[i] {
				address[csvData[i][j]] = " "
			}
		}

	}
	return nil
}
