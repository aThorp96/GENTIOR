package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"sort"

	. "github.com/athorp96/graphs"
)

// if somehow I need to change how this is checked, this provides an interface
type fitnessEvaluator func(*Undirected, []int) int

// A rerecombination is a function that somehow constructs a child Hamiltonian
// from two other Hamiltonians.
type recombination func(*Hamiltonian, *Hamiltonian, *Undirected) *Hamiltonian

// Hamiltonian is a hamiltonian cycle.
// it consists of a cycle and a fitness grade
// the lower the fitness, the shorter the cycle
type Hamiltonian struct {
	// the path
	path []int

	// The optimitality of the solution
	// The lower the number the shorter the path
	fitness int
}


func main() {
	// default values
	defualtPopulation := 50
	defualtGenerations := 100
	defaultBias := 0.5

	// declare input variables
	var filepath string
	var populationSize int
	var generations int
	var bias float64
	var connected bool
	var quiet bool

	// declare flags
	flag.StringVar(&filepath, "f", "data/sgb128/sgb128.dat", "Path to the .wdat graph file")
	flag.IntVar(&populationSize, "p", defualtPopulation, "Population size")
	flag.IntVar(&generations, "g", defualtGenerations, "Number of generations to run")
	flag.Float64Var(&bias, "b", defaultBias, "Fitness bias for parent selection, a number between zero and one **Currently does nothing**")
	flag.BoolVar(&connected, "c", true, "Whether or not the graph is connected: adds a linear speedup since children don't need to be checked that they are a cycle.")
	flag.BoolVar(&quiet, "v", true, "Verbose mode: Outputs percent completion")

	// parse flags
	flag.Parse()

	rand.Seed(time.Now().Unix())

    if connected {
    	ConnectedGentior(filepath, populationSize, generations, bias, quiet)
    } else {
    	Gentior(filepath, populationSize, generations, bias, quiet)
    }

}

func Gentior(filepath string, populationSize, generations int, bias float64, quiet bool) {
	// Create graph and population
	graph := NewWeightedGraphFromFile(filepath)
	population := generatepopulation(graph, populationSize)

	// genetically develop good solutions (hopefully)
	for i := 0; i < generations; i++ {
    	parents := selectParents(len(population), bias)
    	pop := population[parents[0]]
    	mom := population[parents[1]]
        child := edgeRecombination(pop, mom, graph)
		population = reconstructPopulation(population, child)
		if !quiet {
            printProgress(i, generations)
		}
	}

	// print most fit solution
	//showPopulation(population)
	showSolution(population)
}

func printProgress(i, n int) {
    percent := float64(i) / float64(n)
    percent *= 100
    fmt.Println("\033[H\033[2J")
    fmt.Printf("%d%% \n", int(percent))
}

func ConnectedGentior(filepath string, populationSize, generations int, bias float64, quiet bool) {
	// Create graph and population
	graph := NewWeightedGraphFromFile(filepath)
	population := generatepopulation(graph, populationSize)

	// genetically develop good solutions (hopefully)
	for i := 0; i < generations; i++ {
    	parents := selectParents(len(population), bias)
    	pop := population[parents[0]]
    	mom := population[parents[1]]
        child := connectedEdgeRecombination(pop, mom, graph)
		population = reconstructPopulation(population, child)
		if !quiet {
            printProgress(i, generations)
		}
	}

	// print most fit solution
	//showPopulation(population)
	showSolution(population)
}

func showSolution(list []Hamiltonian) {
    fmt.Println()
    fmt.Println("-------------------- Results -------------------- ")
    fmt.Printf("Shortest path:\t%d\n\n", list[0].fitness)
    fmt.Printf("Sholution: %v\n", list[0].path)
    fmt.Println("------------------------------------------------- ")
}

func reconstructPopulation(population []Hamiltonian, offspring *Hamiltonian) []Hamiltonian{
    return binaryInsert(*offspring, population)
}

// Insert inserts an element into the list, maintaining the size
func binaryInsert(el Hamiltonian, data []Hamiltonian) []Hamiltonian{
    index := sort.Search(len(data), func(i int) bool { return data[i].fitness > el.fitness})
    data = append(data, Hamiltonian{})
    copy(data[index+1:], data[index:])
    data[index] = el
    return data[:len(data) - 1]
}

// Add inserts an element into the list, increasing the size
func binaryAdd(el Hamiltonian, data []Hamiltonian) []Hamiltonian{
    index := sort.Search(len(data), func(i int) bool { return data[i].fitness > el.fitness})
    data = append(data, Hamiltonian{})
    copy(data[index+1:], data[index:])
    data[index] = el
    return data
}

func edgeRecombination(pop Hamiltonian, mom Hamiltonian, g *Undirected) *Hamiltonian {
    numVertices := g.Order()

    edgeList := getEdgeList(pop, mom)
    child := new(Hamiltonian)
    attemptCount := 0
    maxAttempts := 6000

    for pathFound := false; !pathFound; attemptCount++ {

        // if you've tried n times with no successful child, adopt a new child
        if attemptCount == maxAttempts {
            child = makeZeroPath(g)
            return child
        }


        start := 0

        // random Start
        start = rand.Intn(numVertices)

        child.path = []int{}
        visited := make([]bool, numVertices)
        for i, m := 0, start; i < numVertices && m >= 0; i++ {
        	rand.Seed(rand.Int63())
            visited[m] = true
            child.path = append(child.path, m)

            // get next index
            nextEdge := smallestAdjecency(m, edgeList, visited)

            if nextEdge >= 0 {
                m = nextEdge
            } else {
                m = getUnvisitedEdge(m, visited, g)
            }
        }
        if len(child.path) == numVertices && isCycle(g, child.path){
            pathFound = true
            child.fitness = fitness(g, child.path)
        }
    }

    return child
}

func connectedEdgeRecombination(pop Hamiltonian, mom Hamiltonian, g *Undirected) *Hamiltonian {
    numVertices := g.Order()

    edgeList := getEdgeList(pop, mom)
    child := new(Hamiltonian)
    attemptCount := 0
    maxAttempts := 6000

    for pathFound := false; !pathFound; attemptCount++ {

        // if you've tried n times with no successful child, adopt a new child
        if attemptCount == maxAttempts {
            child = makeZeroPath(g)
            return child
        }


        start := 0

        // random Start
        start = rand.Intn(numVertices)

        child.path = []int{}
        visited := make([]bool, numVertices)
        for i, m := 0, start; i < numVertices && m >= 0; i++ {
        	rand.Seed(rand.Int63())
            visited[m] = true
            child.path = append(child.path, m)

            // get next index
            nextEdge := smallestAdjecency(m, edgeList, visited)

            if nextEdge >= 0 {
                m = nextEdge
            } else {
                m = getUnvisitedEdge(m, visited, g)
            }
        }
        if len(child.path) == numVertices{
            pathFound = true
            child.fitness = fitness(g, child.path)
        }
    }

    return child
}

// getUnvisitedEdgeaccepts an edge, a list of visited edges, and a graph.
// it returns a random, unvisited, adjecent vertex. If no such edge exists
// the method returns -1
func getUnvisitedEdge(current int, visited []bool, g * Undirected) int {
    adjecents := g.GetEdges(current)
    unvisited := []int{}

    for _, n := range adjecents {
        if !visited[n] {
            unvisited = append(unvisited, n)
        }
    }
    if len(unvisited) > 0 {
        return unvisited[rand.Intn(len(unvisited))]
    } else {
        return -1
    }
}

// getEdgeList takes in two hamiltonian cycles. It builds a list of
// neighbors based on what vertices are neighbors in each cycle.
// if pop = (0 1 2 3 4 5) and mom = (1 2 3 5 0 4), the list will return
//      0 : (1 4 5)
//      1 : (0 2 4)
//      2 : (1 3)
//      3 : (2 4 5)
//      4 : (0 1 3 5)
//      5 : (0 3 4)
func getEdgeList(pop Hamiltonian, mom Hamiltonian) [][]int {
    numEdges := len(pop.path)
    edgeList := make([][]int, numEdges)
    
    // build edge list
    for i, n := range pop.path {
        last := i - 1
        if last < 0 {
            last = numEdges - 1
        }
        next := (i + 1) % len(edgeList)

        edgeList[n] = insert(pop.path[last], edgeList[n])
        edgeList[n] = insert(pop.path[next], edgeList[n])

    }
    for i, n := range mom.path {
        last := i - 1
        if last < 0 {
            last = numEdges - 1
        }
        next := (i + 1) % len(edgeList)

        edgeList[n] = insert(mom.path[last], edgeList[n])
        edgeList[n] = insert(mom.path[next], edgeList[n])
    }
    return edgeList
}

// insert will append an element into a list if that element
// does not currently exist in the list. Used to ensure there are no
// repeated numbers in adjecency lists (introducing edge bias)
func insert(n int, list []int) []int{
    if len(list) == 0 {
        list = []int{n}
    } else {
        for _, m := range list {
            if m == n {
                return list
            }
        }
        list = append(list, n)
    }
    return list
}

// smallestAdjecency finds the lowest degree unvisited edge adjecent to the current edge
// if no adjecent edges are unvisited, the function returns -1
func smallestAdjecency(current int, edges [][]int, visited []bool) int {
    // building list of unvisited
    possibles := []int{}
    for _, n := range edges[current] {
        if !visited[n] {
            possibles = append(possibles, n)
        }
    }

    smallIndex := -1
    smallVal := -1

    // search for smallest unvisisted node
    for i, n := range possibles {
        // if we are starting the list or we've found a smaller vertex, replace it
        if smallIndex < 0 || len(edges[n]) < len(edges[smallVal]) {
            smallIndex = i
            smallVal = n
        // Or if we have found the same degree vertex, randomly replace it 
        } else if len(edges[n]) == len(edges[smallVal]) && (rand.Int() % 2 == 0) {
            smallIndex = i
            smallVal = n
        }
    }

    return smallVal
}

func showPopulation(population []Hamiltonian) {
    fmt.Println("Population: ")
	for _, h := range population {
        fmt.Printf("%v\n", h)
	}
}

func selectParents(populationSize int, bias float64) []int {
    // Select at Random
    // TODO implement bias
    mom := rand.Intn(populationSize)
    pop := rand.Intn(populationSize)

    // ensure parents are different
    for mom == pop {
        pop = rand.Intn(populationSize)
    }

    parentList := []int{pop, mom}
    return parentList
}

//TODO?
func applyBias(i int, max int, b float64) int {
    //probabilities := make(
	n := randomBias(i, b)
	if n > max {
    	return max
	} else {
    	return n
	}
}

//TODO?
func randomBias(i int, b float64) int {
    return -1
}

func makeRandomPath(g * Undirected) *Hamiltonian{
    tour := new(Hamiltonian)
    tour.path = rand.Perm(g.Order())
    tour.fitness = fitness(g, tour.path)
    return tour
}

func generatepopulation(g *Undirected, populationSize int) []Hamiltonian {
	population := make([]Hamiltonian, 0, populationSize)

	for i := 0; i < populationSize; i++ {
		path := makeRandomPath(g)
		population = binaryAdd(*path, population) 
	}
	return population
}

// makeZeroPath makes a path that starts at zero.
// It finds goal using a depth first search.
// and returns a path object
func makeZeroPath(g *Undirected) *Hamiltonian {
	p := new(Hamiltonian)
	//v := rand.Intn(g.Order())
	//edges := g.GetEdges(v)

	// initialize random walk path
	path := randomDFS(0, g)

	p.path = path
	p.fitness = fitness(g, path)

	return p
}

// randomDFS uses dfs to randomly search a graph for a goal.
// it is the master function and should 
func randomDFS(vertex int, g *Undirected) []int {
	visited := make([]bool, g.Order())
	path, found := dfs(vertex, vertex, visited, 0,[]int{}, g)
	if !found {
		fmt.Println("No possible Hamiltonian Cycle")
		panic(fmt.Sprint(""))
	} else {
		return path
	}

}

// dfsa is the helper function to randomDFS and should not
// be called directly.
//
// dfs takes in a current vertex, the goal vertex,
// an array of visited vertices, the depth of the search,
// and a graph.
//
// dfs recursivley searches the graph for the goal.
//
// @return the path taken by the recursive calls
// @return whether or not the goal was found
func dfs(current int, goal int, visited []bool, depth int, soFar []int, g *Undirected) ([]int, bool) {
	// base case
	if current == goal && depth == g.Order() {
		return []int{}, true
	} else {
		visited[current] = true
	}

	edges := g.GetEdges(current)

	// starting at a random index, iterate over edges, looping over the end
	firstIteration := true
	for i, j := rand.Intn(len(edges)), -1; firstIteration || i != j; i = (i + 1) % len(edges) {
		// ensure i can't get back to j
		if firstIteration {
			firstIteration = false
			j = i
		}

        // if at correct depth and adjecent to the goal, return success
		if len(visited) - 1 == depth {
    		for _, n := range edges {
        		if n == goal {
            		return []int{current}, true
        		}
    		}
		}

		// if the the has not been visited OR the next edge is not the goal (unless it is at the coorect depth
		if !visited[edges[i]] || (depth >= len(visited) && edges[i] == goal)  {
			// copy visited array
			visitedCopy := make([]bool, len(visited))
			copy(visitedCopy, visited)
			// visit that edge
			path, found := dfs(edges[i], goal, visitedCopy, depth+1, append(soFar, current),  g)
			if found {
				return append([]int{current}, path...), found
			}
		}
	}
	return nil, false
}

// A fitness evaluator
// Returns the sum weight of the walk
func fitness(g *Undirected, walk []int) int {
	length := 0
	for i := 0; i < len(walk); i++ {
		n := (i + 1) % len(walk)
		length += g.Weight(walk[i], walk[n])
	}
	return length
}

// determines if a walk is a cycle
// by ensuring that every sequential vertex is
// connected.
func isCycle(g *Undirected, walk []int) bool {
	cycle := true
	length := len(walk)

	for i := 0; i < length && cycle; i++ {
		n := walk[(i + 1) % length]
		m := walk[i]
		cycle = g.IsConnected(m, n)
	}
	return cycle
}
