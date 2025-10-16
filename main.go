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

	result, err := processLines(lines, dir, filepath.Dir(path))
	if err != nil {
		fmt.Println("Error processing lines:", err)
		return
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

// Extract meta values from lines before the include
func extractMeta(lines []string, includeIdx int) (title, desc, keywords string, err error) {
	for i := includeIdx - 1; i >= 0; i-- {
		line := lines[i]
		if strings.HasPrefix(line, "{{title:") {
			title = strings.TrimSuffix(strings.TrimPrefix(line, "{{title:"), "}}")
		}
		if strings.HasPrefix(line, "{{description:") {
			desc = strings.TrimSuffix(strings.TrimPrefix(line, "{{description:"), "}}")
		}
		if strings.HasPrefix(line, "{{keywords:") {
			keywords = strings.TrimSuffix(strings.TrimPrefix(line, "{{keywords:"), "}}")
		}
	}
	if title == "" || desc == "" || keywords == "" {
		return "", "", "", fmt.Errorf("Missing meta tag(s) before header include")
	}
	return title, desc, keywords, nil
}

// Replace meta tags in header.html content
func replaceHeaderMeta(headerContent, title, desc, keywords string) (string, error) {
	if !strings.Contains(headerContent, "<title>") ||
		!strings.Contains(headerContent, "name=\"description\"") ||
		!strings.Contains(headerContent, "name=\"keywords\"") {
		return "", fmt.Errorf("Header missing required meta tags")
	}
	headerContent = regexp.MustCompile(`<title>.*?</title>`).ReplaceAllString(headerContent, "<title>"+title+"</title>")
	headerContent = regexp.MustCompile(`<meta name=\"description\" content=\".*?\" ?/?>`).ReplaceAllString(headerContent, `<meta name="description" content="`+desc+`" />`)
	headerContent = regexp.MustCompile(`<meta name=\"keywords\" content=\".*?\" ?/?>`).ReplaceAllString(headerContent, `<meta name="keywords" content="`+keywords+`" />`)
	return headerContent, nil
}

// Recursively process lines for template inclusion
func processLines(lines []string, root string, currentDir string) ([]string, error) {
	includePattern := regexp.MustCompile(`{{include:(.+)}}`)
	var result []string
	for _, line := range lines {
		matches := includePattern.FindStringSubmatch(line)
		if len(matches) == 2 {
			includePath := matches[1]
			absPath := filepath.Join(root, "templates", includePath)

			if !isWithinRoot(root, absPath) {
				result = append(result, fmt.Sprintf("<!-- Error: Attempted directory traversal in include path: %s -->", includePath))
				continue
			}

			// Special handling for header.html to replace meta tags
			if includePath == "header.html" {
				title, desc, keywords, err := extractMeta(lines, func() int {
					for i, l := range lines {
						if l == line {
							return i
						}
					}
					return -1
				}())
				if err != nil {
					result = append(result, fmt.Sprintf("<!-- Error extracting meta tags for header include: %s -->", err.Error()))
					continue
				}
				// todo: do something with these items.
				fmt.Println("Extracted meta - Title:", title, "Description:", desc, "Keywords:", keywords)
			}

			file, err := os.Open(absPath)
			if err != nil {
				result = append(result, fmt.Sprintf("<!-- Error reading included file: %s -->", includePath))
				continue
			}
			var includedLines []string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				includedLines = append(includedLines, scanner.Text())
			}
			file.Close()
			if err := scanner.Err(); err != nil {
				result = append(result, fmt.Sprintf("<!-- Error reading included file: %s -->", includePath))
				continue
			}
			// Recursively process included lines
			recursed, err := processLines(includedLines, root, filepath.Dir(absPath))
			if err != nil {
				result = append(result, fmt.Sprintf("<!-- Error processing included file: %s -->", includePath))
				continue
			}
			result = append(result, recursed...)
		} else {
			result = append(result, line)
		}
	}
	return result, nil
}
