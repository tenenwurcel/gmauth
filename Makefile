runw: main.go
	./gmauth.exe

run: main.go
	go build

	CLIENTID=840058304461406228 REDIRECTURI=https://rel-gmauth-tenenwurcel.cloud.okteto.net/api/v1/auth/callback ./gmauth