package main

import (
	"fmt"
	"testing"

	"github.com/stolostron/go-template-utils/v6/cmd/template-resolver/utils"
)

func doTest(inputFile, expectedOutputFile string) error {
	inputBytes, err := utils.HandleFile(inputFile)
	if err != nil {
		return err
	}

	expectedBytes, err := utils.HandleFile(expectedOutputFile)
	if err != nil {
		return err
	}

	// These last two parameters might need to be actual variables for future tests
	resolvedYAML, err := utils.ProcessTemplate(inputBytes, kubeconfigPath, "local-cluster")
	if err != nil {
		return err
	}

	if string(expectedBytes) != string(resolvedYAML) {
		fmt.Println("Wanted:\n", string(expectedBytes))
		fmt.Println("\nGot:\n", string(resolvedYAML))

		return fmt.Errorf("Mismatch")
	}

	return nil
}

func TestCLI(t *testing.T) {
	for _, num := range []string{"01", "02"} {
		err := doTest("testdata/tests/"+num+".input.yaml", "testdata/tests/"+num+".output.yaml")
		if err != nil {
			t.Fatal(err)
		}
	}
}
