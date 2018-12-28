package wahlin

import (
	"testing"
)

func Test(t *testing.T) {
	words := make(map[string]string)

	// C
	words["Ca"] = "KA"
	words["Co"] = "KO"
	words["Ace"] = "ASE"
	words["Censorn"] = "SENSORN"

	// Ch
	words["Christian"] = "KRISTIAN"

	// D
	words["Ddt"] = "DT"

	// F
	words["Af"] = "AV"
	words["Fa"] = "FA"
	words["zFb"] = "SVB"
	words["Dödsstraff"] = "DÖDSSTRAFF"

	// G
	words["Arg"] = "ARJ"
	words["vRg"] = "VRJ"

	// H
	words["Hand"] = "HAND"

	// I
	words["Hie"] = "HJE"
	words["Ei"] = "EJ"
	words["Iy"] = "JY"

	// K
	words["Skita"] = "*ITA"
	words["Kemi"] = "+EMI"
	words["Körkort"] = "+ÖRKORT"
	words["Krossa"] = "KROSSA"

	// KJ
	words["Skalkeskjulet"] = "SKALKE*ULET"
	words["Kjol"] = "+OL"
	words["Bastkjol"] = "BAST+OL"
	words["Kjäpp"] = "+EPP"

	// L
	words["Ljug"] = "JUG"

	// PH
	words["Philadelphia"] = "FILADELFJA"
	words["Upphettningen"] = "UPPHETTNINGEN"

	// Q, QU
	words["Quinoa"] = "KVINOA"
	words["Quiz"] = "KVIS"

	//
	words["Själ"] = "*EL"
	words["Stjärna"] = "*ERNA"
	words["Sköterska"] = "*ÖTERSKA"
	words["Skidor"] = "*IDOR"
	words["Sked"] = "*ED"
	words["Schema"] = "*EMA"
	words["Skjut"] = "*UT"
	words["Accentförskjutning"] = "ACSENTFÖR*UTNING"

	// TJ
	words["Tjuven"] = "+UVEN"

	// U
	words["Mousserande"] = "MOSSERANDE"

	// W
	words["Wunderbaum"] = "VUNDERBAUM"

	// X
	words["Alex"] = "ALEKS"

	// Y
	words["Aye"] = "AJE"

	// Å
	words["Ångest"] = "ONGEST"
	words["Håv"] = "HOV"

	//
	//words[""] = ""

	for word, encoded := range words {
		result := Encode(word)
		if result != encoded {
			t.Errorf("%s FAILED: wanted %s, got %s\n", word, encoded, result)
		} else {
			t.Logf("%s succeeded: wanted %s, got %s\n", word, encoded, result)
		}
	}
}
