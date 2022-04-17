package marsrover_test

import (
	"testing"

	marsrover "github.com/emrecavdar94/mars-rover/cmd/mars-rover"
	"github.com/emrecavdar94/mars-rover/cmd/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestValidationRoverCoordinate(t *testing.T) {
	roverCoordinatesText := "1 2 N"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	var mockPlateau = mocks.NewMockIPlateau(mockController)
	var rover = marsrover.NewRover(roverCoordinatesText, mockPlateau)
	mockPlateau.EXPECT().GetPlateauSize().Return(5, 5).Times(1)
	err := rover.ValidationRoverCoordinate()
	assert.Nil(t, err)
}
func TestValidationRoverCoordinateWithExceededCoordinateShouldReturnError(t *testing.T) {
	roverCoordinatesText := "7 8 N"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	var mockPlateau = mocks.NewMockIPlateau(mockController)
	var rover = marsrover.NewRover(roverCoordinatesText, mockPlateau)
	mockPlateau.EXPECT().GetPlateauSize().Return(5, 5).Times(1)
	err := rover.ValidationRoverCoordinate()
	assert.Equal(t, "Rover coordinates exceeded plateau", err.Error())
}

func TestValidationRoverCoordinateWitInvalidCoordinateShouldReturnError(t *testing.T) {
	roverCoordinatesText := "1 2 G"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	var mockPlateau = mocks.NewMockIPlateau(mockController)
	var rover = marsrover.NewRover(roverCoordinatesText, mockPlateau)
	mockPlateau.EXPECT().GetPlateauSize().Return(5, 5).Times(1)
	err := rover.ValidationRoverCoordinate()
	assert.Equal(t, "Rover direction is not valid", err.Error())
}

func TestValidationRoverDirection(t *testing.T) {
	roverCoordinatesText := "1 2 N"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	var mockPlateau = mocks.NewMockIPlateau(mockController)
	var rover = marsrover.NewRover(roverCoordinatesText, mockPlateau)
	err := rover.ValidationRoverDirection("L L R R M R M M L M")
	assert.Nil(t, err)
}
func TestValidationRoverDirectionWithEmptyDirectionShouldReturnError(t *testing.T) {
	roverCoordinatesText := "1 2 N"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	var mockPlateau = mocks.NewMockIPlateau(mockController)
	var rover = marsrover.NewRover(roverCoordinatesText, mockPlateau)
	err := rover.ValidationRoverDirection("")
	assert.Equal(t, "command cannot be empty", err.Error())
}

func TestValidationRoverDirectionWithInvalidDirectionShouldReturnError(t *testing.T) {
	roverCoordinatesText := "1 2 N"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPlateau := mocks.NewMockIPlateau(mockController)
	rover := marsrover.NewRover(roverCoordinatesText, mockPlateau)
	err := rover.ValidationRoverDirection("L M R Y L")
	assert.Equal(t, "invalid direction. you can only use  L, R and M key", err.Error())
}

func TestMoveRover(t *testing.T) {
	roverCoordinatesText := "1 2 N"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPlateau := mocks.NewMockIPlateau(mockController)
	rover := marsrover.NewRover(roverCoordinatesText, mockPlateau)
	rover.SetMovingCommand("L M L M L M L M M")
	mockPlateau.EXPECT().GetPlateauSize().Return(5, 5).Times(1)

	result, err := rover.MoveRover()
	assert.Nil(t, err)
	assert.Equal(t, "1 3 N", result)
}

func TestMoveRoverShouldReturnErrorIfRoverExceededPlateau(t *testing.T) {
	roverCoordinatesText := "1 2 N"
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	mockPlateau := mocks.NewMockIPlateau(mockController)
	rover := marsrover.NewRover(roverCoordinatesText, mockPlateau)
	rover.SetMovingCommand("L M L M M M M M M M M M M M")
	mockPlateau.EXPECT().GetPlateauSize().Return(5, 5).Times(1)

	result, err := rover.MoveRover()
	assert.Equal(t, "new coordinates are exeeded plateau", err.Error())
	assert.Equal(t, "", result)
}
