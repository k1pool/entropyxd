#!/bin/bash
rm -rf /tmp/entropyxd-temp

entropyxd --simnet --appdir=/tmp/entropyxd-temp --profile=6061 &
BGAPAD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $BGAPAD_PID

wait $BGAPAD_PID
BGAPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $BGAPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $BGAPAD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
