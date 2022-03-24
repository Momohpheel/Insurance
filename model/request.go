package model

type MarineRequest struct {
	Currency string `validate:"required"`
	Amount   uint64 `validate:"required"`
}

type AllRiskRequest struct {
	Name          string `validate:"required"`
	SerialNumber  string `validate:"required"`
	Configuration string `validate:"required"`
	SumInsured    uint64 `validate:"required"`
}

type GoodsInTransitRequest struct {
	Name                    string `validate:"required"`
	Address                 string `validate:"required"`
	Occupation              string `validate:"required"`
	Phone                   uint64 `validate:"required"`
	Email                   string `validate:"required,email"`
	TypeOfGoods             string `validate:"required"`
	NatureOfGoods           string `validate:"required"`
	LimitAnyoneCarrying     uint32 `validate:"required"`
	EstimatedAnnualCarrying uint64 `validate:"required"`
	ModeOfTransaporting     string `validate:"required"`
	OwnVehicles             bool
	RegisterationNumber     string
}

type ShuttlersRequest struct {
	FullName         string `validate:"required"`
	Address          string `validate:"required"`
	Business         string `validate:"required"`
	MaritalStatus    string `validate:"required"`
	DateOfBirth      string `validate:"required"`
	PhoneNumber      uint64 `validate:"required"`
	Email            string `validate:"required,email"`
	NextOfKin        string `validate:"required"`
	NextOfKinAddress string `validate:"required"`
	NextOfKinPhone   uint64 `validate:"required"`
}
