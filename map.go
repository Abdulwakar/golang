package main
import "fmt"


func main() {
   StudentAge := make(map[string] int)

   StudentAge["Abdul"] = 27
   StudentAge["Abdul1"] = 37
   StudentAge["Abdul2"] = 57
   StudentAge["Abdul3"] = 17
   StudentAge["Abdul4"] = 29
   StudentAge["Abdul5"] = 22
   StudentAge["Abdul6"] = 21

   fmt.Println(StudentAge["Abdul4"])
   fmt.Println(len(StudentAge))


   fmt.Println("Multiple Map example")

   SuperHero := map[string]map[string]string{
        "SuperMan" : map[string]string{
            "RealName": "Clark Kent",
            "City" : "Metropolis",
        },

        "BatMan" : map[string]string{
            "RealName" : "Bruce Wayne",
            "City" : "Gotham City",
        },

        "Kriss" : map[string]string{
            "RealName" : "Hritik Roshan",
            "City" : "Mumbai",
        },


   }
   if temp, hero := SuperHero["BatMan"]; hero{
        fmt.Println(temp["RealName"], temp["City"])
   }
   fmt.Println(len(SuperHero[]))
}