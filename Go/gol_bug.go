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

*/


package main

import (
	"fmt"
)

//Check whether a dead cell will be dead or alive in the next generation
func check_dead(left int, right int, up int, down int, left_up int, left_down int, right_up int, right_down int) int {

	var count int

	count = left + right + up + down + left_up + left_down + right_up + right_down

	if count==3{
		return(1)} else{
			return(0)
	}

}

//check whether an alive cell will be dead or alive in the next generation
func check_alive(left int, right int, up int, down int, left_up int, left_down int, right_up int, right_down int) int {

	var count int

	count = left + right + up + down + left_up + left_down + right_up + right_down

	if count==2 || count==3{
		return(1)} else{
			return(0)
	}

}

//This function is where the next pattern is generated. Find a bug in this function to generate the correct next pattern.
func generate_next_pattern(input_pattern [5][5]int) [5][5]int {

	new_pattern := [5][5]int{
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0},
		{0,0,0,0,0}}
	
		var left int 
		var right int
		var up int
		var down int
		var left_up int
		var left_down int
		var right_up int
		var right_down int
	
	for row:=0; row <= len(new_pattern); row++ {
		for column:=0; column <= len(new_pattern[row]); column++ {
			
			if row-1 < 0 {
				up = 0
				left_up = 0
				right_up = 0
			} else {

				up = input_pattern[row-1][column]

				if column-1 < 0 {
					left_up = 0
				} else {
					left_up = input_pattern[row-1][column -1]}

				if column+1 > 4 {
					right_up = 0
				} else {
					right_up = input_pattern[row-1][column +1]}}
			
			if row+1 > 4 {
				down = 0
				right_down = 0
				left_down = 0
			} else {

				down = input_pattern[row+1][column]

				if column-1 < 0 {
					left_down = 0
				} else {
					left_down = input_pattern[row+1][column-1]}

				if column+1>4 {
					right_down = 0
				} else {
					right_down = input_pattern[row+1][column+1]}}
			
			if column-1 < 0 {
				left = 0
			} else {
				left = input_pattern[row][column-1]}

			if column+1 > 4 {
				right = 0
			} else {
				right = input_pattern[row][column+1]}

			

			if input_pattern[row][column]==0 {
				new_pattern[row][column] = check_alive(left, right, up, down, left_up, left_down, right_up, right_down)} else {
					new_pattern[row][column] = check_dead(left, right, up, down, left_up, left_down, right_up, right_down)}
		}
	}

	return(new_pattern)
}

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
		
		fmt.Println("Generation: ", generation+1)

		var correct bool = test(output_pattern, generation+1)

		fmt.Println()

		if(correct == false){
			break
		} else {

		}
		
		input_pattern = output_pattern
	}


}	

//Test function
func test(pattern [5][5]int, generation int) bool {
	if(generation%2 == 0){
		
		var correct_pattern = [5][5]int{
			{0,0,0,0,0},
			{0,0,1,0,0},
			{0,0,1,0,0},
			{0,0,1,0,0},
			{0,0,0,0,0},
		}


		if(pattern == correct_pattern){

			for i:=0; i<5; i++ {
				for j:=0; j<5; j++ {
					fmt.Print(pattern[i][j])
				}
				fmt.Println();
			}
	
			return(true)
		} else{
	
			fmt.Println("Pattern you generated:")
			for i:=0; i<5; i++ {
				for j:=0; j<5; j++ {
					fmt.Print(pattern[i][j])
				}
				fmt.Println();
			}
			fmt.Println("Correct Pattern:")
			for i:=0; i<5; i++ {
				for j:=0; j<5; j++ {
					fmt.Print(correct_pattern[i][j])
				}
				fmt.Println();
			}
			return(false)
		}

	} else{

		var correct_pattern = [5][5]int{
			{0,0,0,0,0},
			{0,0,0,0,0},
			{0,1,1,1,0},
			{0,0,0,0,0},
			{0,0,0,0,0},
		}

		if(pattern == correct_pattern){

			for i:=0; i<5; i++ {
				for j:=0; j<5; j++ {
					fmt.Print(pattern[i][j])
				}
				fmt.Println();
			}
	
			return(true)
		} else{
	
			fmt.Println("Pattern you generated:")
			for i:=0; i<5; i++ {
				for j:=0; j<5; j++ {
					fmt.Print(pattern[i][j])
				}
				fmt.Println();
			}
			fmt.Println("Correct Pattern:")
			for i:=0; i<5; i++ {
				for j:=0; j<5; j++ {
					fmt.Print(correct_pattern[i][j])
				}
				fmt.Println();
			}
			return(false)
		}

	}

	
}