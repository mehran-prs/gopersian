package gopersian

// Letter holds the Persian/Arabic character with its
// different representation forms (glyphs).
type Letter struct {
	Unicode, Isolated, Initial, Medial, Final rune
}

// Persian/arabic alphabet.
var (
	AlefHamzaAbove = Letter{ // أ
		Unicode:  '\u0623',
		Isolated: '\ufe83',
		Initial:  '\u0623',
		Medial:   '\ufe84',
		Final:    '\ufe84'}

	Alef = Letter{ // ا
		Unicode:  '\u0627',
		Isolated: '\ufe8d',
		Initial:  '\u0627',
		Medial:   '\ufe8e',
		Final:    '\ufe8e'}

	AlefMaddaAbove = Letter{ // آ
		Unicode:  '\u0622',
		Isolated: '\ufe81',
		Initial:  '\u0622',
		Medial:   '\ufe82',
		Final:    '\ufe82'}

	Hamza = Letter{ // ء
		Unicode:  '\u0621',
		Isolated: '\ufe80',
		Initial:  '\u0621',
		Medial:   '\u0621',
		Final:    '\u0621'}

	WasHamzaAbove = Letter{ // ؤ
		Unicode:  '\u0624',
		Isolated: '\ufe85',
		Initial:  '\u0624',
		Medial:   '\ufe86',
		Final:    '\ufe86'}

	AlefHamzaBelow = Letter{ // أ
		Unicode:  '\u0625',
		Isolated: '\ufe87',
		Initial:  '\u0625',
		Medial:   '\ufe88',
		Final:    '\ufe88'}

	YehHamzaAbove = Letter{ // ئ
		Unicode:  '\u0626',
		Isolated: '\ufe89',
		Initial:  '\ufe8b',
		Medial:   '\ufe8c',
		Final:    '\ufe8a'}

	Beh = Letter{ // ب
		Unicode:  '\u0628',
		Isolated: '\ufe8f',
		Initial:  '\ufe91',
		Medial:   '\ufe92',
		Final:    '\ufe90'}

	Peh = Letter{ // پ
		Unicode:  '\u067e',
		Isolated: '\ufb56',
		Initial:  '\ufb58',
		Medial:   '\ufb59',
		Final:    '\ufb57'}

	// nolint:misspell
	Teh = Letter{ // ت
		Unicode:  '\u062A',
		Isolated: '\ufe95',
		Initial:  '\ufe97',
		Medial:   '\ufe98',
		Final:    '\ufe96'}

	TehMarbute = Letter{ // ة
		Unicode:  '\u0629',
		Isolated: '\ufe93',
		Initial:  '\u0629',
		Medial:   '\u0629',
		Final:    '\ufe94'}

	Theh = Letter{ // ث
		Unicode:  '\u062b',
		Isolated: '\ufe99',
		Initial:  '\ufe9b',
		Medial:   '\ufe9c',
		Final:    '\ufe9a'}

	Jeem = Letter{ // ج
		Unicode:  '\u062c',
		Isolated: '\ufe9d',
		Initial:  '\ufe9f',
		Medial:   '\ufea0',
		Final:    '\ufe9e'}

	Tcheh = Letter{ // چ
		Unicode:  '\u0686',
		Isolated: '\ufb7a',
		Initial:  '\ufb7c',
		Medial:   '\ufb7d',
		Final:    '\ufb7b'}

	Hah = Letter{ // ح
		Unicode:  '\u062d',
		Isolated: '\ufea1',
		Initial:  '\ufea3',
		Medial:   '\ufea4',
		Final:    '\ufea2'}

	Khah = Letter{ // خ
		Unicode:  '\u062e',
		Isolated: '\ufea5',
		Initial:  '\ufea7',
		Medial:   '\ufea8',
		Final:    '\ufea6'}

	Dal = Letter{ // د
		Unicode:  '\u062f',
		Isolated: '\ufea9',
		Initial:  '\u062f',
		Medial:   '\ufeaa',
		Final:    '\ufeaa'}

	Thal = Letter{ // ذ
		Unicode:  '\u0630',
		Isolated: '\ufeab',
		Initial:  '\u0630',
		Medial:   '\ufeac',
		Final:    '\ufeac'}

	Reh = Letter{ // ر
		Unicode:  '\u0631',
		Isolated: '\ufead',
		Initial:  '\u0631',
		Medial:   '\ufeae',
		Final:    '\ufeae'}

	Jeh = Letter{ // ژ
		Unicode:  '\u0698',
		Isolated: '\ufb8a',
		Initial:  '\u0698',
		Medial:   '\ufb8b',
		Final:    '\ufb8b',
	}

	Zain = Letter{ // ز
		Unicode:  '\u0632',
		Isolated: '\ufeaf',
		Initial:  '\u0632',
		Medial:   '\ufeb0',
		Final:    '\ufeb0'}

	Seen = Letter{ // س
		Unicode:  '\u0633',
		Isolated: '\ufeb1',
		Initial:  '\ufeb3',
		Medial:   '\ufeb4',
		Final:    '\ufeb2'}

	Sheen = Letter{ // ش
		Unicode:  '\u0634',
		Isolated: '\ufeb5',
		Initial:  '\ufeb7',
		Medial:   '\ufeb8',
		Final:    '\ufeb6'}

	Sad = Letter{ // ص
		Unicode:  '\u0635',
		Isolated: '\ufeb9',
		Initial:  '\ufebb',
		Medial:   '\ufebc',
		Final:    '\ufeba'}

	Dad = Letter{ // ض
		Unicode:  '\u0636',
		Isolated: '\ufebd',
		Initial:  '\ufebf',
		Medial:   '\ufec0',
		Final:    '\ufebe'}

	Tah = Letter{ // ط
		Unicode:  '\u0637',
		Isolated: '\ufec1',
		Initial:  '\ufec3',
		Medial:   '\ufec4',
		Final:    '\ufec2'}

	Zah = Letter{ // ظ
		Unicode:  '\u0638',
		Isolated: '\ufec5',
		Initial:  '\ufec7',
		Medial:   '\ufec8',
		Final:    '\ufec6'}

	Ain = Letter{ // ع
		Unicode:  '\u0639',
		Isolated: '\ufec9',
		Initial:  '\ufecb',
		Medial:   '\ufecc',
		Final:    '\ufeca'}

	Ghain = Letter{ // غ
		Unicode:  '\u063a',
		Isolated: '\ufecd',
		Initial:  '\ufecf',
		Medial:   '\ufed0',
		Final:    '\ufece'}

	Feh = Letter{ // ف
		Unicode:  '\u0641',
		Isolated: '\ufed1',
		Initial:  '\ufed3',
		Medial:   '\ufed4',
		Final:    '\ufed2'}

	Qaf = Letter{ // ق
		Unicode:  '\u0642',
		Isolated: '\ufed5',
		Initial:  '\ufed7',
		Medial:   '\ufed8',
		Final:    '\ufed6'}

	Kaf = Letter{ // ك
		Unicode:  '\u0643',
		Isolated: '\ufed9',
		Initial:  '\ufedb',
		Medial:   '\ufedc',
		Final:    '\ufeda'}

	Keheh = Letter{ // ک
		Unicode:  '\u06a9',
		Isolated: '\ufb8e',
		Initial:  '\ufb90',
		Medial:   '\ufb91',
		Final:    '\ufb8f',
	}

	Gaf = Letter{ // گ
		Unicode:  '\u06af',
		Isolated: '\ufb92',
		Initial:  '\ufb94',
		Medial:   '\ufb95',
		Final:    '\ufb93'}

	Lam = Letter{ // ل
		Unicode:  '\u0644',
		Isolated: '\ufedd',
		Initial:  '\ufedf',
		Medial:   '\ufee0',
		Final:    '\ufede'}

	Meem = Letter{ // م
		Unicode:  '\u0645',
		Isolated: '\ufee1',
		Initial:  '\ufee3',
		Medial:   '\ufee4',
		Final:    '\ufee2'}

	Noon = Letter{ // ن
		Unicode:  '\u0646',
		Isolated: '\ufee5',
		Initial:  '\ufee7',
		Medial:   '\ufee8',
		Final:    '\ufee6'}

	Heh = Letter{ // ه
		Unicode:  '\u0647',
		Isolated: '\ufee9',
		Initial:  '\ufeeb',
		Medial:   '\ufeec',
		Final:    '\ufeea'}

	Waw = Letter{ // و
		Unicode:  '\u0648',
		Isolated: '\ufeed',
		Initial:  '\u0648',
		Medial:   '\ufeee',
		Final:    '\ufeee'}

	Yeh = Letter{ // ی
		Unicode:  '\u06cc',
		Isolated: '\ufbfc',
		Initial:  '\ufbfe',
		Medial:   '\ufbff',
		Final:    '\ufbfd'}

	Arabicyeh = Letter{ // ي
		Unicode:  '\u064a',
		Isolated: '\ufef1',
		Initial:  '\ufef3',
		Medial:   '\ufef4',
		Final:    '\ufef2'}

	AlefMaksura = Letter{ // ى
		Unicode:  '\u0649',
		Isolated: '\ufeef',
		Initial:  '\u0649',
		Medial:   '\ufef0',
		Final:    '\ufef0'}

	Tatweel = Letter{ // ـ
		Unicode:  '\u0640',
		Isolated: '\u0640',
		Initial:  '\u0640',
		Medial:   '\u0640',
		Final:    '\u0640'}

	LamAlef = Letter{ // لا
		Unicode:  '\ufefb',
		Isolated: '\ufefb',
		Initial:  '\ufefb',
		Medial:   '\ufefc',
		Final:    '\ufefc'}

	LamAlefHamzaAbove = Letter{ // ﻷ
		Unicode:  '\ufef7',
		Isolated: '\ufef7',
		Initial:  '\ufef7',
		Medial:   '\ufef8',
		Final:    '\ufef8'}
)

var alphabet = []Letter{
	AlefHamzaAbove,
	Alef,
	AlefMaddaAbove,
	Hamza,
	WasHamzaAbove,
	AlefHamzaBelow,
	YehHamzaAbove,
	Beh,
	Peh,
	Teh, //nolint:misspell
	TehMarbute,
	Theh,
	Jeem,
	Tcheh,
	Hah,
	Khah,
	Dal,
	Thal,
	Reh,
	Jeh,
	Zain,
	Seen,
	Sheen,
	Sad,
	Dad,
	Tah,
	Zah,
	Ain,
	Ghain,
	Feh,
	Qaf,
	Kaf,
	Keheh,
	Gaf,
	Lam,
	Meem,
	Noon,
	Heh,
	Waw,
	Yeh,
	Arabicyeh,
	AlefMaksura,
	Tatweel,
	LamAlef,
	LamAlefHamzaAbove,
}

// use map for faster lookups.
var isolatedAfter = map[Letter]bool{
	AlefHamzaAbove: true,
	AlefMaddaAbove: true,
	Alef:           true,
	Hamza:          true,
	WasHamzaAbove:  true,
	AlefHamzaBelow: true,
	TehMarbute:     true,
	Dal:            true,
	Thal:           true,
	Reh:            true,
	Jeh:            true,
	Zain:           true,
	Waw:            true,
	AlefMaksura:    true,
}

func (c *Letter) equals(char rune) bool {
	switch char {
	case c.Unicode, c.Initial, c.Isolated, c.Medial, c.Final:
		return true
	}

	return false
}

// letterFromRune returns corresponding Letter for the given rune.
func letterFromRune(r rune) Letter {
	for _, s := range alphabet {
		if s.equals(r) {
			return s
		}
	}

	return Letter{Unicode: r, Initial: r, Isolated: r, Medial: r, Final: r}
}
