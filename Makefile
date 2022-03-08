all: bench

TAGS ?= -tags containers_image_openpgp
bench:
	@GOBIN=$(shell pwd)/bin go install github.com/cespare/prettybench@latest
	$(if $(V),,@)CGO_ENABLED=0 go test -bench=. -benchmem -benchtime=$(if $(BT),$(BT),100x) $(TAGS) $(if $(V),-v,) | ./bin/prettybench

cli:
	$(if $(V),,@)CGO_ENABLED=0 go build $(TAGS) -o bin/$@ .
