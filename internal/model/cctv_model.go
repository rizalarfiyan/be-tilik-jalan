package model

import "github.com/google/uuid"

type CCTVItem struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Aspect    string    `json:"aspect"`
	Image     string    `json:"image"`
}

type CCTVs []CCTVItem
