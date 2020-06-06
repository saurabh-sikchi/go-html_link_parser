package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// A Link represents the <a href="..."> tag in an HTML document
type Link struct {
	Href string
	Text string
}

// takes in an HTML document and returns a slice of Links or an error
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	// 1. Find the <a> nodes
	// 2. For each <a> node
	// 			build Link struct
	// Return the Links

	nodes := linkNodes(doc)
	links := make([]Link, len(nodes))
	for i, n := range nodes {
		links[i] = buildLink(n)
	}

	return links, nil
}

func buildLink(n *html.Node) Link {

	var l Link

	for _, attr := range n.Attr {
		if strings.ToLower(attr.Key) == "href" {
			l.Href = attr.Val
			l.Text = buildText(n)
			break
		}
	}
	return l
}

func buildText(n *html.Node) string {

	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var s string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s += buildText(c)
	}

	return strings.Join(strings.Fields(s), " ")
}

func linkNodes(n *html.Node) []*html.Node {

	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "A") {
		return []*html.Node{n}
	}

	var nodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}

	return nodes
}
