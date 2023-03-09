### World map alien invasion

- [Challenge](#challenge)
- [Benchmarks](#benchmarks)


```sh

$ go build -o invasion ./cmd
$ ./invasion -n numberOfAliens -file "path-to-custom-map-file"

```

```
City (Bee) was just been destroyed by alien (David) and alien (Emma) (in iteration 0)
City (Bee) was just been destroyed by alien (Abigail) and alien (Emma) (in iteration 0)
City (Bee) was just been destroyed by alien (Mason) and alien (Emma) (in iteration 0)
City (Foo) was just been destroyed by alien (Mason) and alien (Emily) (in iteration 0)
City (Bar) was just been destroyed by alien (Ava) and alien (Ava) (in iteration 0)
City (Bar) was just been destroyed by alien (Emily) and alien (Ava) (in iteration 4)


+------------+------------------+--------------------------------+
| ALIEN NAME | CURRENT LOCATION |  CURRENT ITERATION AROUND THE  |
|            |                  |              MAP               |
+------------+------------------+--------------------------------+
| Ava        | Bar              | 10000th                        |
+------------+------------------+--------------------------------+


+-----------+-----------------+
| CITY NAME | NUMBER OF PATHS |
+-----------+-----------------+
| Foo       |               1 |
+-----------+-----------------+
```

### <a name="bench"> Benchmarks </a>

```sh

$ go test -bench .

```

```

goos: darwin
goarch: arm64
pkg: alien
BenchmarkNewWorldFileFromReader1000-10            670588              1779 ns/op
PASS
ok      alien   3.446

```


### <a name="challenge"> Challenge </a>

Mad aliens are about to invade the earth and you are tasked with simulating the invasion.

You are given a map containing the names of cities in the non-existent world of X. The map is in a file, with one city per line. The city name is first, followed by 1-4 directions (north, south, east, or west). Each one represents a road to another city that lies in that direction.

For example:

```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```

The city and each of the pairs are separated by a single space, and the directions are separated from their respective cities with an equals (=) sign. You should create N aliens, where N is specified as a command-line argument. These aliens start out at random places on the map, and wander around randomly, following links. Each iteration, the aliens can travel in any of the directions leading out of a city. In our example above, an alien that starts at Foo can go north to Bar, west to Baz, or south to Qu-ux.

When two aliens end up in the same place, they fight, and in the process kill each other and destroy the city. When a city is destroyed, it is removed from the map, and so are any roads that lead into or out of it.

In our example above, if Bar were destroyed the map would now be something like:

```
Foo west=Baz south=Qu-ux
```

Once a city is destroyed, aliens can no longer travel to or through it. This may lead to aliens getting "trapped".

You should create a program that reads in the world map, creates N aliens, and unleashes them. The program should run until all the aliens have been destroyed, or each alien has moved at least 10,000 times. When two aliens fight, print out a message like:

```
Bar has been destroyed by alien 10 and alien 34!
```
