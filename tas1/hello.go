package test

func Hello(name string,  lang string) string {
	if lang == "EN"{
		return "Hello " + name
	}
	if lang == "RU"{
		return "Privet " + name
	}
	if lang == "UZB"{
		return "Salom " + name
	}
	return "Error"
}
