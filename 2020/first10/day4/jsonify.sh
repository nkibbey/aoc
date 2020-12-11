#!/bin/bash

cat samp.txt | awk '
BEGIN {
    print "[ "
    nl = 1
    // a hacked way to deal with trailing commas
    sillyVal = " \"silly\":\"val\" "
}
{
    if (nl) {
        print "  { "
        nl = 0
    } 
    if (!NF) { 
        print sillyVal "  }, "
        nl=1
    } else {
        s = ""
        for (i=1; i <= NF; i++) {
            match($i, /([a-z]*):(.*)/, fchar);
            s = s " \"" fchar[1] "\":\"" fchar[2] "\", "

        }
    
        print s
    }
}
END {
    print sillyVal "  }"
    print "]"
}

'