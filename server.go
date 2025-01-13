package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func executeGoCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Read Go code from request
		code, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		// Create temporary directory
		tempDir, err := os.MkdirTemp("", "tempdir")
		if err != nil {
				http.Error(w, "Failed to create temp directory", http.StatusInternalServerError)
				return
		}
		defer os.RemoveAll(tempDir)
		// Save code to temporary file
		tempFile := tempDir + "/temp.go"
		err = os.WriteFile(tempFile, code, 0644)
		if err != nil {
			http.Error(w, "Failed to save code", http.StatusInternalServerError)
			return
		}

		// Execute Go code
		cmd := exec.Command("go", "run", tempFile)
		output, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Execution error: %v\n%s", err, output), http.StatusInternalServerError)
			return
		}

		// Send execution result
		w.Write(output)
	}
}

func main() {
	http.HandleFunc("/execute", executeGoCode)
	fmt.Println("Server running on http://localhost:4321")
	http.ListenAndServe(":4321", nil)
}
