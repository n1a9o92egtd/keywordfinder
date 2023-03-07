package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func KMP(text, pattern string) bool {
	n, m := len(text), len(pattern)

	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}

	j := 0
	for i := 0; i < n; i++ {
		for j > 0 && text[i] != pattern[j] {
			j = pi[j-1]
		}
		if text[i] == pattern[j] {
			j++
		}
		if j == m {
			return true
		}
	}

	return false
}

func findKeywords(root string, keywords []string) ([]string, error) {
	var result []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == root {
			return nil
		}

		if info.IsDir() {
			return nil
		}
		for _, keyword := range keywords {
			if KMP(info.Name(), keyword) {
				result = append(result, path)
				break
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./keywordfinder <target path> <keywords separated by commas>\nExample: ./keywordfinder /path/to/directory keyword1,keyword2,keyword3")
		return
	}
	start := time.Now()

	root := os.Args[1]
	keywords := strings.Split(os.Args[2], ",")
	if len(keywords) == 0 {
		keywords = append(keywords, os.Args[2])
	}

	result, err := findKeywords(root, keywords)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Files and directories containing keywords:")
	for _, path := range result {
		fmt.Println(path)
	}
	end := time.Now()
	fmt.Printf("The program running time is %v\n", end.Sub(start))
}
