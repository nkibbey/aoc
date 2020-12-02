#!/bin/bash

cat day2.txt | awk '
BEGIN {
  res["good"]=0;
  res["bad"]=0;
  found[""]=0;
}

{
#  print $1
  match($1, /([0-9]*)-([0-9]*)/, crange);
#  print crange[1]
#  print crange[2]

  match($2, /([a-z]*):/, fchar);
#  print fchar[1]
  
  found[fchar[1]] = 0
  plen = length($3)
#  print $3
 
  for (i=1; i<=plen; i++) {
    found[substr($3, i, 1)] += 1
  }
  
  findTotal = found[fchar[1]]
  if (findTotal >= crange[1] && findTotal <= crange[2]) {
    res["good"] += 1
  } else {
    res["bad"] += 1
  }
}

END {
  print "========================================"
  print res["good"]
  print res["bad"]

}


'

