#!/bin/bash

set -euo pipefail

# Remove built packages
rm -f run_day*

DAY_LIST=$(ls -1 -d day*/ )

for day in $DAY_LIST
do
    echo $day
    go build ${day%/}/main/*.go
done

for day in $DAY_LIST
do
    ./run_${day%/}
done
