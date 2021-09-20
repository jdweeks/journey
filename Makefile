SRC        := ./...
TARGET 	   := www
GO         := /usr/bin/env go
ALL 	   := clean lint build test

.PHONY: $(ALL) pretty run

build:
	$(GO) build -o $(TARGET) $(SRC)

clean:
	$(GO) clean; \
	$(RM) $(TARGET);

run:
	$(GO) run $(SRC)

test:
	$(GO) test -v $(SRC)

pretty:
	$(GO) fmt $(SRC)

lint:
	$(GO) vet $(SRC)

tidy:
	$(GO) mod tidy

all: $(ALL)
