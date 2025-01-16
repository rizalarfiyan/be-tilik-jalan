package model

import "github.com/google/uuid"

type CCTVItem struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title" db:"title"`
	Link      string    `json:"link" db:"link"`
	Latitude  float64   `json:"latitude" db:"latitude"`
	Longitude float64   `json:"longitude" db:"longitude"`
}

type CCTVs []CCTVItem
