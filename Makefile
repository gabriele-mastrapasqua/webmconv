# Makefile per webmconv

# Variabili
BINARY_NAME=webmconv
BUILD_DIR=build

# Default target
.PHONY: all
all: build

# Compila il programma
.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

# Esegue i test
.PHONY: test
test:
	go test ./...

# Esegue il programma con argomenti
.PHONY: run
run:
	go run main.go $(ARGS)

# Pulizia
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

# Installa il programma
.PHONY: install
install:
	go install .

# Aiuto
.PHONY: help
help:
	@echo "Comandi disponibili:"
	@echo "  make build    - Compila il programma"
	@echo "  make test     - Esegue i test"
	@echo "  make run      - Esegue il programma (usa ARGS=\"...\" per passare argomenti)"
	@echo "  make clean    - Rimuove i file generati"
	@echo "  make install  - Installa il programma"
	@echo "  make help     - Mostra questo messaggio"