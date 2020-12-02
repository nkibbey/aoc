

class Day1
    def self.hi
        puts "Hello world!!!"
    end
    def fuely(mass)
        return 0 if mass < 9
        fuel_req = mass/3-2
        return fuel_req + fuely(fuel_req) 
    end
```
Opcode 1 adds together numbers read from two positions and stores the result in a third position. 
The three integers immediately after the opcode tell you these three positions - 
    the first two indicate the positions from which you should read the input values, 
    and the third indicates the position at which the output should be stored.

Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them. 
Again, the three integers after the opcode indicate where the inputs and outputs are, not their values.

Once you're done processing an opcode, move to the next one by stepping forward 4 positions.



```
    def self.ejemplo
        arr = []
        f=File.open("resources/day1.txt") or die "no bueno"
        f.each_line { |line| arr.push line.strip }
        arr.map(&:to_i)
        
        arr.inject(0){ |sum, n| sum + fuely(n) }
    end
    def opcodidor(op_type, index_a, index_b, index_res, arr)
        case op_type
        when 1
            arr[index_res] = arr[index_a] + arr[index_b]
        when 2
            arr[index_res] = arr[index_a] * arr[index_b]
        when 99
            puts 'im done'
        else 
            puts 'ahhh'
        end
        return arr
    end
    def operadador(arr) 
        sz = arr.length
        pos = 0
        while pos+3 < sz && arr[pos] != 99
            opcodidor(arr[pos], arr[pos+1], arr[pos+2], arr[pos+3], arr)
            # puts 'round  ' + (pos/4).to_s
            pos+=4
        end

        return arr
    end

    def find_noun_verb(arr, noun, verb, desired)
        arr[1] = noun
        arr[2] = verb
        operadador(arr)
        if arr[0] == desired
            puts 'the answer is: ' + (100 * noun + verb).to_s 
        end       
    end

    def searchy(original, desired)
        for i in 1..99 do
            for j in 1..99 do
                testy = original.clone
                find_noun_verb(testy, i, j, desired)
                #puts 'tried ' + i.to_s + '\t' + j.to_s
            end
        end
    end

    def day2
        realboy = [1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,19,5,23,1,9,23,27,2,27,6,31,1,5,31,35,2,9,35,39,2,6,39,43,2,43,13,47,2,13,47,51,1,10,51,55,1,9,55,59,1,6,59,63,2,63,9,67,1,67,6,71,1,71,13,75,1,6,75,79,1,9,79,83,2,9,83,87,1,87,6,91,1,91,13,95,2,6,95,99,1,10,99,103,2,103,9,107,1,6,107,111,1,10,111,115,2,6,115,119,1,5,119,123,1,123,13,127,1,127,5,131,1,6,131,135,2,135,13,139,1,139,2,143,1,143,10,0,99,2,0,14,0]
        testboy = [1,0,0,0,99]
        practico = realboy.clone
        opcodidor(testboy[0], testboy[1], testboy[2], testboy[3], testboy)
        searchy(practico, 19690720)
    end
end