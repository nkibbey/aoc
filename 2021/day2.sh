#!/bin/bash
echo "part 1"
awk ' BEGIN { c[1] = "forward"; c[2] = "down"; c[3] = "up"; h = 0; d = 0; }
{
 if (c[1] == $1) { h += $2; } 
 else if (c[2] == $1) { d += $2; } 
 else if (c[3] == $1) { d -= $2; }
 else { print "dont know " $0; }
}
END {
 print "depth = " d;
 print "horiz = " h;
 print d * h;
} ' in1.txt

printf "\npart 2\n"
awk ' BEGIN { c[1] = "forward"; c[2] = "down"; c[3] = "up"; h = 0; d = 0; a=0; }
{
 if (c[1] == $1) { h += $2; d += a*$2; } 
 else if (c[2] == $1) { a += $2; } 
 else if (c[3] == $1) { a -= $2; }
 else { print "dont know " $0; }
}
END {
 print "depth = " d;
 print "horiz = " h;
 print d * h;
} ' in1.txt

