build:
	GOARCH=arm64 go build -o ./bin/rasp-web ./cmd
run-background:
	nohup ./rasp-web > /dev/null 2>&1&
find-process:
	pgrep rasp-web
kill-process:
	kill 12345678
