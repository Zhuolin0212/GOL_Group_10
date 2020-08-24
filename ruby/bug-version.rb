require 'set'
=begin
Input Pattern : 

00000000000
00000000000
00011000000
00010100000
00010000000
00000000000
00000000000
00000000000
00000000000
00000000000
00000000000

Output Pattern :

00000000000
00000000000
00011000000
00011000000
00000000000
00000000000
00000000000
00000000000
00000000000
00000000000
00000000000
=end
def display(world,i)
    puts "Generation #{i+1}:"
    ###############print the map###################
    for y in (0..10) do
        for x in (0..10) do
            if world.include?([x,y])
                print "1"
            else
                print "0"
            end
        end
        puts
    end
end 

          
def location(cell,center)
    #################get the pattern and add it location############# 
    temp = Set[]
    (dx,dy) = center
    for x,y in cell
        temp.add([(x+dx),(y+dy)])
    end
    return temp
end

def generation(world,n,neighbor)
    (n-1).times do |i|
        display(world,i)
        counts = Hash.new(0)
        for c in world
            ###############count the neighbor#######################
            location(neighbor,c).each {|h| counts[h] += 1}
        end
        new_world = []
        counts.each do |key,value|
            ################## judge live or die###########################
            if counts[key] != 3 and (counts[key]!= 2 and not world.include?(key))
                new_world.push(key)
            end
        end
        ##############update map######################
        world = new_world
    end
    return world
end

#world = Set[[1,1],[2,2],[2,1],[1,2]]

#######################################
############main program###############
#######################################
glider = Set[[0,1],[1,0],[0,0],[0,2],[2,1]]
neighbor = [[-1,1],[-1,0],[-1,1],[0,-1],[0,1],[1,-1],[1,0],[1,1]]
world = location(glider,[3,2])
user_world = generation(world,4,neighbor)
correct_answer = [[3, 3], [4, 2], [4, 3], [3, 2]]
puts
puts
puts
if user_world == correct_answer
   puts "Test Pass"
else
    puts "Test fail"
end
