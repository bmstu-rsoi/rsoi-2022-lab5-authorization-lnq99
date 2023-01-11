package model

import "time"

type Airport struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type Flight struct {
	ID            int32     `json:"id"`
	FlightNumber  string    `json:"flightNumber"`
	Datetime      time.Time `json:"datetime"`
	FromAirportID int32     `json:"fromAirportID"`
	ToAirportID   int32     `json:"toAirportID"`
	Price         int32     `json:"price"`
}

type FlightResponse struct {
	// Дата и время вылета
	Date time.Time `json:"date,omitempty"`

	// Номер полета
	FlightNumber string `json:"flightNumber,omitempty"`

	// Страна и аэропорт прибытия
	FromAirport string `json:"fromAirport,omitempty"`

	// Стоимость
	Price int32 `json:"price,omitempty"`

	// Страна и аэропорт прибытия
	ToAirport string `json:"toAirport,omitempty"`
}

type PaginationResponse struct {
	Items []FlightResponse `json:"items,omitempty"`

	// Номер страницы
	Page int32 `json:"page,omitempty"`

	// Количество элементов на странице
	PageSize int32 `json:"pageSize,omitempty"`

	// Общее количество элементов
	TotalElements int32 `json:"totalElements,omitempty"`
}
