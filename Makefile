windows:
	@echo "Building for windows"
	GOOS=windows GOARCH=386 go build -o ./bin/windows/apollo.exe
linux:
	@echo "Building for linux"
	go build -o ./bin/linux/apollo
	cd bin/linux && ./apollo
all:
	@echo "Building for every OS and Platform"
	GOOS=windows GOARCH=386 go build -o ./bin/windows/apollo.exe
	GOOS=linux GOARCH=386 go build -o ./bin/linux/apollo
	GOOS=freebsd GOARCH=386 go build -o ./bin/freebsd/apollo
	GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/apollo
	@echo "Zipping for release"
	@tar -czf bin/releases/apollo_linux.tar.gz LICENSE -C bin/linux apollo
	@tar -czf bin/releases/apollo_win.tar.gz LICENSE -C  bin/windows apollo.exe
	@tar -czf bin/releases/apollo_mac_amd64.tar.gz LICENSE -C bin/mac apollo 
	@tar -czf bin/releases/apollo_bsd.tar.gz LICENSE -C bin/freebsd apollo 
run:
	go run .
global:
	go install .

# do not use
release:
	gh release create $v 'bin/releases/apollo_linux.tar.gz' 'bin/releases/apollo_win.tar.gz' 'bin/releases/apollo_bsd.tar.gz' 'bin/releases/apollo_mac_amd64.tar.gz' 

