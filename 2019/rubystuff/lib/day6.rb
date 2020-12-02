
class Space_Body
    attr_reader :parent, :my_name, :children
    attr_accessor :all_children

    def initialize(my_name, parent=nil)
        @parent = parent # parent_spacebody
        @my_name = my_name # just the name
        @children = []
        @all_children = []
    end

    def ==(other_SB)
        if other_SB.class == String
            @my_name == other_SB
        elsif other_SB.class == Space_Body
            @my_name == other_SB.my_name
        end
    end

    def orbits
        num = 0
        papa = @parent
        while papa
            papa = papa.parent
            num +=1
        end
        num
    end

    def add_child(child_sb)
        if children.index(child_sb)
            puts 'oyyy'
            false
        else
            @children << child_sb
            @all_children.push child_sb.my_name
            papa = @parent
            while papa
                papa.all_children.push child_sb.my_name
                papa = papa.parent
            end
            true
        end
    end

    def inspect
        @my_name + " " + all_children.length.to_s
    end
end

def build_a_body(curr, parent_name, child_name) 
    if curr == parent_name # is current
        puts 'is current'
        nieto = Space_Body.new child_name, curr
        curr.add_child nieto
        true
    elsif curr.children.index(parent_name) # is direct child of current
        puts 'is direct'
        c_index = curr.children.index(parent_name)
        curr = curr.children[c_index]
        nieto = Space_Body.new child_name, curr
        curr.add_child nieto
        true
    elsif curr.all_children.index(parent_name) # indirect child of current
        puts 'indirect'
        inserted = false
        i = 0
        while (!inserted) && (i < curr.all_children.length)
            inserted = build_a_body(curr.children[i], parent_name, child_name)
            i += 1
        end
        if !inserted
            puts 'big problem'
        end
        inserted
    else
        false
    end
end 


def get_instructions
    arr = []
    File.open("resources/day6b.txt") do |f|
        f.each_line { |line| arr.push line.strip }
    end
    arr
end

def spacelist_maker(instructions)
    spacelist = []
    i = 0
    for instr in instructions do
        instr = instr.split(")")
        parent_name = instr[0]
        child_name = instr[1] # could do a sanity check
        inserted = false
        # puts 'working on ' + parent_name + ' and child ' + child_name
        for curr_sb in spacelist do # center of masses
           inserted = build_a_body(curr_sb, parent_name, child_name) 
           break if inserted
        end
        if !inserted 
            com = Space_Body.new parent_name
            child = Space_Body.new child_name, com
            com.add_child child
            spacelist << com # inserted = true
        end
        # i += 1
        # break if i == 679
    end

    spacelist
end

def stack_the_orbits(curr)
    stacked = 0
    for child in curr.children do
        stacked += stack_the_orbits(child)
    end
    stacked + curr.orbits
end

def grand_orbits(spacelist)
    grand_total = 0
    for body in spacelist do
        grand_total += stack_the_orbits body
    end
    grand_total
end

some = Space_Body.new "s"
body = Space_Body.new "b", some
some.add_child body

instructions = get_instructions
spacelist = spacelist_maker instructions
grand_orbits spacelist











