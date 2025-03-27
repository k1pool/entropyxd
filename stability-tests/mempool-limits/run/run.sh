#!/bin/bash

APPDIR=/tmp/entropyxd-temp
ENTROPYXD_RPC_PORT=29587

rm -rf "${APPDIR}"

entropyxd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${ENTROPYXD_RPC_PORT}" --profile=6061 &
ENTROPYXD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${ENTROPYXD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $ENTROPYXD_PID

wait $ENTROPYXD_PID
ENTROPYXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $ENTROPYXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ENTROPYXD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
