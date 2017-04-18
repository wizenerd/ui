
all:grid/grid.js footer/footer.js
grid/grid.js: grid/main.go
	gopherjs build -m -o grid/grid.js github.com/wizenerd/ui/grid

footer/footer.js: footer/main.go
	gopherjs build -m -o footer/footer.js github.com/wizenerd/ui/footer