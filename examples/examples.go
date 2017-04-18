package examples

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/wizenerd/grid"
)

type Example struct {
	vecty.Core
	View vecty.MarkupOrComponentOrHTML
	Code string
}

func (e *Example) Render() *vecty.HTML {
	c := make(vecty.ClassMap)
	c["language-go"] = true

	return elem.Div(
		elem.Div(
			e.View,
		),
		elem.Div(
			&grid.Grid{
				Cells: []*grid.Cell{
					{
						Mode: grid.Default,
						Size: 12,
						Children: elem.Preformatted(
							elem.Code(
								prop.Class("language-go"),
								vecty.Text(e.Code)),
						),
					},
				},
			},
		),
	)
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
