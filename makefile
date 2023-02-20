
todo=${PWD}/app/domain/todo/http

test:
	go test ${svc} -v

test-all:
	go test ./... -v -cover

test-todo:
	make test svc=${todo}