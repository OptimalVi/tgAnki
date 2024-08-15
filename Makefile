# Name of the Go executable
# BINARY := myapp
BOT := cmd/bot/bot.go
# Go build flags
# BUILD_FLAGS := -ldflags="-s -w"


# Build the Go executable
# build:
# 	go build $(BUILD_FLAGS) -o $(BINARY)

# Run the Go executable
bot-run: 
	go run -gcflags=all="-N -l" ./$(BOT)


# Clean up the built executable
# clean:
# 	rm -f $(BINARY)

