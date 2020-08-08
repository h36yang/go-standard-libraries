package main

import (
	"fmt"

	"./media"
)

func main() {
	var myFav media.Catalogable = &media.Movie{}
	myFav.NewMovie("Top Gun", media.R, 43.2)

	fmt.Printf("My favourite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %s\n", myFav.GetRating())
	fmt.Printf("It made %f in the box office\n", myFav.GetBoxOffice())

	myFav.SetTitle("Dumb and Dumber")
	fmt.Printf("My favourite movie is %s\n", myFav.GetTitle())
}
