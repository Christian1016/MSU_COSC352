package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// Fetch HTML from URL
func fetchHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Parse HTML and extract tables
func extractTables(htmlContent []byte) ([][]string, error) {
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var tables [][]string
	var currentTable []string
	var isInsideTable bool

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "table" {
				isInsideTable = true
				currentTable = []string{}
			}

			if isInsideTable && n.Data == "tr" {
				var row []string
				for _, child := range n.ChildNodes {
					if child.Data == "td" || child.Data == "th" {
						for _, cell := range child.ChildNodes {
							if cell.Type == html.TextNode {
								row = append(row, cell.Data)
							}
						}
					}
				}
				if len(row) > 0 {
					currentTable = append(currentTable, strings.Join(row, ","))
				}
			}
		}

		if n.Type == html.ElementNode && n.Data == "table" {
			tables = append(tables, currentTable)
			isInsideTable = false
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)
	return tables, nil
}

// Save tables to CSV files
func saveTablesToCSV(tables [][]string) {
	for i, table := range tables {
		fileName := fmt.Sprintf("table_%d.csv", i+1)
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", fileName, err)
			continue
		}
		defer file.Close()

		for _, row := range table {
			_, err := file.WriteString(row + "\n")
			if err != nil {
				fmt.Printf("Error writing to file %s: %v\n", fileName, err)
				break
			}
		}

		fmt.Printf("Table %d saved as %s\n", i+1, fileName)
	}
}

func main() {
	// Download HTML from the Wikipedia page
	url := "https://en.wikipedia.org/wiki/List_of_largest_companies_in_the_United_States_by_revenue"
	htmlContent, err := fetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching HTML:", err)
		return
	}

	// Extract tables from the HTML content
	tables, err := extractTables(htmlContent)
	if err != nil {
		fmt.Println("Error extracting tables:", err)
		return
	}

	// Save each table to a CSV file
	saveTablesToCSV(tables)
}
