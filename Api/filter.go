package api

//check and change the thuth table

func Verifswitch(verifi map[string]bool, motsVerif string) {

    switch motsVerif {
	

    case "verifDanemark":
        if verifi["verifDanemark"] == true {
			verifi["verifDanemark"] = false
        } else {
            verifi["verifDanemark"] = true
			verifi["verifSuède"] = false
			verifi["verifNorvège"] = false
        }
	
    case "verifSuède":
        if verifi["verifSuède"] == true {
            verifi["verifSuède"] = false
        } else {
            verifi["verifSuède"] = true
			verifi["verifDanemark"] = false
			verifi["verifNorvège"] = false
        }

	case "verifNorvège":
        if verifi["verifNorvège"] == true {
            verifi["verifNorvège"] = false
        } else {
            verifi["verifNorvège"] = true
			verifi["verifDanemark"] = false
			verifi["verifSuède"] = false
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
			verifi["verifApoge"] = false
			verifi["verifExpansion"] = false
			verifi["verifDébut"] = false
        }
	

    case "verifApoge":
        if verifi["verifApoge"] == true {
            verifi["verifApoge"] = false
        } else {
            verifi["verifApoge"] = true
			verifi["verifFin"] = false
			verifi["verifExpansion"] = false
			verifi["verifDébut"] = false
        }
	

    case "verifExpansion":
        if verifi["verifExpansion"] == true {
            verifi["verifExpansion"] = false
        } else {
            verifi["verifExpansion"] = true
			verifi["verifFin"] = false
			verifi["verifApoge"] = false
			verifi["verifDébut"] = false
        }

	case "verifDébut":
        if verifi["verifDébut"] == true {
           verifi["verifDébut"] = false
        } else {
            verifi["verifDébut"] = true
			verifi["verifFin"] = false
			verifi["verifApoge"] = false
			verifi["verifExpansion"] = false
        }
    }
}