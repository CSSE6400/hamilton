package service

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	_ "github.com/boombuler/barcode"
	code128 "github.com/boombuler/barcode/code128"
	"image/color"
	"math"
	"strings"
)

type Drawer struct {
}

func NewDrawer() Drawer {
	return Drawer{}
}

func (d Drawer) DrawTicket(ticket Ticket) (string, error) {
	w := &strings.Builder{}

	width := 1200
	height := 518
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Desc(ticket.UUID)
	canvas.Rect(0, 0, width, height, "fill:rgb(255,255,255)")
	canvas.Rect(0, 0, 200, height, "fill:rgb(40,40,40)")
	canvas.Rect(200, 0, 20, height, "fill:rgb(255,0,0)")
	canvas.Circle(210, 100, 50, "fill:rgb(255,0,0)")
	canvas.Text(300, 100, ticket.Concert.Name, "text-anchor:left;font-size:48px;fill:black")
	canvas.Text(300, 150, ticket.Concert.Date, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 200, ticket.Concert.Venue, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 250, ticket.Concert.UUID, "text-anchor:left;font-size:24px;fill:black")
	canvas.Line(300, 300, 900, 300, "stroke:rgb(0,0,0);stroke-width:2")
	canvas.Text(300, 350, ticket.Name, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 400, ticket.Email, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 450, ticket.UUID, "text-anchor:left;font-size:24px;fill:black")
	canvas.Rect(1190, 0, 10, height, "fill:rgb(255,215,0)")
	err := d.barcode(canvas, ticket.UUID)
	if err != nil {
		return "", err
	}
	canvas.End()
	return w.String(), nil
}

func (d Drawer) DrawConcert(concert Concert) (string, error) {
	w := &strings.Builder{}

	width := 1200
	height := 2600

	canvas := svg.New(w)
	canvas.Start(width, height)

	canvas.Desc(fmt.Sprintf("%s|%d|%d", concert.UUID, concert.Seats.Max, concert.Seats.Purchased))
	canvas.Rect(0, 0, width, height, "fill:rgb(255,255,255)")
	canvas.Text(600, 100, concert.Name, "text-anchor:middle;font-size:36px;fill:black")
	canvas.Text(600, 150, concert.Venue, "text-anchor:middle;font-size:24px;fill:black")
	canvas.Line(300, 170, 900, 170, "stroke:rgb(0,0,0);stroke-width:2")

	err := d.seats(canvas, concert.Seats.Max, concert.Seats.Purchased)
	if err != nil {
		return "", err
	}

	canvas.End()
	return w.String(), nil
}

func (d Drawer) barcode(canvas *svg.SVG, uuid string) error {
	bc, err := code128.Encode(uuid)
	if err != nil {
		return err
	}
	for i := 0; i < bc.Bounds().Dx(); i++ {
		if bc.At(i, 0) == color.Black {
			canvas.Rect(1000, 80+i, 150, 1, "fill:rgb(0,0,0)")
		}
	}

	return nil
}

func (d Drawer) seats(canvas *svg.SVG, total int, purchased int) error {
	row := 0
	seat := 0
	shift := 0

	for i := 0; i < total; i++ {

		row = i / 50
		seat = (i % 50) + 1

		if seat%2 == 1 {
			seat = int(math.Ceil(float64(seat)/2.0) - 1)
			seat = seat * -1
		} else {
			seat = seat / 2
		}

		shift = (row % 3) - 1

		if i < purchased {
			canvas.CenterRect(600+(seat*20)+(shift*8), 250+row*20, 20, 20, "fill:rgb(255,0,0);stroke-width:1;stroke:rgb(0,0,0)")
			continue
		}
		canvas.CenterRect(600+(seat*20)+(shift*8), 250+row*20, 20, 20, "fill:rgb(255,255,255);stroke-width:1;stroke:rgb(0,0,0)")
	}
	return nil
}
