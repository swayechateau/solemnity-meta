package meta

import (
	"io"

	"golang.org/x/net/html"
)

// Document represents an HTML document to be manipulated. Unlike jQuery, which
// is loaded as part of a DOM document, and thus acts upon its containing
// document, GoQuery doesn't know which HTML document to act upon. So it needs
// to be told, and that's what the Document class is for. It holds the root
// document node to manipulate, and can make selections on this document.
type Document struct {
	*Selection
	rootNode *html.Node
}

// NewDocumentFromReader returns a Document from an io.Reader.
// It returns an error as second value if the reader's data cannot be parsed
// as html. It does not check if the reader is also an io.Closer, the
// provided reader is never closed by this call. It is the responsibility
// of the caller to close it if required.
func NewDocumentFromReader(r io.Reader) (*Document, error) {
	root, e := html.Parse(r)
	if e != nil {
		return nil, e
	}
	return newDocument(root), nil
}

// Private constructor, make sure all fields are correctly filled.
func newDocument(root *html.Node) *Document {
	// Create and fill the document
	d := &Document{nil, root}
	d.Selection = newSingleSelection(root, d)
	return d
}

// Selection represents a collection of nodes matching some criteria. The
// initial Selection can be created by using Document.Find, and then
// manipulated using the jQuery-like chainable syntax and methods.
type Selection struct {
	Nodes    []*html.Node
	document *Document
	prevSel  *Selection
}

// Helper constructor to create a selection of only one node
func newSingleSelection(node *html.Node, doc *Document) *Selection {
	return &Selection{[]*html.Node{node}, doc, nil}
}

// QueryMatcher is an interface that defines the methods to match
// HTML nodes against a compiled selector string. Cascadia's
// Selector implements this interface.
type QueryMatcher interface {
	Match(*html.Node) bool
	MatchAll(*html.Node) []*html.Node
	Filter([]*html.Node) []*html.Node
}

// SingleMatcher returns a QueryMatcher matches the same nodes as m, but that stops
// after the first match is found.
//
// See the documentation of function Single for more details.
func SingleMatcher(m QueryMatcher) QueryMatcher {
	if _, ok := m.(singleMatcher); ok {
		// m is already a singleMatcher
		return m
	}
	return singleMatcher{m}
}

type singleMatcher struct {
	QueryMatcher
}

// Find gets the descendants of each element in the current set of matched
// elements, filtered by a selector. It returns a new Selection object
// containing these matched elements.
func (s *Selection) Find(selector string) *Selection {
	return pushStack(s, findWithMatcher(s.Nodes, compileMatcher(selector)))
}

// Internal map function used by many traversing methods. Takes the source nodes
// to iterate on and the mapping function that returns an array of nodes.
// Returns an array of nodes mapped by calling the callback function once for
// each node in the source nodes.
func mapNodes(nodes []*html.Node, f func(int, *html.Node) []*html.Node) (result []*html.Node) {
	set := make(map[*html.Node]bool)
	for i, n := range nodes {
		if vals := f(i, n); len(vals) > 0 {
			result = appendWithoutDuplicates(result, vals, set)
		}
	}
	return result
}

// Internal implementation of Find that return raw nodes.
func findWithMatcher(nodes []*html.Node, m QueryMatcher) []*html.Node {
	// Map nodes to find the matches within the children of each node
	return mapNodes(nodes, func(i int, n *html.Node) (result []*html.Node) {
		// Go down one level, becausejQuery's Find selects only within descendants
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				result = append(result, m.MatchAll(c)...)
			}
		}
		return
	})
}

func (m singleMatcher) MatchAll(n *html.Node) []*html.Node {
	// Optimized version - stops finding at the first match (cascadia-compiled
	// matchers all use this code path).
	if mm, ok := m.QueryMatcher.(interface{ MatchFirst(*html.Node) *html.Node }); ok {
		node := mm.MatchFirst(n)
		if node == nil {
			return nil
		}
		return []*html.Node{node}
	}

	// Fallback version, for e.g. test mocks that don't provide the MatchFirst
	// method.
	nodes := m.QueryMatcher.MatchAll(n)
	if len(nodes) > 0 {
		return nodes[:1:1]
	}
	return nil
}

// compileMatcher compiles the selector string s and returns
// the corresponding QueryMatcher. If s is an invalid selector string,
// it returns a QueryMatcher that fails all matches.
func compileMatcher(s string) QueryMatcher {
	// o
	cs, err := Compile(s)
	if err != nil {
		return invalidMatcher{}
	}
	return cs
}

// invalidMatcher is a QueryMatcher that always fails to match.
type invalidMatcher struct{}

func (invalidMatcher) Match(n *html.Node) bool             { return false }
func (invalidMatcher) MatchAll(n *html.Node) []*html.Node  { return nil }
func (invalidMatcher) Filter(ns []*html.Node) []*html.Node { return nil }

// Creates a new Selection object based on the specified nodes, and keeps the
// source Selection object on the stack (linked list).
func pushStack(fromSel *Selection, nodes []*html.Node) *Selection {
	result := &Selection{nodes, fromSel.document, fromSel}
	return result
}

const minNodesForSet = 1000

// Appends the new nodes to the target slice, making sure no duplicate is added.
// There is no check to the original state of the target slice, so it may still
// contain duplicates. The target slice is returned because append() may create
// a new underlying array. If targetSet is nil, a local set is created with the
// target if len(target) + len(nodes) is greater than minNodesForSet.
func appendWithoutDuplicates(target []*html.Node, nodes []*html.Node, targetSet map[*html.Node]bool) []*html.Node {
	// if there are not that many nodes, don't use the map, faster to just use nested loops
	// (unless a non-nil targetSet is passed, in which case the caller knows better).
	if targetSet == nil && len(target)+len(nodes) < minNodesForSet {
		for _, n := range nodes {
			if !isInSlice(target, n) {
				target = append(target, n)
			}
		}
		return target
	}

	// if a targetSet is passed, then assume it is reliable, otherwise create one
	// and initialize it with the current target contents.
	if targetSet == nil {
		targetSet = make(map[*html.Node]bool, len(target))
		for _, n := range target {
			targetSet[n] = true
		}
	}
	for _, n := range nodes {
		if !targetSet[n] {
			target = append(target, n)
			targetSet[n] = true
		}
	}

	return target
}

// Checks if the target node is in the slice of nodes.
func isInSlice(slice []*html.Node, node *html.Node) bool {
	return indexInSlice(slice, node) > -1
}

// Returns the index of the target node in the slice, or -1.
func indexInSlice(slice []*html.Node, node *html.Node) int {
	if node != nil {
		for i, n := range slice {
			if n == node {
				return i
			}
		}
	}
	return -1
}

// Each iterates over a Selection object, executing a function for each
// matched element. It returns the current Selection object. The function
// f is called for each element in the selection with the index of the
// element in that selection starting at 0, and a *Selection that contains
// only that element.
func (s *Selection) Each(f func(int, *Selection)) *Selection {
	for i, n := range s.Nodes {
		f(i, newSingleSelection(n, s.document))
	}
	return s
}

// Attr gets the specified attribute's value for the first element in the
// Selection. To get the value for each element individually, use a looping
// construct such as Each or Map method.
func (s *Selection) Attr(attrName string) (val string, exists bool) {
	if len(s.Nodes) == 0 {
		return
	}
	return getAttributeValue(attrName, s.Nodes[0])
}

// Private function to get the specified attribute's value from a node.
func getAttributeValue(attrName string, n *html.Node) (val string, exists bool) {
	if a := getAttributePtr(attrName, n); a != nil {
		val = a.Val
		exists = true
	}
	return
}

func getAttributePtr(attrName string, n *html.Node) *html.Attribute {
	if n == nil {
		return nil
	}

	for i, a := range n.Attr {
		if a.Key == attrName {
			return &n.Attr[i]
		}
	}
	return nil
}
