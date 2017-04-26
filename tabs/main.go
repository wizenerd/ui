package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/wizenerd/tabs"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
	vecty.AddStylesheet("https://code.getmdl.io/1.3.0/material.teal-red.min.css")
	vecty.RenderBody(&app{})
}

type app struct {
	vecty.Core
}

func (app) Render() *vecty.HTML {
	return elem.Body(
		&tabs.Tabs{
			IsJS:    true,
			Rippled: true,
			Panels:  panes()},
	)
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
			Children: vecty.List{
				elem.Paragraph(
					vecty.Text(a1),
				),
				elem.Paragraph(
					vecty.Text(a2),
				),
			},
		},
		{
			ID:   "#members-panel",
			Name: "Members",
			Children: vecty.List{
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
			},
		},
	}
}
