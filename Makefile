runw: main.go
	./gmauth.exe

run: main.go
	go build

	CLIENTID= REDIRECTURI= ./gmauth