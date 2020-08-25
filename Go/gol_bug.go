/*

Input Pattern : 

0 0 0 0 0 
0 0 1 0 0
0 0 1 0 0
0 0 1 0 0 
0 0 0 0 0

Output Pattern :

If the generation number is odd : | If the generation number is even :
								  |
0 0 0 0 0                         | 0 0 0 0 0
0 0 0 0 0                         | 0 0 1 0 0
0 1 1 1 0                         | 0 0 1 0 0
0 0 0 0 0                         | 0 0 1 0 0
0 0 0 0 0                         | 0 0 0 0 0
								  |
To print for debugging use :

fmt.Print("String", variable)
fmt.Println("String", variable) - adds a newline at the end

*/

package main

import (
	"fmt"
)

//main function
func main() {
	
	//Initial pattern
	var input_pattern = [5][5]int{
		{0,0,0,0,0},
		{0,0,1,0,0},
		{0,0,1,0,0},
		{0,0,1,0,0},
		{0,0,0,0,0},
	}

	//Defining the number of generations
	var number_of_generations int = 3 

	//Loop which calls the generate_next_pattern function 
	for generation:=0; generation < number_of_generations; generation++ {

	var output_pattern [5][5]int= generate_next_pattern(input_pattern)
	input_pattern = output_pattern
	
	}

	var correct bool = test()

	if(correct == false){
		fmt.Println("Wrong patterns generated")
	} else {
		fmt.Println("Tests passed")
	}
}

//This function is where the next pattern is generated. Find bugs in this function to generate the correct next pattern.
func generate_next_pattern(input_pattern [5][5]int) [5][5]int {

	//initializing new_pattern array  which will be the next generation
	new_pattern := [5][5]int{
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0}}
	
	//looping through every element in input_pattern
	for row:=0; row < input_pattern.length; row++ {
		for column:=0; column < input_pattern[row].length; column++ {
			
			var count int = 0

			//looping through neighbours to calculate sum
			for i:=(row-1); i<=(row+1); i++ {
				for j:= (column-1); j<=(column + 1); j++ {
					if (i>=0 && i<=4 && j>=0 && j<=4 && !(i==row && j==column)) {
						count = count + input_pattern[i][j]	
					} 
				}
			}
			
			if (input_pattern[row][column] == 0) {
				if (count==3 || count == 2) {
					new_pattern[row][column] = 1} else{
				}
			} else {
				if (count==3) {
					new_pattern[row][column] = 1} else{
						new_pattern[row][column] = 0
				}
			}
		}
	}

	return(new_pattern)
}	

//Test function
func test() bool {

	var cur_pattern_1 = [5][5]int{
		{0,0,0,0,0},
		{0,0,1,0,0},
		{0,0,1,0,0},
		{0,0,1,0,0},
		{0,0,0,0,0},
	}

	for i:= 1; i<=4; i++ {
		cur_pattern_1 = generate_next_pattern(cur_pattern_1)

		if(i%2 == 0){
		
			var correct_pattern_1 = [5][5]int{
				{0,0,0,0,0},
				{0,0,1,0,0},
				{0,0,1,0,0},
				{0,0,1,0,0},
				{0,0,0,0,0},
			}
	
	
			if(cur_pattern_1 == correct_pattern_1){
				continue
			} else{
				return(false)
			}
	
		} else{
	
			var correct_pattern_1 = [5][5]int{
				{0,0,0,0,0},
				{0,0,0,0,0},
				{0,1,1,1,0},
				{0,0,0,0,0},
				{0,0,0,0,0},
			}
	
			if(cur_pattern_1 == correct_pattern_1){
				continue
			} else{
				return(false)
			}
	
		}

	}

	var cur_pattern_2 = [5][5]int{
		{0,0,0,0,0},
		{0,1,1,0,0},
		{0,1,1,0,0},
		{0,0,0,0,0},
		{0,0,0,0,1},
	}

	for i:= 1; i<=2; i++ {
		cur_pattern_2 = generate_next_pattern(cur_pattern_2)

		
		
			var correct_pattern_1 = [5][5]int{
				{0,0,0,0,0},
				{0,0,1,0,0},
				{0,0,1,0,0},
				{0,0,1,0,0},
				{0,0,0,0,0},
			}
	
	
			if(cur_pattern_1 == correct_pattern_1){
				continue
			} else{
				return(false)
			}
	

	}

	return(true)
	
}