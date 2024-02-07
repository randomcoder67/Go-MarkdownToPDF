normal:
	go mod tidy
	go build -o mdtogroff main.go definitions.go helpers.go markdownParser.go renderers.go
	
clean:
	rm mdtogroff

install:
	cp mdtogroff ~/.local/bin

uninstall:
	rm ~/.local/bin/mdtogroff

full: normal install clean
