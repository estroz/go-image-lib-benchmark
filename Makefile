all: bench

TAGS ?= -tags containers_image_openpgp
bench:
	$(if $(V),,@)CGO_ENABLED=0 go test -bench=. -benchmem -benchtime=$(if $(BT),$(BT),10s) $(TAGS) $(if $(V),-v,)

cli:
	$(if $(V),,@)CGO_ENABLED=0 go build $(TAGS) -o bin/$@ .
