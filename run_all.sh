#!/bin/bash

set -euo pipefail

DAY_LIST=$(ls -1 -d day*/ )

start=`date +%s.%N`



for day in $DAY_LIST
do
    echo $day
    go build ${day%/}/main/*.go
done
copil_end=`date +%s.%N`


for day in $DAY_LIST
do
    ./run_${day%/}
done
end=`date +%s.%N`

echo
echo "Compilation time: $( echo "$copil_end - $start" | bc -l )s"
echo "Execution time: $( echo "$end - $copil_end" | bc -l )s"
echo "Total run time: $( echo "$end - $start" | bc -l )s"

# Remove built packages
rm -f run_day*
