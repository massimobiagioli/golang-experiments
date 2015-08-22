package utility

import (	
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatePersona(t *testing.T) {
	persona := CreatePersona("Gianna", "Azzurri")
	assert.True(t, persona != nil, "Impossibile creare persona")
}

func TestInterface(t *testing.T) {
	quadrato := Quadrato{X: 4, Y: 3}
	forma := FormaGeometrica(quadrato)
	area := forma.Area()
	assert.Equal(t, 12, area, "Calcolo area errato")
}