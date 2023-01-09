package model

const UsernameHeader = "X-User-Name"

type GetFlightsParams struct {
	Page *float32 `form:"page,omitempty" json:"page,omitempty"`
	Size *float32 `form:"size,omitempty" json:"size,omitempty"`
}

type UsernameParam struct {
	// Имя пользователя
	XUserName string `json:"X-User-Name"`
}

type GetMeParams = UsernameParam

type GetPrivilegeParams = UsernameParam

type GetTicketsParams = UsernameParam

type PostTicketsJSONBody = TicketPurchaseRequest

type PostTicketsParams = UsernameParam

type DeleteTicketsTicketUidParams = UsernameParam

type GetTicketsTicketUidParams = UsernameParam

// PostTicketsJSONRequestBody defines body for PostTickets for application/json ContentType.
type PostTicketsJSONRequestBody = PostTicketsJSONBody
