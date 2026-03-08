#!/bin/bash
# HTTP API Test Runner
# Usage: ./scripts/test-api.sh [test-name]

cd "$(dirname "$0")/.."

# Start chatlog server in background
echo "Starting chatlog server..."
./chatlog server -d ~/Library/Containers/com.tencent.xinWeChat/Data/Library/Application\ Support/com.tencent.xinWeChat/2.0.0/ -p darwin -v 3 &
SERVER_PID=$!

# Wait for server to start
sleep 3

# Run tests
if [ -z "$1" ]; then
    echo "Running all HTTP API tests..."
    go test ./internal/chatlog/http/... -v
else
    echo "Running test: $1"
    go test ./internal/chatlog/http/... -v -run "$1"
fi

TEST_RESULT=$?

# Cleanup
echo "Stopping chatlog server..."
kill $SERVER_PID 2>/dev/null

exit $TEST_RESULT
