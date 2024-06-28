LGOROOT=$$HOME/go

gen-godoc:
	- mkdir -p ~/go/src
	- rm ${LGOROOT}/src/$$(basename $$(pwd))
	ln -s $$(pwd) ${LGOROOT}/src/
	godoc -http=:6060 -goroot=${LGOROOT} &
	RUNNING_PID=$$!
	- wget -np -k -p -q -r -E -P ./docs http://localhost:6060/pkg/analyze/
	kill $${RUNNING_PID}

make run:
	go run cmd/main.go
