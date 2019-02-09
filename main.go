package gentior

import (
    callgraph

func Main(){

    graph := readgraph(filepath)
    defualtPopulation := 50
    defualtgenerations := 100
    populationSize int
    generations int

    flag.StringVar(&populationSize, "population", defualtPopulation, "Population size")
    flag.StringVar(&generations, "generations", defualtgenerations, "Number of generations to run")

    population := generatepopulation(graph, n)

    for (i := 0; i < generations; i++) {
        parents := selectParents(graph, population, fitnessEvaluator)
        offspring := createOffspring(parents, recombination)
        reconstructPopulation(population, offspring, fitnessEvaluator)
    }

    showSolution(population)
}
