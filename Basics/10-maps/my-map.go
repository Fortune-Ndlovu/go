package main
import "fmt"

func main () {
	superHero := map[string]map[string]string{
		"Superman" : map[string]string {
			"RealName": "Clark Kent",
			"City": "Metropolis",
		},

		"Batman" : map[string] string {
			"RealName": "Bruce Wayne",
			"City": "Gotham",
		},
	}
	if temp, hero := superHero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}