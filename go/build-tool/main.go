package main

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type PageData struct {
	Title     string
	Items     []string
	Condition bool
}

func main() {
	// tmpl := template.Must(template.ParseFiles("index.html", "partial.html"))
	// tmpl := template.Must(template.ParseGlob("src/**/*.html"))

	// data := PageData{
	// 	Title:     "Main HTML",
	// 	Items:     []string{"Item 1", "Item 2", "Item 3"},
	// 	Condition: true,
	// }

	// // Create a file to write the output
	// file, err := os.Create("dist/index.html")
	// if err != nil {
	// 	fmt.Printf("Failed to create file: %v\n", err)
	// }
	// defer file.Close()

	// err = tmpl.Execute(file, data)
	// if err != nil {
	// 	fmt.Printf("Failed to execute template: %v\n", err)
	// }

	// fmt.Printf("Built done\n")

	// Parse all HTML files in the src/components directory as partial templates
	partialTmpl := template.Must(template.ParseGlob("src/components/*.html"))

	inputDir := "src/app"
	outputDir := "dist"

	// Traverse the src/app directory and its subdirectories to find HTML files
	filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".html" {

			fmt.Printf("Processing %s\n", path)

			// Parse the current HTML file and include the partial templates
			tmpl := template.Must(partialTmpl.Clone())
			tmpl = template.Must(tmpl.ParseFiles(path))

			data := PageData{
				Title:     "Main HTML",
				Items:     []string{"Item 1", "Item 2", "Item 3"},
				Condition: true,
			}

			// Create the output directory if it doesn't exist
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				fmt.Printf("Failed to create output directory: %v\n", err)
				return err
			}

			// Get the relative path from src/app
			relPath, err := filepath.Rel(inputDir, path)
			if err != nil {
				fmt.Printf("Failed to get relative path: %v\n", err)
				return err
			}

			// Create the output path in the dist directory
			outputPath := filepath.Join(outputDir, relPath)

			// Ensure the output directory structure exists
			if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
				fmt.Printf("Failed to create output directory structure: %v\n", err)
				return err
			}

			// Create a file to write the output
			// fmt.Printf("Output %s\n", outputPath)
			file, err := os.Create(outputPath)
			if err != nil {
				fmt.Printf("Failed to create file: %v\n", err)
				return err
			}
			defer file.Close()

			// Execute the template and write the output to the file
			err = tmpl.ExecuteTemplate(file, filepath.Base(path), data)
			if err != nil {
				fmt.Printf("Failed to execute template: %v\n", err)
				return err
			}

			fmt.Printf("Built %s\n", outputPath)
		}
		return nil
	})

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
		} else if !isHTMLorJS(path) {
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
		} else if !isHTMLorJS(path) {
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

// isHTMLorJS checks if the file is an HTML or JS file
func isHTMLorJS(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".html" || ext == ".js" || ext == ".ts"
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
