package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gbrlsnchs/radix"
)

type fileMatcher struct {
	tree    *radix.Tree
	zipTree *radix.Tree
}

func New() *fileMatcher {
	tree := radix.New(radix.Tdebug)
	tree.SetBoundaries('{', '}')

	matcher := &fileMatcher{
		tree:    tree,
		zipTree: radix.New(radix.Tdebug),
	}

	// PNG
	matcher.Register("89 50 4E 47 0D 0A 1A 0A", "png")

	// WebP
	matcher.Register("52 49 46 46", "webp")
	//matcher.Register("52 49 46 46 {fourBytes} 57 45 42 50 56 50 38", "webp")

	// PDF
	matcher.Register("25 50 44 46 2D", "pdf")

	// DOC
	matcher.Register("D0 CF 11 E0", "doc")

	matcher.Register("50 4B 03 04", "zip, aar, apk, docx, epub, ipa, jar, kmz, maff, msix, odp, ods, odt, pk3, pk4, pptx, usdz, vsdx, xlsx, xpi")
	matcher.Register("50 4B 05 06", "zip (empty)")
	matcher.Register("50 4B 07 08", "zip (spanned)")

	// TODO: Check header in ZIP 512 bytes
	// https://sceweb.sce.uhcl.edu/abeysekera/itec3831/labs/FILE%20SIGNATURES%20TABLE.pdf
	// matcher.zipTree.Add(label string, v interface{})

	matcher.RegisterJPEG()

	return matcher
}

func (m *fileMatcher) Register(hexHeader string, ext string) {
	hex := strings.ReplaceAll(hexHeader, " ", "")
	m.tree.Add(hex, ext)
	m.tree.Add(hex+"{end}", ext)
}

func (m *fileMatcher) MatchString(str string) string {
	str = strings.ReplaceAll(str, " ", "")

	node, submatches := m.tree.Get(str)

	if end, ok := submatches["end"]; ok {
		newStr := str[:len(str)-len(end)]
		node, submatches = m.tree.Get(newStr)
	}

	if node == nil {
		return spew.Sdump(str)
	}

	val, _ := node.Value.(string)
	return val
}

func (m *fileMatcher) String() string {
	return m.tree.String()
}
