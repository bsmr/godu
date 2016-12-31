package interactive

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/viktomas/godu/core"
)

func TestDoesntGoPastRoot(t *testing.T) {
	testTree := &core.File{"b", 180, []*core.File{
		&core.File{"d", 100, []*core.File{
			&core.File{"e", 10, []*core.File{}},
		}},
	}}
	input := "0\nb\nb\n"
	expected := "b/\n"
	expected += "---\n"
	expected += "0) d/ 100B\n"
	expected += "b/d/\n"
	expected += "---\n"
	expected += "0) e 10B\n"
	expected += "b/\n"
	expected += "---\n"
	expected += "0) d/ 100B\n"
	expected += "b/\n"
	expected += "---\n"
	expected += "0) d/ 100B\n"
	testInteractive(testTree, input, expected, t)

}

func testInteractive(tree *core.File, input string, expected string, t *testing.T) {
	reader := strings.NewReader(input)
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	InteractiveTree(tree, writer, reader, 0)
	writer.Flush()
	result := buffer.String()
	if result != expected {
		t.Errorf("expected:\n%sbut got:\n%s", expected, result)
	}
}