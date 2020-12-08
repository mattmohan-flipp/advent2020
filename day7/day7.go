package day7

import (
	"regexp"
	"strconv"
)

// Vertex in a DAG
type Vertex struct {
	parents    map[string]Link
	children   map[string]Link
	name       string
	totalScore int
}

// Link between two Verticies
type Link struct {
	parent *Vertex
	child  *Vertex
	score  int
}

func (parent *Vertex) addChild(child *Vertex, score int) {
	parent.children[child.name] = Link{parent, child, score}
	child.parents[parent.name] = Link{parent, child, score}
}

var lineRegexp = regexp.MustCompile(`^(.+?) bags contain ((?:[\d]|no+) (?:.+?) bags?,?)+\.$`)
var childrenRegexp = regexp.MustCompile(`([\d]+|no) (.+?) bags?,?`)

// Day7a solves the puzzle and returns a string
func Day7a(a []string) string {
	nodes := buildTree(a)

	// Walk tree and build a set of touched nodes
	total := len(walkTree(nodes["shiny gold"]))
	// Remove the node itself
	return strconv.Itoa(total - 1)
}

// Day7b solves the puzzle and returns a string
func Day7b(a []string) string {
	nodes := buildTree(a)

	// Walk tree counting bags
	total := walkTreeB(nodes["shiny gold"])
	// Remove the root bag
	return strconv.Itoa(total - 1)
}

func findOrNewVertex(nodes map[string]*Vertex, key string) *Vertex {
	node, foundNode := nodes[key]
	if !foundNode {
		node = &Vertex{make(map[string]Link), make(map[string]Link), key, -1}
		nodes[key] = node
	}
	return node
}

func buildTree(a []string) map[string]*Vertex {
	nodes := make(map[string]*Vertex)
	// Parse input
	for _, i := range a {
		matches := lineRegexp.FindStringSubmatch(i)

		parentName := matches[1]
		subMatch := childrenRegexp.FindAllStringSubmatch(matches[2], -1)

		parent := findOrNewVertex(nodes, parentName)
		for _, j := range subMatch {
			for k := 1; k < len(j); k += 2 {
				childName, linkCost := j[k+1], j[k]
				if childName == "other" {
					continue
				}

				child := findOrNewVertex(nodes, childName)
				cost, _ := strconv.Atoi(linkCost)
				parent.addChild(child, cost)
			}
		}

	}
	return nodes
}

func walkTree(node *Vertex) map[string]bool {
	passed := make(map[string]bool)

	passed[node.name] = true

	if len(node.parents) > 0 {
		for _, i := range node.parents {
			grandparents := walkTree(i.parent)
			for j := range grandparents {
				passed[j] = true
			}
		}
	}

	return passed
}

func walkTreeB(node *Vertex) int {
	// memoize result
	if node.totalScore > 0 {
		return node.totalScore
	}

	// One for the bag itself, plus the recursive sum of all children
	node.totalScore = 1
	if len(node.children) > 0 {
		for _, i := range node.children {
			node.totalScore += i.score * walkTreeB(i.child)
		}
	}
	return node.totalScore
}
