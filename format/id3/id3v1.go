package id3

import (
	"github.com/wader/fq/format"
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/interp"
	"github.com/wader/fq/pkg/scalar"
)

// TODO: comment 28 long, zero byte, track number

func init() {
	interp.RegisterFormat(decode.Format{
		Name:        format.ID3V1,
		Description: "ID3v1 metadata",
		DecodeFn:    id3v1Decode,
	})
}

// Decode ID3v1 tag
func id3v1Decode(d *decode.D, in any) any {
	d.AssertAtLeastBitsLeft(128 * 8)
	d.FieldUTF8("magic", 3, d.AssertStr("TAG"))
	if d.PeekBits(8) == uint64('+') {
		d.Errorf("looks like id3v11")
	}
	d.FieldUTF8NullFixedLen("song_name", 30)
	d.FieldUTF8NullFixedLen("artist", 30)
	d.FieldUTF8NullFixedLen("album_name", 30)
	d.FieldUTF8NullFixedLen("year", 4)
	d.FieldUTF8NullFixedLen("comment", 30)
	// from https://en.wikipedia.org/wiki/List_of_ID3v1_Genres
	d.FieldU8("genre", scalar.UToScalar{
		0:   {Sym: "blues", Description: "Blues"},
		1:   {Sym: "classic_rock", Description: "Classic Rock"},
		2:   {Sym: "country", Description: "Country"},
		3:   {Sym: "dance", Description: "Dance"},
		4:   {Sym: "disco", Description: "Disco"},
		5:   {Sym: "funk", Description: "Funk"},
		6:   {Sym: "grunge", Description: "Grunge"},
		7:   {Sym: "hiphop", Description: "Hip-Hop"},
		8:   {Sym: "jazz", Description: "Jazz"},
		9:   {Sym: "metal", Description: "Metal"},
		10:  {Sym: "new_age", Description: "New Age"},
		11:  {Sym: "oldies", Description: "Oldies"},
		12:  {Sym: "other", Description: "Other"},
		13:  {Sym: "pop", Description: "Pop"},
		14:  {Sym: "rhythm_and_blues", Description: "Rhythm and Blues"},
		15:  {Sym: "rap", Description: "Rap"},
		16:  {Sym: "reggae", Description: "Reggae"},
		17:  {Sym: "rock", Description: "Rock"},
		18:  {Sym: "techno", Description: "Techno"},
		19:  {Sym: "industrial", Description: "Industrial"},
		20:  {Sym: "alternative", Description: "Alternative"},
		21:  {Sym: "ska", Description: "Ska"},
		22:  {Sym: "death_metal", Description: "Death Metal"},
		23:  {Sym: "pranks", Description: "Pranks"},
		24:  {Sym: "soundtrack", Description: "Soundtrack"},
		25:  {Sym: "euro_techno", Description: "Euro-Techno"},
		26:  {Sym: "ambient", Description: "Ambient"},
		27:  {Sym: "triphop", Description: "Trip-Hop"},
		28:  {Sym: "vocal", Description: "Vocal"},
		29:  {Sym: "jazz_funk", Description: "Jazz & Funk"},
		30:  {Sym: "fusion", Description: "Fusion"},
		31:  {Sym: "trance", Description: "Trance"},
		32:  {Sym: "classical", Description: "Classical"},
		33:  {Sym: "instrumental", Description: "Instrumental"},
		34:  {Sym: "acid", Description: "Acid"},
		35:  {Sym: "house", Description: "House"},
		36:  {Sym: "game", Description: "Game"},
		37:  {Sym: "sound_clip", Description: "Sound clip"},
		38:  {Sym: "gospel", Description: "Gospel"},
		39:  {Sym: "noise", Description: "Noise"},
		40:  {Sym: "alternative_rock", Description: "Alternative Rock"},
		41:  {Sym: "bass", Description: "Bass"},
		42:  {Sym: "soul", Description: "Soul"},
		43:  {Sym: "punk", Description: "Punk"},
		44:  {Sym: "space", Description: "Space"},
		45:  {Sym: "meditative", Description: "Meditative"},
		46:  {Sym: "instrumental_pop", Description: "Instrumental Pop"},
		47:  {Sym: "instrumental_rock", Description: "Instrumental Rock"},
		48:  {Sym: "ethnic", Description: "Ethnic"},
		49:  {Sym: "gothic", Description: "Gothic"},
		50:  {Sym: "darkwave", Description: "Darkwave"},
		51:  {Sym: "techno-industrial", Description: "Techno-Industrial"},
		52:  {Sym: "electronic", Description: "Electronic"},
		53:  {Sym: "pop-folk", Description: "Pop-Folk"},
		54:  {Sym: "eurodance", Description: "Eurodance"},
		55:  {Sym: "dream", Description: "Dream"},
		56:  {Sym: "southern_rock", Description: "Southern Rock"},
		57:  {Sym: "comedy", Description: "Comedy"},
		58:  {Sym: "cult", Description: "Cult"},
		59:  {Sym: "gangsta", Description: "Gangsta"},
		60:  {Sym: "top_40", Description: "Top 40"},
		61:  {Sym: "christian_rap", Description: "Christian Rap"},
		62:  {Sym: "pop_funk", Description: "Pop/Funk"},
		63:  {Sym: "jungle_music", Description: "Jungle music"},
		64:  {Sym: "native_us", Description: "Native US"},
		65:  {Sym: "cabaret", Description: "Cabaret"},
		66:  {Sym: "new_wave", Description: "New Wave"},
		67:  {Sym: "psychedelic", Description: "Psychedelic"},
		68:  {Sym: "rave", Description: "Rave"},
		69:  {Sym: "showtunes", Description: "Showtunes"},
		70:  {Sym: "trailer", Description: "Trailer"},
		71:  {Sym: "lo-fi", Description: "Lo-Fi"},
		72:  {Sym: "tribal", Description: "Tribal"},
		73:  {Sym: "acid_punk", Description: "Acid Punk"},
		74:  {Sym: "acid_jazz", Description: "Acid Jazz"},
		75:  {Sym: "polka", Description: "Polka"},
		76:  {Sym: "retro", Description: "Retro"},
		77:  {Sym: "musical", Description: "Musical"},
		78:  {Sym: "rock_n_roll", Description: "Rock ’n’ Roll"},
		79:  {Sym: "hard_rock", Description: "Hard Rock"},
		80:  {Sym: "folk", Description: "Folk"},
		81:  {Sym: "folk-rock", Description: "Folk-Rock"},
		82:  {Sym: "national_folk", Description: "National Folk"},
		83:  {Sym: "swing", Description: "Swing"},
		84:  {Sym: "fast_fusion", Description: "Fast Fusion"},
		85:  {Sym: "bebop", Description: "Bebop"},
		86:  {Sym: "latin", Description: "Latin"},
		87:  {Sym: "revival", Description: "Revival"},
		88:  {Sym: "celtic", Description: "Celtic"},
		89:  {Sym: "bluegrass", Description: "Bluegrass"},
		90:  {Sym: "avantgarde", Description: "Avantgarde"},
		91:  {Sym: "gothic_rock", Description: "Gothic Rock"},
		92:  {Sym: "progressive_rock", Description: "Progressive Rock"},
		93:  {Sym: "psychedelic_rock", Description: "Psychedelic Rock"},
		94:  {Sym: "symphonic_rock", Description: "Symphonic Rock"},
		95:  {Sym: "slow_rock", Description: "Slow Rock"},
		96:  {Sym: "big_band", Description: "Big Band"},
		97:  {Sym: "chorus", Description: "Chorus"},
		98:  {Sym: "easy_listening", Description: "Easy Listening"},
		99:  {Sym: "acoustic", Description: "Acoustic"},
		100: {Sym: "humour", Description: "Humour"},
		101: {Sym: "speech", Description: "Speech"},
		102: {Sym: "chanson", Description: "Chanson"},
		103: {Sym: "opera", Description: "Opera"},
		104: {Sym: "chamber_music", Description: "Chamber Music"},
		105: {Sym: "sonata", Description: "Sonata"},
		106: {Sym: "symphony", Description: "Symphony"},
		107: {Sym: "booty_bass", Description: "Booty Bass"},
		108: {Sym: "primus", Description: "Primus"},
		109: {Sym: "porn_groove", Description: "Porn Groove"},
		110: {Sym: "satire", Description: "Satire"},
		111: {Sym: "slow_jam", Description: "Slow Jam"},
		112: {Sym: "club", Description: "Club"},
		113: {Sym: "tango", Description: "Tango"},
		114: {Sym: "samba", Description: "Samba"},
		115: {Sym: "folklore", Description: "Folklore"},
		116: {Sym: "ballad", Description: "Ballad"},
		117: {Sym: "power_ballad", Description: "Power Ballad"},
		118: {Sym: "rhythmic_soul", Description: "Rhythmic Soul"},
		119: {Sym: "freestyle", Description: "Freestyle"},
		120: {Sym: "duet", Description: "Duet"},
		121: {Sym: "punk_rock", Description: "Punk Rock"},
		122: {Sym: "drum_solo", Description: "Drum Solo"},
		123: {Sym: "a_cappella", Description: "A cappella"},
		124: {Sym: "euro_house", Description: "Euro-House"},
		125: {Sym: "dance_hall", Description: "Dance Hall"},
		126: {Sym: "goa_music", Description: "Goa music"},
		127: {Sym: "drum_n_bass", Description: "Drum & Bass"},
		128: {Sym: "club_house", Description: "Club-House"},
		129: {Sym: "hardcore_techno", Description: "Hardcore Techno"},
		130: {Sym: "terror", Description: "Terror"},
		131: {Sym: "indie", Description: "Indie"},
		132: {Sym: "britpop", Description: "BritPop"},
		133: {Sym: "negerpunk", Description: "Negerpunk"},
		134: {Sym: "polsk_punk", Description: "Polsk Punk"},
		135: {Sym: "beat", Description: "Beat"},
		136: {Sym: "christian_gangsta_rap", Description: "Christian Gangsta Rap"},
		137: {Sym: "heavy_metal", Description: "Heavy Metal"},
		138: {Sym: "black_metal", Description: "Black Metal"},
		139: {Sym: "crossover", Description: "Crossover"},
		140: {Sym: "contemporary_christian", Description: "Contemporary Christian"},
		141: {Sym: "christian_rock", Description: "Christian Rock"},
		142: {Sym: "merengue", Description: "Merengue"},
		143: {Sym: "salsa", Description: "Salsa"},
		144: {Sym: "thrash_metal", Description: "Thrash Metal"},
		145: {Sym: "anime", Description: "Anime"},
		146: {Sym: "jpop", Description: "Jpop"},
		147: {Sym: "synthpop", Description: "Synthpop"},
		148: {Sym: "abstract", Description: "Abstract"},
		149: {Sym: "art_rock", Description: "Art Rock"},
		150: {Sym: "baroque", Description: "Baroque"},
		151: {Sym: "bhangra", Description: "Bhangra"},
		152: {Sym: "big_beat", Description: "Big beat"},
		153: {Sym: "breakbeat", Description: "Breakbeat"},
		154: {Sym: "chillout", Description: "Chillout"},
		155: {Sym: "downtempo", Description: "Downtempo"},
		156: {Sym: "dub", Description: "Dub"},
		157: {Sym: "ebm", Description: "EBM"},
		158: {Sym: "eclectic", Description: "Eclectic"},
		159: {Sym: "electro", Description: "Electro"},
		160: {Sym: "electroclash", Description: "Electroclash"},
		161: {Sym: "emo", Description: "Emo"},
		162: {Sym: "experimental", Description: "Experimental"},
		163: {Sym: "garage", Description: "Garage"},
		164: {Sym: "global", Description: "Global"},
		165: {Sym: "idm", Description: "IDM"},
		166: {Sym: "illbient", Description: "Illbient"},
		167: {Sym: "industro_goth", Description: "Industro-Goth"},
		168: {Sym: "jam_band", Description: "Jam Band"},
		169: {Sym: "krautrock", Description: "Krautrock"},
		170: {Sym: "leftfield", Description: "Leftfield"},
		171: {Sym: "lounge", Description: "Lounge"},
		172: {Sym: "math_rock", Description: "Math Rock"},
		173: {Sym: "new_romantic", Description: "New Romantic"},
		174: {Sym: "nu_breakz", Description: "Nu-Breakz"},
		175: {Sym: "post_punk", Description: "Post-Punk"},
		176: {Sym: "post_rock", Description: "Post-Rock"},
		177: {Sym: "psytrance", Description: "Psytrance"},
		178: {Sym: "shoegaze", Description: "Shoegaze"},
		179: {Sym: "space_rock", Description: "Space Rock"},
		180: {Sym: "trop_rock", Description: "Trop Rock"},
		181: {Sym: "world_music", Description: "World Music"},
		182: {Sym: "neoclassical", Description: "Neoclassical"},
		183: {Sym: "audiobook", Description: "Audiobook"},
		184: {Sym: "audio_theatre", Description: "Audio Theatre"},
		185: {Sym: "neue_deutsche_welle", Description: "Neue Deutsche Welle"},
		186: {Sym: "podcast", Description: "Podcast"},
		187: {Sym: "indie_rock", Description: "Indie-Rock"},
		188: {Sym: "g_funk", Description: "G-Funk"},
		189: {Sym: "dubstep", Description: "Dubstep"},
		190: {Sym: "garage_rock", Description: "Garage Rock"},
		191: {Sym: "psybient", Description: "Psybient"},
	})

	return nil
}
