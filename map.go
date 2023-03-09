package alien

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
)

const (
	minTurnsPerAlien = 10000
)

type City struct {
	Name  string
	Paths []*Path
}

func (c *City) Add(p *Path) { c.Paths = append(c.Paths, p) }

type Path struct {
	Direction string
	City      string
}

type WorldWideMap struct {
	Cities []*City
	Aliens []*Alien
}

func (m *WorldWideMap) Simulate(numOfAliens int) {
	m.PopulateMapWithAliens(numOfAliens)

	for i := 0; i < minTurnsPerAlien; i++ {
		m.WanderToNewLocation()
		m.Fight(i)
	}
}

func NewMap(f string) (*WorldWideMap, error) {

	fi, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	defer fi.Close()

	return NewWorldWideMapFromReader(fi)
}

func NewWorldWideMapFromReader(f io.Reader) (*WorldWideMap, error) {
	m := &WorldWideMap{}

	// asumption that you could load a 200mb or 1GB file
	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		city := &City{Paths: make([]*Path, 0)}

		cityDefinition := strings.Split(scanner.Text(), " ")

		if len(cityDefinition) < 2 {
			return nil, errors.New("a valid city must have at least one connection")
		}

		for _, definition := range cityDefinition[1:] {
			directions := strings.Split(definition, "=")
			city.Add(&Path{
				Direction: directions[0],
				City:      directions[1],
			})
		}

		city.Name = cityDefinition[0]

		m.AddCity(city)
	}

	return m, scanner.Err()
}

func (m *WorldWideMap) AddCity(c *City) { m.Cities = append(m.Cities, c) }

func (m *WorldWideMap) WanderToNewLocation() {

	// we could have 1K aliens, this makes all aliens wander at the same
	// time
	var g sync.WaitGroup

	for _, a := range m.Aliens {
		go func([]*City) {
			g.Add(1)
			defer g.Done()

			for _, c := range m.Cities {

				if len(c.Paths) == 0 {
					continue
				}

				// make sure we move this alien to a  totally new city
				for c.Name == a.City {
					a.City = c.Paths[rand.Intn(len(c.Paths))].City
					return
				}
			}

		}(m.Cities)

		a.CurrentIteration++
	}

	g.Wait()
}

func (m *WorldWideMap) Fight(turn int) {

	// use just a string here to keep the first alien, when the next alien
	// gets here, they fight
	mapToCheckMultipleAliens := make(map[string]string, len(m.Aliens))

	for _, a := range m.Aliens {
		if _, pres := mapToCheckMultipleAliens[a.City]; !pres {
			mapToCheckMultipleAliens[a.City] = a.Name
			continue
		}

		// if we get here it means the map already exists. So this is
		// the second alien here
		s := fmt.Sprintf("City (%s) was just been destroyed by alien (%s) and alien (%s) (in iteration %d) \n", a.City, a.Name, mapToCheckMultipleAliens[a.City], turn)

		print(s)

		for _, alien := range m.Aliens {
			// make sure to kill both aliens
			if alien.Name == a.Name || alien.Name == mapToCheckMultipleAliens[a.City] {
				m.Aliens = m.Aliens[:len(m.Aliens)-1]
			}
		}

		for _, c := range m.Cities {
			// if this is the current city, delete it directly
			if c.Name == a.City {
				m.Cities = m.Cities[:len(m.Cities)-1]
				continue
			}

			// for other cities, check if they have the path to the recently
			// deleted city as they have to go too
			for _, v := range c.Paths {
				if v.City == a.City {
					c.Paths = c.Paths[:len(c.Paths)-1]
				}
			}
		}
	}
}

func (m *WorldWideMap) PopulateMapWithAliens(n int) {
	for i := 0; i < n; i++ {
		city := m.Cities[rand.Intn(len(m.Cities))]
		m.Aliens = append(m.Aliens, NewAlien(city))
	}
}
