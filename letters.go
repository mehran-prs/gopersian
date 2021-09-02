package gopersian

// Letter holds the Persian/Arabic character with its
// different representation forms (glyphs).
type Letter struct {
	Unicode, Isolated, Initial, Medial, Final rune
}

// Persian/arabic alphabet.
var (
	ALEF_HAMZA_ABOVE = Letter{ // أ
		Unicode:  '\u0623',
		Isolated: '\ufe83',
		Initial:  '\u0623',
		Medial:   '\ufe84',
		Final:    '\ufe84'}

	ALEF = Letter{ // ا
		Unicode:  '\u0627',
		Isolated: '\ufe8d',
		Initial:  '\u0627',
		Medial:   '\ufe8e',
		Final:    '\ufe8e'}

	ALEF_MADDA_ABOVE = Letter{ // آ
		Unicode:  '\u0622',
		Isolated: '\ufe81',
		Initial:  '\u0622',
		Medial:   '\ufe82',
		Final:    '\ufe82'}

	HAMZA = Letter{ // ء
		Unicode:  '\u0621',
		Isolated: '\ufe80',
		Initial:  '\u0621',
		Medial:   '\u0621',
		Final:    '\u0621'}

	WAW_HAMZA_ABOVE = Letter{ // ؤ
		Unicode:  '\u0624',
		Isolated: '\ufe85',
		Initial:  '\u0624',
		Medial:   '\ufe86',
		Final:    '\ufe86'}

	ALEF_HAMZA_BELOW = Letter{ // أ
		Unicode:  '\u0625',
		Isolated: '\ufe87',
		Initial:  '\u0625',
		Medial:   '\ufe88',
		Final:    '\ufe88'}

	YEH_HAMZA_ABOVE = Letter{ // ئ
		Unicode:  '\u0626',
		Isolated: '\ufe89',
		Initial:  '\ufe8b',
		Medial:   '\ufe8c',
		Final:    '\ufe8a'}

	BEH = Letter{ // ب
		Unicode:  '\u0628',
		Isolated: '\ufe8f',
		Initial:  '\ufe91',
		Medial:   '\ufe92',
		Final:    '\ufe90'}

	PEH = Letter{ // پ
		Unicode:  '\u067e',
		Isolated: '\ufb56',
		Initial:  '\ufb58',
		Medial:   '\ufb59',
		Final:    '\ufb57'}

	TEH = Letter{ // ت
		Unicode:  '\u062A',
		Isolated: '\ufe95',
		Initial:  '\ufe97',
		Medial:   '\ufe98',
		Final:    '\ufe96'}

	TEH_MARBUTA = Letter{ // ة
		Unicode:  '\u0629',
		Isolated: '\ufe93',
		Initial:  '\u0629',
		Medial:   '\u0629',
		Final:    '\ufe94'}

	THEH = Letter{ // ث
		Unicode:  '\u062b',
		Isolated: '\ufe99',
		Initial:  '\ufe9b',
		Medial:   '\ufe9c',
		Final:    '\ufe9a'}

	JEEM = Letter{ // ج
		Unicode:  '\u062c',
		Isolated: '\ufe9d',
		Initial:  '\ufe9f',
		Medial:   '\ufea0',
		Final:    '\ufe9e'}

	TCHEH = Letter{ // چ
		Unicode:  '\u0686',
		Isolated: '\ufb7a',
		Initial:  '\ufb7c',
		Medial:   '\ufb7d',
		Final:    '\ufb7b'}

	HAH = Letter{ // ح
		Unicode:  '\u062d',
		Isolated: '\ufea1',
		Initial:  '\ufea3',
		Medial:   '\ufea4',
		Final:    '\ufea2'}

	KHAH = Letter{ // خ
		Unicode:  '\u062e',
		Isolated: '\ufea5',
		Initial:  '\ufea7',
		Medial:   '\ufea8',
		Final:    '\ufea6'}

	DAL = Letter{ // د
		Unicode:  '\u062f',
		Isolated: '\ufea9',
		Initial:  '\u062f',
		Medial:   '\ufeaa',
		Final:    '\ufeaa'}

	THAL = Letter{ // ذ
		Unicode:  '\u0630',
		Isolated: '\ufeab',
		Initial:  '\u0630',
		Medial:   '\ufeac',
		Final:    '\ufeac'}

	REH = Letter{ // ر
		Unicode:  '\u0631',
		Isolated: '\ufead',
		Initial:  '\u0631',
		Medial:   '\ufeae',
		Final:    '\ufeae'}

	JEH = Letter{ // ژ
		Unicode:  '\u0698',
		Isolated: '\ufb8a',
		Initial:  '\u0698',
		Medial:   '\ufb8b',
		Final:    '\ufb8b',
	}

	ZAIN = Letter{ // ز
		Unicode:  '\u0632',
		Isolated: '\ufeaf',
		Initial:  '\u0632',
		Medial:   '\ufeb0',
		Final:    '\ufeb0'}

	SEEN = Letter{ // س
		Unicode:  '\u0633',
		Isolated: '\ufeb1',
		Initial:  '\ufeb3',
		Medial:   '\ufeb4',
		Final:    '\ufeb2'}

	SHEEN = Letter{ // ش
		Unicode:  '\u0634',
		Isolated: '\ufeb5',
		Initial:  '\ufeb7',
		Medial:   '\ufeb8',
		Final:    '\ufeb6'}

	SAD = Letter{ // ص
		Unicode:  '\u0635',
		Isolated: '\ufeb9',
		Initial:  '\ufebb',
		Medial:   '\ufebc',
		Final:    '\ufeba'}

	DAD = Letter{ // ض
		Unicode:  '\u0636',
		Isolated: '\ufebd',
		Initial:  '\ufebf',
		Medial:   '\ufec0',
		Final:    '\ufebe'}

	TAH = Letter{ // ط
		Unicode:  '\u0637',
		Isolated: '\ufec1',
		Initial:  '\ufec3',
		Medial:   '\ufec4',
		Final:    '\ufec2'}

	ZAH = Letter{ // ظ
		Unicode:  '\u0638',
		Isolated: '\ufec5',
		Initial:  '\ufec7',
		Medial:   '\ufec8',
		Final:    '\ufec6'}

	AIN = Letter{ // ع
		Unicode:  '\u0639',
		Isolated: '\ufec9',
		Initial:  '\ufecb',
		Medial:   '\ufecc',
		Final:    '\ufeca'}

	GHAIN = Letter{ // غ
		Unicode:  '\u063a',
		Isolated: '\ufecd',
		Initial:  '\ufecf',
		Medial:   '\ufed0',
		Final:    '\ufece'}

	FEH = Letter{ // ف
		Unicode:  '\u0641',
		Isolated: '\ufed1',
		Initial:  '\ufed3',
		Medial:   '\ufed4',
		Final:    '\ufed2'}

	QAF = Letter{ // ق
		Unicode:  '\u0642',
		Isolated: '\ufed5',
		Initial:  '\ufed7',
		Medial:   '\ufed8',
		Final:    '\ufed6'}

	KAF = Letter{ // ك
		Unicode:  '\u0643',
		Isolated: '\ufed9',
		Initial:  '\ufedb',
		Medial:   '\ufedc',
		Final:    '\ufeda'}

	KEHEH = Letter{ // ک
		Unicode:  '\u06a9',
		Isolated: '\ufb8e',
		Initial:  '\ufb90',
		Medial:   '\ufb91',
		Final:    '\ufb8f',
	}

	GAF = Letter{ // گ
		Unicode:  '\u06af',
		Isolated: '\ufb92',
		Initial:  '\ufb94',
		Medial:   '\ufb95',
		Final:    '\ufb93'}

	LAM = Letter{ // ل
		Unicode:  '\u0644',
		Isolated: '\ufedd',
		Initial:  '\ufedf',
		Medial:   '\ufee0',
		Final:    '\ufede'}

	MEEM = Letter{ // م
		Unicode:  '\u0645',
		Isolated: '\ufee1',
		Initial:  '\ufee3',
		Medial:   '\ufee4',
		Final:    '\ufee2'}

	NOON = Letter{ // ن
		Unicode:  '\u0646',
		Isolated: '\ufee5',
		Initial:  '\ufee7',
		Medial:   '\ufee8',
		Final:    '\ufee6'}

	HEH = Letter{ // ه
		Unicode:  '\u0647',
		Isolated: '\ufee9',
		Initial:  '\ufeeb',
		Medial:   '\ufeec',
		Final:    '\ufeea'}

	WAW = Letter{ // و
		Unicode:  '\u0648',
		Isolated: '\ufeed',
		Initial:  '\u0648',
		Medial:   '\ufeee',
		Final:    '\ufeee'}

	YEH = Letter{ // ی
		Unicode:  '\u06cc',
		Isolated: '\ufbfc',
		Initial:  '\ufbfe',
		Medial:   '\ufbff',
		Final:    '\ufbfd'}

	ARABICYEH = Letter{ // ي
		Unicode:  '\u064a',
		Isolated: '\ufef1',
		Initial:  '\ufef3',
		Medial:   '\ufef4',
		Final:    '\ufef2'}

	ALEF_MAKSURA = Letter{ // ى
		Unicode:  '\u0649',
		Isolated: '\ufeef',
		Initial:  '\u0649',
		Medial:   '\ufef0',
		Final:    '\ufef0'}

	TATWEEL = Letter{ // ـ
		Unicode:  '\u0640',
		Isolated: '\u0640',
		Initial:  '\u0640',
		Medial:   '\u0640',
		Final:    '\u0640'}

	LAM_ALEF = Letter{ // لا
		Unicode:  '\ufefb',
		Isolated: '\ufefb',
		Initial:  '\ufefb',
		Medial:   '\ufefc',
		Final:    '\ufefc'}

	LAM_ALEF_HAMZA_ABOVE = Letter{ // ﻷ
		Unicode:  '\ufef7',
		Isolated: '\ufef7',
		Initial:  '\ufef7',
		Medial:   '\ufef8',
		Final:    '\ufef8'}
)

var alphabet = []Letter{
	ALEF_HAMZA_ABOVE,
	ALEF,
	ALEF_MADDA_ABOVE,
	HAMZA,
	WAW_HAMZA_ABOVE,
	ALEF_HAMZA_BELOW,
	YEH_HAMZA_ABOVE,
	BEH,
	PEH,
	TEH,
	TEH_MARBUTA,
	THEH,
	JEEM,
	TCHEH,
	HAH,
	KHAH,
	DAL,
	THAL,
	REH,
	JEH,
	ZAIN,
	SEEN,
	SHEEN,
	SAD,
	DAD,
	TAH,
	ZAH,
	AIN,
	GHAIN,
	FEH,
	QAF,
	KAF,
	KEHEH,
	GAF,
	LAM,
	MEEM,
	NOON,
	HEH,
	WAW,
	YEH,
	ARABICYEH,
	ALEF_MAKSURA,
	TATWEEL,
	LAM_ALEF,
	LAM_ALEF_HAMZA_ABOVE,
}

// use map for faster lookups.
var isolatedAfter = map[Letter]bool{
	ALEF_HAMZA_ABOVE: true,
	ALEF_MADDA_ABOVE: true,
	ALEF:             true,
	HAMZA:            true,
	WAW_HAMZA_ABOVE:  true,
	ALEF_HAMZA_BELOW: true,
	TEH_MARBUTA:      true,
	DAL:              true,
	THAL:             true,
	REH:              true,
	JEH:              true,
	ZAIN:             true,
	WAW:              true,
	ALEF_MAKSURA:     true,
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
