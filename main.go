package gentior

import (
	"flag"
	"github.com/athorp96/graphs"
	"math/rand"
)

func Main() {
	// default values
	defualtPopulation := 50
	defualtgenerations := 100

	// declare input variables
	var filepath string
	var populationSize int
	var generations int

	// declare flags
	flag.StringVar(&filepath, "file", "", "Path to the .dat graph file")
	flag.IntVar(&populationSize, "population", defualtPopulation, "Population size")
	flag.IntVar(&generations, "generations", defualtgenerations, "Number of generations to run")

	// parse flags
	flag.Parse()

	// Create graph and population
	graph := graphs.NewGraphFromFile(filepath)
	population := generatepopulation(graph, populationSize)

	// genetically develop good solutions (hopefully)
	for i := 0; i < generations; i++ {
		parents := selectParents(graph, population, fitnessEvaluator)
		offspring := createOffspring(parents, recombination)
		reconstructPopulation(population, offspring, fitnessEvaluator)
	}

	// print most fit solution
	showSolution(population, fitnessEvaluator)
}

func generatepopulation(g *graph, populationSize int) [][]int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	population := make([][]int, populationSize)

	for i := 0; i < populationSize; i++ {
		path := makeRandomPath(g)
	}
}

type simplePath struct {
	path    []int // the path
	fitness int
}

func makeRandomPath(g) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	n := g.Order()
	p := new(simplePath)
	edges = g.GetEdges(v)
	v = rand.Intn(populationSize)

	// initialize random walk path
	path := r.Perm(n)

	// randomize until there is a valid cycle
	for !isCycle(g, path) {
		path = r.Perm(n)
	}

	p.path = walk


func isCycle(g *graph, walk []int) bool {
	cycle := true
	length := len(walk)

	for i := 0; i < length && cycle; i++ {
		n := (i + 1) % length
		cycle = g.IsConnected(i, n)
	}
	return cycle
}
