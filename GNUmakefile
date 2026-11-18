all: check test

check: vet

vet:
	go vet $(CURDIR)/...

test:
	go test -race -count 1 $(CURDIR)

.PHONY: all check vet test
