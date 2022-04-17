package marsrover_test

import (
	"testing"

	marsrover "github.com/emrecavdar94/mars-rover/cmd/mars-rover"
	"github.com/stretchr/testify/assert"
)

func TestValidatePlateauSize(t *testing.T) {
	var plateauSizeTxt = "5 5"
	var plateau = marsrover.NewPlateau(plateauSizeTxt)
	var err = plateau.ValidatePlateauSize()
	assert.Nil(t, err)
}
func TestValidatePlateauSizeWithInvalidSizeShouldReturnError(t *testing.T) {
	var plateauSizeTxt = "5 -5"
	var plateau = marsrover.NewPlateau(plateauSizeTxt)
	var err = plateau.ValidatePlateauSize()
	assert.Equal(t, "Plateau size is not valid", err.Error())
}

func TestGetPlateauSize(t *testing.T) {
	var plateauSizeTxt = "5 5"
	var plateau = marsrover.NewPlateau(plateauSizeTxt)
	var plateauSizeX, plateauSizeY = plateau.GetPlateauSize()
	assert.Equal(t, 5, plateauSizeX)
	assert.Equal(t, 5, plateauSizeY)
}
