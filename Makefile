html:
	cd rpq &&\
	go test -cover -coverprofile=c.out &&\
	go tool cover -html=c.out -o coverage.html &&\
	cd .. &&\
	firefox rpq/coverage.html
