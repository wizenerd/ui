package examples

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
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
						Children: elem.Code(
							c, elem.Preformatted(vecty.Text(e.Code)),
						),
					},
				},
			},
		),
	)
}
