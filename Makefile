
all:grid/grid.js footer/footer.js icons/icons.js tabs/tabs.js
grid/grid.js: grid/main.go
	gopherjs build -m -o grid/grid.js github.com/wizenerd/ui/grid

footer/footer.js: footer/main.go
	gopherjs build -m -o footer/footer.js github.com/wizenerd/ui/footer

icons/icons.js: icons/main.go
	gopherjs build -m -o icons/icons.js github.com/wizenerd/ui/icons

tabs/tabs.js: tabs/main.go
	gopherjs build -m -o tabs/tabs.js github.com/wizenerd/ui/tabs