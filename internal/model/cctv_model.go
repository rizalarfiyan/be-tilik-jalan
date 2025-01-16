package model

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rizalarfiyan/be-tilik-jalan/config"
)

type CCTVItem struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Link      string    `json:"link"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Aspect    string    `json:"aspect"`
	Image     CCTVImage `json:"image"`
}

type CCTVImage struct {
	Src   string `json:"src"`
	Thumb string `json:"thumb"`
}

func (c *CCTVItem) FillImage() {
	conf := config.Get()
	c.Image = CCTVImage{
		Src:   conf.PublicUrl.JoinPath(fmt.Sprintf("/cctv/%s.jpg", c.Id)).String(),
		Thumb: conf.PublicUrl.JoinPath(fmt.Sprintf("/cctv/thumb/%s.jpg", c.Id)).String(),
	}
}

type CCTVs []CCTVItem
