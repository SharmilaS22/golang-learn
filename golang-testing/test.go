package golangtesting

import (
	"go/format"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_format(t *testing.T) {
	tests := []struct {
		// test cases
		testCaseName string
		testingInput int
	}{{"test with number 10", 10}, {"test with number 200", 200}}

	for _, tt := range tests {
		t.Run(tt.testCaseName, func(t *testing.T) {
			output, err := format(tt.testingInput)

			assert.NotNil(t, err)
			assert.Equal(t, output)
		})
	}
}

// https://medium.com/goingogo/why-use-testmain-for-testing-in-go-dafb52b406bc
func TestMain(m *testing.M) {
    // setup() //init
    code := m.Run() // run unit test
    // shutdown() //close later
    os.Exit(code)
}