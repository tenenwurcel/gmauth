runw: main.go
	./gmauth.exe

run: main.go
	go build

	CLIENTID=123123123123 REDIRECTURI=http://localhost:8080 ./gmauth