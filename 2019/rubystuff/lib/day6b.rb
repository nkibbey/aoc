

def orbits(biglist, com, steps)
    children = biglist[com]
    sum = steps
    Array(children).each do | child |
        sum += orbits biglist, child, steps+1
    end
    sum
end

def pathtoplanet(biglist, curr_map, target)
    if target == "COM"
        curr_map
    else
        biglist.each do | parent, children |
            if children.index(target)
                curr_map << parent 
                pathtoplanet biglist, curr_map, parent
                break
            end
        end
        curr_map
    end
end

def lca(m1, m2)
    lowest = 9999999

    m1.each do | papa |
        m1_index =  m1.index(papa)
        if m1_index > lowest
            break
        end
        m2_index = m2.index(papa)
        if m2_index
            steps = m1_index + m2_index
            puts papa + ": " + m1_index.to_s + ", " + m2_index.to_s
            if steps < lowest
                lowest = steps
                puts lowest
            end
        end        
    end
    lowest
end

def get_instructions
    arr = []
    File.open("resources/day6.txt") do |f|
        f.each_line { |line| arr.push line.strip }
    end
    arr
end

def spacelist_maker(instructions)
    spacelist = {}
    i = 0
    for instr in instructions do
        instr = instr.split(")")
        parent_name = instr[0]
        child_name = instr[1]
        if !spacelist[parent_name]
            spacelist[parent_name] = []
        end
        spacelist[parent_name].push child_name
    end

    spacelist
end


def com_maker(spacelist)
    child_list = []
    parent_list = []
    com_list = []
    spacelist.each do | parent, children |
        children.each { | child | child_list << child }
        parent_list << parent
    end
    parent_list.each do | planet | 
        if !child_list.index(planet)
            com_list << planet
        end
    end 
    com_list
end


instructions = get_instructions
spacelist = spacelist_maker instructions
com_list = com_maker spacelist
sum = 0
for com in com_list do
    sum += orbits spacelist, com, 0
end
sum

my_map = pathtoplanet spacelist, [], "YOU"
san_map = pathtoplanet spacelist, [], "SAN"
lca my_map, san_map