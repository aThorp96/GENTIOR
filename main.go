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

type path struct {
	order   []int // the actualy path
	fitness int
}

func makeRandomPath(g) {
	n := g.Order()
	p := new(path)
	v = rand.Intn(populationSize)
	edges = g.GetEdges(v)

}

func generatepopulation(g *graph, populationSize int) [][]int {
	population := make([][]int, populationSize)
	for i := 0; i < populationSize; i++ {
		path := makeRandomPath(g)

		// generate a random starting vertex
		v = rand.Intn(populationSize)
	}
}
