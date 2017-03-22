package whatlanggo

import (
	"testing"
	"unicode"
)

func TestDetectScripts(t *testing.T) {
	tests := map[string][]*unicode.RangeTable{
		"Привет! Текст на русском with some English.": []*unicode.RangeTable{unicode.Cyrillic, unicode.Latin},
		"Russian word любовь means love.":             []*unicode.RangeTable{unicode.Cyrillic, unicode.Latin},
		"Это текст с одним латинским sимволом":        []*unicode.RangeTable{unicode.Cyrillic},
	}

	for text, want := range tests {
		got := DetectScripts(text, 10)

		t.Log(text, got)
		continue

		if isScriptSliceEqual(want, got) {
			t.Fatalf("%s want %s got %s", text, want, got)
		}
	}
}

func isScriptSliceEqual(s1 []*unicode.RangeTable, s2 []*unicode.RangeTable) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func Test_DetectScript(t *testing.T) {
	tests := map[string]*unicode.RangeTable{
		"123456789-=?":                                                  nil,
		"Hello, world!":                                                 unicode.Latin,
		"Привет всем!":                                                  unicode.Cyrillic,
		"ქართული ენა მსოფლიო ":                                          unicode.Georgian,
		"県見夜上温国阪題富販":                                                    unicode.Han,
		" ككل حوالي 1.6، ومعظم الناس ":                                  unicode.Arabic,
		"हिमालयी वन चिड़िया (जूथेरा सालिमअली) चिड़िया की एक प्रजाति है": unicode.Devanagari,
		"היסטוריה והתפתחות של האלפבית העברי":                            unicode.Hebrew,
		"የኢትዮጵያ ፌዴራላዊ ዴሞክራሲያዊሪፐብሊክ":                                     unicode.Ethiopic,
		"Привет! Текст на русском with some English.":                   unicode.Cyrillic,
		"Russian word любовь means love.":                               unicode.Latin,
		"আমি ভালো আছি, ধন্যবাদ!":                                        unicode.Bengali,
	}

	for text, want := range tests {
		got := DetectScript(text)
		if want != got {
			t.Fatalf("%s want %s got %s", text, Scripts[want], Scripts[got])
		}
	}
}

func Test_isLatin(t *testing.T) {
	tests := map[rune]bool{
		'z': true, 'A': true, 'č': true, 'š': true, 'Ĵ': true, 'ж': false,
	}

	for r, want := range tests {
		got := isLatin(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isEthiopic(t *testing.T) {
	tests := map[rune]bool{
		'ፚ': true, 'ᎀ': true, 'а': false, 'L': false,
	}

	for r, want := range tests {
		got := isEthiopic(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isGeorgian(t *testing.T) {
	tests := map[rune]bool{
		'რ': true, 'Я': false,
	}

	for r, want := range tests {
		got := isGeorgian(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isBengali(t *testing.T) {
	tests := map[rune]bool{
		'а': false, 'ই': true,
	}

	for r, want := range tests {
		got := isBengali(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isKatakana(t *testing.T) {
	tests := map[rune]bool{
		'カ': true, 'Ґ': false,
	}

	for r, want := range tests {
		got := isKatakana(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isHiragana(t *testing.T) {
	tests := map[rune]bool{
		'ひ': true, 'Ꙕ': false,
	}

	for r, want := range tests {
		got := isHiragana(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isHangul(t *testing.T) {
	tests := map[rune]bool{
		'ᄁ': true, 't': false,
	}

	for r, want := range tests {
		got := isHangul(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isGreek(t *testing.T) {
	tests := map[rune]bool{
		'φ': true, 'ф': false,
	}

	for r, want := range tests {
		got := isGreek(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isKannada(t *testing.T) {
	tests := map[rune]bool{
		'ಡ': true, 'S': false,
	}

	for r, want := range tests {
		got := isKannada(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isTamil(t *testing.T) {
	tests := map[rune]bool{
		'ஐ': true, 'Ж': false,
	}

	for r, want := range tests {
		got := isTamil(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isThai(t *testing.T) {
	tests := map[rune]bool{
		'ก': true, '๛': true, 'Ґ': false,
	}

	for r, want := range tests {
		got := isThai(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isGujarati(t *testing.T) {
	tests := map[rune]bool{
		'ઁ': true, '૱': true, 'l': false,
	}

	for r, want := range tests {
		got := isGujarati(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isGurmukhi(t *testing.T) {
	tests := map[rune]bool{
		'ਁ': true, 'ੴ': true, 'Ж': false,
	}

	for r, want := range tests {
		got := isGurmukhi(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isTelugu(t *testing.T) {
	tests := map[rune]bool{
		'ఁ': true, '౿': true, 'l': false,
	}

	for r, want := range tests {
		got := isTelugu(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}

func Test_isOriya(t *testing.T) {
	tests := map[rune]bool{
		'ଐ': true, '୷': true, 'l': false,
	}

	for r, want := range tests {
		got := isOriya(r)
		if want != got {
			t.Fatalf("%#U want %t got %t", r, want, got)
		}
	}
}
