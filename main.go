package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

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

	// declare flags
	flag.StringVar(&filepath, "file", "quilt_small.dat", "Path to the .dat graph file")
	flag.IntVar(&populationSize, "population", defualtPopulation, "Population size")
	flag.IntVar(&generations, "generations", defualtGenerations, "Number of generations to run")
	flag.Float64Var(&bias, "bias", defaultBias, "Fitness bias for parent selection, a number between zero and one")

	// parse flags
	flag.Parse()

	rand.Seed(time.Now().Unix())

	Gentior(filepath, populationSize, generations, bias)

}

func Gentior(filepath string, populationSize, generations int, bias float64) {
	// Create graph and population
	graph := NewGraphFromFile(filepath)
	population := generatepopulation(graph, populationSize)

	// genetically develop good solutions (hopefully)
	//	for i := 0; i < generations; i++ {
	//	parents := selectParents(len(population), bias)
     edgeRecombination(population[0], population[1], graph)
		//reconstructPopulation(population, offspring)
	//}

	// print most fit solution
	showPopulation(population)
}

func edgeRecombination(pop Hamiltonian, mom Hamiltonian, g *Undirected) *Hamiltonian {
    numVertices := g.Order()
    visited := make([]bool, numVertices)

    edgeList := getEdgeList(pop, mom)
    fmt.Printf("%v\n", edgeList)

    // start with 0
    start := 0

    child := new(Hamiltonian)
    child.path = []int{start}

    for i, m := 0, start; i < numVertices; i++ {
        visited[m] = true
        // get next index
        nextEdge := smallestAdjecency(m, edgeList, visited)

        if nextEdge >= 0 {
            m = nextEdge
        } else {
            fmt.Printf("nextEdge: %d, who's neighbors are %v, currently: %v\n", m, edgeList[m], child.path)
            m = getUnvisitedEdge(m, visited, g)
        }
        child.path = append(child.path, m)
    }

    fmt.Printf("%v\n", child)
    return child
}

func getUnvisitedEdge(current int, visited []bool, g * Undirected) int {
    adjecents := g.GetEdges(current)
    unvisited := []int{}

    for _, n := range adjecents {
        if !visited[n] {
            unvisited = append(unvisited, n)
        }
    }
    fmt.Printf("unvisited: %v\n", unvisited)
    fmt.Printf("adjecents: %v\n", adjecents)
    return unvisited[rand.Intn(len(unvisited))]

}

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
	for _, h := range population {
        fmt.Printf("%v\n", h)
	}
}

// TODO 
func selectParents(populationSize int, bias float64) []int {
	return make([]int, 2)
}

//TODO
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

func generatepopulation(g *Undirected, populationSize int) []Hamiltonian {
	population := make([]Hamiltonian, populationSize)

	for i := 0; i < populationSize; i++ {
		path := makeZeroPath(g)
		population[i] = *path
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
		fmt.Printf("Found path! %v\n", path)
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
	for i := 0; i < length; i++ {
		n := (i + 1) % length
		length += g.Weight(i, n)
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
		n := (i + 1) % length
		cycle = g.IsConnected(i, n)
	}
	return cycle
}
