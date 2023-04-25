#!/bin/sh
#

rm -f for.txt
touch for.txt

for i in `seq 5000`
do
  echo -n $i >> for.txt
  ./main >> for.txt
done
