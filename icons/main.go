package main

import (
	"strings"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/wizenerd/grid"
	icon "github.com/wizenerd/icons"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://code.getmdl.io/1.3.0/material.teal-red.min.css")
	vecty.AddStylesheet("https://google.github.io/material-design-icons/www/css/material.css")
	vecty.RenderBody(&app{})
}

type app struct {
	vecty.Core
}

func (a *app) Render() *vecty.HTML {
	var list vecty.List
	for _, v := range breakNames() {
		g := &grid.Grid{}
		for _, c := range v {
			println(c)
			g.Cells = append(g.Cells,
				&grid.Cell{
					Children: elem.Div(
						prop.Class("icons-preview-code"),
						elem.Div(
							prop.Class("icons-preview"),
							&icon.Icon{Name: c},
						),
						prop.Class("icons-preview-code"),
						elem.Div(
							prop.Class("icons-code"),
							elem.Code(
								vecty.Text("&icon.Icon{Name:"+properName(c)+"}"),
							),
						),
					),
					Size: 2,
				},
			)
		}
		list = append(list, g)
	}

	return elem.Body(
		elem.Div(list),
	)
}

func properName(n icon.Name) string {
	p := strings.Split(string(n), "_")
	for i := 0; i < len(p); i++ {
		if p[i][0] == '3' {
			p[i] = "ThreeDRotation"
		} else {
			p[i] = strings.Title(p[i])
		}
	}
	return "icon." + strings.Join(p, "")
}
func breakNames() [][]icon.Name {
	size := len(names)
	var c [][]icon.Name
	i := 0
stop:
	for i < size {
		var v []icon.Name
		for k := 0; k < 4; k++ {
			if k+i >= size {
				break stop
			}
			v = append(v, names[k+i])
			i++
		}
		c = append(c, v)
		println(len(v))
	}
	return c
}
