#!/bin/bash

APPDIR=/tmp/entropyxd-temp
ENXPAD_RPC_PORT=29587

rm -rf "${APPDIR}"

entropyxd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${ENXPAD_RPC_PORT}" --profile=6061 &
ENXPAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${ENXPAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $ENXPAD_PID

wait $ENXPAD_PID
ENXPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $ENXPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ENXPAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
