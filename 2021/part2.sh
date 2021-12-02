#!/bin/bash
cat in.txt | awk 'BEGIN { a = 0; b = 0; c =0; ups= -3; } { ps = a+b+c; a = b; b = c; c = $1; cs = a+b+c; if (cs > ps) { ups++; }  } END { print ups } '
