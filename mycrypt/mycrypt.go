package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; ")

// Krypter funksjon
func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
    kryptertMelding := make([]rune, len(melding))
    for i := 0; i < len(melding); i++ {
        indeks := sokIAlfabetet(melding[i], alphabet)
        kryptertMelding[i] = alphabet[(indeks+chiffer)%len(alphabet)]
    }
    return kryptertMelding
}

// DeKrypter funksjon
func DeKrypter(kryptertMelding []rune, alphabet []rune, chiffer int) []rune {
    melding := make([]rune, len(kryptertMelding))
    for i := 0; i < len(kryptertMelding); i++ {
        indeks := sokIAlfabetet(kryptertMelding[i], alphabet)
        melding[i] = alphabet[(indeks+(len(alphabet)-chiffer))%len(alphabet)]
    }
    return melding
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
