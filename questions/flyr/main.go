// Given the following possible flight legs, write a function, findRoutes(origin, destination) that efficiently outputs all possible routes from origin to destination.
//
// # Origin | Destination
//
// Seattle | LA
// LA | Orlando
// LA | Augusta
// Orlando | Seattle
// Seattle | Orlando
// LA | Portland
// Portland | Orlando
// An example output could be:
//
// > findRoutes('Seattle', 'Orlando')
// ['Seattle -> Orlando', 'Seattle -> LA -> Orlando', 'Seattle -> LA -> Portland -> Orlando']
package main

import (
	"errors"
	"fmt"
	"strings"
)

var ods [][]string

var errorNoDirectRoute = errors.New("no direct route")

var allDirectRoutes = make(map[string][]string)

func findRoutes(origin string, destination string) []string {
	routes := make([]string, 0)

	directHops, err := findDirectRoutes(origin)
	if err != nil {
		fmt.Println(err)
		return routes
	}

	for _, directHop := range directHops {
		route := []string{origin}
		route, err = findRoute(directHop, destination, route)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(route) > 1 {
			routes = append(routes, strings.Join(route, " -> "))
		}
	}

	return routes
}

func findRoute(hop string, destination string, route []string) ([]string, error) {
	route = append(route, hop)
	if hop == destination {
		return route, nil
	}

	hops, err := findDirectRoutes(hop)
	if err != nil {
		if errors.Is(err, errorNoDirectRoute) {
			return route, fmt.Errorf("%w: %s", errorNoDirectRoute, hop)
		}
		panic("don't want to get to this point")
	}

	for _, newHop := range hops {
		route, err = findRoute(newHop, destination, route)
		if err != nil {
			continue
		}
		if newHop == destination {
			break
		}
	}

	return route, nil
}

func findDirectRoutes(origin string) ([]string, error) {
	if destinations, ok := allDirectRoutes[origin]; ok {
		return destinations, nil
	}

	return nil, errorNoDirectRoute
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
