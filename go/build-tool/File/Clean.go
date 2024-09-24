package File

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Define a struct that matches the JSON structure
type Config struct {
	InputFolder      string   `json:"inputFolder"`
	OutputFolder     string   `json:"outputFolder"`
	CompileExtension []string `json:"compileExtension"`
	PublicExtension  []string `json:"publicExtension"`
}

// LoadConfig reads the JSON configuration file and returns a Config struct
func LoadConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// isHTMLorJS checks if the file is an HTML, JS, or TS file based on the configuration
func isHTMLorJS(path string, config Config) bool {
	ext := strings.ToLower(filepath.Ext(path))
	for _, allowedExt := range config.CompileExtension {
		if ext == "."+allowedExt {
			return true
		}
	}
	return false
}

func isPublicExtension(path string, config Config) bool {
	ext := strings.ToLower(filepath.Ext(path))
	for _, allowedExt := range config.PublicExtension {
		if ext == "."+allowedExt {
			return true
		}
	}
	return false
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}

// filesAreEqual checks if two files are identical
func filesAreEqual(file1, file2 string) bool {
	f1, err := os.Open(file1)
	if err != nil {
		return false
	}
	defer f1.Close()

	f2, err := os.Open(file2)
	if err != nil {
		return false
	}
	defer f2.Close()

	fi1, err := f1.Stat()
	if err != nil {
		return false
	}

	fi2, err := f2.Stat()
	if err != nil {
		return false
	}

	if fi1.Size() != fi2.Size() {
		return false
	}

	buf1 := make([]byte, fi1.Size())
	buf2 := make([]byte, fi2.Size())

	_, err = f1.Read(buf1)
	if err != nil {
		return false
	}

	_, err = f2.Read(buf2)
	if err != nil {
		return false
	}

	return string(buf1) == string(buf2)
}

func Clean() {
	// Load the configuration from the JSON file
	config, err := LoadConfig("conf.json")
	if err != nil {
		fmt.Printf("failed to load config: %s\n", err)
		return
	}

	inputDir := config.InputFolder
	outputDir := config.OutputFolder

	// Step 1: Walk through the outputDir folder and its subfolders
	filepath.Walk(outputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Get the relative path from outputDir
			relPath, err := filepath.Rel(outputDir, path)
			if err != nil {
				return err
			}
			// Check if the subfolder exists in inputDir
			srcPath := filepath.Join(inputDir, relPath)
			if _, err := os.Stat(srcPath); os.IsNotExist(err) {
				// If not, delete the subfolder in outputDir
				fmt.Printf("Deleting non-existent subfolder in src/app: %s\n", path)
				os.RemoveAll(path)
			}
			// } else if !isHTMLorJS(path, config) {
		} else if !isHTMLorJS(path, config) {
			// Get the relative path from outputDir
			relPath, err := filepath.Rel(outputDir, path)
			if err != nil {
				return err
			}
			// Check if the file exists in inputDir
			srcPath := filepath.Join(inputDir, relPath)
			if _, err := os.Stat(srcPath); os.IsNotExist(err) {
				// If not, delete the file in outputDir
				fmt.Printf("Deleting non-existent file in src/app: %s\n", path)
				os.Remove(path)
			} else if !filesAreEqual(srcPath, path) {
				// If the file exists but is different, delete the file in outputDir
				fmt.Printf("Deleting different file in dist: %s\n", path)
				os.Remove(path)
			}
		}
		return nil
	})

	// Step 2: Walk through the inputDir folder and its subfolders
	filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Get the relative path from inputDir
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}
			// Create the equivalent folder in outputDir if it doesn't exist
			distPath := filepath.Join(outputDir, relPath)
			if _, err := os.Stat(distPath); os.IsNotExist(err) {
				fmt.Printf("Creating missing subfolder in dist: %s\n", distPath)
				os.MkdirAll(distPath, 0755)
			}
		} else if !isHTMLorJS(path, config) && isPublicExtension(path, config) {
			// Copy all non-HTML and non-JS files to the equivalent folder in outputDir
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}
			distPath := filepath.Join(outputDir, relPath)
			if _, err := os.Stat(distPath); os.IsNotExist(err) {
				// If the file does not exist in outputDir, copy it
				fmt.Printf("Copying file to %s: %s\n", outputDir, distPath)
				copyFile(path, distPath)
			} else if !filesAreEqual(path, distPath) {
				fmt.Printf("Copying file to %s: %s\n", outputDir, distPath)
				copyFile(path, distPath)
			}
		}
		return nil
	})
}
