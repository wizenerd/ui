package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/style"
	"github.com/wizenerd/color"
	"github.com/wizenerd/grid"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://code.getmdl.io/1.3.0/material.teal-red.min.css")
	a := New(gridExamples()...)
	vecty.RenderBody(a)
}

type App struct {
	vecty.Core
	grids []*Example
}

func New(e ...*Example) *App {
	return &App{
		grids: e,
	}
}

func (a *App) Render() *vecty.HTML {
	var v []vecty.MarkupOrComponentOrHTML
	for _, e := range a.grids {
		v = append(v, e)
	}
	return elem.Body(v...)
}

type Example struct {
	vecty.Core
	View vecty.MarkupOrComponentOrHTML
	Code string
}

func (e *Example) Render() *vecty.HTML {
	return elem.Div(
		elem.Div(
			e.View,
		),
		elem.Div(
			&grid.Grid{
				Cells: []*grid.Cell{
					{
						Mode:     grid.Default,
						Size:     12,
						Children: elem.Code(elem.Preformatted(vecty.Text(e.Code))),
					},
				},
			},
		),
	)
}

func gridExamples() []*Example {
	return []*Example{
		&Example{
			View: grid1(),
			Code: grid1Txt,
		},
	}
}

var cellColor = color.Class(color.Blue, color.A400)

func cellChild(m ...vecty.MarkupOrComponentOrHTML) vecty.MarkupOrComponentOrHTML {
	var v []vecty.MarkupOrComponentOrHTML
	v = append(v, m...)
	v = append(v, style.Height(style.Px(50)))
	v = append(v, style.Color("white"))
	v = append(v, vecty.Style("background-color", "#BDBDBD"))
	v = append(v, vecty.Style("box-sizing", "border-box"))
	v = append(v, vecty.Style("padding-leftr", string(style.Px(8))))
	v = append(v, vecty.Style("padding-top", string(style.Px(4))))
	return vecty.List(v)
}

func grid1() *grid.Grid {
	return &grid.Grid{
		Cells: []*grid.Cell{
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
			{
				Mode:     grid.Default,
				Size:     1,
				Children: cellChild(vecty.Text("1")),
			},
		},
	}
}

var grid1Txt = `

`
