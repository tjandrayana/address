package main

type AddressDetail struct {
	CityID     int64  `json:"city_id"`
	ProvinceID int64  `json:"province_id"`
	DistrictID int64  `json:"district_id"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}
