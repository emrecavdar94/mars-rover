package marsrover

import (
	"errors"
	"strconv"
	"strings"
)

type Plateau struct {
	plateauSizeX int
	plateauSizeY int
}

func NewPlateau(plateauSizeTxt string) *Plateau {
	plateauSizeX, _ := strconv.Atoi(strings.Fields(plateauSizeTxt)[0])
	plateauSizeY, _ := strconv.Atoi(strings.Fields(plateauSizeTxt)[1])

	return &Plateau{
		plateauSizeX: plateauSizeX,
		plateauSizeY: plateauSizeY,
	}
}

func (p *Plateau) ValidatePlateauSize() error {
	if p.plateauSizeX < 0 || p.plateauSizeY < 0 {
		return errors.New("Plateau size is not valid")
	}

	return nil
}

func (p *Plateau) GetPlateauSize() (int, int) {
	return p.plateauSizeX, p.plateauSizeY
}
