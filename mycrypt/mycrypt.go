package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

func Krypter(melding []rune, alphabet []rune, chiffer ...int) []rune {
    c := 1 // default value
    if len(chiffer) > 0 {
        c = chiffer[0]
    }

    kryptertMelding := make([]rune, len(melding))
    for i := 0; i < len(melding); i++ {
        indeks := sokIAlfabetet(melding[i], alphabet)
        if indeks+c >= len(alphabet) {
            kryptertMelding[i] = alphabet[indeks+c-len(alphabet)]
        } else {
            kryptertMelding[i] = alphabet[indeks+c]
        }
    }
    return kryptertMelding
}

func sokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
			break
		}
	}
	return -1
}
