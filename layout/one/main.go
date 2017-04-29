package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/wizenerd/layout"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://code.getmdl.io/1.3.0/material.teal-red.min.css")
	vecty.RenderBody(&app{})
}

type app struct {
	vecty.Core
}

func (app) Render() *vecty.HTML {
	l := layout.New(false)
	l.Header = &layout.Header{
		Icon: &layout.Icon{},
		Row: &layout.HeaderRow{
			Title: &layout.Title{
				Text: "Title",
			},
			AddSpacer: true,
			Nav: &layout.Nav{
				Links: []*layout.NavLink{
					{
						Text: "Link",
					},
					{
						Text: "Link",
					},
					{
						Text: "Link",
					},
					{
						Text: "Link",
					},
					{
						Text: "Link",
					},
				},
			},
		},
	}
	l.Drawer = &layout.Drawer{
		Title: &layout.Title{
			Text: "Title",
		},
		Nav: &layout.Nav{
			Links: []*layout.NavLink{
				{
					Text: "Link",
				},
				{
					Text: "Link",
				},
				{
					Text: "Link",
				},
				{
					Text: "Link",
				},
				{
					Text: "Link",
				},
			},
		},
	}
	return elem.Body(l)
}
