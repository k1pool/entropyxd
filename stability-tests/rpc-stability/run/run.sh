#!/bin/bash
rm -rf /tmp/entropyxd-temp

entropyxd --devnet --appdir=/tmp/entropyxd-temp --profile=6061 --loglevel=debug &
ENTROPYXD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $ENTROPYXD_PID

wait $ENTROPYXD_PID
ENTROPYXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $ENTROPYXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ENTROPYXD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
