package api

//check and change the thuth table

func Verifswitch(verifi map[string]bool, motsVerif string) {

    switch motsVerif {
    case "verifAlbumDate":
        if verifi["verifAlbumDate"] == true {
            verifi["verifAlbumDate"] = false
        } else {
            verifi["verifAlbumDate"] = true
        }
	

    case "verifDate":
        if verifi["verifDate"] == true {
            verifi["verifDate"] = false
        } else {
            verifi["verifDate"] = true
        }
	

    case "verifMemberNumber":
        if verifi["verifMemberNumber"] == true {
            verifi["verifMemberNumber"] = false
        } else {
            verifi["verifMemberNumber"] = true
        }

	case "verifLocation":
        if verifi["verifLocation"] == true {
            verifi["verifLocation"] = false
        } else {
            verifi["verifLocation"] = true
        }



	case "réinitialiser":
        if verifi["réinitialiser"] == true {
            verifi["réinitialiser"] = false
        } else {
            verifi["réinitialiser"] = true
        }
	



    case "verifNonAlbumDate":
        if verifi["verifNonAlbumDate"] == true {
            verifi["verifNonAlbumDate"] = false
        } else {
            verifi["verifNonAlbumDate"] = true
        }
	

    case "verifNonDate":
        if verifi["verifNonDate"] == true {
            verifi["verifNonDate"] = false
        } else {
            verifi["verifNonDate"] = true
        }
	

    case "verifNonMemberNumber":
        if verifi["verifNonMemberNumber"] == true {
            verifi["verifNonMemberNumber"] = false
        } else {
            verifi["verifNonMemberNumber"] = true
        }

	case "verifNonLocation":
        if verifi["verifNonLocation"] == true {
            verifi["verifNonLocation"] = false
        } else {
            verifi["verifNonLocation"] = true
        }
    }
}
