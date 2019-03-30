all: cmangos-website

cmangos-website:
	go get github.com/gorilla/mux
	go get gopkg.in/ini.v1
	go build

clean:
	rm -rf cmangos-website
