package main

import (
	"errors"
	"fmt"
	"log"
)

func area(length float64, width float64) (float64, error) {
	//check if length is greater than 0
	if length < 0 {
		err := errors.New("negative lenght not allowed")
		return -1, err
	}
	if width < 0 {
		err := errors.New("negative width not allowed")
		return -1, err
	}
	result := length * width
	return result, nil
}

func main() {
	length := 5.5
	width := 10
	result, err := area(length, float64(width))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

/*
git status
git commit -m "comment"
git log -oneline
git push -u origin main
git clone
git push git pull
*/
/*
create a repository on github called errors
*/

/*
do radius pi r2*/
