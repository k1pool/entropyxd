#!/bin/bash
rm -rf /tmp/entropyxd-temp

NUM_CLIENTS=128
entropyxd --devnet --appdir=/tmp/entropyxd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
BGAPAD_PID=$!
BGAPAD_KILLED=0
function killEntropyxdIfNotKilled() {
  if [ $BGAPAD_KILLED -eq 0 ]; then
    kill $BGAPAD_PID
  fi
}
trap "killEntropyxdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $BGAPAD_PID

wait $BGAPAD_PID
BGAPAD_EXIT_CODE=$?
BGAPAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $BGAPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $BGAPAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
