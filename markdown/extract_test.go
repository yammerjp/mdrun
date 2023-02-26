package markdown

import (
	"testing"
)

func TestExtractFile(t *testing.T) {
	commandBlocks, err := ExtractFile("extract_test.go.md")
	if err != nil {
		t.Errorf("Error occuerd: %s", err)
	}
	if len(commandBlocks) != 3 {
		t.Errorf("len(commandBlocks) must be 3, but %d", len(commandBlocks))
	}
	onelineStrings := []string{
		"go install github.com/yammerjp/mdrun",
		"mdrun list",
		"echo \"hello, world\"; echo \"hello, world\"; ls -al; echo \"hello\"; ls -al",
	}
	for i, commandBlock := range commandBlocks {
		expected := onelineStrings[i]
		actual := commandBlock.CommandOneLineString()
		if expected != actual {
			t.Errorf("commandBlocks[%d] must be\n%s\nbut\n%s\n", i, expected, actual)
		}
	}
}
