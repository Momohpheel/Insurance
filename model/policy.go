package model

import "gorm.io/gorm"

type Marine struct {
	Value          uint64  `json:"value"`
	Rate           float64 `json:"rate"`
	TotalInvoice   float64 `json:"total_invoice"`
	RatePercentage float64 `json:"rate_percentage"`
	gorm.Model
}

type GoodInTransit struct {
	Name                    string  `json:"name"`
	Address                 string  `json:"address"`
	Occupation              string  `json:"occupation"`
	Phone                   string  `json:"phone"`
	Email                   string  `json:"email"`
	TypeOfGoods             string  `json:"type_of_goods"`
	NatureOfGoods           string  `json:"nature_of_goods"`
	LimitAnyoneCarrying     uint32  `json:"limit_anyone_carrying"`
	EstimatedAnnualCarrying uint64  `json:"estimated_annual_carying"`
	ModeOfTransaporting     string  `json:"mode_of_transporting"`
	OwnVehicles             bool    `json:"own_vehicles"`
	RegisterationNumber     string  `json:"registeration_number"`
	Rate                    float64 `json:"rate"`
	RatePercentage          float64 `json:"rate_percentage"`
	gorm.Model
}

type AllRisk struct {
	Name           string  `json:"name"`
	SerialNumber   string  `json:"serial_number"`
	Configuration  string  `json:"configuration"`
	SumInsured     uint64  `json:"sum_insured"`
	Rate           float64 `json:"rate"`
	RatePercentage float64 `json:"rate_percentage"`
	gorm.Model
}

type Shuttlers struct {
	FullName         string `json:"full_name"`
	Address          string `json:"address"`
	Business         string `json:"business"`
	MaritalStatus    string `json:"marital_status"`
	DateOfBirth      string `json:"dob"`
	PhoneNumber      uint64 `json:"phone_number"`
	Email            string `json:"email"`
	NextOfKin        string `json:"next_of_kin"`
	NextOfKinAddress string `json:"next_of_kin_address"`
	NextOfKinPhone   string `json:"next_of_kin_phone"`
	gorm.Model
}
