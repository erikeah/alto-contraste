GENERATOR = ./generator
PALETTE_NAME := $(PALETTE)

default: light dark

templates: alacritty kakoune editor

light:
	$(MAKE) PALETTE_NAME=light PALETTE=./light_palette.json templates

dark:
	$(MAKE) PALETTE_NAME=dark PALETTE=./dark_palette.json templates

build-generator: $(GENERATOR)

$(GENERATOR): cmd/generator/*.go
	go build -o $@ ./cmd/generator

gen/$(PALETTE_NAME)/%: templates/% $(PALETTE) $(GENERATOR)
	$(eval TEMP := $(shell mktemp))
	@mkdir -p gen/$(PALETTE_NAME)
	$(GENERATOR) $(PALETTE) < templates/$(notdir $*) > $(TEMP)
	mv $(TEMP) $@

alacritty: gen/$(PALETTE_NAME)/alacritty-theme.toml

kakoune: gen/$(PALETTE_NAME)/kakoune-theme.kak

editor: gen/$(PALETTE_NAME)/editor.html

.PHONY: $(templates) templates light
