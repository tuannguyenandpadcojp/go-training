package domain

type ClientID string
type UserID string

type UserAttributes struct {
	ClientID        ClientID //  the main Tenant
	UserID          UserID
	CurrentClientID ClientID // the current Tenant which the user is accessing
}
