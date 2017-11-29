package parser

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"text/scanner"
)

func (p *Parser) CleanAddress(detailAddress []byte) string {
	var result []string

	var data struct {
		Address    string `json:"address"`
		CityID     int64  `json:"city_id"`
		ProvinceID int64  `json:"province_id"`
		DistrictID int64  `json:"district_id"`
		PostalCode string `json:"postal_code"`
	}

	if err := json.Unmarshal(detailAddress, &data); err != nil {
		log.Println(err)
		return ""
	}

	addr := data.Address
	provinceParam := fmt.Sprintf("%d", data.ProvinceID)
	cityParam := fmt.Sprintf("%d", data.CityID)
	districtParam := fmt.Sprintf("%d", data.DistrictID)
	addr = p.newLine.ReplaceAllString(addr, " ")
	addr = p.special.ReplaceAllString(addr, " ")
	addr = p.ilegalOctalNumber.ReplaceAllString(addr, " ")
	addr = strings.ToLower(addr)

	// remove district,city,province name
	if provinceParam != "" {
		addr = strings.Replace(addr, province[provinceParam], "", -1)
	}
	if cityParam != "" {
		addr = strings.Replace(addr, city[cityParam], "", -1)
	}
	if districtParam != "" {
		addr = strings.Replace(addr, district[districtParam], "", -1)
	}
	if data.PostalCode != "" {
		addr = strings.Replace(addr, data.PostalCode, "", -1)
	}

	if addr != "" {
		words := tokenize(addr)
		for _, word := range words {
			if _, ok := address[word]; !ok {
				result = append(result, word)
			}
		}
	}

	return strings.Join(result, " ")
}

func tokenize(inp string) []string {
	var (
		s      scanner.Scanner
		tok    rune
		result []string
	)

	s.Init(strings.NewReader(inp))
	for tok != scanner.EOF {
		tok = s.Scan()
		result = append(result, s.TokenText())
	}

	return result
}
