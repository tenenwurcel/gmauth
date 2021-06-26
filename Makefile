runw: main.go
	./gmauth.exe

run: main.go
	go build

	CLIENTID=840058304461406228 REDIRECTURI=https://tenenwurcel.com.br/api/v1/auth/callback ./gmauth