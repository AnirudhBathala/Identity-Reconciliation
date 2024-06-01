package models

import "time"

type Contact struct {
	ID             int    `json:"id"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	LinkedId       int    `json:"linked_id"`
	LinkPrecedence string `json:"link_precedence"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type IndentityPayload struct{
	Email string 
	PhoneNumber string 
}

type ServiceStore interface{
	
}