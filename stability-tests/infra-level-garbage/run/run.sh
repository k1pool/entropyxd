#!/bin/bash
rm -rf /tmp/entropyxd-temp

entropyxd --devnet --appdir=/tmp/entropyxd-temp --profile=6061 &
ENXPAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $ENXPAD_PID

wait $ENXPAD_PID
ENXPAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Entropyxd exit code: $ENXPAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $ENXPAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
