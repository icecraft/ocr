package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestITaxCard(t *testing.T) {

	t.Run("itax card", func(t *testing.T) {
		t1 := []string{"fT", "3eR", "中口", "ER", "INCOMETAXDEPARTMENT", "GOVTOFINDLA", "taRel teaal", "Permanent Account Number Card", "ATHPL6094K", "/Name", "LAKHADE CHANDOJI NAMDEC", "uaTTE/FathersName", "NAMDEO LAKHADE", "Harfrg/Date of Birll", "分", "01/01/1992", "Chda", "RH8R/Slgnalure"}
		it, err := NewITaxCardFromArr(t1)
		assert.NoError(t, err)
		assert.Equal(t, it.ID, "ATHPL6094K")
		assert.Equal(t, it.Name, "LAKHADE CHANDOJI NAMDEC")
		assert.Equal(t, it.FatherName, "NAMDEO LAKHADE")
	})
}
