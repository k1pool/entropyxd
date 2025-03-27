#!/bin/bash
rm -rf /tmp/entropyxd-temp

NUM_CLIENTS=128
entropyxd --devnet --appdir=/tmp/entropyxd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
ENTROPYXD_PID=$!
ENTROPYXD_KILLED=0
function killEntropyxdIfNotKilled() {
  if [ $ENTROPYXD_KILLED -eq 0 ]; then
    kill $ENTROPYXD_PID
  fi
}
trap "killEntropyxdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $ENTROPYXD_PID

wait $ENTROPYXD_PID
ENTROPYXD_EXIT_CODE=$?
ENTROPYXD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $ENTROPYXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ENTROPYXD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
