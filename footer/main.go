package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/wizenerd/footer"
	"github.com/wizenerd/ui/examples"
)

func main() {
	vecty.AddStylesheet("https://fonts.googleapis.com/css?family=Roboto:400,300,500|Roboto+Mono|Roboto+Condensed:400,700&subset=latin,latin-ext")
	vecty.AddStylesheet("https://fonts.googleapis.com/icon?family=Material+Icons")
	vecty.AddStylesheet("https://code.getmdl.io/1.3.0/material.teal-red.min.css")
	a := examples.New(footerExamples()...)
	vecty.RenderBody(a)
}

func footerExamples() []*examples.Example {
	return []*examples.Example{
		&examples.Example{
			View: mega(),
			Code: megaText,
		},
	}
}

func mega() *footer.Footer {
	return &footer.Footer{
		Kind: footer.Mega,
		Children: vecty.List{
			&footer.Section{
				Kind: footer.Mega,
				Top:  true,
				Children: vecty.List{
					&footer.Section{
						Kind: footer.Mega,
						Left: true,
						Children: vecty.List{
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("1"),
							},
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("2"),
							},
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("3"),
							},
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("4"),
							},
						},
					},
					&footer.Section{
						Kind:  footer.Mega,
						Right: true,
						Children: vecty.List{
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Introduction"),
							),
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("App Status Dashboard"),
							),
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Terms of Service"),
							),
						},
					},
				},
			},
			&footer.Section{
				Kind:   footer.Mega,
				Middle: true,
				Children: vecty.List{
					&footer.DropDown{
						Kind: footer.Mega,
						Children: vecty.List{
							footer.H1(
								footer.Mega,
								vecty.Text("Learning and Support"),
							),
							&footer.LinkList{
								Kind: footer.Mega,
								Children: vecty.List{
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Resource Center"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Help Center"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Community"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Learn with Google"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Small Business Community"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Think Insights"),
										),
									),
								},
							},
						},
					},
					&footer.DropDown{
						Kind: footer.Mega,
						Children: vecty.List{
							footer.H1(
								footer.Mega,
								vecty.Text("Just for Developers"),
							),
							&footer.LinkList{
								Kind: footer.Mega,
								Children: vecty.List{
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Google Developers"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("AdWords API"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("AdWords Scipts"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("AdWords Remarketing Tag"),
										),
									),
								},
							},
						},
					},
				},
			},
			&footer.Section{
				Kind:   footer.Mega,
				Bottom: true,
				Children: vecty.List{
					&footer.Logo{
						Children: vecty.Text(" More Information"),
					},
					&footer.LinkList{
						Kind: footer.Mega,
						Children: vecty.List{
							elem.ListItem(
								elem.Anchor(
									prop.Href("#"),
									vecty.Text("Help"),
								),
							),
							elem.ListItem(
								elem.Anchor(
									prop.Href("#"),
									vecty.Text("Privacy and Terms"),
								),
							),
						},
					},
				},
			},
		},
	}
}

var megaText = `
func mega() *footer.Footer {
	return &footer.Footer{
		Kind: footer.Mega,
		Children: vecty.List{
			&footer.Section{
				Kind: footer.Mega,
				Top:  true,
				Children: vecty.List{
					&footer.Section{
						Kind: footer.Mega,
						Left: true,
						Children: vecty.List{
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("1"),
							},
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("2"),
							},
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("3"),
							},
							&footer.SociatButton{
								Kind:     footer.Mega,
								Children: vecty.Text("4"),
							},
						},
					},
					&footer.Section{
						Kind:  footer.Mega,
						Right: true,
						Children: vecty.List{
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Introduction"),
							),
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("App Status Dashboard"),
							),
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Terms of Service"),
							),
						},
					},
				},
			},
			&footer.Section{
				Kind:   footer.Mega,
				Middle: true,
				Children: vecty.List{
					&footer.DropDown{
						Kind: footer.Mega,
						Children: vecty.List{
							footer.H1(
								footer.Mega,
								vecty.Text("Learning and Support"),
							),
							&footer.LinkList{
								Kind: footer.Mega,
								Children: vecty.List{
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Resource Center"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Help Center"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Community"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Learn with Google"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Small Business Community"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Think Insights"),
										),
									),
								},
							},
						},
					},
					&footer.DropDown{
						Kind: footer.Mega,
						Children: vecty.List{
							footer.H1(
								footer.Mega,
								vecty.Text("Just for Developers"),
							),
							&footer.LinkList{
								Kind: footer.Mega,
								Children: vecty.List{
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("Google Developers"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("AdWords API"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("AdWords Scipts"),
										),
									),
									elem.ListItem(
										elem.Anchor(
											prop.Href("#"),
											vecty.Text("AdWords Remarketing Tag"),
										),
									),
								},
							},
						},
					},
				},
			},
			&footer.Section{
				Kind:   footer.Mega,
				Bottom: true,
				Children: vecty.List{
					&footer.Logo{
						Children: vecty.Text(" More Information"),
					},
					&footer.LinkList{
						Kind: footer.Mega,
						Children: vecty.List{
							elem.ListItem(
								elem.Anchor(
									prop.Href("#"),
									vecty.Text("Help"),
								),
							),
							elem.ListItem(
								elem.Anchor(
									prop.Href("#"),
									vecty.Text("Privacy and Terms"),
								),
							),
						},
					},
				},
			},
		},
	}
}
`
