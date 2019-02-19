.PHONY: all update

setup:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

update:
	dep ensure -update -v

