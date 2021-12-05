#!/bin/bash

set -euo pipefail

DAY_LIST=$(ls -1 -d day*/ )

for day in $DAY_LIST
do
    echo $day
    go build ${day%/}/main/*.go
done

for day in $DAY_LIST
do
    time ./run_${day%/}
done

# Remove built packages
rm -f run_day*
