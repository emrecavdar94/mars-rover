package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	marsrover "github.com/emrecavdar94/mars-rover/cmd/mars-rover"
)

func main() {
	var err error
	var rover *marsrover.Rover
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the plateau size: ")
	plateauSizeTxt, _ := reader.ReadString('\n')
	plateau := marsrover.NewPlateau(plateauSizeTxt)
	err = plateau.ValidatePlateauSize()
	for err != nil {
		fmt.Println(err.Error(), ". Please enter the valid plateau size.")
		plateauSizeTxt, _ = reader.ReadString('\n')
		plateau = marsrover.NewPlateau(plateauSizeTxt)
		err = plateau.ValidatePlateauSize()
	}

	isPlateauDiscoveryEnough := false
	for !isPlateauDiscoveryEnough {
		fmt.Println("Please enter the rover coordinates: ")
		roverCoordinatesText, _ := reader.ReadString('\n')
		rover = marsrover.NewRover(roverCoordinatesText, plateau)

		for err != nil {
			fmt.Println(err.Error(), ". Please enter the valid rover coordinates.")
			roverCoordinatesText, _ = reader.ReadString('\n')
			rover = marsrover.NewRover(roverCoordinatesText, plateau)
			err = rover.ValidationRoverCoordinate()
		}

		fmt.Println("Please enter the direction: ")
		directionText, _ := reader.ReadString('\n')
		err = rover.ValidationRoverDirection(directionText)
		for err != nil {
			fmt.Println(err.Error(), ". Please enter the valid direction.")
			directionText, _ = reader.ReadString('\n')
			err = rover.ValidationRoverDirection(directionText)
		}

		rover.SetMovingCommand(directionText)
		result, err := rover.MoveRover()
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("NEW COORDINATE => %s \n", result)
		fmt.Println("Do you want to add another rover : ( Y / N) :")
		isDiscoverChoosingValid := false
		for !isDiscoverChoosingValid {
			userAnswer, _ := reader.ReadString('\n')
			trimmedVal := strings.
				ToLower(strings.
					TrimSpace(strings.
						ReplaceAll(userAnswer, " ", "")))
			if trimmedVal != "y" && trimmedVal != "n" {
				fmt.Print("\n Invalid answer please choose: ( Y / N) :")
				isDiscoverChoosingValid = false
			} else {
				if trimmedVal == "y" {
					isPlateauDiscoveryEnough = false
				} else {
					isPlateauDiscoveryEnough = true
					fmt.Println("\n Thank you for using Mars Rover")
				}
				isDiscoverChoosingValid = true
			}
		}
	}
}
