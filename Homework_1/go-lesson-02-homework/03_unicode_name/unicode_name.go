package unicode_name

import "unicode"

// AnalyzeName считает байты, Unicode-кодовые точки, латинские и кириллические
// буквы и возвращает байтовые позиции начала каждой Unicode-кодовой точки.
// Предполагается, что входная строка содержит корректный UTF-8.
func AnalyzeName(name string) (byteCount, runeCount, latinCount, cyrillicCount int, positions []int) {
	byteCount = len(name)
	for i, v := range name {
		runeCount++
		switch {
		case unicode.Is(unicode.Cyrillic, v) && unicode.IsLetter(v):
			cyrillicCount++
		case unicode.Is(unicode.Latin, v) && unicode.IsLetter(v):
			latinCount++
		}
		positions = append(positions, i)
	}

	return 
}
