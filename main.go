package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func plan(root string) (map[string][]string, error) {
	moves := map[string][]string{}
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(e.Name())), ".")
		if ext == "" {
			ext = "misc"
		}
		moves[ext] = append(moves[ext], e.Name())
	}
	return moves, nil
}

func apply(root string, moves map[string][]string) error {
	for ext, names := range moves {
		dir := filepath.Join(root, ext)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		for _, n := range names {
			src := filepath.Join(root, n)
			dst := filepath.Join(dir, n)
			if err := os.Rename(src, dst); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	dry := flag.Bool("dry-run", false, "print plan only")
	flag.Parse()
	root := flag.Arg(0)
	if root == "" {
		fmt.Println("usage: tidydesk <dir> [--dry-run]")
		os.Exit(2)
	}
	moves, err := plan(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	total := 0
	for ext, names := range moves {
		fmt.Printf("%s/  <-  %d files\n", ext, len(names))
		total += len(names)
	}
	if *dry || total == 0 {
		return
	}
	if err := apply(root, moves); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Printf("moved %d files\n", total)
}
