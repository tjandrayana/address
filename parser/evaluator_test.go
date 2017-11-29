package parser

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AddressDetail struct {
	CityID     int64  `json:"city_id"`
	ProvinceID int64  `json:"province_id"`
	DistrictID int64  `json:"district_id"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

func TestCleanAddress(t *testing.T) {

	p := New()

	inputJson := `[
		{
			"city_id":200,
			"province_id":14,
			"district_id":2314,
			"address":"Jalan Kolonel Sugiono 87 A Purworejo Jawa Tengah",
			"postal_code":"54111",
			"country":"indonesia"
		},
		{
			"city_id":167,
			"province_id":12,
			"district_id":2017,
			"address":"Perumahan Taman Naroging Indah Jalan Bojong Molek IV Blok F23/12 RT 004 RW 014 Kecamatan Rawalumbu Kelurahan Bojong Rawalumbu Bekasi",
			"postal_code":"54111",
			"country":"indonesia"
		},
		{
			"city_id":175,
			"province_id":13,
			"district_id":2270,
			"address":"JL.H.R. RASUNA SAID",
			"postal_code":"12940",
			"country":""
		},
		{
			"city_id":463,
			"province_id":11,
			"district_id":1613,
			"address":"GD. GRIYA NIAGA 2 LT.5 JL. KH. WAHID HASYIM BLOK B.IV/3 BINTARO JAYA SEKTOR 7",
			"postal_code":"15224",
			"country":""
		},
		{
			"city_id":175,
			"province_id":13,
			"district_id":2264,
			"address":"Jl. Jend. Gatot Subroto Kav 36-38 Jakarta 12190",
			"postal_code":"12190",
			"country":""
		},
		{
			"city_id":167,
			"province_id":12,
			"district_id":2207,
			"address":"Jln Setia 1 N Rt. 008/004 no. 43 pintu no. 6 jatiwaringin pondok gede bekasi",
			"postal_code":"17411",
			"country":""
		},
		{
			"city_id":174,
			"province_id":13,
			"district_id":2258,
			"address":"Jl bhakti iv no 17 Rt 03 rw 09 kemanggisan palmerah jakbar(pagar hijau)",
			"postal_code":"11480",
			"country":""
		},
		{
			"city_id":165,
			"province_id":12,
			"district_id":2176,
			"address":"Jl. Soekarno Hatta No. 236 (samping BCA terminal Lw. Panjang), Kel. Kebonlega",
			"postal_code":"40235",
			"country":""
		},
		{
			"city_id":182,
			"province_id":14,
			"district_id":2353,
			"address":"Dukuh Sidorejo RT 027 RW 007 Desa Ketanggan Kec. Gringsing, Kab. Batang, Jawa Tengah 51281",
			"postal_code":"51281",
			"country":""
		},
		{
			"city_id":171,
			"province_id":12,
			"district_id":2228,
			"address":"Jl.Pekapuran Gg.H Camat No.27 Rt/RW 02/24 Kel.Sukatani / Kec.Tapos (Depan Alfamart Bakti ABRI)",
			"postal_code":"16454",
			"country":""
		}]`
	var input []AddressDetail

	if err := json.Unmarshal([]byte(inputJson), &input); err != nil {
		log.Println(err)
	}

	output := []string{
		"kolonel sugiono 87 a",
		"taman naroging molek iv f23 12 4 14 rawalumbu rawalumbu",
		"h r rasuna said",
		"gd niaga 2 5 kh wahid hasyim b iv 3 bintaro jaya 7",
		"gatot subroto 36 38",
		"setia 1 n 8 4 43 pintu 6 jatiwaringin",
		"bhakti iv 17 3 9 kemanggisan jakbar",
		"soekarno hatta 236 bca terminal lw panjang kebonlega",
		"dukuh sidorejo 27 7 ketanggan",
		"pekapuran h camat 27 2 24 sukatani tapos alfamart bakti abri",
	}

	size := len(input)
	for i := 0; i < size; i++ {
		addressJson, err := json.Marshal(input[i])
		if err != nil {
			log.Println(err)
		}
		assert.Equal(t, output[i], p.CleanAddress(addressJson))
	}
}

func TestTokenize(t *testing.T) {

	input := "jakarta barat"
	output := []string{
		"jakarta",
		"barat",
		"",
	}

	assert.Equal(t, output, tokenize(input))
}
