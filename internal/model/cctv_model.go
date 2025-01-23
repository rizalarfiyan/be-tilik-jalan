package model

import (
	"fmt"

	"github.com/google/uuid"
)

type CCTVItem struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Thumbnail string    `json:"thumbnail"`
}

func (c *CCTVItem) FillImage() {
	c.Thumbnail = fmt.Sprintf("cctv/thumb/%s.jpg", c.Id)
}

type CCTVs []CCTVItem
