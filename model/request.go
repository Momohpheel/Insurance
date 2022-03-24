package model

type MarineRequest struct {
	Currency string `binding:"required"`
	Amount   uint64 `binding:"required"`
}

type AllRiskRequest struct {
	Name          string
	SerialNumber  string
	Configuration string
	SumInsured    uint64
}

type GoodsInTransitRequest struct {
	Name                    string
	Address                 string
	Occupation              string
	Phone                   string
	Email                   string
	TypeOfGoods             string
	NatureOfGoods           string
	LimitAnyoneCarrying     uint32
	EstimatedAnnualCarrying uint64
	ModeOfTransaporting     string
	OwnVehicles             bool
	RegisterationNumber     string
}

type ShuttlersRequest struct {
	FullName         string
	Address          string
	Business         string
	MaritalStatus    string
	DateOfBirth      string
	PhoneNumber      uint64
	Email            string
	NextOfKin        string
	NextOfKinAddress string
	NextOfKinPhone   string
}
