TARGET=demo

define message
	@echo "### $(1)"
endef

all: build run

build: $(TARGET)

demo: $(shell find . -name '*.go')
	$(call message,Building $@)
	go build -o $@

run:
	$(call message,Run $@)
	@./$(TARGET)

clean:
	rm -f $(TARGET)

.PHONY: all build run clean
