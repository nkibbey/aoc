#!/bin/bash

cat day2.txt | awk '
BEGIN {
  res["good"]=0;
  res["bad"]=0;
}

{
#  print $1
  match($1, /([0-9]*)-([0-9]*)/, crange);
#  print crange[1]
#  print crange[2]

  match($2, /([a-z]*):/, fchar);
#  print fchar[1]
  
  plen = length($3)
  print $0

  c1 = substr($3, crange[1], 1)
  c2 = substr($3, crange[2], 1)
  #print gc " " bc "\t" fchar[1]
  #print (fchar[1] == gc) " " (fchar[1] != bc) "\t" ((gc == fchar[1]) && (bc != fchar[1]))
  y1 = (fchar[1] == c1)
  y2 = (fchar[1] == c2)
  #if ((gc == fchar[1]) && (bc != fchar[1])) {
  if ((y1 || y2) && (y1 != y2)) {
    res["good"] += 1
    print "good"
  } else {
    res["bad"] += 1
    print "bad"
  }
}

END {
  print "========================================"
  print res["good"]
  print res["bad"]

}


'

