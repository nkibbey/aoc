#!/bin/bash

cat s.txt | sort -n | awk ' BEGIN { prev = 0; diffs[""]=0 } { d = $1 - prev; prev = $1; diffs[d]++; } END { print "1s: " diffs[1]; print "2s: " diffs[2]; print "3s: " diffs[3] + 1; print "ans " diffs[1] * (diffs[3] + 1)}'
