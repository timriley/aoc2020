// --- Day 7: Handy Haversacks ---
//
// You land at the regional airport in time for your next flight. In fact, it
// looks like you'll even have time to grab some food: all flights are currently
// delayed due to issues in luggage processing.
//
// Due to recent aviation regulations, many rules (your puzzle input) are being
// enforced about bags and their contents; bags must be color-coded and must
// contain specific quantities of other color-coded bags. Apparently, nobody
// responsible for these regulations considered how long they would take to
// enforce!
//
// For example, consider the following rules:
//
// - light red bags contain 1 bright white bag, 2 muted yellow bags.
// - dark orange bags contain 3 bright white bags, 4 muted yellow bags.
// - bright white bags contain 1 shiny gold bag.
// - muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
// - shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
// - dark olive bags contain 3 faded blue bags, 4 dotted black bags.
// - vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
// - faded blue bags contain no other bags.
// - dotted black bags contain no other bags.
//
// These rules specify the required contents for 9 bag types. In this example,
// every faded blue bag is empty, every vibrant plum bag contains 11 bags (5
// faded blue and 6 dotted black), and so on.
//
// You have a shiny gold bag. If you wanted to carry it in at least one other
// bag, how many different bag colors would be valid for the outermost bag? (In
// other words: how many colors can, eventually, contain at least one shiny gold
// bag?)
//
// In the above rules, the following options would be available to you:
//
// - A bright white bag, which can hold your shiny gold bag directly.
// - A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
// - A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
// - A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
//
// So, in this example, the number of bag colors that can eventually contain at
// least one shiny gold bag is 4.
//
// How many bag colors can eventually contain at least one shiny gold bag? (The
// list of rules is quite long; make sure you get all of it.)
//
// --- Part Two ---
//
// It's getting pretty expensive to fly these days - not because of ticket
// prices, but because of the ridiculous number of bags you need to buy!
//
// Consider again your shiny gold bag and the rules from the above example:
//
// - faded blue bags contain 0 other bags.
// - dotted black bags contain 0 other bags.
// - vibrant plum bags contain 11 other bags: 5 faded blue bags and 6 dotted black bags.
// - dark olive bags contain 7 other bags: 3 faded blue bags and 4 dotted black bags.
//
// So, a single shiny gold bag must contain 1 dark olive bag (and the 7 bags
// within it) plus 2 vibrant plum bags (and the 11 bags within each of those): 1
// + 1*7 + 2 + 2*11 = 32 bags!
//
// Of course, the actual rules have a small chance of going several levels
// deeper than this example; be sure to count all of the bags, even if the
// nesting becomes topologically impractical!
//
// Here's another example:
//
// - shiny gold bags contain 2 dark red bags.
// - dark red bags contain 2 dark orange bags.
// - dark orange bags contain 2 dark yellow bags.
// - dark yellow bags contain 2 dark green bags.
// - dark green bags contain 2 dark blue bags.
// - dark blue bags contain 2 dark violet bags.
// - dark violet bags contain no other bags.
//
// In this example, a single shiny gold bag must contain 126 other bags.
//
// How many individual bags are required inside your single shiny gold bag?

package main

import (
	"aoc2020/utils/fileinput"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	parent    string
	contained []Contained
}

type Contained struct {
	descriptor string
	quantity   int
}

type Graph struct {
	Vertices map[string]*Vertex
}

type Vertex struct {
	Key      string
	Vertices map[string]*Vertex
}

func main() {
	var rules []Rule
	err := fileinput.LoadThen("day07/input.txt", "\n", func(s string) {
		rules = append(rules, ruleFromString(s))
	})
	if err != nil {
		log.Fatal(err)
	}

	numberContainingShinyGold := part1(rules)

	fmt.Printf("Number of bags that can contain a shiny gold bag (part 1): %v\n", numberContainingShinyGold)
}

func part1(rules []Rule) int {
	// Populate graph in direction of contained->container
	g := newGraph()
	for _, rule := range rules {
		g.AddVertex(rule.parent)
		for _, c := range rule.contained {
			g.AddVertex(c.descriptor)
			g.AddEdge(c.descriptor, rule.parent)
		}
	}

	startBag := "shiny gold"
	containers := map[string]bool{}
	g.Traverse(g.Vertices[startBag], func(bagType string) {
		if bagType == startBag {
			return
		}
		containers[bagType] = true
	})

	return len(containers)
}

func part2(rules []Rule) int {
	return 0
}

func newGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func newVertex(key string) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[string]*Vertex{},
	}
}

func (g *Graph) AddVertex(key string) {
	if _, ok := g.Vertices[key]; ok {
		return
	}

	v := newVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 string) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		panic("vertices do not exist")
	}

	// Do nothing if vertices are already connected
	if _, ok := v1.Vertices[v2.Key]; ok {
		return
	}

	v1.Vertices[v2.Key] = v2

	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

// Traverse is a depth-first traversal implementation
func (g *Graph) Traverse(start *Vertex, visitor func(string)) {
	if start == nil {
		return
	}

	visited := map[string]bool{}

	visited[start.Key] = true
	visitor(start.Key)

	for _, v := range start.Vertices {
		if visited[v.Key] {
			continue
		}

		g.Traverse(v, visitor)
	}
}

func ruleFromString(s string) Rule {
	// Example: "striped fuchsia bags contain 3 dotted green bags, 2 plaid maroon bags."

	// Strip spurious detail: the word "bag"/"bags" and trailing "."
	s = regexp.MustCompile(`\sbags?\.?`).ReplaceAllLiteralString(s, "")

	parts := strings.Split(s, " contain ")
	if len(parts) != 2 {
		log.Fatalf("expected only 2 parts in %v", parts)
	}

	var contained []Contained
	for _, s := range strings.Split(parts[1], ", ") {
		if s == "no other" {
			continue
		}

		parts := strings.Split(s, " ")
		if len(parts) != 3 {
			log.Fatalf("expected only 3 parts in %v", parts)
		}

		qty, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		descriptor := strings.Join(parts[1:], " ")

		contained = append(contained, Contained{
			quantity:   qty,
			descriptor: descriptor,
		})
	}

	return Rule{
		parent:    parts[0],
		contained: contained,
	}
}
