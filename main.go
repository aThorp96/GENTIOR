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

var random *rand.Rand

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

	random = rand.New(rand.NewSource(time.Now().Unix()))

	Gentior(filepath, populationSize, generations, bias)
}

func Gentior(filepath string, populationSize, generations int, bias float64) {
	// Create graph and population
	graph := NewGraphFromFile(filepath)
	population := generatepopulation(graph, populationSize)

	// genetically develop good solutions (hopefully)
	/*	for i := 0; i < generations; i++ {
		parents := selectParents(len(population), bias)
		offspring := createOffspring(population[parents[0]], population[parents[1]], graph, recombination)
		reconstructPopulation(population, offspring)
	}*/

	// print most fit solution
	showPopulation(population)
}

func showPopulation(population []Hamiltonian) {
	for _, h := range population {
        fmt.Printf("%v\n", h)
	}
}

func selectParents(populationSize int, bias float64) []int {
	//int
	return make([]int, 2)
}

func applyBias(i int, b float64) int {
	return randomBias(i, b)
}

func randomBias(i int, b float64) int {
	addend := random.Float64()
	biased := float64(i) * (1.0 + addend)
	return int(biased)
}

func createOffspring(parents []Hamiltonian, g *Undirected, combine recombination) *Hamiltonian {
	pop := parents[0]
	mom := parents[1]
	return combine(&pop, &mom, g)
}

func generatepopulation(g *Undirected, populationSize int) []Hamiltonian {
	population := make([]Hamiltonian, populationSize)

	for i := 0; i < populationSize; i++ {
		path := makeZeroPath(g)
		population[i] = *path
	}
	return population
}

// makemakeRandomPath creates random permutations from [0,n) until
// the permutation is a cycle, a hamiltonian cycle since it includes
// each vertex exaclty once.
func makeRandomPath(g *Undirected) *Hamiltonian {
	n := g.Order()
	p := new(Hamiltonian)
	//v := rand.Intn(g.Order())
	//edges := g.GetEdges(v)

	// initialize random walk path
	path := random.Perm(n)

	fmt.Println("Trying Path")
	// randomize until there is a valid cycle
	for !isCycle(g, path) {

		for _, n := range path {
			fmt.Printf("%d ", n)
		}
		fmt.Println()

		path = random.Perm(n)
	}

	p.path = path
	p.fitness = fitness(g, path)

	return p
}

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

	// initialize a random object

	edges := g.GetEdges(current)

	// starting at a random index, iterate over edges, looping over the end
	firstIteration := true
	for i, j := random.Intn(len(edges)), -1; firstIteration || i != j; i = (i + 1) % len(edges) {
		// ensure i can't get back to j
		if firstIteration {
			firstIteration = false
			j = i
		}

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
				return append(path, current), found
			}
		}


		//fmt.Printf("i: %d\n\n", i)
		//fmt.Printf("j: %d\n\n", j)
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
