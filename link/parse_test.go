package link

import (
	"testing"

	"golang.org/x/net/html"
)

func TestText(t *testing.T) {
	var n html.Node
	n.Type = html.TextNode
	n.Data = "Testing text"
	got := text(&n)
	if got != "Testing text" {
		t.Errorf("Error in text - got: %v, want: %v.\n", got, "Testing text")
	}

	var e html.Node
	e.Type = html.ElementNode
	e.Data = "Not suppose to be here."
	gotE := text(&e)
	if gotE != "" {
		t.Errorf("Error in text - got: %v, want: %v.\n", gotE, "")
	}
}

func TestBuildLink(t *testing.T) {
	var n html.Node
	n.Type = html.ElementNode
	n.Attr = []html.Attribute{html.Attribute{Key: "href", Val: "/helloWorld"}}
	var s html.Node
	s.Type = html.TextNode
	s.Data = "Testing Build Link"
	n.FirstChild = &s
	got := buildLink(&n)
	want := Link{"/helloWorld", "Testing Build Link"}
	if got != want {
		t.Errorf("Error in build link - got: %v, want: %v.\n", got, want)
	}
}

func TestLinkNode(t *testing.T) {
	var n html.Node
	n.Type = html.ElementNode
	n.Data = "a"
	got := linkNodes(&n)
	want := []*html.Node{&n}
	for i, p := range got {
		if p != want[i] {
			t.Errorf("Error in link node - got: %v, want: %v.\n", p, want[i])
		}
	}
}
