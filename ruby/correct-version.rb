# Hello World Program in Ruby
require 'set'
#file = File.open('input.txt')
def display(world,i)
    puts "Generation #{i+1}:"
    x_list = []
    y_list = []
    world.each do |item|
       x_list.push(item[0])
       y_list.push(item[1])
    end
    ###############print the map###################
    for y in (y_list.min-3..y_list.max+3) do
        for x in (x_list.min-3..x_list.max+3) do
            if world.include?([x,y])
                print "*"
            else
                print "."
            end
        end
        puts
    end
end 


neighbor = [[-1,1],[-1,0],[-1,1],[0,-1],[0,1],[1,-1],[1,0],[1,1]]
          
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
    n.times do |i|
        display(world,i)
        counts = Hash.new(0)
        for c in world
            ###############count the neighbor#######################
            location(neighbor,c).each {|h| counts[h] += 1}
        end
        new_world = []
        counts.each do |key,value|
            ################## judge live or die###########################
            if counts[key] == 3 or (counts[key]== 2 and world.include?(key))
                new_world.push(key)
            end
        end
        ##############update map######################
        world = new_world
    end
end

#######################################
############main program###############
#######################################
glider = Set[[0,1],[1,0],[0,0],[0,2],[2,1]]
world = location(glider,[3,2])
generation(world,10,neighbor)

