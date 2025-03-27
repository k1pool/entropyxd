#!/bin/bash
rm -rf /tmp/entropyxd-temp

entropyxd --devnet --appdir=/tmp/entropyxd-temp --profile=6061 --loglevel=debug &
ENTROPYXD_PID=$!
ENTROPYXD_KILLED=0
function killEntropyxdIfNotKilled() {
    if [ $ENTROPYXD_KILLED -eq 0 ]; then
      kill $ENTROPYXD_PID
    fi
}
trap "killEntropyxdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ENTROPYXD_PID

wait $ENTROPYXD_PID
ENTROPYXD_KILLED=1
ENTROPYXD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $ENTROPYXD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ENTROPYXD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
