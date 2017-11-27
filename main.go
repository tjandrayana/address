package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/tjandrayana/normalize-address/helper/parser"
)

func main() {

	input := `[{"city_id":200,"province_id":14,"district_id":2314,"address":"Jalan Kolonel Sugiono 87 A Purworejo Jawa Tengah","postal_code":"54111","country":""},{"city_id":167,"province_id":12,"district_id":2017,"address":"Perumahan Taman Naroging Indah Jalan Bojong Molek IV Blok F23/12 RT 004 RW 014 Kecamatan Rawalumbu Kelurahan Bojong Rawalumbu Bekasi","postal_code":"54111","country":""},{"city_id":175,"province_id":13,"district_id":2270,"address":"JL.H.R. RASUNA SAID","postal_code":"12940","country":""},{"city_id":463,"province_id":11,"district_id":1613,"address":"GD. GRIYA NIAGA 2 LT.5 JL. KH. WAHID HASYIM BLOK B.IV/3 BINTARO JAYA SEKTOR 7","postal_code":"15224","country":""},{"city_id":175,"province_id":13,"district_id":2264,"address":"Jl. Jend. Gatot Subroto Kav 36-38 Jakarta 12190","postal_code":"12190","country":""},{"city_id":167,"province_id":12,"district_id":2207,"address":"Jln Setia 1 N Rt. 008/004 no. 43 pintu no. 6 jatiwaringin pondok gede bekasi","postal_code":"17411","country":""},{"city_id":174,"province_id":13,"district_id":2258,"address":"Jl bhakti iv no 17 Rt 03 rw 09 kemanggisan palmerah jakbar(pagar hijau)","postal_code":"11480","country":""},{"city_id":165,"province_id":12,"district_id":2176,"address":"Jl. Soekarno Hatta No. 236 (samping BCA terminal Lw. Panjang), Kel. Kebonlega","postal_code":"40235","country":""},{"city_id":182,"province_id":14,"district_id":2353,"address":"Dukuh Sidorejo RT 027 RW 007 Desa Ketanggan Kec. Gringsing, Kab. Batang, Jawa Tengah 51281","postal_code":"51281","country":""},{"city_id":171,"province_id":12,"district_id":2228,"address":"Jl.Pekapuran Gg.H Camat No.27 Rt/RW 02/24 Kel.Sukatani / Kec.Tapos (Depan Alfamart Bakti ABRI)","postal_code":"16454","country":""}]`

	p := parser.New()

	var arrAddress []AddressDetail

	if err := json.Unmarshal([]byte(input), &arrAddress); err != nil {
		log.Println(err)

	}
	for i := range arrAddress {
		addressJson, err := json.Marshal(arrAddress[i])
		if err != nil {
			log.Println(err)
		}
		normalizeAddress := p.CleanAddress(addressJson)
		fmt.Printf("\n%d.\nBefore : %s\nAfter : %s\n", i+1, arrAddress[i].Address, normalizeAddress)
	}

}
