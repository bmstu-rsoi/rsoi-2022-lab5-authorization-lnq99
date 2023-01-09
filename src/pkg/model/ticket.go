package model

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID           int32     `json:"id"`
	TicketUid    uuid.UUID `json:"ticketUid"`
	Username     string    `json:"username"`
	FlightNumber string    `json:"flightNumber"`
	Price        int32     `json:"price"`
	Status       string    `json:"status"`
}

type TicketPurchaseStatus string

const (
	TicketPurchaseStatusCANCELED TicketPurchaseStatus = "CANCELED"
	TicketPurchaseStatusPAID     TicketPurchaseStatus = "PAID"
)

type TicketPurchaseRequest struct {
	// Номер полета
	FlightNumber string `json:"flightNumber,omitempty"`

	// Выполнить списание бонусных баллов в учет покупки билета
	PaidFromBalance bool `json:"paidFromBalance,omitempty"`

	// Стоимость
	Price int32 `json:"price,omitempty"`
}

type TicketPurchaseResponse struct {
	// Время вылета
	Date time.Time `json:"date"`

	// Номер полета
	FlightNumber string `json:"flightNumber"`

	// Страна и аэропорт прибытия
	FromAirport string `json:"fromAirport"`

	// Сумма оплаченная бонусами
	PaidByBonuses int32 `json:"paidByBonuses"`

	// Сумма оплаченная деньгами
	PaidByMoney int32 `json:"paidByMoney"`

	// Стоимость
	Price int32 `json:"price"`

	Privilege PrivilegeShortInfo `json:"privilege"`

	// Статус билета
	Status TicketPurchaseStatus `json:"status"`

	// UUID билета
	TicketUid string `json:"ticketUid"`

	// Страна и аэропорт прибытия
	ToAirport string `json:"toAirport"`
}

type TicketResponse struct {
	// Дата и время вылета
	Date time.Time `json:"date,omitempty"`

	// Номер полета
	FlightNumber string `json:"flightNumber,omitempty"`

	// Страна и аэропорт прибытия
	FromAirport string `json:"fromAirport,omitempty"`

	// Стоимость
	Price int32 `json:"price,omitempty"`

	// Статус билета
	Status TicketPurchaseStatus `json:"status,omitempty"`

	// UUID билета
	TicketUid string `json:"ticketUid,omitempty"`

	// Страна и аэропорт прибытия
	ToAirport string `json:"toAirport,omitempty"`
}
