


def good_code(a, b, c, d, e, f)
    if a > b || b > c || c > d || d > e || e > f
        false
    end
    double = false
    if (a==b && a!=c)
        double = true
    elsif (a!=b &&  b==c && b!=d)
        double = true
    elsif (b!=c && c==d && c!=e)
        double = true
    elsif (c!=d && d==e && d!=f)
        double = true
    elsif (d!=e && e == f)
        double = true
    end

    return double

        
end


def test_range(starty)

    goods = 0

    for a in starty..5 do
        for b in a..9 do
            for c in b..9 do
                for d in c..9 do
                    for e in d..9 do 
                        for f in e..9 do
                            if good_code a,b,c,d,e,f
                                goods+=1
                            end
                        end
                    end
                end
            end
        end
    end
    goods
end
good_code 2,2,2,2,2,2

test_range(2)