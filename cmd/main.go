package main

import (
	"alien"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/olekukonko/tablewriter"
)

var (
	mapFile   string
	numAliens int
)

func main() {

	flag.StringVar(&mapFile, "map", "testdata/map.txt", "file containing the map definition")
	flag.IntVar(&numAliens, "n", 2, "number of aliens to use in the simulation")

	flag.Parse()

	if numAliens < 2 {
		log.Fatal("Aliens must be more than 2")
	}

	// no need to use crypto/rand here
	// but we make sure to seed this
	rand.Seed(time.Now().Unix())

	f, err := os.Open(mapFile)
	if err != nil {
		log.Fatalf("could not open map file... %v", err)
	}

	defer f.Close()

	m, err := alien.NewWorldWideMapFromReader(f)
	if err != nil {
		log.Fatalf("could not create alien map... %v", err)
	}

	m.Simulate(numAliens)

	fmt.Print("\n \n")

	data := [][]string{}

	for _, v := range m.Aliens {
		data = append(data, []string{
			v.Name,
			v.City,
			humanize.Ordinal(v.CurrentIteration),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Alien name", "Current location", "Current iteration around the map"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	fmt.Print("\n \n")

	data = [][]string{}

	for _, v := range m.Cities {
		data = append(data, []string{
			v.Name,
			humanize.Comma(int64(len(v.Paths))),
		})
	}

	table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"city name", "number of paths"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()
}
