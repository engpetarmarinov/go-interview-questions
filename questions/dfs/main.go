package main

import (
	"fmt"
	"strings"
)

var ods [][]string
var allDirectRoutes = map[string][]string{}

func findRoutes(origin string, destination string) []string {
	var routes []string
	var dfs func(from string, path []string)
	dfs = func(current string, path []string) {
		path = append(path, current)

		if current == destination {
			routes = append(routes, strings.Join(path, " -> "))
			return
		}

		if hops, ok := allDirectRoutes[current]; ok {
			for _, hop := range hops {
				dfs(hop, path)
			}
		}
	}

	dfs(origin, []string{})

	return routes
}

func main() {
	routes := findRoutes("Seattle", "Orlando")
	for _, route := range routes {
		fmt.Println(route)
	}
}

func init() {
	ods = [][]string{
		{"Seattle", "LA"},
		{"LA", "Orlando"},
		{"LA", "Augusta"},
		{"Orlando", "Seattle"},
		{"Seattle", "Orlando"},
		{"LA", "Portland"},
		{"Portland", "Orlando"},
	}

	initMapWithDirectRoutes(ods)
}

func initMapWithDirectRoutes(ods [][]string) {
	for _, route := range ods {
		origin := route[0]
		dest := route[1]
		allDirectRoutes[origin] = append(allDirectRoutes[origin], dest)
	}
}
