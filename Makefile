debug: gotbyond.go
	go build -o gotbyond -buildmode=c-shared gotbyond.go

release: clean gotbyond.go
	go build -ldflags=-w -o gotbyond -buildmode=c-shared gotbyond.go
	
release-windows: gotbyond.go
	- rm gotbyond.dll
	go build -ldflags=-w -o gotbyond.dll -buildmode=c-shared gotbyond.go

clean:
	- rm gotbyond