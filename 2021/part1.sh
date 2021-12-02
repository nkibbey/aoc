#!/bin/bash
awk 'BEGIN { prev = 99999999; ups= 0; } { if ($1 > prev) { ups++; } prev = $1; } END { print ups } ' in.txt
