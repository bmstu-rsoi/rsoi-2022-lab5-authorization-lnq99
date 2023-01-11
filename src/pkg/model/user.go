package model

type UserInfoResponse struct {
	Privilege PrivilegeShortInfo `json:"privilege,omitempty"`

	// Информация о билетах пользоватлея
	Tickets []TicketResponse `json:"tickets,omitempty"`
}
