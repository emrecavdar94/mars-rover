package marsrover

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var DirectionRotateNumbers = map[string]int{
	"N": 1,
	"W": 2,
	"S": 3,
	"E": 4,
}

type Rover struct {
	xCoordinate   int
	yCoordinate   int
	direction     string
	plateau       IPlateau
	movingCommand []rune
}

type IPlateau interface {
	GetPlateauSize() (int, int)
}

func NewRover(roverCoordinatesText string,
	plateau IPlateau) *Rover {
	var xCoordinate int
	var yCoordinate int
	var direction string
	data := strings.Fields(roverCoordinatesText)
	xCoordinate, _ = strconv.Atoi(data[0])
	yCoordinate, _ = strconv.Atoi(data[1])
	direction = data[2]
	return &Rover{
		xCoordinate: xCoordinate,
		yCoordinate: yCoordinate,
		direction:   direction,
		plateau:     plateau,
	}
}

func (r *Rover) ValidationRoverCoordinate() error {
	plateauSizeX, plateauSizeY := r.plateau.GetPlateauSize()
	var isDirectionValid bool
	if r.direction == "E" || r.direction == "W" ||
		r.direction == "S" || r.direction == "N" {
		isDirectionValid = true
	}

	if !isDirectionValid {
		return errors.New("Rover direction is not valid")
	}

	if r.xCoordinate > plateauSizeX || r.yCoordinate > plateauSizeY {
		return errors.New("Rover coordinates exceeded plateau")
	}

	return nil
}

func (r *Rover) SetMovingCommand(directionText string) {
	trimmedVal := strings.ReplaceAll(directionText, " ", "")
	roverRotateAndMovingInfo := strings.TrimSpace(trimmedVal)
	r.movingCommand = []rune(roverRotateAndMovingInfo)
}

func (r *Rover) ValidationRoverDirection(val string) error {
	roverDirections := strings.Fields(val)
	if len(roverDirections) == 0 {
		return errors.New("command cannot be empty")
	}

	for _, direction := range roverDirections {
		directionString := string(direction)
		if directionString != "L" && directionString != "R" && directionString != "M" {
			return errors.New("invalid direction. you can only use  L, R and M key")
		}
	}
	return nil

}

func (r *Rover) findRoversNewDirection(rotate string) (newDirection string) {
	var newDirectionNum int

	directionNum := DirectionRotateNumbers[r.direction]

	if rotate == "L" && directionNum != DirectionRotateNumbers["E"] {
		newDirectionNum = directionNum + 1
	} else if rotate == "L" && directionNum == DirectionRotateNumbers["E"] {
		newDirectionNum = DirectionRotateNumbers["N"]
	} else if directionNum == DirectionRotateNumbers["N"] {
		newDirectionNum = DirectionRotateNumbers["E"]
	} else {
		newDirectionNum = directionNum - 1
	}
	for key, value := range DirectionRotateNumbers {
		if value == newDirectionNum {
			newDirection = key
		}
	}
	return newDirection
}

func (r *Rover) findNewCoordinates(newDirection string) (newXcoordinate, newYcoordinate int) {
	newXcoordinate = r.xCoordinate
	newYcoordinate = r.yCoordinate
	switch r.direction {
	case "N":
		newYcoordinate = r.yCoordinate + 1
	case "W":
		newXcoordinate = r.xCoordinate - 1
	case "E":
		newXcoordinate = r.xCoordinate + 1
	case "S":
		newYcoordinate = r.yCoordinate - 1
	}

	return newXcoordinate, newYcoordinate
}

func (r *Rover) MoveRover() (string, error) {
	newDirection := r.direction
	plateauSizeX, plateauSizeY := r.plateau.GetPlateauSize()
	for i := 0; i < len(r.movingCommand); i++ {
		currentVal := string(r.movingCommand[i])
		if currentVal == "L" || currentVal == "R" {
			newDirection = r.findRoversNewDirection(string(r.movingCommand[i]))
			r.direction = newDirection
		} else {
			newXCoordinate, newYCoordinate := r.findNewCoordinates(newDirection)
			if newXCoordinate > plateauSizeX || newYCoordinate > plateauSizeY ||
				newXCoordinate < 0 || newYCoordinate < 0 {
				return "", errors.New("new coordinates are exeeded plateau")

			} else {
				r.xCoordinate = newXCoordinate
				r.yCoordinate = newYCoordinate
			}
		}

	}

	return fmt.Sprintf("%d %d %s", r.xCoordinate, r.yCoordinate, newDirection), nil
}
