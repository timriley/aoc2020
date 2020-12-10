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

type rule struct {
	descriptor string
	contained  []contained
}

type contained struct {
	descriptor string
	quantity   int
}

type graph struct {
	nodes map[string]*node
}

type node struct {
	key         string
	parentEdges []*edge
	childEdges  []*edge
}

type edge struct {
	parent *node
	child  *node
	weight int
}

func main() {
	var rules []rule
	err := fileinput.LoadThen("day07/input.txt", "\n", func(s string) {
		rules = append(rules, ruleFromString(s))
	})
	if err != nil {
		log.Fatal(err)
	}

	g := newGraph()
	for _, rule := range rules {
		for _, c := range rule.contained {
			g.addEdge(rule.descriptor, c.descriptor, c.quantity)
		}
	}

	numberContainingShinyGold := part1(g)
	numberContainedByShinyGold := part2(g)

	fmt.Printf("Number of bags that can contain a shiny gold bag (part 1): %v\n", numberContainingShinyGold)
	fmt.Printf("Number of bags contained by shiny gold bag (part 2): %v\n", numberContainedByShinyGold)
}

func part1(g *graph) int {
	return len(g.allAscendants(g.nodes["shiny gold"]))
}

func part2(g *graph) int {
	return g.combinedChildrenWeight(g.nodes["shiny gold"])
}

func newGraph() *graph {
	return &graph{
		nodes: map[string]*node{},
	}
}

func (g *graph) addNode(key string) *node {
	if existing, ok := g.nodes[key]; ok {
		return existing
	}

	node := &node{key: key}
	g.nodes[key] = node
	return node
}

func (g *graph) addEdge(parentKey, childKey string, weight int) {
	parent, ok := g.nodes[parentKey]
	if !ok {
		parent = g.addNode(parentKey)
	}

	child, ok := g.nodes[childKey]
	if !ok {
		child = g.addNode(childKey)
	}

	edge := &edge{parent: parent, child: child, weight: weight}
	parent.childEdges = append(parent.childEdges, edge)
	child.parentEdges = append(child.parentEdges, edge)
}

func (g *graph) allAscendants(n *node) []*node {
	set := make(map[*node]bool)

	for _, pe := range n.parentEdges {
		set[pe.parent] = true

		for _, pn := range g.allAscendants(pe.parent) {
			set[pn] = true
		}
	}

	nodes := make([]*node, 0, len(set))
	for k := range set {
		nodes = append(nodes, k)
	}
	return nodes
}

func (g *graph) combinedChildrenWeight(n *node) int {
	weight := 0
	for _, ce := range n.childEdges {
		weight += ce.weight                                      // contains n bags directly
		weight += ce.weight * g.combinedChildrenWeight(ce.child) // weight of all combined children
	}
	return weight
}

func ruleFromString(s string) rule {
	// Example: "striped fuchsia bags contain 3 dotted green bags, 2 plaid maroon bags."

	// Strip spurious detail: the word "bag"/"bags" and trailing "."
	s = regexp.MustCompile(`\sbags?\.?`).ReplaceAllLiteralString(s, "")

	parts := strings.Split(s, " contain ")
	if len(parts) != 2 {
		log.Fatalf("expected only 2 parts in %v", parts)
	}

	var c []contained
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

		c = append(c, contained{
			quantity:   qty,
			descriptor: descriptor,
		})
	}

	return rule{
		descriptor: parts[0],
		contained:  c,
	}
}
