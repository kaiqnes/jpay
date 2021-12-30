package model

type Customer struct {
	Id    uint
	Name  string
	Phone string
}

func (Customer) TableName() string {
	return "customer"
}
