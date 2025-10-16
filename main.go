package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	dir := "./templates/pages"
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			fmt.Println(path)
			lines(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func lines(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//fmt.Println("Current working directory:", dir)

	includePattern := regexp.MustCompile(`{{include:(.+)}}`)

	var result []string

	for _, line := range lines {
		matches := includePattern.FindStringSubmatch(line)
		if len(matches) == 2 {
			includePath := matches[1]
			absPath := filepath.Join(filepath.Join(dir, "templates"), includePath)
			//fmt.Println("Include path:", includePath, "->", absPath)
			if !isWithinRoot(dir, absPath) {
				fmt.Println("Error: Attempted directory traversal in include path:", includePath)
				continue
			}
			content, err := os.ReadFile(absPath)
			if err != nil {
				fmt.Println("Error reading included file:", err)
				result = append(result, line) // fallback: keep original line
			} else {
				fmt.Println("Included file:", absPath)
				//fmt.Println("Replacing line:", line, "with content of length", len(content))
				result = append(result, string(content))
			}
		} else {
			result = append(result, line)
		}
	}

	writeLines(result, "./public", filepath.Base(path))
}

func writeLines(lines []string, dir, filename string) error {
	// Ensure the directory exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	path := filepath.Join(dir, filename)
	content := strings.Join(lines, "\n")
	return os.WriteFile(path, []byte(content), 0644)
}

func isWithinRoot(root, target string) bool {
	rel, err := filepath.Rel(root, target)
	if err != nil {
		return false
	}
	// Split the relative path and check if it starts with ".."
	parts := strings.Split(rel, string(filepath.Separator))
	return len(parts) == 0 || parts[0] != ".."
}

func startsWithDotDot(rel string) bool {
	return rel == ".." || filepath.HasPrefix(rel, ".."+string(filepath.Separator))
}

func main11234() {
	mainFile := "pages/about.html"
	includePattern := regexp.MustCompile(`{{include:(.+)}}`)

	file, err := os.Open(mainFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := includePattern.FindStringSubmatch(line)
		if len(matches) == 2 {
			includePath := matches[1]
			absPath := filepath.Join(filepath.Dir(mainFile), includePath)
			content, err := os.ReadFile(absPath)
			if err != nil {
				fmt.Println("Error reading included file:", err)
				result = append(result, line) // fallback: keep original line
			} else {
				result = append(result, string(content))
			}
		} else {
			result = append(result, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Example: print the result
	for _, line := range result {
		fmt.Println(line)
	}
}
