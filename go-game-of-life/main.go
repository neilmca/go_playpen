package main //this is the name of the package

import (
	"fmt" //import a package named "fmt"
    "strconv"
)

func main() {
	fmt.Println("Go - Game of Life")


	/*
	Simple version that creates a 2D bounded grid as a slice
	*/

	grid  :=[][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, false, true, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}

	generations_to_run := 5 //set number of generations to play out

	grid_size := len(grid[0]) //we assue square!

	print_results(grid, 0)

	

	for g:=0; g < generations_to_run; g++ {
		
		//loop through the grid and see what will be dead in the next generation

		//Any live cell with fewer than two live neighbours dies, as if by underpopulation.
        //Any live cell with two or three live neighbours lives on to the next generation.
		//Any live cell with more than three live neighbours dies, as if by overpopulation. 
		//Any dead cell with exactly three neighbours will come alive
		
		

		grid_next:= make([][]bool, grid_size)      
		deep_copy(grid_next, grid)

		
		for i :=0; i < grid_size; i++ {
			for j :=0; j < grid_size; j++ {
		//for i :=2; i <= 2; i++ {
		//	for j :=2; j <= 2; j++ {

				neighbouring_live_count :=0
				//look in all directions - start at bottom left
				if i-1 >= 0 {
					if j-1 >= 0 {
						if grid[i-1][j-1] == true {
							neighbouring_live_count += 1
						}
					}
					
					if grid[i-1][j] == true {
						neighbouring_live_count += 1
					}
					
					if j+1 < grid_size {
						if grid[i-1][j+1] == true {
							neighbouring_live_count += 1
						}
					}
				}
				

				if j-1 >= 0 {
					if grid[i][j-1] == true {
						neighbouring_live_count += 1
					}
				}
				
				if j+1 < grid_size {
					if grid[i][j+1] == true {
						neighbouring_live_count += 1
					}
				}
				
				if i+1 < grid_size {
					if j-1 >= 0 {
						if grid[i+1][j-1] == true {
							neighbouring_live_count += 1
						}
					}
					
					if grid[i+1][j] == true {
						neighbouring_live_count += 1
					}
					
					if j+1 < grid_size {
						if grid[i+1][j+1] == true {
							neighbouring_live_count += 1
						}
					}
				}


				//fmt.Printf("%d %d %s neighbouring_live_count = %d \n", i, j, strconv.FormatBool( grid[i][j]), neighbouring_live_count)
				
				if grid[i][j] == true && neighbouring_live_count > 3 {
					//it dies
					grid_next[i][j] = false
				} else if grid[i][j] == true && neighbouring_live_count < 2 {
					//its dies
					grid_next[i][j] = false
				}
				if grid[i][j] == false && neighbouring_live_count == 3 {
					//it comes alive
					grid_next[i][j] = true
				}

				//fmt.Println("-----")
				//print_grid_with_hghlight(grid_next, i, j)
				//fmt.Println("-----")

				//fmt.Println("-----")
				//print_grid_with_hghlight(grid, i,j)
				//fmt.Println("-----")
				
			}
		}

	

		//swap grids
		deep_copy(grid, grid_next)

		print_results(grid, g+1)
		
	}
	

	
}

func deep_copy(dest [][]bool, src [][]bool){
	//need to manually do this as built-in copy makes shallow copy of rows when used with 2D slice
	for i := range src {
		dest[i] = make([]bool, len(src[i]))
		copy(dest[i], src[i])
	}
}

func print_results(grid [][]bool, generation int){
	fmt.Printf("After %d generations life is....\n", generation)
	print_grid(grid)
	fmt.Println("\n")
}

func print_grid_with_hghlight(grid [][]bool, x, y int)  {

	
	var line string
	for  i := range grid {
		line = ""
		for j := range grid[i] {
			highlighter := ""
			if i == x && j == y {
				highlighter = "!"
			}
			line += " " + highlighter + strconv.FormatBool( grid[i][j]) + highlighter
		   
		}
		fmt.Printf("%s\n", line )
	 }
	 
}

func print_grid(grid [][]bool)  {

	
	var line string
	for  i := range grid {
		line = ""
		for j := range grid[i] {
			line += " " + strconv.FormatBool( grid[i][j])
		   
		}
		fmt.Printf("%s\n", line )
	 }
	 
}
