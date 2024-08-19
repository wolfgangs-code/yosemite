#?	Yosemite Makefile
#	for both the frontend and the backend

#? 	Variables
BIN_DIR := bin
# 	Binary paths
FRONTEND_BIN := $(BIN_DIR)/frontend
BACKEND_BIN := $(BIN_DIR)/backend
#	Source paths
FRONTEND_SRC := cmd/frontend/main.go
BACKEND_SRC := cmd/backend/main.go

#?	Build targets
all: frontend backend
frontend: $(FRONTEND_BIN)
backend: $(BACKEND_BIN)

#?	Build rules
#	Frontend
$(FRONTEND_BIN): $(FRONTEND_SRC)
	@mkdir -p $(BIN_DIR)
	@go build -o $(FRONTEND_BIN) $(FRONTEND_SRC)
	@echo "$(FRONTEND_BIN) created"
#	Backend
$(BACKEND_BIN): $(BACKEND_SRC)
	@mkdir -p $(BIN_DIR)
	@go build -o $(BACKEND_BIN) $(BACKEND_SRC)
	@echo "$(BACKEND_BIN) created"

#?	Clean targets
clean:
	@rm -rf $(BIN_DIR)
	@echo "Cleaned up binaries"

#	Phony targets
.PHONY: all frontend backend clean