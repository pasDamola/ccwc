package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)


var (
	binName = "ccwc"
	fileName = "test.txt"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
	  binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
	  fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
	  os.Exit(1)
	}
	fmt.Println("Running tests....")
	result := m.Run()
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Exit(result)
}

func TestCountLines(t *testing.T) {

}

func TestWCCLI(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
	  t.Fatal(err)
	}
  
	cmdPath := filepath.Join(dir, binName)

	t.Run("Count lines from flag arguments", func(t *testing.T) {
		expectedNumOfLines := 7145
		cmd := exec.Command(cmdPath, "-l", fileName)
	
		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%d %s\n", expectedNumOfLines, fileName)

		
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	
	  })

	  t.Run("Count lines from STDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-l", fileName)
		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
		 t.Fatal(err)
		}

	  	io.WriteString(cmdStdIn, fileName)
	  	cmdStdIn.Close()
	  
	  if err := cmd.Run(); err != nil {
		 t.Fatal(err)
	   }
	  })

	  t.Run("Count bytes from flag arguments", func(t *testing.T) {
		expectedNumOfBytes := 342190
		cmd := exec.Command(cmdPath, "-c", fileName)
	
		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%d %s\n", expectedNumOfBytes, fileName)

		
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	
	  })

	  t.Run("Count words from flag arguments", func(t *testing.T) {
		expectedNumOfBytes := 58164
		cmd := exec.Command(cmdPath, "-w", fileName)
	
		out, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatal(err)
		}

		expected := fmt.Sprintf("%d %s\n", expectedNumOfBytes, fileName)

		
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	
	  })
}