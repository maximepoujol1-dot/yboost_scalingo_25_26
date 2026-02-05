package api

//check and change the thuth table

func Verifswitch(verifi map[string]bool, motsVerif string) {

    switch motsVerif {
	

    case "verifDanemark":
        if verifi["verifDanemark"] == true {
			verifi["verifDanemark"] = false
        } else {
            verifi["verifDanemark"] = true
        }
	
    case "verifSuède":
        if verifi["verifSuède"] == true {
            verifi["verifSuède"] = false
        } else {
            verifi["verifSuède"] = true
        }

	case "verifNorvège":
        if verifi["verifNorvège"] == true {
            verifi["verifNorvège"] = false
        } else {
            verifi["verifNorvège"] = true
        }

	case "réinitialiser":
        if verifi["réinitialiser"] == true {
            verifi["réinitialiser"] = false
        } else {
            verifi["réinitialiser"] = true
        }
	
    case "verifFin":
        if verifi["verifFin"] == true {
            verifi["verifFin"] = false
        } else {
            verifi["verifFin"] = true
        }
	

    case "verifApoge":
        if verifi["verifApoge"] == true {
            verifi["verifApoge"] = false
        } else {
            verifi["verifApoge"] = true
        }
	

    case "verifExpansion":
        if verifi["verifExpansion"] == true {
            verifi["verifExpansion"] = false
        } else {
            verifi["verifExpansion"] = true
        }

	case "verifDébut":
        if verifi["verifDébut"] == true {
           verifi["verifDébut"] = false
        } else {
            verifi["verifDébut"] = true
        }
    }
}