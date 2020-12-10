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

package main

import (
	"aoc2020/utils/fileinput"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Rule struct {
	parent   string
	children []string
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
		for _, c := range rule.children {
			g.AddVertex(c)
			g.AddEdge(c, rule.parent)
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

	// Ignore quantities for now, since we don't need them
	var children []string
	for _, s := range strings.Split(parts[1], ", ") {
		s = regexp.MustCompile(`\d+\s+`).ReplaceAllLiteralString(s, "")
		children = append(children, s)
	}

	return Rule{
		parent:   parts[0],
		children: children,
	}
}
