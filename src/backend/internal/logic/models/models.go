package models

import (
	"github.com/google/uuid"
)

const (
	Customer int = iota
	Admin
)

const PlacedOrder = "placed"
const ProcessOrder = "in processing"

const PaidBill = "paid"
const NotPaidBill = "not paid"

type User struct {
	ID       uuid.UUID
	Login    string
	Password string
	Fio      string
	Email    string
	Points   int
	Status   int
}

type Order struct {
	ID         uuid.UUID
	IDUser     uuid.UUID
	TotalPrice int
	IsPoints   bool
	Status     string
}

type OrderElement struct {
	ID      uuid.UUID
	IDOrder uuid.UUID
	IDWine  uuid.UUID
	Count   int
}

type Bill struct {
	ID      uuid.UUID
	IDOrder uuid.UUID
	Price   int
	Status  string
}

type Wine struct {
	ID       uuid.UUID
	Name     string
	Count    int
	Year     int
	Strength int
	Price    int
	Type     string
}

type UserWine struct {
	IDUser uuid.UUID
	IDWine uuid.UUID
}
