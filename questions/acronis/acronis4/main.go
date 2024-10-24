// Order-of-dependency. We have a dependency mapping like this
//
// {
// "a": [],
// "b": ["a"],
// "c": [],
// "d": ["c", "b"],
// "e": ["a", "d"],
// "f": ["a", "f"],
// "g": ["d", "h"],
// "h": ["a", "g"],
// "i": ["h", "a"],
// "k": ["a", "z"]
// }
//
// Implement a function which for given item returns ordered list of items it depends on.
//
// For example:
//
// +-------------+-------------------+
//
// | input       | output            |
//
// +-------------+-------------------+
//
// | a           | []                |
//
// | b           | a                 |
//
// | c           | []                |
//
// | d           | c, a, b           |
//
// | e           | a, c, b, d        |
//
// | f           | ?                 |
//
// | g           | ?                 |
//
// | h           | ?                 |
//
// | i           | ?                 |
//
// | k           | ?                 |
//
// +-------------+-------------------+
package main

import (
	"fmt"
	"log/slog"
)

var deps = map[string][]string{
	"a": {},
	"b": {"a"},
	"c": {},
	"d": {"c", "b"},
	"e": {"a", "d"},
	"f": {"a", "f"},
	"g": {"d", "h"},
	"h": {"a", "g"},
	"i": {"h", "a"},
	"k": {"a", "z"},
}

func main() {
	for _, letter := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k"} {
		fmt.Printf("getDeps(%s): ", letter)
		dependencies, err := getDeps(letter)
		if err != nil {
			slog.Error("getDeps", slog.String("error", err.Error()))
			continue
		}

		fmt.Println(dependencies)
	}
}

func getDeps(in string) ([]string, error) {
	dependencies := []string{}
	visited := map[string]any{}
	circular := map[string]bool{}

	var dfs func(in string, routes []string) error
	dfs = func(in string, routes []string) error {
		if circularDep, ok := circular[in]; ok && circularDep {
			return fmt.Errorf("circular dependency of %s", in)
		}

		if _, ok := visited[in]; ok {
			return nil
		}

		visited[in] = nil
		circular[in] = true
		defer func() {
			circular[in] = false
			routes = append(routes, in)
			dependencies = append(dependencies, routes...)
		}()

		if _, ok := deps[in]; !ok {
			return fmt.Errorf("dependency %s not found", in)
		}

		for _, dep := range deps[in] {
			if err := dfs(dep, routes); err != nil {
				return err
			}
		}

		return nil
	}

	err := dfs(in, []string{})
	if err != nil {
		return nil, err
	}

	return dependencies[:len(dependencies)-1], nil
}
