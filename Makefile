GENERATOR = ./generator
PALETTE_NAME := $(PALETTE)

templates: alacritty kakoune

build-generator: $(GENERATOR)

$(GENERATOR): cmd/generator/*.go
	go build -o $@ ./cmd/generator

gen/$(PALETTE_NAME)/%: templates/%.tmpl $(PALETTE) $(GENERATOR)
	@mkdir -p gen/$(PALETTE_NAME)
	$(GENERATOR) $(PALETTE) < templates/$(notdir $*).tmpl > $@


alacritty: gen/$(PALETTE_NAME)/alacritty-theme.toml

kakoune: gen/$(PALETTE_NAME)/kakoune-theme.kak

.PHONY: alacritty kakoune
