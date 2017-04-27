package main

import (
	"context"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/wizenerd/tabs"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://code.getmdl.io/1.3.0/material.teal-red.min.css")
	ctx, _ := context.WithCancel(context.Background())
	// defer cancel()
	a := newApp(ctx)
	vecty.RenderBody(a)
}

type app struct {
	vecty.Core
	t *tabs.Tabs
}

func newApp(ctx context.Context) *app {
	t := tabs.New(ctx)
	t.IsJS = true
	t.Rippled = true
	t.Panels = panes()
	return &app{t: t}
}
func (a *app) Render() *vecty.HTML {
	return elem.Body(a.t)
}

func panes() []*tabs.Panel {
	return []*tabs.Panel{
		{
			ID:       "starks-panel",
			Name:     "Starks",
			IsActive: true,
			Children: elem.UnorderedList(
				elem.ListItem(
					vecty.Text("Eddard"),
				),
				elem.ListItem(
					vecty.Text("Eddard"),
				),
				elem.ListItem(
					vecty.Text("Catelyn"),
				),
				elem.ListItem(
					vecty.Text("Robb"),
				),
				elem.ListItem(
					vecty.Text("Sansa"),
				),
				elem.ListItem(
					vecty.Text("Brandon"),
				),
				elem.ListItem(
					vecty.Text("Arya"),
				),
				elem.ListItem(
					vecty.Text("Rickon"),
				),
			),
		},
		{
			ID:   "lannisters-panel",
			Name: "Starks",
			Children: elem.UnorderedList(
				elem.ListItem(
					vecty.Text("Tywin"),
				),
				elem.ListItem(
					vecty.Text("Cersei"),
				),
				elem.ListItem(
					vecty.Text("Jamie"),
				),
				elem.ListItem(
					vecty.Text("Tyrion"),
				),
			),
		},
		{
			ID:   "targaryens-panel",
			Name: "Starks",
			Children: elem.UnorderedList(
				elem.ListItem(
					vecty.Text("Viserys"),
				),
				elem.ListItem(
					vecty.Text("Daenerys"),
				),
			),
		},
	}
}
