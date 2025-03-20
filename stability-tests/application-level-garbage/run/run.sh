#!/bin/bash
rm -rf /tmp/entropyxd-temp

entropyxd --devnet --appdir=/tmp/entropyxd-temp --profile=6061 --loglevel=debug &
ENXPAD_PID=$!
ENXPAD_KILLED=0
function killEntropyxdIfNotKilled() {
    if [ $ENXPAD_KILLED -eq 0 ]; then
      kill $ENXPAD_PID
    fi
}
trap "killEntropyxdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ENXPAD_PID

wait $ENXPAD_PID
ENXPAD_KILLED=1
ENXPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $ENXPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ENXPAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
