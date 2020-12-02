
```
segA=p1,p2
segB=p3,p4
want to find point of intersect
p1+tA(p2-p1) = p3+tB(p4-p3)
in xy
x1+tA(x2-x1)=x3+tB(x4-x3)
y1+tA(y2-y1)=y3+tB(y4-y3)

x1-x3 = tB (x4-x3) -tA(x2-x1)

begin
tA= (0.0 + (y3-y4)*(x1-x3) + (x4-x3)*(y1-y3)) /
        ((x4-x3)*(y1-y2) - (x1-x2)*(y4-y3))

tB= (0.0 + (y1-y2)*(x1-x3) + (x2-x1)*(y1-y3)) /
        ((x4-x3)*(y1-y2) - (x1-x2)*(y4-y3))

if (tA>=0 && tA<=1) && (tB>=0 && tB<=1)
    h_dist = x1+tA(x2-x1)
    v_dist = y1+tA(y2-y1)
else
    puts 'didnt intersect'
end

rescue ZeroDivisionError # colinear handling
    puts 'oy'

end

```



class Spot
    def initialize(curr_direction, curr_distance, h_ds=0, v_ds=0)
        case curr_direction
        when "R"
            h_ds+=curr_distance
        when "L"
            h_ds-=curr_distance
        when "U"
            v_ds+=curr_distance
        when "D"
            v_ds-=curr_distance
        else
            puts 'asdfasdfasdfasdfasdfasdfasdf'
        end

        @h_ds = h_ds
        @v_ds = v_ds
    end
    
    # attr accessors
    def h_ds
        @h_ds
    end
    
    def v_ds
        @v_ds
    end
    
    def ==(other_spot) 
        other_spot.h_ds == @h_ds && other_spot.v_ds == @v_ds
    end

    # provide quick next child
    def next_spot(next_dir, next_dist) 
        Spot.new next_dir, next_dist, @h_ds, @v_ds
    end

    def distance_to_home
        @v_ds.abs() + @h_ds.abs()
    end

    def inspect
        "h" + @h_ds.to_s + "v" + @v_ds.to_s
    end
end


def spotfinder(spot_list1, spot_list2)
    lowest = 99999999 # should check but i no care
    for s1 in spot_list1 do
        for s2 in spot_list2 do
            if s1 == s2 && s1.distance_to_home < lowest
                lowest = s1.distance_to_home
                puts 'lowest'
            end
        end
    end
    lowest
end

def spotlist_maker(direction_arr) 
    spotlist = []
    for direction in direction_arr do
        dir = direction[0]
        dist = direction[1..-1].to_i

        if spotlist.length == 0
            spotlist.push(Spot.new dir, dist )
        else
            spotlist.push(spotlist[-1].next_spot dir, dist)
        end
    end
    spotlist
end

# bob = Spot.new "R", 100
# bill = bob.next_spot "R", 200
# joe = bill.next_spot "U", 15
# jim = Spot.new "U", 15
# joejoe = jim.next_spot "R", 300

# spot_list1 = [bob, bill, joe, (joe.next_spot "D",4)]
# spot_list2 = [jim, joejoe]

# spotfinder spot_list1, spot_list2
# z = ["R300", "R30"]
# spotlist_maker z

def min_point(a, b) 
    if a.abs() < b.abs()
        a
    else
        b
    end
end



class Wire
    def initialize(x1, y1, x2=0, y2=0, offset=0)
        @x1 = x1
        @y1 = y1
        @x2 = x2
        @y2 = y2
        @offset = offset
    end
    def inspect
        "p1(" + @x1.to_s + ", " + @y1.to_s + "); p2("  + @x2.to_s + ", " + @y2.to_s + ") "
    end
    def x1
        @x1
    end
    def y1
        @y1
    end
    def x2
        @x2
    end
    def y2
        @y2
    end
    def offset 
        @offset
    end
    def offset_topoint(x, y)
        @offset + (@x1 - x).abs() + (@y1 - y).abs()
    end
    def closest_dist
        min_x = 9999
        min_y = 9999
        if (@x1 < 0 && @x2 > 0) || (@x1 > 0 && @x2 < 0)
            min_x = 0
        else
            min_x = min_point @x1, @x2
        end

        if (@y1 < 0 && @y2 > 0) || (@y1 > 0 && @y2 < 0)
            min_y = 0
        else 
            min_y = min_point @y1, @y2
        end
        min_x + min_y
    end

    
    def intersect_dist(other_wire)
        dist = 99999999999
        # if other_wire.is_vertical # check to see if our x intercepts it
        #     if (@x.1 <= other_wire.x1 && @x.2 >= other_wire.x1)  
        #     end
        # end
        x1=@x1
        x2=@x2
        x3=other_wire.x1
        x4=other_wire.x2

        y1=@y1
        y2=@y2
        y3=other_wire.y1
        y4=other_wire.y2
        denominator =  ((x4-x3)*(y1-y2) - (x1-x2)*(y4-y3))
        
        if denominator != 0 
            tA = (0.0 + (y3-y4)*(x1-x3) + (x4-x3)*(y1-y3)) / denominator
            tB = (0.0 + (y1-y2)*(x1-x3) + (x2-x1)*(y1-y3)) / denominator
                
            if (tA>=0 && tA<=1) && (tB>=0 && tB<=1)
                x_dist = x1+tA*(x2-x1)
                y_dist = y1+tA*(y2-y1)
                dist = self.offset_topoint(x_dist, y_dist) 
                dist += other_wire.offset_topoint(x_dist, y_dist)
                # dist = x_dist.abs() + y_dist.abs()
                puts dist.to_s + 'point x ' + x_dist.to_s + ' y ' + y_dist.to_s 
            else
                # puts 'didnt intersect ' + tA.to_s + " " + tB.to_s
            end
        else 
            puts 'colinear me ' + self.inspect + ' and ' + other_wire.inspect
        end

        dist
    end
end

def wirelist_maker(direction_arr) 
    wirelist = []
    prev_x = 0
    prev_y = 0
    curr_x = prev_x
    curr_y = prev_y
    offset = 0
    for direction in direction_arr do

        dir = direction[0]
        dist = direction[1..-1].to_i
        case dir
        when "R"
            curr_x+=dist
        when "L"
            curr_x-=dist
        when "U"
            curr_y+=dist
        when "D"
            curr_y-=dist
        else
            puts 'asdfasdfasdfasdfasdfasdfasdf'
        end

        wirelist.push Wire.new prev_x, prev_y, curr_x, curr_y, offset
        offset += dist
        prev_x = curr_x
        prev_y = curr_y
    end
    wirelist
end

def compare_boys(wirelist1, wirelist2)
    curr_low = 9999999
    
    # for wire1 in wirelist1 do 
    #     for wire2 in wirelist2 do
    #         if wire1.closest_dist < curr_low && wire2.closest_dist < curr_low
    #             curr_low = min_point curr_low, (wire1.intersect_dist wire2)
    #         end
    #     end
    # end
    for wire1 in wirelist1 do 
        for wire2 in wirelist2 do
            if (wire1.offset + wire2.offset) < curr_low
                curr_low = min_point curr_low, (wire1.intersect_dist wire2)
            end
        end
    end

    curr_low

end


# seg1 = Wire.new 0,0,20,0
# seg2 = Wire.new 5,0,5,10
# seg3 = Wire.new 0,10,10,10
# seg1.intersect_dist seg1

# seg4 = Wire.new 10,-10,10,20

# z = ["R300", "R30"]
# wirelist_maker z


dirlist1 = "R990,U796,R784,U604,R6,U437,L96,U285,L361,U285,L339,D512,L389,D840,L425,U444,L485,D528,L262,U178,L80,U2,R952,U459,L361,D985,R56,U135,R953,D913,L361,U120,L329,U965,L294,U890,L126,U214,R232,D444,L714,U791,R888,U923,R378,U233,L654,D703,R902,D715,R469,D60,R990,U238,R755,U413,L409,D601,R452,U504,R472,D874,L766,D594,R696,U398,R593,D889,R609,D405,L962,U176,L237,U642,L393,D91,L463,U936,R199,D136,R601,D8,R359,D863,L410,U598,L444,D34,R664,D323,R72,D98,L565,D476,L197,D132,R510,U665,R936,U3,R385,U144,L284,D713,L605,U106,R543,D112,R528,D117,R762,U330,R722,U459,L229,U375,L870,D81,R623,U95,L148,D530,L622,D62,R644,D365,L214,U847,R31,D832,L648,D293,R79,D748,L270,U159,L8,U83,R195,U912,L409,D649,L750,D286,L623,D956,R81,U775,R44,D437,L199,U698,L42,U419,L883,U636,L323,U89,L246,D269,L992,U739,R62,U47,R63,U17,L234,U135,R126,D208,L69,U550,L123,D66,R463,U992,R411,D276,L851,U520,R805,D300,L894,U171,L922,D901,R637,U907,R328,U433,L316,D644,L398,U10,L648,D190,R884,U474,R397,D718,L925,D578,R249,U959,L697,D836,R231,U806,R982,U827,R579,U830,L135,D666,R818,D502,L898,D585,R91,D190,L255,U535,R56,U390,R619,D815,L300,D81,R432,D70,L940,D587,L259,D196,R241,U4,R440,U678,R185,U451,R733,D984,R464,D298,L738,U600,R353,D44,L458,U559,L726,D786,L307,D333,L226,D463,R138,D142,L521,D201,R51,D202,L204,U130,L333,U597,R298,U42,L951,U66,R312,U707,L555,D225,L360,D12,L956,D361,L989,D625,L944,D398,L171,D982,L377,U114,L339,U164,R39,D793,R992,U834,R675,U958,R334,D697,L734,D40,L149,U394,R976".split(",")
wirelist1 = wirelist_maker dirlist1

dirlist2 = "L1005,D52,L125,U787,L761,U262,L466,D966,R895,U789,R6,U2,R870,U971,R238,D946,L752,D240,R721,U349,L679,D518,L104,U417,L462,U544,L519,U797,R873,U70,R298,U45,L779,D921,R468,D421,R803,U108,L812,D498,R226,D309,R766,U724,L961,U472,R940,U944,R418,D682,R328,U55,R737,U961,L343,U397,R112,D292,L155,U162,R398,U445,L524,U256,R323,D587,L862,D726,R624,D230,R460,U539,R723,U93,L507,U608,L150,U159,R35,U458,R208,U546,L495,D835,L636,U960,L322,U408,L78,D250,L994,U818,R107,U978,R401,D147,R574,D549,R983,U698,L99,D63,L772,U409,R975,U990,L893,U467,L860,D721,R504,U102,R678,D672,L406,D933,R743,D788,R142,D44,R208,D424,R28,D674,R331,D968,L154,U206,R222,D354,R687,D331,L539,D390,L373,D514,L622,U673,R345,U943,L508,D337,R265,D785,L189,U429,R344,D719,R622,U199,L765,U350,R833,U309,R95,U911,R548,U746,R107,D867,L648,D680,R28,U596,L891,U168,R933,U571,R365,U267,R916,D130,L149,D898,L513,D167,R587,U799,R134,D328,R562,D929,L399,U568,R565,U241,L395,U822,L624,D145,L995,U516,R474,D609,R153,U52,R561,D15,R283,U321,L850,U218,L225,D635,L630,U102,L84,D672,L128,D885,L506,U911,R355,D132,R155,D120,L110,U368,R149,D343,L708,U378,R591,D585,L381,D517,R852,U432,R342,U273,R893,D277,L548,U859,L891,U311,L901,U255,R421,U90,L72,D474,L654,U12,L146,D867,L485,D663,R123,D82,L21,U408,L38,D864,L114,D645,R936,U765,L832,D668,L482,U79,L594,U276,L559,D469,R314,D162,R621,U230,L688,U82,R605,U191,L381,D327,L91,D217,R714,D942,R851,U358,R22,U952,R279,D897,R485,D867,L940,U891,L223,D264,L634,D560,R645,D705,L289,U584,R97,U920,R41".split(",")
wirelist2 = wirelist_maker dirlist2

compare_boys(wirelist1, wirelist2)


