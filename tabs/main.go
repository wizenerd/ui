package main

import (
	"context"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/wizenerd/tabs"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
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
	t.Panels = panes()
	return &app{t: t}
}
func (a *app) Render() *vecty.HTML {
	return elem.Body(a.t)
}

var a1 = `
			<b>The Beatles</b> were a four-piece musical group from Liverpool, England.
    Formed in 1960, their career spanned just over a decade, yet they are widely
    regarded as the most influential band in history.
`
var a2 = `
Their songs are among the best-loved music of all time. It is said that every
    minute of every day, a radio station somewhere is playing a Beatles song.
`

func panes() []*tabs.Panel {
	return []*tabs.Panel{
		{
			ID:       "#about-panel",
			Name:     "About the Beatles",
			IsActive: true,
			Children: elem.Div(
				elem.Paragraph(
					vecty.Text(a1),
				),
				elem.Paragraph(
					vecty.Text(a2),
				),
			),
		},
		{
			ID:   "#members-panel",
			Name: "Members",
			Children: elem.Div(
				elem.Paragraph(vecty.Text("The Beatles' members were:")),
				elem.UnorderedList(
					vecty.List{
						elem.ListItem(
							vecty.Text("John Lennon (1940-1980)"),
						),
						elem.ListItem(
							vecty.Text("Paul McCartney (1942-)"),
						),
						elem.ListItem(
							vecty.Text("George Harrison (1943-2001)"),
						),
						elem.ListItem(
							vecty.Text("Ringo Starr (1940-)"),
						),
					},
				),
			),
		},
	}
}
