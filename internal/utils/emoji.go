package utils

import (
	"fmt"
	"strings"
)

type Emoji struct {
	Name    string
	Unicode string
	Alias   string
}

func EmojiKeys() []string {
	keys := []string{}

	for k := range EmojiMap {
		keys = append(keys, k)
	}
	return keys
}

func ConvertToHex(emoji string) string {
	converted := fmt.Sprintf("%+q", emoji)
	unicodeHex := strings.Split(converted, " ")[0]
	return strings.Trim(unicodeHex, `"`)
}

var EmojiMap = map[string]Emoji{
	"#\uFE0F\u20E3   :hash:": {
		Name:    "Hash Key",
		Unicode: "#\uFE0F\u20E3",
		Alias:   ":hash:",
	},
	"*\u002A\uFE0F\u20E3   :keycap_star:": {
		Name:    "Keycap: *",
		Unicode: "*\u002A\uFE0F\u20E3",
		Alias:   ":keycap_star:",
	},
	"0\uFE0F\u20E3   :zero:": {
		Name:    "Keycap 0",
		Unicode: "0\uFE0F\u20E3",
		Alias:   ":zero:",
	},
	"1\uFE0F\u20E3  :one:": {
		Name:    "Keycap 1",
		Unicode: "1\uFE0F\u20E3",
		Alias:   ":one:",
	},
	"2\uFE0F\u20E3  :two:": {
		Name:    "Keycap 2",
		Unicode: "2\uFE0F\u20E3",
		Alias:   ":two:",
	},
	"3\uFE0F\u20E3  :three:": {
		Name:    "Keycap 3",
		Unicode: "3\uFE0F\u20E3",
		Alias:   ":three:",
	},
	"4\uFE0F\u20E3  :four:": {
		Name:    "Keycap 4",
		Unicode: "4\uFE0F\u20E3",
		Alias:   ":four:",
	},
	"5\uFE0F\u20E3  :five:": {
		Name:    "Keycap 5",
		Unicode: "5\uFE0F\u20E3",
		Alias:   ":five:",
	},
	"6\uFE0F\u20E3  :six:": {
		Name:    "Keycap 6",
		Unicode: "6\uFE0F\u20E3",
		Alias:   ":six:",
	},
	"7\uFE0F\u20E3  :seven:": {
		Name:    "Keycap 7",
		Unicode: "7\uFE0F\u20E3",
		Alias:   ":seven:",
	},
	"8\uFE0F\u20E3  :eight:": {
		Name:    "Keycap 8",
		Unicode: "8\uFE0F\u20E3",
		Alias:   ":eight:",
	},
	"9\uFE0F\u20E3  :nine:": {
		Name:    "Keycap 9",
		Unicode: "9\uFE0F\u20E3",
		Alias:   ":nine:",
	},
	"\u00A9\uFE0F  :copyright:": {
		Name:    "Copyright Sign",
		Unicode: "\u00A9\uFE0F",
		Alias:   ":copyright:",
	},
	"\u00AE\uFE0F  :registered:": {
		Name:    "Registered Sign",
		Unicode: "\u00AE\uFE0F",
		Alias:   ":registered:",
	},
	"\U0001F004  :mahjong:": {
		Name:    "Mahjong Tile Red Dragon",
		Unicode: "\U0001F004",
		Alias:   ":mahjong:",
	},
	"\U0001F0CF  :black_joker:": {
		Name:    "Playing Card Black Joker",
		Unicode: "\U0001F0CF",
		Alias:   ":black_joker:",
	},
	"\U0001F170\uFE0F  :a:": {
		Name:    "Negative Squared Latin Capital Letter A",
		Unicode: "\U0001F170\uFE0F",
		Alias:   ":a:",
	},
	"\U0001F171\uFE0F  :b:": {
		Name:    "Negative Squared Latin Capital Letter B",
		Unicode: "\U0001F171\uFE0F",
		Alias:   ":b:",
	},
	"\U0001F17E\uFE0F  :o2:": {
		Name:    "Negative Squared Latin Capital Letter O",
		Unicode: "\U0001F17E\uFE0F",
		Alias:   ":o2:",
	},
	"\U0001F17F\uFE0F  :parking:": {
		Name:    "Negative Squared Latin Capital Letter P",
		Unicode: "\U0001F17F\uFE0F",
		Alias:   ":parking:",
	},
	"\U0001F18E  :ab:": {
		Name:    "Negative Squared Ab",
		Unicode: "\U0001F18E",
		Alias:   ":ab:",
	},
	"\U0001F191  :cl:": {
		Name:    "Squared Cl",
		Unicode: "\U0001F191",
		Alias:   ":cl:",
	},
	"\U0001F192  :cool:": {
		Name:    "Squared Cool",
		Unicode: "\U0001F192",
		Alias:   ":cool:",
	},
	"\U0001F193  :free:": {
		Name:    "Squared Free",
		Unicode: "\U0001F193",
		Alias:   ":free:",
	},
	"\U0001F194  :id:": {
		Name:    "Squared Id",
		Unicode: "\U0001F194",
		Alias:   ":id:",
	},
	"\U0001F195  :new:": {
		Name:    "Squared New",
		Unicode: "\U0001F195",
		Alias:   ":new:",
	},
	"\U0001F196  :ng:": {
		Name:    "Squared Ng",
		Unicode: "\U0001F196",
		Alias:   ":ng:",
	},
	"\U0001F197  :ok:": {
		Name:    "Squared Ok",
		Unicode: "\U0001F197",
		Alias:   ":ok:",
	},
	"\U0001F198  :sos:": {
		Name:    "Squared Sos",
		Unicode: "\U0001F198",
		Alias:   ":sos:",
	},
	"\U0001F199  :up:": {
		Name:    "Squared Up With Exclamation Mark",
		Unicode: "\U0001F199",
		Alias:   ":up:",
	},
	"\U0001F19A  :vs:": {
		Name:    "Squared Vs",
		Unicode: "\U0001F19A",
		Alias:   ":vs:",
	},
	"\U0001F1E6\u1F1E8 :flag-ac:": {
		Name:    "Ascension Island Flag",
		Unicode: "\U0001F1E6\u1F1E8",
		Alias:   ":flag-ac:",
	},
	"\U0001F1E6\u1F1E9 :flag-ad:": {
		Name:    "Andorra Flag",
		Unicode: "\U0001F1E6\u1F1E9",
		Alias:   ":flag-ad:",
	},
	"\U0001F1E6\u1F1EA :flag-ae:": {
		Name:    "United Arab Emirates Flag",
		Unicode: "\U0001F1E6\u1F1EA",
		Alias:   ":flag-ae:",
	},
	"\U0001F1E6\u1F1EB :flag-af:": {
		Name:    "Afghanistan Flag",
		Unicode: "\U0001F1E6\u1F1EB",
		Alias:   ":flag-af:",
	},
	"\U0001F1E6\u1F1EC :flag-ag:": {
		Name:    "Antigua & Barbuda Flag",
		Unicode: "\U0001F1E6\u1F1EC",
		Alias:   ":flag-ag:",
	},
	"\U0001F1E6\u1F1EE :flag-ai:": {
		Name:    "Anguilla Flag",
		Unicode: "\U0001F1E6\u1F1EE",
		Alias:   ":flag-ai:",
	},
	"\U0001F1E6\u1F1F1 :flag-al:": {
		Name:    "Albania Flag",
		Unicode: "\U0001F1E6\u1F1F1",
		Alias:   ":flag-al:",
	},
	"\U0001F1E6\u1F1F2 :flag-am:": {
		Name:    "Armenia Flag",
		Unicode: "\U0001F1E6\u1F1F2",
		Alias:   ":flag-am:",
	},
	"\U0001F1E6\u1F1F4 :flag-ao:": {
		Name:    "Angola Flag",
		Unicode: "\U0001F1E6\u1F1F4",
		Alias:   ":flag-ao:",
	},
	"\U0001F1E6\u1F1F6 :flag-aq:": {
		Name:    "Antarctica Flag",
		Unicode: "\U0001F1E6\u1F1F6",
		Alias:   ":flag-aq:",
	},
	"\U0001F1E6\u1F1F7 :flag-ar:": {
		Name:    "Argentina Flag",
		Unicode: "\U0001F1E6\u1F1F7",
		Alias:   ":flag-ar:",
	},
	"\U0001F1E6\u1F1F8 :flag-as:": {
		Name:    "American Samoa Flag",
		Unicode: "\U0001F1E6\u1F1F8",
		Alias:   ":flag-as:",
	},
	"\U0001F1E6\u1F1F9 :flag-at:": {
		Name:    "Austria Flag",
		Unicode: "\U0001F1E6\u1F1F9",
		Alias:   ":flag-at:",
	},
	"\U0001F1E6\u1F1FA :flag-au:": {
		Name:    "Australia Flag",
		Unicode: "\U0001F1E6\u1F1FA",
		Alias:   ":flag-au:",
	},
	"\U0001F1E6\u1F1FC :flag-aw:": {
		Name:    "Aruba Flag",
		Unicode: "\U0001F1E6\u1F1FC",
		Alias:   ":flag-aw:",
	},
	"\U0001F1E6\u1F1FD :flag-ax:": {
		Name:    "\u00c5land Islands Flag",
		Unicode: "\U0001F1E6\u1F1FD",
		Alias:   ":flag-ax:",
	},
	"\U0001F1E6\u1F1FF :flag-az:": {
		Name:    "Azerbaijan Flag",
		Unicode: "\U0001F1E6\u1F1FF",
		Alias:   ":flag-az:",
	},
	"\U0001F1E7\u1F1E6 :flag-ba:": {
		Name:    "Bosnia & Herzegovina Flag",
		Unicode: "\U0001F1E7\u1F1E6",
		Alias:   ":flag-ba:",
	},
	"\U0001F1E7\u1F1E7 :flag-bb:": {
		Name:    "Barbados Flag",
		Unicode: "\U0001F1E7\u1F1E7",
		Alias:   ":flag-bb:",
	},
	"\U0001F1E7\u1F1E9 :flag-bd:": {
		Name:    "Bangladesh Flag",
		Unicode: "\U0001F1E7\u1F1E9",
		Alias:   ":flag-bd:",
	},
	"\U0001F1E7\u1F1EA :flag-be:": {
		Name:    "Belgium Flag",
		Unicode: "\U0001F1E7\u1F1EA",
		Alias:   ":flag-be:",
	},
	"\U0001F1E7\u1F1EB :flag-bf:": {
		Name:    "Burkina Faso Flag",
		Unicode: "\U0001F1E7\u1F1EB",
		Alias:   ":flag-bf:",
	},
	"\U0001F1E7\u1F1EC :flag-bg:": {
		Name:    "Bulgaria Flag",
		Unicode: "\U0001F1E7\u1F1EC",
		Alias:   ":flag-bg:",
	},
	"\U0001F1E7\u1F1ED :flag-bh:": {
		Name:    "Bahrain Flag",
		Unicode: "\U0001F1E7\u1F1ED",
		Alias:   ":flag-bh:",
	},
	"\U0001F1E7\u1F1EE :flag-bi:": {
		Name:    "Burundi Flag",
		Unicode: "\U0001F1E7\u1F1EE",
		Alias:   ":flag-bi:",
	},
	"\U0001F1E7\u1F1EF :flag-bj:": {
		Name:    "Benin Flag",
		Unicode: "\U0001F1E7\u1F1EF",
		Alias:   ":flag-bj:",
	},
	"\U0001F1E7\u1F1F1 :flag-bl:": {
		Name:    "St. Barth\u00e9lemy Flag",
		Unicode: "\U0001F1E7\u1F1F1",
		Alias:   ":flag-bl:",
	},
	"\U0001F1E7\u1F1F2 :flag-bm:": {
		Name:    "Bermuda Flag",
		Unicode: "\U0001F1E7\u1F1F2",
		Alias:   ":flag-bm:",
	},
	"\U0001F1E7\u1F1F3 :flag-bn:": {
		Name:    "Brunei Flag",
		Unicode: "\U0001F1E7\u1F1F3",
		Alias:   ":flag-bn:",
	},
	"\U0001F1E7\u1F1F4 :flag-bo:": {
		Name:    "Bolivia Flag",
		Unicode: "\U0001F1E7\u1F1F4",
		Alias:   ":flag-bo:",
	},
	"\U0001F1E7\u1F1F6 :flag-bq:": {
		Name:    "Caribbean Netherlands Flag",
		Unicode: "\U0001F1E7\u1F1F6",
		Alias:   ":flag-bq:",
	},
	"\U0001F1E7\u1F1F7 :flag-br:": {
		Name:    "Brazil Flag",
		Unicode: "\U0001F1E7\u1F1F7",
		Alias:   ":flag-br:",
	},
	"\U0001F1E7\u1F1F8 :flag-bs:": {
		Name:    "Bahamas Flag",
		Unicode: "\U0001F1E7\u1F1F8",
		Alias:   ":flag-bs:",
	},
	"\U0001F1E7\u1F1F9 :flag-bt:": {
		Name:    "Bhutan Flag",
		Unicode: "\U0001F1E7\u1F1F9",
		Alias:   ":flag-bt:",
	},
	"\U0001F1E7\u1F1FB :flag-bv:": {
		Name:    "Bouvet Island Flag",
		Unicode: "\U0001F1E7\u1F1FB",
		Alias:   ":flag-bv:",
	},
	"\U0001F1E7\u1F1FC :flag-bw:": {
		Name:    "Botswana Flag",
		Unicode: "\U0001F1E7\u1F1FC",
		Alias:   ":flag-bw:",
	},
	"\U0001F1E7\u1F1FE :flag-by:": {
		Name:    "Belarus Flag",
		Unicode: "\U0001F1E7\u1F1FE",
		Alias:   ":flag-by:",
	},
	"\U0001F1E7\u1F1FF :flag-bz:": {
		Name:    "Belize Flag",
		Unicode: "\U0001F1E7\u1F1FF",
		Alias:   ":flag-bz:",
	},
	"\U0001F1E8\u1F1E6 :flag-ca:": {
		Name:    "Canada Flag",
		Unicode: "\U0001F1E8\u1F1E6",
		Alias:   ":flag-ca:",
	},
	"\U0001F1E8\u1F1E8 :flag-cc:": {
		Name:    "Cocos (Keeling) Islands Flag",
		Unicode: "\U0001F1E8\u1F1E8",
		Alias:   ":flag-cc:",
	},
	"\U0001F1E8\u1F1E9 :flag-cd:": {
		Name:    "Congo - Kinshasa Flag",
		Unicode: "\U0001F1E8\u1F1E9",
		Alias:   ":flag-cd:",
	},
	"\U0001F1E8\u1F1EB :flag-cf:": {
		Name:    "Central African Republic Flag",
		Unicode: "\U0001F1E8\u1F1EB",
		Alias:   ":flag-cf:",
	},
	"\U0001F1E8\u1F1EC :flag-cg:": {
		Name:    "Congo - Brazzaville Flag",
		Unicode: "\U0001F1E8\u1F1EC",
		Alias:   ":flag-cg:",
	},
	"\U0001F1E8\u1F1ED :flag-ch:": {
		Name:    "Switzerland Flag",
		Unicode: "\U0001F1E8\u1F1ED",
		Alias:   ":flag-ch:",
	},
	"\U0001F1E8\u1F1EE :flag-ci:": {
		Name:    "C\u00f4te D\u2019Ivoire Flag",
		Unicode: "\U0001F1E8\u1F1EE",
		Alias:   ":flag-ci:",
	},
	"\U0001F1E8\u1F1F0 :flag-ck:": {
		Name:    "Cook Islands Flag",
		Unicode: "\U0001F1E8\u1F1F0",
		Alias:   ":flag-ck:",
	},
	"\U0001F1E8\u1F1F1 :flag-cl:": {
		Name:    "Chile Flag",
		Unicode: "\U0001F1E8\u1F1F1",
		Alias:   ":flag-cl:",
	},
	"\U0001F1E8\u1F1F2 :flag-cm:": {
		Name:    "Cameroon Flag",
		Unicode: "\U0001F1E8\u1F1F2",
		Alias:   ":flag-cm:",
	},
	"\U0001F1E8\u1F1F3  :cn:": {
		Name:    "China Flag",
		Unicode: "\U0001F1E8\u1F1F3",
		Alias:   ":cn:",
	},
	"\U0001F1E8\u1F1F4 :flag-co:": {
		Name:    "Colombia Flag",
		Unicode: "\U0001F1E8\u1F1F4",
		Alias:   ":flag-co:",
	},
	"\U0001F1E8\u1F1F5 :flag-cp:": {
		Name:    "Clipperton Island Flag",
		Unicode: "\U0001F1E8\u1F1F5",
		Alias:   ":flag-cp:",
	},
	"\U0001F1E8\u1F1F7 :flag-cr:": {
		Name:    "Costa Rica Flag",
		Unicode: "\U0001F1E8\u1F1F7",
		Alias:   ":flag-cr:",
	},
	"\U0001F1E8\u1F1FA :flag-cu:": {
		Name:    "Cuba Flag",
		Unicode: "\U0001F1E8\u1F1FA",
		Alias:   ":flag-cu:",
	},
	"\U0001F1E8\u1F1FB :flag-cv:": {
		Name:    "Cape Verde Flag",
		Unicode: "\U0001F1E8\u1F1FB",
		Alias:   ":flag-cv:",
	},
	"\U0001F1E8\u1F1FC :flag-cw:": {
		Name:    "Cura\u00e7ao Flag",
		Unicode: "\U0001F1E8\u1F1FC",
		Alias:   ":flag-cw:",
	},
	"\U0001F1E8\u1F1FD :flag-cx:": {
		Name:    "Christmas Island Flag",
		Unicode: "\U0001F1E8\u1F1FD",
		Alias:   ":flag-cx:",
	},
	"\U0001F1E8\u1F1FE :flag-cy:": {
		Name:    "Cyprus Flag",
		Unicode: "\U0001F1E8\u1F1FE",
		Alias:   ":flag-cy:",
	},
	"\U0001F1E8\u1F1FF :flag-cz:": {
		Name:    "Czechia Flag",
		Unicode: "\U0001F1E8\u1F1FF",
		Alias:   ":flag-cz:",
	},
	"\U0001F1E9\u1F1EA  :de:": {
		Name:    "Germany Flag",
		Unicode: "\U0001F1E9\u1F1EA",
		Alias:   ":de:",
	},
	"\U0001F1E9\u1F1EC :flag-dg:": {
		Name:    "Diego Garcia Flag",
		Unicode: "\U0001F1E9\u1F1EC",
		Alias:   ":flag-dg:",
	},
	"\U0001F1E9\u1F1EF :flag-dj:": {
		Name:    "Djibouti Flag",
		Unicode: "\U0001F1E9\u1F1EF",
		Alias:   ":flag-dj:",
	},
	"\U0001F1E9\u1F1F0 :flag-dk:": {
		Name:    "Denmark Flag",
		Unicode: "\U0001F1E9\u1F1F0",
		Alias:   ":flag-dk:",
	},
	"\U0001F1E9\u1F1F2 :flag-dm:": {
		Name:    "Dominica Flag",
		Unicode: "\U0001F1E9\u1F1F2",
		Alias:   ":flag-dm:",
	},
	"\U0001F1E9\u1F1F4 :flag-do:": {
		Name:    "Dominican Republic Flag",
		Unicode: "\U0001F1E9\u1F1F4",
		Alias:   ":flag-do:",
	},
	"\U0001F1E9\u1F1FF :flag-dz:": {
		Name:    "Algeria Flag",
		Unicode: "\U0001F1E9\u1F1FF",
		Alias:   ":flag-dz:",
	},
	"\U0001F1EA\u1F1E6 :flag-ea:": {
		Name:    "Ceuta & Melilla Flag",
		Unicode: "\U0001F1EA\u1F1E6",
		Alias:   ":flag-ea:",
	},
	"\U0001F1EA\u1F1E8 :flag-ec:": {
		Name:    "Ecuador Flag",
		Unicode: "\U0001F1EA\u1F1E8",
		Alias:   ":flag-ec:",
	},
	"\U0001F1EA\u1F1EA :flag-ee:": {
		Name:    "Estonia Flag",
		Unicode: "\U0001F1EA\u1F1EA",
		Alias:   ":flag-ee:",
	},
	"\U0001F1EA\u1F1EC :flag-eg:": {
		Name:    "Egypt Flag",
		Unicode: "\U0001F1EA\u1F1EC",
		Alias:   ":flag-eg:",
	},
	"\U0001F1EA\u1F1ED :flag-eh:": {
		Name:    "Western Sahara Flag",
		Unicode: "\U0001F1EA\u1F1ED",
		Alias:   ":flag-eh:",
	},
	"\U0001F1EA\u1F1F7 :flag-er:": {
		Name:    "Eritrea Flag",
		Unicode: "\U0001F1EA\u1F1F7",
		Alias:   ":flag-er:",
	},
	"\U0001F1EA\u1F1F8  :es:": {
		Name:    "Spain Flag",
		Unicode: "\U0001F1EA\u1F1F8",
		Alias:   ":es:",
	},
	"\U0001F1EA\u1F1F9 :flag-et:": {
		Name:    "Ethiopia Flag",
		Unicode: "\U0001F1EA\u1F1F9",
		Alias:   ":flag-et:",
	},
	"\U0001F1EA\u1F1FA :flag-eu:": {
		Name:    "European Union Flag",
		Unicode: "\U0001F1EA\u1F1FA",
		Alias:   ":flag-eu:",
	},
	"\U0001F1EB\u1F1EE :flag-fi:": {
		Name:    "Finland Flag",
		Unicode: "\U0001F1EB\u1F1EE",
		Alias:   ":flag-fi:",
	},
	"\U0001F1EB\u1F1EF :flag-fj:": {
		Name:    "Fiji Flag",
		Unicode: "\U0001F1EB\u1F1EF",
		Alias:   ":flag-fj:",
	},
	"\U0001F1EB\u1F1F0 :flag-fk:": {
		Name:    "Falkland Islands Flag",
		Unicode: "\U0001F1EB\u1F1F0",
		Alias:   ":flag-fk:",
	},
	"\U0001F1EB\u1F1F2 :flag-fm:": {
		Name:    "Micronesia Flag",
		Unicode: "\U0001F1EB\u1F1F2",
		Alias:   ":flag-fm:",
	},
	"\U0001F1EB\u1F1F4 :flag-fo:": {
		Name:    "Faroe Islands Flag",
		Unicode: "\U0001F1EB\u1F1F4",
		Alias:   ":flag-fo:",
	},
	"\U0001F1EB\u1F1F7  :fr:": {
		Name:    "France Flag",
		Unicode: "\U0001F1EB\u1F1F7",
		Alias:   ":fr:",
	},
	"\U0001F1EC\u1F1E6 :flag-ga:": {
		Name:    "Gabon Flag",
		Unicode: "\U0001F1EC\u1F1E6",
		Alias:   ":flag-ga:",
	},
	"\U0001F1EC\u1F1E7  :gb:": {
		Name:    "United Kingdom Flag",
		Unicode: "\U0001F1EC\u1F1E7",
		Alias:   ":gb:",
	},
	"\U0001F1EC\u1F1E9 :flag-gd:": {
		Name:    "Grenada Flag",
		Unicode: "\U0001F1EC\u1F1E9",
		Alias:   ":flag-gd:",
	},
	"\U0001F1EC\u1F1EA :flag-ge:": {
		Name:    "Georgia Flag",
		Unicode: "\U0001F1EC\u1F1EA",
		Alias:   ":flag-ge:",
	},
	"\U0001F1EC\u1F1EB :flag-gf:": {
		Name:    "French Guiana Flag",
		Unicode: "\U0001F1EC\u1F1EB",
		Alias:   ":flag-gf:",
	},
	"\U0001F1EC\u1F1EC :flag-gg:": {
		Name:    "Guernsey Flag",
		Unicode: "\U0001F1EC\u1F1EC",
		Alias:   ":flag-gg:",
	},
	"\U0001F1EC\u1F1ED :flag-gh:": {
		Name:    "Ghana Flag",
		Unicode: "\U0001F1EC\u1F1ED",
		Alias:   ":flag-gh:",
	},
	"\U0001F1EC\u1F1EE :flag-gi:": {
		Name:    "Gibraltar Flag",
		Unicode: "\U0001F1EC\u1F1EE",
		Alias:   ":flag-gi:",
	},
	"\U0001F1EC\u1F1F1 :flag-gl:": {
		Name:    "Greenland Flag",
		Unicode: "\U0001F1EC\u1F1F1",
		Alias:   ":flag-gl:",
	},
	"\U0001F1EC\u1F1F2 :flag-gm:": {
		Name:    "Gambia Flag",
		Unicode: "\U0001F1EC\u1F1F2",
		Alias:   ":flag-gm:",
	},
	"\U0001F1EC\u1F1F3 :flag-gn:": {
		Name:    "Guinea Flag",
		Unicode: "\U0001F1EC\u1F1F3",
		Alias:   ":flag-gn:",
	},
	"\U0001F1EC\u1F1F5 :flag-gp:": {
		Name:    "Guadeloupe Flag",
		Unicode: "\U0001F1EC\u1F1F5",
		Alias:   ":flag-gp:",
	},
	"\U0001F1EC\u1F1F6 :flag-gq:": {
		Name:    "Equatorial Guinea Flag",
		Unicode: "\U0001F1EC\u1F1F6",
		Alias:   ":flag-gq:",
	},
	"\U0001F1EC\u1F1F7 :flag-gr:": {
		Name:    "Greece Flag",
		Unicode: "\U0001F1EC\u1F1F7",
		Alias:   ":flag-gr:",
	},
	"\U0001F1EC\u1F1F8 :flag-gs:": {
		Name:    "South Georgia & South Sandwich Islands Flag",
		Unicode: "\U0001F1EC\u1F1F8",
		Alias:   ":flag-gs:",
	},
	"\U0001F1EC\u1F1F9 :flag-gt:": {
		Name:    "Guatemala Flag",
		Unicode: "\U0001F1EC\u1F1F9",
		Alias:   ":flag-gt:",
	},
	"\U0001F1EC\u1F1FA :flag-gu:": {
		Name:    "Guam Flag",
		Unicode: "\U0001F1EC\u1F1FA",
		Alias:   ":flag-gu:",
	},
	"\U0001F1EC\u1F1FC :flag-gw:": {
		Name:    "Guinea-Bissau Flag",
		Unicode: "\U0001F1EC\u1F1FC",
		Alias:   ":flag-gw:",
	},
	"\U0001F1EC\u1F1FE :flag-gy:": {
		Name:    "Guyana Flag",
		Unicode: "\U0001F1EC\u1F1FE",
		Alias:   ":flag-gy:",
	},
	"\U0001F1ED\u1F1F0 :flag-hk:": {
		Name:    "Hong Kong Sar China Flag",
		Unicode: "\U0001F1ED\u1F1F0",
		Alias:   ":flag-hk:",
	},
	"\U0001F1ED\u1F1F2 :flag-hm:": {
		Name:    "Heard & Mcdonald Islands Flag",
		Unicode: "\U0001F1ED\u1F1F2",
		Alias:   ":flag-hm:",
	},
	"\U0001F1ED\u1F1F3 :flag-hn:": {
		Name:    "Honduras Flag",
		Unicode: "\U0001F1ED\u1F1F3",
		Alias:   ":flag-hn:",
	},
	"\U0001F1ED\u1F1F7 :flag-hr:": {
		Name:    "Croatia Flag",
		Unicode: "\U0001F1ED\u1F1F7",
		Alias:   ":flag-hr:",
	},
	"\U0001F1ED\u1F1F9 :flag-ht:": {
		Name:    "Haiti Flag",
		Unicode: "\U0001F1ED\u1F1F9",
		Alias:   ":flag-ht:",
	},
	"\U0001F1ED\u1F1FA :flag-hu:": {
		Name:    "Hungary Flag",
		Unicode: "\U0001F1ED\u1F1FA",
		Alias:   ":flag-hu:",
	},
	"\U0001F1EE\u1F1E8 :flag-ic:": {
		Name:    "Canary Islands Flag",
		Unicode: "\U0001F1EE\u1F1E8",
		Alias:   ":flag-ic:",
	},
	"\U0001F1EE\u1F1E9 :flag-id:": {
		Name:    "Indonesia Flag",
		Unicode: "\U0001F1EE\u1F1E9",
		Alias:   ":flag-id:",
	},
	"\U0001F1EE\u1F1EA :flag-ie:": {
		Name:    "Ireland Flag",
		Unicode: "\U0001F1EE\u1F1EA",
		Alias:   ":flag-ie:",
	},
	"\U0001F1EE\u1F1F1 :flag-il:": {
		Name:    "Israel Flag",
		Unicode: "\U0001F1EE\u1F1F1",
		Alias:   ":flag-il:",
	},
	"\U0001F1EE\u1F1F2 :flag-im:": {
		Name:    "Isle Of Man Flag",
		Unicode: "\U0001F1EE\u1F1F2",
		Alias:   ":flag-im:",
	},
	"\U0001F1EE\u1F1F3 :flag-in:": {
		Name:    "India Flag",
		Unicode: "\U0001F1EE\u1F1F3",
		Alias:   ":flag-in:",
	},
	"\U0001F1EE\u1F1F4 :flag-io:": {
		Name:    "British Indian Ocean Territory Flag",
		Unicode: "\U0001F1EE\u1F1F4",
		Alias:   ":flag-io:",
	},
	"\U0001F1EE\u1F1F6 :flag-iq:": {
		Name:    "Iraq Flag",
		Unicode: "\U0001F1EE\u1F1F6",
		Alias:   ":flag-iq:",
	},
	"\U0001F1EE\u1F1F7 :flag-ir:": {
		Name:    "Iran Flag",
		Unicode: "\U0001F1EE\u1F1F7",
		Alias:   ":flag-ir:",
	},
	"\U0001F1EE\u1F1F8 :flag-is:": {
		Name:    "Iceland Flag",
		Unicode: "\U0001F1EE\u1F1F8",
		Alias:   ":flag-is:",
	},
	"\U0001F1EE\u1F1F9  :it:": {
		Name:    "Italy Flag",
		Unicode: "\U0001F1EE\u1F1F9",
		Alias:   ":it:",
	},
	"\U0001F1EF\u1F1EA :flag-je:": {
		Name:    "Jersey Flag",
		Unicode: "\U0001F1EF\u1F1EA",
		Alias:   ":flag-je:",
	},
	"\U0001F1EF\u1F1F2 :flag-jm:": {
		Name:    "Jamaica Flag",
		Unicode: "\U0001F1EF\u1F1F2",
		Alias:   ":flag-jm:",
	},
	"\U0001F1EF\u1F1F4 :flag-jo:": {
		Name:    "Jordan Flag",
		Unicode: "\U0001F1EF\u1F1F4",
		Alias:   ":flag-jo:",
	},
	"\U0001F1EF\u1F1F5  :jp:": {
		Name:    "Japan Flag",
		Unicode: "\U0001F1EF\u1F1F5",
		Alias:   ":jp:",
	},
	"\U0001F1F0\u1F1EA :flag-ke:": {
		Name:    "Kenya Flag",
		Unicode: "\U0001F1F0\u1F1EA",
		Alias:   ":flag-ke:",
	},
	"\U0001F1F0\u1F1EC :flag-kg:": {
		Name:    "Kyrgyzstan Flag",
		Unicode: "\U0001F1F0\u1F1EC",
		Alias:   ":flag-kg:",
	},
	"\U0001F1F0\u1F1ED :flag-kh:": {
		Name:    "Cambodia Flag",
		Unicode: "\U0001F1F0\u1F1ED",
		Alias:   ":flag-kh:",
	},
	"\U0001F1F0\u1F1EE :flag-ki:": {
		Name:    "Kiribati Flag",
		Unicode: "\U0001F1F0\u1F1EE",
		Alias:   ":flag-ki:",
	},
	"\U0001F1F0\u1F1F2 :flag-km:": {
		Name:    "Comoros Flag",
		Unicode: "\U0001F1F0\u1F1F2",
		Alias:   ":flag-km:",
	},
	"\U0001F1F0\u1F1F3 :flag-kn:": {
		Name:    "St. Kitts & Nevis Flag",
		Unicode: "\U0001F1F0\u1F1F3",
		Alias:   ":flag-kn:",
	},
	"\U0001F1F0\u1F1F5 :flag-kp:": {
		Name:    "North Korea Flag",
		Unicode: "\U0001F1F0\u1F1F5",
		Alias:   ":flag-kp:",
	},
	"\U0001F1F0\u1F1F7  :kr:": {
		Name:    "South Korea Flag",
		Unicode: "\U0001F1F0\u1F1F7",
		Alias:   ":kr:",
	},
	"\U0001F1F0\u1F1FC :flag-kw:": {
		Name:    "Kuwait Flag",
		Unicode: "\U0001F1F0\u1F1FC",
		Alias:   ":flag-kw:",
	},
	"\U0001F1F0\u1F1FE :flag-ky:": {
		Name:    "Cayman Islands Flag",
		Unicode: "\U0001F1F0\u1F1FE",
		Alias:   ":flag-ky:",
	},
	"\U0001F1F0\u1F1FF :flag-kz:": {
		Name:    "Kazakhstan Flag",
		Unicode: "\U0001F1F0\u1F1FF",
		Alias:   ":flag-kz:",
	},
	"\U0001F1F1\u1F1E6 :flag-la:": {
		Name:    "Laos Flag",
		Unicode: "\U0001F1F1\u1F1E6",
		Alias:   ":flag-la:",
	},
	"\U0001F1F1\u1F1E7 :flag-lb:": {
		Name:    "Lebanon Flag",
		Unicode: "\U0001F1F1\u1F1E7",
		Alias:   ":flag-lb:",
	},
	"\U0001F1F1\u1F1E8 :flag-lc:": {
		Name:    "St. Lucia Flag",
		Unicode: "\U0001F1F1\u1F1E8",
		Alias:   ":flag-lc:",
	},
	"\U0001F1F1\u1F1EE :flag-li:": {
		Name:    "Liechtenstein Flag",
		Unicode: "\U0001F1F1\u1F1EE",
		Alias:   ":flag-li:",
	},
	"\U0001F1F1\u1F1F0 :flag-lk:": {
		Name:    "Sri Lanka Flag",
		Unicode: "\U0001F1F1\u1F1F0",
		Alias:   ":flag-lk:",
	},
	"\U0001F1F1\u1F1F7 :flag-lr:": {
		Name:    "Liberia Flag",
		Unicode: "\U0001F1F1\u1F1F7",
		Alias:   ":flag-lr:",
	},
	"\U0001F1F1\u1F1F8 :flag-ls:": {
		Name:    "Lesotho Flag",
		Unicode: "\U0001F1F1\u1F1F8",
		Alias:   ":flag-ls:",
	},
	"\U0001F1F1\u1F1F9 :flag-lt:": {
		Name:    "Lithuania Flag",
		Unicode: "\U0001F1F1\u1F1F9",
		Alias:   ":flag-lt:",
	},
	"\U0001F1F1\u1F1FA :flag-lu:": {
		Name:    "Luxembourg Flag",
		Unicode: "\U0001F1F1\u1F1FA",
		Alias:   ":flag-lu:",
	},
	"\U0001F1F1\u1F1FB :flag-lv:": {
		Name:    "Latvia Flag",
		Unicode: "\U0001F1F1\u1F1FB",
		Alias:   ":flag-lv:",
	},
	"\U0001F1F1\u1F1FE :flag-ly:": {
		Name:    "Libya Flag",
		Unicode: "\U0001F1F1\u1F1FE",
		Alias:   ":flag-ly:",
	},
	"\U0001F1F2\u1F1E6 :flag-ma:": {
		Name:    "Morocco Flag",
		Unicode: "\U0001F1F2\u1F1E6",
		Alias:   ":flag-ma:",
	},
	"\U0001F1F2\u1F1E8 :flag-mc:": {
		Name:    "Monaco Flag",
		Unicode: "\U0001F1F2\u1F1E8",
		Alias:   ":flag-mc:",
	},
	"\U0001F1F2\u1F1E9 :flag-md:": {
		Name:    "Moldova Flag",
		Unicode: "\U0001F1F2\u1F1E9",
		Alias:   ":flag-md:",
	},
	"\U0001F1F2\u1F1EA :flag-me:": {
		Name:    "Montenegro Flag",
		Unicode: "\U0001F1F2\u1F1EA",
		Alias:   ":flag-me:",
	},
	"\U0001F1F2\u1F1EB :flag-mf:": {
		Name:    "St. Martin Flag",
		Unicode: "\U0001F1F2\u1F1EB",
		Alias:   ":flag-mf:",
	},
	"\U0001F1F2\u1F1EC :flag-mg:": {
		Name:    "Madagascar Flag",
		Unicode: "\U0001F1F2\u1F1EC",
		Alias:   ":flag-mg:",
	},
	"\U0001F1F2\u1F1ED :flag-mh:": {
		Name:    "Marshall Islands Flag",
		Unicode: "\U0001F1F2\u1F1ED",
		Alias:   ":flag-mh:",
	},
	"\U0001F1F2\u1F1F0 :flag-mk:": {
		Name:    "North Macedonia Flag",
		Unicode: "\U0001F1F2\u1F1F0",
		Alias:   ":flag-mk:",
	},
	"\U0001F1F2\u1F1F1 :flag-ml:": {
		Name:    "Mali Flag",
		Unicode: "\U0001F1F2\u1F1F1",
		Alias:   ":flag-ml:",
	},
	"\U0001F1F2\u1F1F2 :flag-mm:": {
		Name:    "Myanmar (Burma) Flag",
		Unicode: "\U0001F1F2\u1F1F2",
		Alias:   ":flag-mm:",
	},
	"\U0001F1F2\u1F1F3 :flag-mn:": {
		Name:    "Mongolia Flag",
		Unicode: "\U0001F1F2\u1F1F3",
		Alias:   ":flag-mn:",
	},
	"\U0001F1F2\u1F1F4 :flag-mo:": {
		Name:    "Macao Sar China Flag",
		Unicode: "\U0001F1F2\u1F1F4",
		Alias:   ":flag-mo:",
	},
	"\U0001F1F2\u1F1F5 :flag-mp:": {
		Name:    "Northern Mariana Islands Flag",
		Unicode: "\U0001F1F2\u1F1F5",
		Alias:   ":flag-mp:",
	},
	"\U0001F1F2\u1F1F6 :flag-mq:": {
		Name:    "Martinique Flag",
		Unicode: "\U0001F1F2\u1F1F6",
		Alias:   ":flag-mq:",
	},
	"\U0001F1F2\u1F1F7 :flag-mr:": {
		Name:    "Mauritania Flag",
		Unicode: "\U0001F1F2\u1F1F7",
		Alias:   ":flag-mr:",
	},
	"\U0001F1F2\u1F1F8 :flag-ms:": {
		Name:    "Montserrat Flag",
		Unicode: "\U0001F1F2\u1F1F8",
		Alias:   ":flag-ms:",
	},
	"\U0001F1F2\u1F1F9 :flag-mt:": {
		Name:    "Malta Flag",
		Unicode: "\U0001F1F2\u1F1F9",
		Alias:   ":flag-mt:",
	},
	"\U0001F1F2\u1F1FA :flag-mu:": {
		Name:    "Mauritius Flag",
		Unicode: "\U0001F1F2\u1F1FA",
		Alias:   ":flag-mu:",
	},
	"\U0001F1F2\u1F1FB :flag-mv:": {
		Name:    "Maldives Flag",
		Unicode: "\U0001F1F2\u1F1FB",
		Alias:   ":flag-mv:",
	},
	"\U0001F1F2\u1F1FC :flag-mw:": {
		Name:    "Malawi Flag",
		Unicode: "\U0001F1F2\u1F1FC",
		Alias:   ":flag-mw:",
	},
	"\U0001F1F2\u1F1FD :flag-mx:": {
		Name:    "Mexico Flag",
		Unicode: "\U0001F1F2\u1F1FD",
		Alias:   ":flag-mx:",
	},
	"\U0001F1F2\u1F1FE :flag-my:": {
		Name:    "Malaysia Flag",
		Unicode: "\U0001F1F2\u1F1FE",
		Alias:   ":flag-my:",
	},
	"\U0001F1F2\u1F1FF :flag-mz:": {
		Name:    "Mozambique Flag",
		Unicode: "\U0001F1F2\u1F1FF",
		Alias:   ":flag-mz:",
	},
	"\U0001F1F3\u1F1E6 :flag-na:": {
		Name:    "Namibia Flag",
		Unicode: "\U0001F1F3\u1F1E6",
		Alias:   ":flag-na:",
	},
	"\U0001F1F3\u1F1E8 :flag-nc:": {
		Name:    "New Caledonia Flag",
		Unicode: "\U0001F1F3\u1F1E8",
		Alias:   ":flag-nc:",
	},
	"\U0001F1F3\u1F1EA :flag-ne:": {
		Name:    "Niger Flag",
		Unicode: "\U0001F1F3\u1F1EA",
		Alias:   ":flag-ne:",
	},
	"\U0001F1F3\u1F1EB :flag-nf:": {
		Name:    "Norfolk Island Flag",
		Unicode: "\U0001F1F3\u1F1EB",
		Alias:   ":flag-nf:",
	},
	"\U0001F1F3\u1F1EC :flag-ng:": {
		Name:    "Nigeria Flag",
		Unicode: "\U0001F1F3\u1F1EC",
		Alias:   ":flag-ng:",
	},
	"\U0001F1F3\u1F1EE :flag-ni:": {
		Name:    "Nicaragua Flag",
		Unicode: "\U0001F1F3\u1F1EE",
		Alias:   ":flag-ni:",
	},
	"\U0001F1F3\u1F1F1 :flag-nl:": {
		Name:    "Netherlands Flag",
		Unicode: "\U0001F1F3\u1F1F1",
		Alias:   ":flag-nl:",
	},
	"\U0001F1F3\u1F1F4 :flag-no:": {
		Name:    "Norway Flag",
		Unicode: "\U0001F1F3\u1F1F4",
		Alias:   ":flag-no:",
	},
	"\U0001F1F3\u1F1F5 :flag-np:": {
		Name:    "Nepal Flag",
		Unicode: "\U0001F1F3\u1F1F5",
		Alias:   ":flag-np:",
	},
	"\U0001F1F3\u1F1F7 :flag-nr:": {
		Name:    "Nauru Flag",
		Unicode: "\U0001F1F3\u1F1F7",
		Alias:   ":flag-nr:",
	},
	"\U0001F1F3\u1F1FA :flag-nu:": {
		Name:    "Niue Flag",
		Unicode: "\U0001F1F3\u1F1FA",
		Alias:   ":flag-nu:",
	},
	"\U0001F1F3\u1F1FF :flag-nz:": {
		Name:    "New Zealand Flag",
		Unicode: "\U0001F1F3\u1F1FF",
		Alias:   ":flag-nz:",
	},
	"\U0001F1F4\u1F1F2 :flag-om:": {
		Name:    "Oman Flag",
		Unicode: "\U0001F1F4\u1F1F2",
		Alias:   ":flag-om:",
	},
	"\U0001F1F5\u1F1E6 :flag-pa:": {
		Name:    "Panama Flag",
		Unicode: "\U0001F1F5\u1F1E6",
		Alias:   ":flag-pa:",
	},
	"\U0001F1F5\u1F1EA :flag-pe:": {
		Name:    "Peru Flag",
		Unicode: "\U0001F1F5\u1F1EA",
		Alias:   ":flag-pe:",
	},
	"\U0001F1F5\u1F1EB :flag-pf:": {
		Name:    "French Polynesia Flag",
		Unicode: "\U0001F1F5\u1F1EB",
		Alias:   ":flag-pf:",
	},
	"\U0001F1F5\u1F1EC :flag-pg:": {
		Name:    "Papua New Guinea Flag",
		Unicode: "\U0001F1F5\u1F1EC",
		Alias:   ":flag-pg:",
	},
	"\U0001F1F5\u1F1ED :flag-ph:": {
		Name:    "Philippines Flag",
		Unicode: "\U0001F1F5\u1F1ED",
		Alias:   ":flag-ph:",
	},
	"\U0001F1F5\u1F1F0 :flag-pk:": {
		Name:    "Pakistan Flag",
		Unicode: "\U0001F1F5\u1F1F0",
		Alias:   ":flag-pk:",
	},
	"\U0001F1F5\u1F1F1 :flag-pl:": {
		Name:    "Poland Flag",
		Unicode: "\U0001F1F5\u1F1F1",
		Alias:   ":flag-pl:",
	},
	"\U0001F1F5\u1F1F2 :flag-pm:": {
		Name:    "St. Pierre & Miquelon Flag",
		Unicode: "\U0001F1F5\u1F1F2",
		Alias:   ":flag-pm:",
	},
	"\U0001F1F5\u1F1F3 :flag-pn:": {
		Name:    "Pitcairn Islands Flag",
		Unicode: "\U0001F1F5\u1F1F3",
		Alias:   ":flag-pn:",
	},
	"\U0001F1F5\u1F1F7 :flag-pr:": {
		Name:    "Puerto Rico Flag",
		Unicode: "\U0001F1F5\u1F1F7",
		Alias:   ":flag-pr:",
	},
	"\U0001F1F5\u1F1F8 :flag-ps:": {
		Name:    "Palestinian Territories Flag",
		Unicode: "\U0001F1F5\u1F1F8",
		Alias:   ":flag-ps:",
	},
	"\U0001F1F5\u1F1F9 :flag-pt:": {
		Name:    "Portugal Flag",
		Unicode: "\U0001F1F5\u1F1F9",
		Alias:   ":flag-pt:",
	},
	"\U0001F1F5\u1F1FC :flag-pw:": {
		Name:    "Palau Flag",
		Unicode: "\U0001F1F5\u1F1FC",
		Alias:   ":flag-pw:",
	},
	"\U0001F1F5\u1F1FE :flag-py:": {
		Name:    "Paraguay Flag",
		Unicode: "\U0001F1F5\u1F1FE",
		Alias:   ":flag-py:",
	},
	"\U0001F1F6\u1F1E6 :flag-qa:": {
		Name:    "Qatar Flag",
		Unicode: "\U0001F1F6\u1F1E6",
		Alias:   ":flag-qa:",
	},
	"\U0001F1F7\u1F1EA :flag-re:": {
		Name:    "R\u00e9union Flag",
		Unicode: "\U0001F1F7\u1F1EA",
		Alias:   ":flag-re:",
	},
	"\U0001F1F7\u1F1F4 :flag-ro:": {
		Name:    "Romania Flag",
		Unicode: "\U0001F1F7\u1F1F4",
		Alias:   ":flag-ro:",
	},
	"\U0001F1F7\u1F1F8 :flag-rs:": {
		Name:    "Serbia Flag",
		Unicode: "\U0001F1F7\u1F1F8",
		Alias:   ":flag-rs:",
	},
	"\U0001F1F7\u1F1FA  :ru:": {
		Name:    "Russia Flag",
		Unicode: "\U0001F1F7\u1F1FA",
		Alias:   ":ru:",
	},
	"\U0001F1F7\u1F1FC :flag-rw:": {
		Name:    "Rwanda Flag",
		Unicode: "\U0001F1F7\u1F1FC",
		Alias:   ":flag-rw:",
	},
	"\U0001F1F8\u1F1E6 :flag-sa:": {
		Name:    "Saudi Arabia Flag",
		Unicode: "\U0001F1F8\u1F1E6",
		Alias:   ":flag-sa:",
	},
	"\U0001F1F8\u1F1E7 :flag-sb:": {
		Name:    "Solomon Islands Flag",
		Unicode: "\U0001F1F8\u1F1E7",
		Alias:   ":flag-sb:",
	},
	"\U0001F1F8\u1F1E8 :flag-sc:": {
		Name:    "Seychelles Flag",
		Unicode: "\U0001F1F8\u1F1E8",
		Alias:   ":flag-sc:",
	},
	"\U0001F1F8\u1F1E9 :flag-sd:": {
		Name:    "Sudan Flag",
		Unicode: "\U0001F1F8\u1F1E9",
		Alias:   ":flag-sd:",
	},
	"\U0001F1F8\u1F1EA :flag-se:": {
		Name:    "Sweden Flag",
		Unicode: "\U0001F1F8\u1F1EA",
		Alias:   ":flag-se:",
	},
	"\U0001F1F8\u1F1EC :flag-sg:": {
		Name:    "Singapore Flag",
		Unicode: "\U0001F1F8\u1F1EC",
		Alias:   ":flag-sg:",
	},
	"\U0001F1F8\u1F1ED :flag-sh:": {
		Name:    "St. Helena Flag",
		Unicode: "\U0001F1F8\u1F1ED",
		Alias:   ":flag-sh:",
	},
	"\U0001F1F8\u1F1EE :flag-si:": {
		Name:    "Slovenia Flag",
		Unicode: "\U0001F1F8\u1F1EE",
		Alias:   ":flag-si:",
	},
	"\U0001F1F8\u1F1EF :flag-sj:": {
		Name:    "Svalbard & Jan Mayen Flag",
		Unicode: "\U0001F1F8\u1F1EF",
		Alias:   ":flag-sj:",
	},
	"\U0001F1F8\u1F1F0 :flag-sk:": {
		Name:    "Slovakia Flag",
		Unicode: "\U0001F1F8\u1F1F0",
		Alias:   ":flag-sk:",
	},
	"\U0001F1F8\u1F1F1 :flag-sl:": {
		Name:    "Sierra Leone Flag",
		Unicode: "\U0001F1F8\u1F1F1",
		Alias:   ":flag-sl:",
	},
	"\U0001F1F8\u1F1F2 :flag-sm:": {
		Name:    "San Marino Flag",
		Unicode: "\U0001F1F8\u1F1F2",
		Alias:   ":flag-sm:",
	},
	"\U0001F1F8\u1F1F3 :flag-sn:": {
		Name:    "Senegal Flag",
		Unicode: "\U0001F1F8\u1F1F3",
		Alias:   ":flag-sn:",
	},
	"\U0001F1F8\u1F1F4 :flag-so:": {
		Name:    "Somalia Flag",
		Unicode: "\U0001F1F8\u1F1F4",
		Alias:   ":flag-so:",
	},
	"\U0001F1F8\u1F1F7 :flag-sr:": {
		Name:    "Suriname Flag",
		Unicode: "\U0001F1F8\u1F1F7",
		Alias:   ":flag-sr:",
	},
	"\U0001F1F8\u1F1F8 :flag-ss:": {
		Name:    "South Sudan Flag",
		Unicode: "\U0001F1F8\u1F1F8",
		Alias:   ":flag-ss:",
	},
	"\U0001F1F8\u1F1F9 :flag-st:": {
		Name:    "S\u00e3o Tom\u00e9 & Pr\u00edncipe Flag",
		Unicode: "\U0001F1F8\u1F1F9",
		Alias:   ":flag-st:",
	},
	"\U0001F1F8\u1F1FB :flag-sv:": {
		Name:    "El Salvador Flag",
		Unicode: "\U0001F1F8\u1F1FB",
		Alias:   ":flag-sv:",
	},
	"\U0001F1F8\u1F1FD :flag-sx:": {
		Name:    "Sint Maarten Flag",
		Unicode: "\U0001F1F8\u1F1FD",
		Alias:   ":flag-sx:",
	},
	"\U0001F1F8\u1F1FE :flag-sy:": {
		Name:    "Syria Flag",
		Unicode: "\U0001F1F8\u1F1FE",
		Alias:   ":flag-sy:",
	},
	"\U0001F1F8\u1F1FF :flag-sz:": {
		Name:    "Eswatini Flag",
		Unicode: "\U0001F1F8\u1F1FF",
		Alias:   ":flag-sz:",
	},
	"\U0001F1F9\u1F1E6 :flag-ta:": {
		Name:    "Tristan Da Cunha Flag",
		Unicode: "\U0001F1F9\u1F1E6",
		Alias:   ":flag-ta:",
	},
	"\U0001F1F9\u1F1E8 :flag-tc:": {
		Name:    "Turks & Caicos Islands Flag",
		Unicode: "\U0001F1F9\u1F1E8",
		Alias:   ":flag-tc:",
	},
	"\U0001F1F9\u1F1E9 :flag-td:": {
		Name:    "Chad Flag",
		Unicode: "\U0001F1F9\u1F1E9",
		Alias:   ":flag-td:",
	},
	"\U0001F1F9\u1F1EB :flag-tf:": {
		Name:    "French Southern Territories Flag",
		Unicode: "\U0001F1F9\u1F1EB",
		Alias:   ":flag-tf:",
	},
	"\U0001F1F9\u1F1EC :flag-tg:": {
		Name:    "Togo Flag",
		Unicode: "\U0001F1F9\u1F1EC",
		Alias:   ":flag-tg:",
	},
	"\U0001F1F9\u1F1ED :flag-th:": {
		Name:    "Thailand Flag",
		Unicode: "\U0001F1F9\u1F1ED",
		Alias:   ":flag-th:",
	},
	"\U0001F1F9\u1F1EF :flag-tj:": {
		Name:    "Tajikistan Flag",
		Unicode: "\U0001F1F9\u1F1EF",
		Alias:   ":flag-tj:",
	},
	"\U0001F1F9\u1F1F0 :flag-tk:": {
		Name:    "Tokelau Flag",
		Unicode: "\U0001F1F9\u1F1F0",
		Alias:   ":flag-tk:",
	},
	"\U0001F1F9\u1F1F1 :flag-tl:": {
		Name:    "Timor-Leste Flag",
		Unicode: "\U0001F1F9\u1F1F1",
		Alias:   ":flag-tl:",
	},
	"\U0001F1F9\u1F1F2 :flag-tm:": {
		Name:    "Turkmenistan Flag",
		Unicode: "\U0001F1F9\u1F1F2",
		Alias:   ":flag-tm:",
	},
	"\U0001F1F9\u1F1F3 :flag-tn:": {
		Name:    "Tunisia Flag",
		Unicode: "\U0001F1F9\u1F1F3",
		Alias:   ":flag-tn:",
	},
	"\U0001F1F9\u1F1F4 :flag-to:": {
		Name:    "Tonga Flag",
		Unicode: "\U0001F1F9\u1F1F4",
		Alias:   ":flag-to:",
	},
	"\U0001F1F9\u1F1F7 :flag-tr:": {
		Name:    "Turkey Flag",
		Unicode: "\U0001F1F9\u1F1F7",
		Alias:   ":flag-tr:",
	},
	"\U0001F1F9\u1F1F9 :flag-tt:": {
		Name:    "Trinidad & Tobago Flag",
		Unicode: "\U0001F1F9\u1F1F9",
		Alias:   ":flag-tt:",
	},
	"\U0001F1F9\u1F1FB :flag-tv:": {
		Name:    "Tuvalu Flag",
		Unicode: "\U0001F1F9\u1F1FB",
		Alias:   ":flag-tv:",
	},
	"\U0001F1F9\u1F1FC :flag-tw:": {
		Name:    "Taiwan Flag",
		Unicode: "\U0001F1F9\u1F1FC",
		Alias:   ":flag-tw:",
	},
	"\U0001F1F9\u1F1FF :flag-tz:": {
		Name:    "Tanzania Flag",
		Unicode: "\U0001F1F9\u1F1FF",
		Alias:   ":flag-tz:",
	},
	"\U0001F1FA\u1F1E6 :flag-ua:": {
		Name:    "Ukraine Flag",
		Unicode: "\U0001F1FA\u1F1E6",
		Alias:   ":flag-ua:",
	},
	"\U0001F1FA\u1F1EC :flag-ug:": {
		Name:    "Uganda Flag",
		Unicode: "\U0001F1FA\u1F1EC",
		Alias:   ":flag-ug:",
	},
	"\U0001F1FA\u1F1F2 :flag-um:": {
		Name:    "U.S. Outlying Islands Flag",
		Unicode: "\U0001F1FA\u1F1F2",
		Alias:   ":flag-um:",
	},
	"\U0001F1FA\u1F1F3 :flag-un:": {
		Name:    "United Nations Flag",
		Unicode: "\U0001F1FA\u1F1F3",
		Alias:   ":flag-un:",
	},
	"\U0001F1FA\u1F1F8  :us:": {
		Name:    "United States Flag",
		Unicode: "\U0001F1FA\u1F1F8",
		Alias:   ":us:",
	},
	"\U0001F1FA\u1F1FE :flag-uy:": {
		Name:    "Uruguay Flag",
		Unicode: "\U0001F1FA\u1F1FE",
		Alias:   ":flag-uy:",
	},
	"\U0001F1FA\u1F1FF :flag-uz:": {
		Name:    "Uzbekistan Flag",
		Unicode: "\U0001F1FA\u1F1FF",
		Alias:   ":flag-uz:",
	},
	"\U0001F1FB\u1F1E6 :flag-va:": {
		Name:    "Vatican City Flag",
		Unicode: "\U0001F1FB\u1F1E6",
		Alias:   ":flag-va:",
	},
	"\U0001F1FB\u1F1E8 :flag-vc:": {
		Name:    "St. Vincent & Grenadines Flag",
		Unicode: "\U0001F1FB\u1F1E8",
		Alias:   ":flag-vc:",
	},
	"\U0001F1FB\u1F1EA :flag-ve:": {
		Name:    "Venezuela Flag",
		Unicode: "\U0001F1FB\u1F1EA",
		Alias:   ":flag-ve:",
	},
	"\U0001F1FB\u1F1EC :flag-vg:": {
		Name:    "British Virgin Islands Flag",
		Unicode: "\U0001F1FB\u1F1EC",
		Alias:   ":flag-vg:",
	},
	"\U0001F1FB\u1F1EE :flag-vi:": {
		Name:    "U.S. Virgin Islands Flag",
		Unicode: "\U0001F1FB\u1F1EE",
		Alias:   ":flag-vi:",
	},
	"\U0001F1FB\u1F1F3 :flag-vn:": {
		Name:    "Vietnam Flag",
		Unicode: "\U0001F1FB\u1F1F3",
		Alias:   ":flag-vn:",
	},
	"\U0001F1FB\u1F1FA :flag-vu:": {
		Name:    "Vanuatu Flag",
		Unicode: "\U0001F1FB\u1F1FA",
		Alias:   ":flag-vu:",
	},
	"\U0001F1FC\u1F1EB :flag-wf:": {
		Name:    "Wallis & Futuna Flag",
		Unicode: "\U0001F1FC\u1F1EB",
		Alias:   ":flag-wf:",
	},
	"\U0001F1FC\u1F1F8 :flag-ws:": {
		Name:    "Samoa Flag",
		Unicode: "\U0001F1FC\u1F1F8",
		Alias:   ":flag-ws:",
	},
	"\U0001F1FD\u1F1F0 :flag-xk:": {
		Name:    "Kosovo Flag",
		Unicode: "\U0001F1FD\u1F1F0",
		Alias:   ":flag-xk:",
	},
	"\U0001F1FE\u1F1EA :flag-ye:": {
		Name:    "Yemen Flag",
		Unicode: "\U0001F1FE\u1F1EA",
		Alias:   ":flag-ye:",
	},
	"\U0001F1FE\u1F1F9 :flag-yt:": {
		Name:    "Mayotte Flag",
		Unicode: "\U0001F1FE\u1F1F9",
		Alias:   ":flag-yt:",
	},
	"\U0001F1FF\u1F1E6 :flag-za:": {
		Name:    "South Africa Flag",
		Unicode: "\U0001F1FF\u1F1E6",
		Alias:   ":flag-za:",
	},
	"\U0001F1FF\u1F1F2 :flag-zm:": {
		Name:    "Zambia Flag",
		Unicode: "\U0001F1FF\u1F1F2",
		Alias:   ":flag-zm:",
	},
	"\U0001F1FF\u1F1FC :flag-zw:": {
		Name:    "Zimbabwe Flag",
		Unicode: "\U0001F1FF\u1F1FC",
		Alias:   ":flag-zw:",
	},
	"\U0001F201  :koko:": {
		Name:    "Squared Katakana Koko",
		Unicode: "\U0001F201",
		Alias:   ":koko:",
	},
	"\U0001F202\uFE0F  :sa:": {
		Name:    "Squared Katakana Sa",
		Unicode: "\U0001F202\uFE0F",
		Alias:   ":sa:",
	},
	"\U0001F21A  :u7121:": {
		Name:    "Squared Cjk Unified Ideograph-7121",
		Unicode: "\U0001F21A",
		Alias:   ":u7121:",
	},
	"\U0001F22F  :u6307:": {
		Name:    "Squared Cjk Unified Ideograph-6307",
		Unicode: "\U0001F22F",
		Alias:   ":u6307:",
	},
	"\U0001F232  :u7981:": {
		Name:    "Squared Cjk Unified Ideograph-7981",
		Unicode: "\U0001F232",
		Alias:   ":u7981:",
	},
	"\U0001F233  :u7a7a:": {
		Name:    "Squared Cjk Unified Ideograph-7A7A",
		Unicode: "\U0001F233",
		Alias:   ":u7a7a:",
	},
	"\U0001F234  :u5408:": {
		Name:    "Squared Cjk Unified Ideograph-5408",
		Unicode: "\U0001F234",
		Alias:   ":u5408:",
	},
	"\U0001F235  :u6e80:": {
		Name:    "Squared Cjk Unified Ideograph-6E80",
		Unicode: "\U0001F235",
		Alias:   ":u6e80:",
	},
	"\U0001F236  :u6709:": {
		Name:    "Squared Cjk Unified Ideograph-6709",
		Unicode: "\U0001F236",
		Alias:   ":u6709:",
	},
	"\U0001F237\uFE0F  :u6708:": {
		Name:    "Squared Cjk Unified Ideograph-6708",
		Unicode: "\U0001F237\uFE0F",
		Alias:   ":u6708:",
	},
	"\U0001F238  :u7533:": {
		Name:    "Squared Cjk Unified Ideograph-7533",
		Unicode: "\U0001F238",
		Alias:   ":u7533:",
	},
	"\U0001F239  :u5272:": {
		Name:    "Squared Cjk Unified Ideograph-5272",
		Unicode: "\U0001F239",
		Alias:   ":u5272:",
	},
	"\U0001F23A  :u55b6:": {
		Name:    "Squared Cjk Unified Ideograph-55B6",
		Unicode: "\U0001F23A",
		Alias:   ":u55b6:",
	},
	"\U0001F250  :ideograph_advantage:": {
		Name:    "Circled Ideograph Advantage",
		Unicode: "\U0001F250",
		Alias:   ":ideograph_advantage:",
	},
	"\U0001F251  :accept:": {
		Name:    "Circled Ideograph Accept",
		Unicode: "\U0001F251",
		Alias:   ":accept:",
	},
	"\U0001F300  :cyclone:": {
		Name:    "Cyclone",
		Unicode: "\U0001F300",
		Alias:   ":cyclone:",
	},
	"\U0001F301  :foggy:": {
		Name:    "Foggy",
		Unicode: "\U0001F301",
		Alias:   ":foggy:",
	},
	"\U0001F302  :closed_umbrella:": {
		Name:    "Closed Umbrella",
		Unicode: "\U0001F302",
		Alias:   ":closed_umbrella:",
	},
	"\U0001F303  :night_with_stars:": {
		Name:    "Night With Stars",
		Unicode: "\U0001F303",
		Alias:   ":night_with_stars:",
	},
	"\U0001F304  :sunrise_over_mountains:": {
		Name:    "Sunrise Over Mountains",
		Unicode: "\U0001F304",
		Alias:   ":sunrise_over_mountains:",
	},
	"\U0001F305  :sunrise:": {
		Name:    "Sunrise",
		Unicode: "\U0001F305",
		Alias:   ":sunrise:",
	},
	"\U0001F306  :city_sunset:": {
		Name:    "Cityscape At Dusk",
		Unicode: "\U0001F306",
		Alias:   ":city_sunset:",
	},
	"\U0001F307  :city_sunrise:": {
		Name:    "Sunset Over Buildings",
		Unicode: "\U0001F307",
		Alias:   ":city_sunrise:",
	},
	"\U0001F308  :rainbow:": {
		Name:    "Rainbow",
		Unicode: "\U0001F308",
		Alias:   ":rainbow:",
	},
	"\U0001F309  :bridge_at_night:": {
		Name:    "Bridge At Night",
		Unicode: "\U0001F309",
		Alias:   ":bridge_at_night:",
	},
	"\U0001F30A  :ocean:": {
		Name:    "Water Wave",
		Unicode: "\U0001F30A",
		Alias:   ":ocean:",
	},
	"\U0001F30B  :volcano:": {
		Name:    "Volcano",
		Unicode: "\U0001F30B",
		Alias:   ":volcano:",
	},
	"\U0001F30C  :milky_way:": {
		Name:    "Milky Way",
		Unicode: "\U0001F30C",
		Alias:   ":milky_way:",
	},
	"\U0001F30D  :earth_africa:": {
		Name:    "Earth Globe Europe-Africa",
		Unicode: "\U0001F30D",
		Alias:   ":earth_africa:",
	},
	"\U0001F30E  :earth_americas:": {
		Name:    "Earth Globe Americas",
		Unicode: "\U0001F30E",
		Alias:   ":earth_americas:",
	},
	"\U0001F30F  :earth_asia:": {
		Name:    "Earth Globe Asia-Australia",
		Unicode: "\U0001F30F",
		Alias:   ":earth_asia:",
	},
	"\U0001F310  :globe_with_meridians:": {
		Name:    "Globe With Meridians",
		Unicode: "\U0001F310",
		Alias:   ":globe_with_meridians:",
	},
	"\U0001F311  :new_moon:": {
		Name:    "New Moon Symbol",
		Unicode: "\U0001F311",
		Alias:   ":new_moon:",
	},
	"\U0001F312  :waxing_crescent_moon:": {
		Name:    "Waxing Crescent Moon Symbol",
		Unicode: "\U0001F312",
		Alias:   ":waxing_crescent_moon:",
	},
	"\U0001F313  :first_quarter_moon:": {
		Name:    "First Quarter Moon Symbol",
		Unicode: "\U0001F313",
		Alias:   ":first_quarter_moon:",
	},
	"\U0001F314  :moon:": {
		Name:    "Waxing Gibbous Moon Symbol",
		Unicode: "\U0001F314",
		Alias:   ":moon:",
	},
	"\U0001F315  :full_moon:": {
		Name:    "Full Moon Symbol",
		Unicode: "\U0001F315",
		Alias:   ":full_moon:",
	},
	"\U0001F316  :waning_gibbous_moon:": {
		Name:    "Waning Gibbous Moon Symbol",
		Unicode: "\U0001F316",
		Alias:   ":waning_gibbous_moon:",
	},
	"\U0001F317  :last_quarter_moon:": {
		Name:    "Last Quarter Moon Symbol",
		Unicode: "\U0001F317",
		Alias:   ":last_quarter_moon:",
	},
	"\U0001F318  :waning_crescent_moon:": {
		Name:    "Waning Crescent Moon Symbol",
		Unicode: "\U0001F318",
		Alias:   ":waning_crescent_moon:",
	},
	"\U0001F319  :crescent_moon:": {
		Name:    "Crescent Moon",
		Unicode: "\U0001F319",
		Alias:   ":crescent_moon:",
	},
	"\U0001F31A  :new_moon_with_face:": {
		Name:    "New Moon With Face",
		Unicode: "\U0001F31A",
		Alias:   ":new_moon_with_face:",
	},
	"\U0001F31B  :first_quarter_moon_with_face:": {
		Name:    "First Quarter Moon With Face",
		Unicode: "\U0001F31B",
		Alias:   ":first_quarter_moon_with_face:",
	},
	"\U0001F31C  :last_quarter_moon_with_face:": {
		Name:    "Last Quarter Moon With Face",
		Unicode: "\U0001F31C",
		Alias:   ":last_quarter_moon_with_face:",
	},
	"\U0001F31D  :full_moon_with_face:": {
		Name:    "Full Moon With Face",
		Unicode: "\U0001F31D",
		Alias:   ":full_moon_with_face:",
	},
	"\U0001F31E  :sun_with_face:": {
		Name:    "Sun With Face",
		Unicode: "\U0001F31E",
		Alias:   ":sun_with_face:",
	},
	"\U0001F31F  :star2:": {
		Name:    "Glowing Star",
		Unicode: "\U0001F31F",
		Alias:   ":star2:",
	},
	"\U0001F320  :stars:": {
		Name:    "Shooting Star",
		Unicode: "\U0001F320",
		Alias:   ":stars:",
	},
	"\U0001F321\uFE0F  :thermometer:": {
		Name:    "Thermometer",
		Unicode: "\U0001F321\uFE0F",
		Alias:   ":thermometer:",
	},
	"\U0001F324\uFE0F  :mostly_sunny:": {
		Name:    "Sun Behind Small Cloud",
		Unicode: "\U0001F324\uFE0F",
		Alias:   ":mostly_sunny:",
	},
	"\U0001F325\uFE0F  :barely_sunny:": {
		Name:    "Sun Behind Large Cloud",
		Unicode: "\U0001F325\uFE0F",
		Alias:   ":barely_sunny:",
	},
	"\U0001F326\uFE0F  :partly_sunny_rain:": {
		Name:    "Sun Behind Rain Cloud",
		Unicode: "\U0001F326\uFE0F",
		Alias:   ":partly_sunny_rain:",
	},
	"\U0001F327\uFE0F  :rain_cloud:": {
		Name:    "Cloud With Rain",
		Unicode: "\U0001F327\uFE0F",
		Alias:   ":rain_cloud:",
	},
	"\U0001F328\uFE0F  :snow_cloud:": {
		Name:    "Cloud With Snow",
		Unicode: "\U0001F328\uFE0F",
		Alias:   ":snow_cloud:",
	},
	"\U0001F329\uFE0F  :lightning:": {
		Name:    "Cloud With Lightning",
		Unicode: "\U0001F329\uFE0F",
		Alias:   ":lightning:",
	},
	"\U0001F32A\uFE0F  :tornado:": {
		Name:    "Tornado",
		Unicode: "\U0001F32A\uFE0F",
		Alias:   ":tornado:",
	},
	"\U0001F32B\uFE0F  :fog:": {
		Name:    "Fog",
		Unicode: "\U0001F32B\uFE0F",
		Alias:   ":fog:",
	},
	"\U0001F32C\uFE0F  :wind_blowing_face:": {
		Name:    "Wind Face",
		Unicode: "\U0001F32C\uFE0F",
		Alias:   ":wind_blowing_face:",
	},
	"\U0001F32D  :hotdog:": {
		Name:    "Hot Dog",
		Unicode: "\U0001F32D",
		Alias:   ":hotdog:",
	},
	"\U0001F32E  :taco:": {
		Name:    "Taco",
		Unicode: "\U0001F32E",
		Alias:   ":taco:",
	},
	"\U0001F32F  :burrito:": {
		Name:    "Burrito",
		Unicode: "\U0001F32F",
		Alias:   ":burrito:",
	},
	"\U0001F330  :chestnut:": {
		Name:    "Chestnut",
		Unicode: "\U0001F330",
		Alias:   ":chestnut:",
	},
	"\U0001F331  :seedling:": {
		Name:    "Seedling",
		Unicode: "\U0001F331",
		Alias:   ":seedling:",
	},
	"\U0001F332  :evergreen_tree:": {
		Name:    "Evergreen Tree",
		Unicode: "\U0001F332",
		Alias:   ":evergreen_tree:",
	},
	"\U0001F333  :deciduous_tree:": {
		Name:    "Deciduous Tree",
		Unicode: "\U0001F333",
		Alias:   ":deciduous_tree:",
	},
	"\U0001F334  :palm_tree:": {
		Name:    "Palm Tree",
		Unicode: "\U0001F334",
		Alias:   ":palm_tree:",
	},
	"\U0001F335  :cactus:": {
		Name:    "Cactus",
		Unicode: "\U0001F335",
		Alias:   ":cactus:",
	},
	"\U0001F336\uFE0F  :hot_pepper:": {
		Name:    "Hot Pepper",
		Unicode: "\U0001F336\uFE0F",
		Alias:   ":hot_pepper:",
	},
	"\U0001F337  :tulip:": {
		Name:    "Tulip",
		Unicode: "\U0001F337",
		Alias:   ":tulip:",
	},
	"\U0001F338  :cherry_blossom:": {
		Name:    "Cherry Blossom",
		Unicode: "\U0001F338",
		Alias:   ":cherry_blossom:",
	},
	"\U0001F339  :rose:": {
		Name:    "Rose",
		Unicode: "\U0001F339",
		Alias:   ":rose:",
	},
	"\U0001F33A  :hibiscus:": {
		Name:    "Hibiscus",
		Unicode: "\U0001F33A",
		Alias:   ":hibiscus:",
	},
	"\U0001F33B  :sunflower:": {
		Name:    "Sunflower",
		Unicode: "\U0001F33B",
		Alias:   ":sunflower:",
	},
	"\U0001F33C  :blossom:": {
		Name:    "Blossom",
		Unicode: "\U0001F33C",
		Alias:   ":blossom:",
	},
	"\U0001F33D  :corn:": {
		Name:    "Ear Of Maize",
		Unicode: "\U0001F33D",
		Alias:   ":corn:",
	},
	"\U0001F33E  :ear_of_rice:": {
		Name:    "Ear Of Rice",
		Unicode: "\U0001F33E",
		Alias:   ":ear_of_rice:",
	},
	"\U0001F33F  :herb:": {
		Name:    "Herb",
		Unicode: "\U0001F33F",
		Alias:   ":herb:",
	},
	"\U0001F340  :four_leaf_clover:": {
		Name:    "Four Leaf Clover",
		Unicode: "\U0001F340",
		Alias:   ":four_leaf_clover:",
	},
	"\U0001F341  :maple_leaf:": {
		Name:    "Maple Leaf",
		Unicode: "\U0001F341",
		Alias:   ":maple_leaf:",
	},
	"\U0001F342  :fallen_leaf:": {
		Name:    "Fallen Leaf",
		Unicode: "\U0001F342",
		Alias:   ":fallen_leaf:",
	},
	"\U0001F343  :leaves:": {
		Name:    "Leaf Fluttering In Wind",
		Unicode: "\U0001F343",
		Alias:   ":leaves:",
	},
	"\U0001F344  :mushroom:": {
		Name:    "Mushroom",
		Unicode: "\U0001F344",
		Alias:   ":mushroom:",
	},
	"\U0001F345  :tomato:": {
		Name:    "Tomato",
		Unicode: "\U0001F345",
		Alias:   ":tomato:",
	},
	"\U0001F346  :eggplant:": {
		Name:    "Aubergine",
		Unicode: "\U0001F346",
		Alias:   ":eggplant:",
	},
	"\U0001F347  :grapes:": {
		Name:    "Grapes",
		Unicode: "\U0001F347",
		Alias:   ":grapes:",
	},
	"\U0001F348  :melon:": {
		Name:    "Melon",
		Unicode: "\U0001F348",
		Alias:   ":melon:",
	},
	"\U0001F349  :watermelon:": {
		Name:    "Watermelon",
		Unicode: "\U0001F349",
		Alias:   ":watermelon:",
	},
	"\U0001F34A  :tangerine:": {
		Name:    "Tangerine",
		Unicode: "\U0001F34A",
		Alias:   ":tangerine:",
	},
	"\U0001F34B  :lemon:": {
		Name:    "Lemon",
		Unicode: "\U0001F34B",
		Alias:   ":lemon:",
	},
	"\U0001F34C  :banana:": {
		Name:    "Banana",
		Unicode: "\U0001F34C",
		Alias:   ":banana:",
	},
	"\U0001F34D  :pineapple:": {
		Name:    "Pineapple",
		Unicode: "\U0001F34D",
		Alias:   ":pineapple:",
	},
	"\U0001F34E  :apple:": {
		Name:    "Red Apple",
		Unicode: "\U0001F34E",
		Alias:   ":apple:",
	},
	"\U0001F34F  :green_apple:": {
		Name:    "Green Apple",
		Unicode: "\U0001F34F",
		Alias:   ":green_apple:",
	},
	"\U0001F350  :pear:": {
		Name:    "Pear",
		Unicode: "\U0001F350",
		Alias:   ":pear:",
	},
	"\U0001F351  :peach:": {
		Name:    "Peach",
		Unicode: "\U0001F351",
		Alias:   ":peach:",
	},
	"\U0001F352  :cherries:": {
		Name:    "Cherries",
		Unicode: "\U0001F352",
		Alias:   ":cherries:",
	},
	"\U0001F353  :strawberry:": {
		Name:    "Strawberry",
		Unicode: "\U0001F353",
		Alias:   ":strawberry:",
	},
	"\U0001F354  :hamburger:": {
		Name:    "Hamburger",
		Unicode: "\U0001F354",
		Alias:   ":hamburger:",
	},
	"\U0001F355  :pizza:": {
		Name:    "Slice Of Pizza",
		Unicode: "\U0001F355",
		Alias:   ":pizza:",
	},
	"\U0001F356  :meat_on_bone:": {
		Name:    "Meat On Bone",
		Unicode: "\U0001F356",
		Alias:   ":meat_on_bone:",
	},
	"\U0001F357  :poultry_leg:": {
		Name:    "Poultry Leg",
		Unicode: "\U0001F357",
		Alias:   ":poultry_leg:",
	},
	"\U0001F358  :rice_cracker:": {
		Name:    "Rice Cracker",
		Unicode: "\U0001F358",
		Alias:   ":rice_cracker:",
	},
	"\U0001F359  :rice_ball:": {
		Name:    "Rice Ball",
		Unicode: "\U0001F359",
		Alias:   ":rice_ball:",
	},
	"\U0001F35A  :rice:": {
		Name:    "Cooked Rice",
		Unicode: "\U0001F35A",
		Alias:   ":rice:",
	},
	"\U0001F35B  :curry:": {
		Name:    "Curry And Rice",
		Unicode: "\U0001F35B",
		Alias:   ":curry:",
	},
	"\U0001F35C  :ramen:": {
		Name:    "Steaming Bowl",
		Unicode: "\U0001F35C",
		Alias:   ":ramen:",
	},
	"\U0001F35D  :spaghetti:": {
		Name:    "Spaghetti",
		Unicode: "\U0001F35D",
		Alias:   ":spaghetti:",
	},
	"\U0001F35E  :bread:": {
		Name:    "Bread",
		Unicode: "\U0001F35E",
		Alias:   ":bread:",
	},
	"\U0001F35F  :fries:": {
		Name:    "French Fries",
		Unicode: "\U0001F35F",
		Alias:   ":fries:",
	},
	"\U0001F360  :sweet_potato:": {
		Name:    "Roasted Sweet Potato",
		Unicode: "\U0001F360",
		Alias:   ":sweet_potato:",
	},
	"\U0001F361  :dango:": {
		Name:    "Dango",
		Unicode: "\U0001F361",
		Alias:   ":dango:",
	},
	"\U0001F362  :oden:": {
		Name:    "Oden",
		Unicode: "\U0001F362",
		Alias:   ":oden:",
	},
	"\U0001F363  :sushi:": {
		Name:    "Sushi",
		Unicode: "\U0001F363",
		Alias:   ":sushi:",
	},
	"\U0001F364  :fried_shrimp:": {
		Name:    "Fried Shrimp",
		Unicode: "\U0001F364",
		Alias:   ":fried_shrimp:",
	},
	"\U0001F365  :fish_cake:": {
		Name:    "Fish Cake With Swirl Design",
		Unicode: "\U0001F365",
		Alias:   ":fish_cake:",
	},
	"\U0001F366  :icecream:": {
		Name:    "Soft Ice Cream",
		Unicode: "\U0001F366",
		Alias:   ":icecream:",
	},
	"\U0001F367  :shaved_ice:": {
		Name:    "Shaved Ice",
		Unicode: "\U0001F367",
		Alias:   ":shaved_ice:",
	},
	"\U0001F368  :ice_cream:": {
		Name:    "Ice Cream",
		Unicode: "\U0001F368",
		Alias:   ":ice_cream:",
	},
	"\U0001F369  :doughnut:": {
		Name:    "Doughnut",
		Unicode: "\U0001F369",
		Alias:   ":doughnut:",
	},
	"\U0001F36A  :cookie:": {
		Name:    "Cookie",
		Unicode: "\U0001F36A",
		Alias:   ":cookie:",
	},
	"\U0001F36B  :chocolate_bar:": {
		Name:    "Chocolate Bar",
		Unicode: "\U0001F36B",
		Alias:   ":chocolate_bar:",
	},
	"\U0001F36C  :candy:": {
		Name:    "Candy",
		Unicode: "\U0001F36C",
		Alias:   ":candy:",
	},
	"\U0001F36D  :lollipop:": {
		Name:    "Lollipop",
		Unicode: "\U0001F36D",
		Alias:   ":lollipop:",
	},
	"\U0001F36E  :custard:": {
		Name:    "Custard",
		Unicode: "\U0001F36E",
		Alias:   ":custard:",
	},
	"\U0001F36F  :honey_pot:": {
		Name:    "Honey Pot",
		Unicode: "\U0001F36F",
		Alias:   ":honey_pot:",
	},
	"\U0001F370  :cake:": {
		Name:    "Shortcake",
		Unicode: "\U0001F370",
		Alias:   ":cake:",
	},
	"\U0001F371  :bento:": {
		Name:    "Bento Box",
		Unicode: "\U0001F371",
		Alias:   ":bento:",
	},
	"\U0001F372  :stew:": {
		Name:    "Pot Of Food",
		Unicode: "\U0001F372",
		Alias:   ":stew:",
	},
	"\U0001F373  :fried_egg:": {
		Name:    "Cooking",
		Unicode: "\U0001F373",
		Alias:   ":fried_egg:",
	},
	"\U0001F374  :fork_and_knife:": {
		Name:    "Fork And Knife",
		Unicode: "\U0001F374",
		Alias:   ":fork_and_knife:",
	},
	"\U0001F375  :tea:": {
		Name:    "Teacup Without Handle",
		Unicode: "\U0001F375",
		Alias:   ":tea:",
	},
	"\U0001F376  :sake:": {
		Name:    "Sake Bottle And Cup",
		Unicode: "\U0001F376",
		Alias:   ":sake:",
	},
	"\U0001F377  :wine_glass:": {
		Name:    "Wine Glass",
		Unicode: "\U0001F377",
		Alias:   ":wine_glass:",
	},
	"\U0001F378  :cocktail:": {
		Name:    "Cocktail Glass",
		Unicode: "\U0001F378",
		Alias:   ":cocktail:",
	},
	"\U0001F379  :tropical_drink:": {
		Name:    "Tropical Drink",
		Unicode: "\U0001F379",
		Alias:   ":tropical_drink:",
	},
	"\U0001F37A  :beer:": {
		Name:    "Beer Mug",
		Unicode: "\U0001F37A",
		Alias:   ":beer:",
	},
	"\U0001F37B  :beers:": {
		Name:    "Clinking Beer Mugs",
		Unicode: "\U0001F37B",
		Alias:   ":beers:",
	},
	"\U0001F37C  :baby_bottle:": {
		Name:    "Baby Bottle",
		Unicode: "\U0001F37C",
		Alias:   ":baby_bottle:",
	},
	"\U0001F37D\uFE0F  :knife_fork_plate:": {
		Name:    "Fork And Knife With Plate",
		Unicode: "\U0001F37D\uFE0F",
		Alias:   ":knife_fork_plate:",
	},
	"\U0001F37E  :champagne:": {
		Name:    "Bottle With Popping Cork",
		Unicode: "\U0001F37E",
		Alias:   ":champagne:",
	},
	"\U0001F37F  :popcorn:": {
		Name:    "Popcorn",
		Unicode: "\U0001F37F",
		Alias:   ":popcorn:",
	},
	"\U0001F380  :ribbon:": {
		Name:    "Ribbon",
		Unicode: "\U0001F380",
		Alias:   ":ribbon:",
	},
	"\U0001F381  :gift:": {
		Name:    "Wrapped Present",
		Unicode: "\U0001F381",
		Alias:   ":gift:",
	},
	"\U0001F382  :birthday:": {
		Name:    "Birthday Cake",
		Unicode: "\U0001F382",
		Alias:   ":birthday:",
	},
	"\U0001F383  :jack_o_lantern:": {
		Name:    "Jack-O-Lantern",
		Unicode: "\U0001F383",
		Alias:   ":jack_o_lantern:",
	},
	"\U0001F384  :christmas_tree:": {
		Name:    "Christmas Tree",
		Unicode: "\U0001F384",
		Alias:   ":christmas_tree:",
	},
	"\U0001F385  :santa:": {
		Name:    "Father Christmas",
		Unicode: "\U0001F385",
		Alias:   ":santa:",
	},
	"\U0001F386  :fireworks:": {
		Name:    "Fireworks",
		Unicode: "\U0001F386",
		Alias:   ":fireworks:",
	},
	"\U0001F387  :sparkler:": {
		Name:    "Firework Sparkler",
		Unicode: "\U0001F387",
		Alias:   ":sparkler:",
	},
	"\U0001F388  :balloon:": {
		Name:    "Balloon",
		Unicode: "\U0001F388",
		Alias:   ":balloon:",
	},
	"\U0001F389  :tada:": {
		Name:    "Party Popper",
		Unicode: "\U0001F389",
		Alias:   ":tada:",
	},
	"\U0001F38A  :confetti_ball:": {
		Name:    "Confetti Ball",
		Unicode: "\U0001F38A",
		Alias:   ":confetti_ball:",
	},
	"\U0001F38B  :tanabata_tree:": {
		Name:    "Tanabata Tree",
		Unicode: "\U0001F38B",
		Alias:   ":tanabata_tree:",
	},
	"\U0001F38C  :crossed_flags:": {
		Name:    "Crossed Flags",
		Unicode: "\U0001F38C",
		Alias:   ":crossed_flags:",
	},
	"\U0001F38D  :bamboo:": {
		Name:    "Pine Decoration",
		Unicode: "\U0001F38D",
		Alias:   ":bamboo:",
	},
	"\U0001F38E  :dolls:": {
		Name:    "Japanese Dolls",
		Unicode: "\U0001F38E",
		Alias:   ":dolls:",
	},
	"\U0001F38F  :flags:": {
		Name:    "Carp Streamer",
		Unicode: "\U0001F38F",
		Alias:   ":flags:",
	},
	"\U0001F390  :wind_chime:": {
		Name:    "Wind Chime",
		Unicode: "\U0001F390",
		Alias:   ":wind_chime:",
	},
	"\U0001F391  :rice_scene:": {
		Name:    "Moon Viewing Ceremony",
		Unicode: "\U0001F391",
		Alias:   ":rice_scene:",
	},
	"\U0001F392  :school_satchel:": {
		Name:    "School Satchel",
		Unicode: "\U0001F392",
		Alias:   ":school_satchel:",
	},
	"\U0001F393  :mortar_board:": {
		Name:    "Graduation Cap",
		Unicode: "\U0001F393",
		Alias:   ":mortar_board:",
	},
	"\U0001F396\uFE0F  :medal:": {
		Name:    "Military Medal",
		Unicode: "\U0001F396\uFE0F",
		Alias:   ":medal:",
	},
	"\U0001F397\uFE0F  :reminder_ribbon:": {
		Name:    "Reminder Ribbon",
		Unicode: "\U0001F397\uFE0F",
		Alias:   ":reminder_ribbon:",
	},
	"\U0001F399\uFE0F  :studio_microphone:": {
		Name:    "Studio Microphone",
		Unicode: "\U0001F399\uFE0F",
		Alias:   ":studio_microphone:",
	},
	"\U0001F39A\uFE0F  :level_slider:": {
		Name:    "Level Slider",
		Unicode: "\U0001F39A\uFE0F",
		Alias:   ":level_slider:",
	},
	"\U0001F39B\uFE0F  :control_knobs:": {
		Name:    "Control Knobs",
		Unicode: "\U0001F39B\uFE0F",
		Alias:   ":control_knobs:",
	},
	"\U0001F39E\uFE0F  :film_frames:": {
		Name:    "Film Frames",
		Unicode: "\U0001F39E\uFE0F",
		Alias:   ":film_frames:",
	},
	"\U0001F39F\uFE0F  :admission_tickets:": {
		Name:    "Admission Tickets",
		Unicode: "\U0001F39F\uFE0F",
		Alias:   ":admission_tickets:",
	},
	"\U0001F3A0  :carousel_horse:": {
		Name:    "Carousel Horse",
		Unicode: "\U0001F3A0",
		Alias:   ":carousel_horse:",
	},
	"\U0001F3A1  :ferris_wheel:": {
		Name:    "Ferris Wheel",
		Unicode: "\U0001F3A1",
		Alias:   ":ferris_wheel:",
	},
	"\U0001F3A2  :roller_coaster:": {
		Name:    "Roller Coaster",
		Unicode: "\U0001F3A2",
		Alias:   ":roller_coaster:",
	},
	"\U0001F3A3  :fishing_pole_and_fish:": {
		Name:    "Fishing Pole And Fish",
		Unicode: "\U0001F3A3",
		Alias:   ":fishing_pole_and_fish:",
	},
	"\U0001F3A4  :microphone:": {
		Name:    "Microphone",
		Unicode: "\U0001F3A4",
		Alias:   ":microphone:",
	},
	"\U0001F3A5  :movie_camera:": {
		Name:    "Movie Camera",
		Unicode: "\U0001F3A5",
		Alias:   ":movie_camera:",
	},
	"\U0001F3A6  :cinema:": {
		Name:    "Cinema",
		Unicode: "\U0001F3A6",
		Alias:   ":cinema:",
	},
	"\U0001F3A7  :headphones:": {
		Name:    "Headphone",
		Unicode: "\U0001F3A7",
		Alias:   ":headphones:",
	},
	"\U0001F3A8  :art:": {
		Name:    "Artist Palette",
		Unicode: "\U0001F3A8",
		Alias:   ":art:",
	},
	"\U0001F3A9  :tophat:": {
		Name:    "Top Hat",
		Unicode: "\U0001F3A9",
		Alias:   ":tophat:",
	},
	"\U0001F3AA  :circus_tent:": {
		Name:    "Circus Tent",
		Unicode: "\U0001F3AA",
		Alias:   ":circus_tent:",
	},
	"\U0001F3AB  :ticket:": {
		Name:    "Ticket",
		Unicode: "\U0001F3AB",
		Alias:   ":ticket:",
	},
	"\U0001F3AC  :clapper:": {
		Name:    "Clapper Board",
		Unicode: "\U0001F3AC",
		Alias:   ":clapper:",
	},
	"\U0001F3AD  :performing_arts:": {
		Name:    "Performing Arts",
		Unicode: "\U0001F3AD",
		Alias:   ":performing_arts:",
	},
	"\U0001F3AE  :video_game:": {
		Name:    "Video Game",
		Unicode: "\U0001F3AE",
		Alias:   ":video_game:",
	},
	"\U0001F3AF  :dart:": {
		Name:    "Direct Hit",
		Unicode: "\U0001F3AF",
		Alias:   ":dart:",
	},
	"\U0001F3B0  :slot_machine:": {
		Name:    "Slot Machine",
		Unicode: "\U0001F3B0",
		Alias:   ":slot_machine:",
	},
	"\U0001F3B1  :8ball:": {
		Name:    "Billiards",
		Unicode: "\U0001F3B1",
		Alias:   ":8ball:",
	},
	"\U0001F3B2  :game_die:": {
		Name:    "Game Die",
		Unicode: "\U0001F3B2",
		Alias:   ":game_die:",
	},
	"\U0001F3B3  :bowling:": {
		Name:    "Bowling",
		Unicode: "\U0001F3B3",
		Alias:   ":bowling:",
	},
	"\U0001F3B4  :flower_playing_cards:": {
		Name:    "Flower Playing Cards",
		Unicode: "\U0001F3B4",
		Alias:   ":flower_playing_cards:",
	},
	"\U0001F3B5  :musical_note:": {
		Name:    "Musical Note",
		Unicode: "\U0001F3B5",
		Alias:   ":musical_note:",
	},
	"\U0001F3B6  :notes:": {
		Name:    "Multiple Musical Notes",
		Unicode: "\U0001F3B6",
		Alias:   ":notes:",
	},
	"\U0001F3B7  :saxophone:": {
		Name:    "Saxophone",
		Unicode: "\U0001F3B7",
		Alias:   ":saxophone:",
	},
	"\U0001F3B8  :guitar:": {
		Name:    "Guitar",
		Unicode: "\U0001F3B8",
		Alias:   ":guitar:",
	},
	"\U0001F3B9  :musical_keyboard:": {
		Name:    "Musical Keyboard",
		Unicode: "\U0001F3B9",
		Alias:   ":musical_keyboard:",
	},
	"\U0001F3BA  :trumpet:": {
		Name:    "Trumpet",
		Unicode: "\U0001F3BA",
		Alias:   ":trumpet:",
	},
	"\U0001F3BB  :violin:": {
		Name:    "Violin",
		Unicode: "\U0001F3BB",
		Alias:   ":violin:",
	},
	"\U0001F3BC  :musical_score:": {
		Name:    "Musical Score",
		Unicode: "\U0001F3BC",
		Alias:   ":musical_score:",
	},
	"\U0001F3BD  :running_shirt_with_sash:": {
		Name:    "Running Shirt With Sash",
		Unicode: "\U0001F3BD",
		Alias:   ":running_shirt_with_sash:",
	},
	"\U0001F3BE  :tennis:": {
		Name:    "Tennis Racquet And Ball",
		Unicode: "\U0001F3BE",
		Alias:   ":tennis:",
	},
	"\U0001F3BF  :ski:": {
		Name:    "Ski And Ski Boot",
		Unicode: "\U0001F3BF",
		Alias:   ":ski:",
	},
	"\U0001F3C0  :basketball:": {
		Name:    "Basketball And Hoop",
		Unicode: "\U0001F3C0",
		Alias:   ":basketball:",
	},
	"\U0001F3C1  :checkered_flag:": {
		Name:    "Chequered Flag",
		Unicode: "\U0001F3C1",
		Alias:   ":checkered_flag:",
	},
	"\U0001F3C2  :snowboarder:": {
		Name:    "Snowboarder",
		Unicode: "\U0001F3C2",
		Alias:   ":snowboarder:",
	},
	"\U0001F3C3\u200D\u2640\uFE0F :woman-running:": {
		Name:    "Woman Running",
		Unicode: "\U0001F3C3\u200D\u2640\uFE0F",
		Alias:   ":woman-running:",
	},
	"\U0001F3C3\u200D\u2642\uFE0F :man-running:": {
		Name:    "Man Running",
		Unicode: "\U0001F3C3\u200D\u2642\uFE0F",
		Alias:   ":man-running:",
	},
	"\U0001F3C3  :runner:": {
		Name:    "Runner",
		Unicode: "\U0001F3C3",
		Alias:   ":runner:",
	},
	"\U0001F3C4\u200D\u2640\uFE0F :woman-surfing:": {
		Name:    "Woman Surfing",
		Unicode: "\U0001F3C4\u200D\u2640\uFE0F",
		Alias:   ":woman-surfing:",
	},
	"\U0001F3C4\u200D\u2642\uFE0F :man-surfing:": {
		Name:    "Man Surfing",
		Unicode: "\U0001F3C4\u200D\u2642\uFE0F",
		Alias:   ":man-surfing:",
	},
	"\U0001F3C4  :surfer:": {
		Name:    "Surfer",
		Unicode: "\U0001F3C4",
		Alias:   ":surfer:",
	},
	"\U0001F3C5  :sports_medal:": {
		Name:    "Sports Medal",
		Unicode: "\U0001F3C5",
		Alias:   ":sports_medal:",
	},
	"\U0001F3C6  :trophy:": {
		Name:    "Trophy",
		Unicode: "\U0001F3C6",
		Alias:   ":trophy:",
	},
	"\U0001F3C7  :horse_racing:": {
		Name:    "Horse Racing",
		Unicode: "\U0001F3C7",
		Alias:   ":horse_racing:",
	},
	"\U0001F3C8  :football:": {
		Name:    "American Football",
		Unicode: "\U0001F3C8",
		Alias:   ":football:",
	},
	"\U0001F3C9  :rugby_football:": {
		Name:    "Rugby Football",
		Unicode: "\U0001F3C9",
		Alias:   ":rugby_football:",
	},
	"\U0001F3CA\u200D\u2640\uFE0F :woman-swimming:": {
		Name:    "Woman Swimming",
		Unicode: "\U0001F3CA\u200D\u2640\uFE0F",
		Alias:   ":woman-swimming:",
	},
	"\U0001F3CA\u200D\u2642\uFE0F :man-swimming:": {
		Name:    "Man Swimming",
		Unicode: "\U0001F3CA\u200D\u2642\uFE0F",
		Alias:   ":man-swimming:",
	},
	"\U0001F3CA  :swimmer:": {
		Name:    "Swimmer",
		Unicode: "\U0001F3CA",
		Alias:   ":swimmer:",
	},
	"\U0001F3CB\uFE0F\u200D\u2640\uFE0F :woman-lifting-weights:": {
		Name:    "Woman Lifting Weights",
		Unicode: "\U0001F3CB\uFE0F\u200D\u2640\uFE0F",
		Alias:   ":woman-lifting-weights:",
	},
	"\U0001F3CB\uFE0F\u200D\u2642\uFE0F :man-lifting-weights:": {
		Name:    "Man Lifting Weights",
		Unicode: "\U0001F3CB\uFE0F\u200D\u2642\uFE0F",
		Alias:   ":man-lifting-weights:",
	},
	"\U0001F3CB\uFE0F  :weight_lifter:": {
		Name:    "Person Lifting Weights",
		Unicode: "\U0001F3CB\uFE0F",
		Alias:   ":weight_lifter:",
	},
	"\U0001F3CC\uFE0F\u200D\u2640\uFE0F :woman-golfing:": {
		Name:    "Woman Golfing",
		Unicode: "\U0001F3CC\uFE0F\u200D\u2640\uFE0F",
		Alias:   ":woman-golfing:",
	},
	"\U0001F3CC\uFE0F\u200D\u2642\uFE0F :man-golfing:": {
		Name:    "Man Golfing",
		Unicode: "\U0001F3CC\uFE0F\u200D\u2642\uFE0F",
		Alias:   ":man-golfing:",
	},
	"\U0001F3CC\uFE0F  :golfer:": {
		Name:    "Person Golfing",
		Unicode: "\U0001F3CC\uFE0F",
		Alias:   ":golfer:",
	},
	"\U0001F3CD\uFE0F  :racing_motorcycle:": {
		Name:    "Motorcycle",
		Unicode: "\U0001F3CD\uFE0F",
		Alias:   ":racing_motorcycle:",
	},
	"\U0001F3CE\uFE0F  :racing_car:": {
		Name:    "Racing Car",
		Unicode: "\U0001F3CE\uFE0F",
		Alias:   ":racing_car:",
	},
	"\U0001F3CF  :cricket_bat_and_ball:": {
		Name:    "Cricket Bat And Ball",
		Unicode: "\U0001F3CF",
		Alias:   ":cricket_bat_and_ball:",
	},
	"\U0001F3D0  :volleyball:": {
		Name:    "Volleyball",
		Unicode: "\U0001F3D0",
		Alias:   ":volleyball:",
	},
	"\U0001F3D1  :field_hockey_stick_and_ball:": {
		Name:    "Field Hockey Stick And Ball",
		Unicode: "\U0001F3D1",
		Alias:   ":field_hockey_stick_and_ball:",
	},
	"\U0001F3D2  :ice_hockey_stick_and_puck:": {
		Name:    "Ice Hockey Stick And Puck",
		Unicode: "\U0001F3D2",
		Alias:   ":ice_hockey_stick_and_puck:",
	},
	"\U0001F3D3  :table_tennis_paddle_and_ball:": {
		Name:    "Table Tennis Paddle And Ball",
		Unicode: "\U0001F3D3",
		Alias:   ":table_tennis_paddle_and_ball:",
	},
	"\U0001F3D4\uFE0F  :snow_capped_mountain:": {
		Name:    "Snow-Capped Mountain",
		Unicode: "\U0001F3D4\uFE0F",
		Alias:   ":snow_capped_mountain:",
	},
	"\U0001F3D5\uFE0F  :camping:": {
		Name:    "Camping",
		Unicode: "\U0001F3D5\uFE0F",
		Alias:   ":camping:",
	},
	"\U0001F3D6\uFE0F  :beach_with_umbrella:": {
		Name:    "Beach With Umbrella",
		Unicode: "\U0001F3D6\uFE0F",
		Alias:   ":beach_with_umbrella:",
	},
	"\U0001F3D7\uFE0F  :building_construction:": {
		Name:    "Building Construction",
		Unicode: "\U0001F3D7\uFE0F",
		Alias:   ":building_construction:",
	},
	"\U0001F3D8\uFE0F  :house_buildings:": {
		Name:    "Houses",
		Unicode: "\U0001F3D8\uFE0F",
		Alias:   ":house_buildings:",
	},
	"\U0001F3D9\uFE0F  :cityscape:": {
		Name:    "Cityscape",
		Unicode: "\U0001F3D9\uFE0F",
		Alias:   ":cityscape:",
	},
	"\U0001F3DA\uFE0F  :derelict_house_building:": {
		Name:    "Derelict House",
		Unicode: "\U0001F3DA\uFE0F",
		Alias:   ":derelict_house_building:",
	},
	"\U0001F3DB\uFE0F  :classical_building:": {
		Name:    "Classical Building",
		Unicode: "\U0001F3DB\uFE0F",
		Alias:   ":classical_building:",
	},
	"\U0001F3DC\uFE0F  :desert:": {
		Name:    "Desert",
		Unicode: "\U0001F3DC\uFE0F",
		Alias:   ":desert:",
	},
	"\U0001F3DD\uFE0F  :desert_island:": {
		Name:    "Desert Island",
		Unicode: "\U0001F3DD\uFE0F",
		Alias:   ":desert_island:",
	},
	"\U0001F3DE\uFE0F  :national_park:": {
		Name:    "National Park",
		Unicode: "\U0001F3DE\uFE0F",
		Alias:   ":national_park:",
	},
	"\U0001F3DF\uFE0F  :stadium:": {
		Name:    "Stadium",
		Unicode: "\U0001F3DF\uFE0F",
		Alias:   ":stadium:",
	},
	"\U0001F3E0  :house:": {
		Name:    "House Building",
		Unicode: "\U0001F3E0",
		Alias:   ":house:",
	},
	"\U0001F3E1  :house_with_garden:": {
		Name:    "House With Garden",
		Unicode: "\U0001F3E1",
		Alias:   ":house_with_garden:",
	},
	"\U0001F3E2  :office:": {
		Name:    "Office Building",
		Unicode: "\U0001F3E2",
		Alias:   ":office:",
	},
	"\U0001F3E3  :post_office:": {
		Name:    "Japanese Post Office",
		Unicode: "\U0001F3E3",
		Alias:   ":post_office:",
	},
	"\U0001F3E4  :european_post_office:": {
		Name:    "European Post Office",
		Unicode: "\U0001F3E4",
		Alias:   ":european_post_office:",
	},
	"\U0001F3E5  :hospital:": {
		Name:    "Hospital",
		Unicode: "\U0001F3E5",
		Alias:   ":hospital:",
	},
	"\U0001F3E6  :bank:": {
		Name:    "Bank",
		Unicode: "\U0001F3E6",
		Alias:   ":bank:",
	},
	"\U0001F3E7  :atm:": {
		Name:    "Automated Teller Machine",
		Unicode: "\U0001F3E7",
		Alias:   ":atm:",
	},
	"\U0001F3E8  :hotel:": {
		Name:    "Hotel",
		Unicode: "\U0001F3E8",
		Alias:   ":hotel:",
	},
	"\U0001F3E9  :love_hotel:": {
		Name:    "Love Hotel",
		Unicode: "\U0001F3E9",
		Alias:   ":love_hotel:",
	},
	"\U0001F3EA  :convenience_store:": {
		Name:    "Convenience Store",
		Unicode: "\U0001F3EA",
		Alias:   ":convenience_store:",
	},
	"\U0001F3EB  :school:": {
		Name:    "School",
		Unicode: "\U0001F3EB",
		Alias:   ":school:",
	},
	"\U0001F3EC  :department_store:": {
		Name:    "Department Store",
		Unicode: "\U0001F3EC",
		Alias:   ":department_store:",
	},
	"\U0001F3ED  :factory:": {
		Name:    "Factory",
		Unicode: "\U0001F3ED",
		Alias:   ":factory:",
	},
	"\U0001F3EE  :izakaya_lantern:": {
		Name:    "Izakaya Lantern",
		Unicode: "\U0001F3EE",
		Alias:   ":izakaya_lantern:",
	},
	"\U0001F3EF  :japanese_castle:": {
		Name:    "Japanese Castle",
		Unicode: "\U0001F3EF",
		Alias:   ":japanese_castle:",
	},
	"\U0001F3F0  :european_castle:": {
		Name:    "European Castle",
		Unicode: "\U0001F3F0",
		Alias:   ":european_castle:",
	},
	"\U0001F3F3\uFE0F\u200D\u1F308 :rainbow-flag:": {
		Name:    "Rainbow Flag",
		Unicode: "\U0001F3F3\uFE0F\u200D\u1F308",
		Alias:   ":rainbow-flag:",
	},
	"\U0001F3F3\uFE0F\u200D\u26A7\uFE0F  :transgender_flag:": {
		Name:    "Transgender Flag",
		Unicode: "\U0001F3F3\uFE0F\u200D\u26A7\uFE0F",
		Alias:   ":transgender_flag:",
	},
	"\U0001F3F3\uFE0F  :waving_white_flag:": {
		Name:    "White Flag",
		Unicode: "\U0001F3F3\uFE0F",
		Alias:   ":waving_white_flag:",
	},
	"\U0001F3F4\u200D\u2620\uFE0F  :pirate_flag:": {
		Name:    "Pirate Flag",
		Unicode: "\U0001F3F4\u200D\u2620\uFE0F",
		Alias:   ":pirate_flag:",
	},
	"\U0001F3F4\uE0067\uE0062\uE0065\uE006E\uE0067\uE007F :flag-england:": {
		Name:    "England Flag",
		Unicode: "\U0001F3F4\uE0067\uE0062\uE0065\uE006E\uE0067\uE007F",
		Alias:   ":flag-england:",
	},
	"\U0001F3F4\uE0067\uE0062\uE0073\uE0063\uE0074\uE007F :flag-scotland:": {
		Name:    "Scotland Flag",
		Unicode: "\U0001F3F4\uE0067\uE0062\uE0073\uE0063\uE0074\uE007F",
		Alias:   ":flag-scotland:",
	},
	"\U0001F3F4\uE0067\uE0062\uE0077\uE006C\uE0073\uE007F :flag-wales:": {
		Name:    "Wales Flag",
		Unicode: "\U0001F3F4\uE0067\uE0062\uE0077\uE006C\uE0073\uE007F",
		Alias:   ":flag-wales:",
	},
	"\U0001F3F4  :waving_black_flag:": {
		Name:    "Waving Black Flag",
		Unicode: "\U0001F3F4",
		Alias:   ":waving_black_flag:",
	},
	"\U0001F3F5\uFE0F  :rosette:": {
		Name:    "Rosette",
		Unicode: "\U0001F3F5\uFE0F",
		Alias:   ":rosette:",
	},
	"\U0001F3F7\uFE0F  :label:": {
		Name:    "Label",
		Unicode: "\U0001F3F7\uFE0F",
		Alias:   ":label:",
	},
	"\U0001F3F8  :badminton_racquet_and_shuttlecock:": {
		Name:    "Badminton Racquet And Shuttlecock",
		Unicode: "\U0001F3F8",
		Alias:   ":badminton_racquet_and_shuttlecock:",
	},
	"\U0001F3F9  :bow_and_arrow:": {
		Name:    "Bow And Arrow",
		Unicode: "\U0001F3F9",
		Alias:   ":bow_and_arrow:",
	},
	"\U0001F3FA  :amphora:": {
		Name:    "Amphora",
		Unicode: "\U0001F3FA",
		Alias:   ":amphora:",
	},
	"\U0001F3FB :skin-tone-2:": {
		Name:    "Emoji Modifier Fitzpatrick Type-1-2",
		Unicode: "\U0001F3FB",
		Alias:   ":skin-tone-2:",
	},
	"\U0001F3FC :skin-tone-3:": {
		Name:    "Emoji Modifier Fitzpatrick Type-3",
		Unicode: "\U0001F3FC",
		Alias:   ":skin-tone-3:",
	},
	"\U0001F3FD :skin-tone-4:": {
		Name:    "Emoji Modifier Fitzpatrick Type-4",
		Unicode: "\U0001F3FD",
		Alias:   ":skin-tone-4:",
	},
	"\U0001F3FE :skin-tone-5:": {
		Name:    "Emoji Modifier Fitzpatrick Type-5",
		Unicode: "\U0001F3FE",
		Alias:   ":skin-tone-5:",
	},
	"\U0001F3FF :skin-tone-6:": {
		Name:    "Emoji Modifier Fitzpatrick Type-6",
		Unicode: "\U0001F3FF",
		Alias:   ":skin-tone-6:",
	},
	"\U0001F400  :rat:": {
		Name:    "Rat",
		Unicode: "\U0001F400",
		Alias:   ":rat:",
	},
	"\U0001F401  :mouse2:": {
		Name:    "Mouse",
		Unicode: "\U0001F401",
		Alias:   ":mouse2:",
	},
	"\U0001F402  :ox:": {
		Name:    "Ox",
		Unicode: "\U0001F402",
		Alias:   ":ox:",
	},
	"\U0001F403  :water_buffalo:": {
		Name:    "Water Buffalo",
		Unicode: "\U0001F403",
		Alias:   ":water_buffalo:",
	},
	"\U0001F404  :cow2:": {
		Name:    "Cow",
		Unicode: "\U0001F404",
		Alias:   ":cow2:",
	},
	"\U0001F405  :tiger2:": {
		Name:    "Tiger",
		Unicode: "\U0001F405",
		Alias:   ":tiger2:",
	},
	"\U0001F406  :leopard:": {
		Name:    "Leopard",
		Unicode: "\U0001F406",
		Alias:   ":leopard:",
	},
	"\U0001F407  :rabbit2:": {
		Name:    "Rabbit",
		Unicode: "\U0001F407",
		Alias:   ":rabbit2:",
	},
	"\U0001F408\u200D\u2B1B  :black_cat:": {
		Name:    "Black Cat",
		Unicode: "\U0001F408\u200D\u2B1B",
		Alias:   ":black_cat:",
	},
	"\U0001F408  :cat2:": {
		Name:    "Cat",
		Unicode: "\U0001F408",
		Alias:   ":cat2:",
	},
	"\U0001F409  :dragon:": {
		Name:    "Dragon",
		Unicode: "\U0001F409",
		Alias:   ":dragon:",
	},
	"\U0001F40A  :crocodile:": {
		Name:    "Crocodile",
		Unicode: "\U0001F40A",
		Alias:   ":crocodile:",
	},
	"\U0001F40B  :whale2:": {
		Name:    "Whale",
		Unicode: "\U0001F40B",
		Alias:   ":whale2:",
	},
	"\U0001F40C  :snail:": {
		Name:    "Snail",
		Unicode: "\U0001F40C",
		Alias:   ":snail:",
	},
	"\U0001F40D  :snake:": {
		Name:    "Snake",
		Unicode: "\U0001F40D",
		Alias:   ":snake:",
	},
	"\U0001F40E  :racehorse:": {
		Name:    "Horse",
		Unicode: "\U0001F40E",
		Alias:   ":racehorse:",
	},
	"\U0001F40F  :ram:": {
		Name:    "Ram",
		Unicode: "\U0001F40F",
		Alias:   ":ram:",
	},
	"\U0001F410  :goat:": {
		Name:    "Goat",
		Unicode: "\U0001F410",
		Alias:   ":goat:",
	},
	"\U0001F411  :sheep:": {
		Name:    "Sheep",
		Unicode: "\U0001F411",
		Alias:   ":sheep:",
	},
	"\U0001F412  :monkey:": {
		Name:    "Monkey",
		Unicode: "\U0001F412",
		Alias:   ":monkey:",
	},
	"\U0001F413  :rooster:": {
		Name:    "Rooster",
		Unicode: "\U0001F413",
		Alias:   ":rooster:",
	},
	"\U0001F414  :chicken:": {
		Name:    "Chicken",
		Unicode: "\U0001F414",
		Alias:   ":chicken:",
	},
	"\U0001F415\u200D\u1F9BA  :service_dog:": {
		Name:    "Service Dog",
		Unicode: "\U0001F415\u200D\u1F9BA",
		Alias:   ":service_dog:",
	},
	"\U0001F415  :dog2:": {
		Name:    "Dog",
		Unicode: "\U0001F415",
		Alias:   ":dog2:",
	},
	"\U0001F416  :pig2:": {
		Name:    "Pig",
		Unicode: "\U0001F416",
		Alias:   ":pig2:",
	},
	"\U0001F417  :boar:": {
		Name:    "Boar",
		Unicode: "\U0001F417",
		Alias:   ":boar:",
	},
	"\U0001F418  :elephant:": {
		Name:    "Elephant",
		Unicode: "\U0001F418",
		Alias:   ":elephant:",
	},
	"\U0001F419  :octopus:": {
		Name:    "Octopus",
		Unicode: "\U0001F419",
		Alias:   ":octopus:",
	},
	"\U0001F41A  :shell:": {
		Name:    "Spiral Shell",
		Unicode: "\U0001F41A",
		Alias:   ":shell:",
	},
	"\U0001F41B  :bug:": {
		Name:    "Bug",
		Unicode: "\U0001F41B",
		Alias:   ":bug:",
	},
	"\U0001F41C  :ant:": {
		Name:    "Ant",
		Unicode: "\U0001F41C",
		Alias:   ":ant:",
	},
	"\U0001F41D  :bee:": {
		Name:    "Honeybee",
		Unicode: "\U0001F41D",
		Alias:   ":bee:",
	},
	"\U0001F41E  :ladybug:": {
		Name:    "Lady Beetle",
		Unicode: "\U0001F41E",
		Alias:   ":ladybug:",
	},
	"\U0001F41F  :fish:": {
		Name:    "Fish",
		Unicode: "\U0001F41F",
		Alias:   ":fish:",
	},
	"\U0001F420  :tropical_fish:": {
		Name:    "Tropical Fish",
		Unicode: "\U0001F420",
		Alias:   ":tropical_fish:",
	},
	"\U0001F421  :blowfish:": {
		Name:    "Blowfish",
		Unicode: "\U0001F421",
		Alias:   ":blowfish:",
	},
	"\U0001F422  :turtle:": {
		Name:    "Turtle",
		Unicode: "\U0001F422",
		Alias:   ":turtle:",
	},
	"\U0001F423  :hatching_chick:": {
		Name:    "Hatching Chick",
		Unicode: "\U0001F423",
		Alias:   ":hatching_chick:",
	},
	"\U0001F424  :baby_chick:": {
		Name:    "Baby Chick",
		Unicode: "\U0001F424",
		Alias:   ":baby_chick:",
	},
	"\U0001F425  :hatched_chick:": {
		Name:    "Front-Facing Baby Chick",
		Unicode: "\U0001F425",
		Alias:   ":hatched_chick:",
	},
	"\U0001F426  :bird:": {
		Name:    "Bird",
		Unicode: "\U0001F426",
		Alias:   ":bird:",
	},
	"\U0001F427  :penguin:": {
		Name:    "Penguin",
		Unicode: "\U0001F427",
		Alias:   ":penguin:",
	},
	"\U0001F428  :koala:": {
		Name:    "Koala",
		Unicode: "\U0001F428",
		Alias:   ":koala:",
	},
	"\U0001F429  :poodle:": {
		Name:    "Poodle",
		Unicode: "\U0001F429",
		Alias:   ":poodle:",
	},
	"\U0001F42A  :dromedary_camel:": {
		Name:    "Dromedary Camel",
		Unicode: "\U0001F42A",
		Alias:   ":dromedary_camel:",
	},
	"\U0001F42B  :camel:": {
		Name:    "Bactrian Camel",
		Unicode: "\U0001F42B",
		Alias:   ":camel:",
	},
	"\U0001F42C  :dolphin:": {
		Name:    "Dolphin",
		Unicode: "\U0001F42C",
		Alias:   ":dolphin:",
	},
	"\U0001F42D  :mouse:": {
		Name:    "Mouse Face",
		Unicode: "\U0001F42D",
		Alias:   ":mouse:",
	},
	"\U0001F42E  :cow:": {
		Name:    "Cow Face",
		Unicode: "\U0001F42E",
		Alias:   ":cow:",
	},
	"\U0001F42F  :tiger:": {
		Name:    "Tiger Face",
		Unicode: "\U0001F42F",
		Alias:   ":tiger:",
	},
	"\U0001F430  :rabbit:": {
		Name:    "Rabbit Face",
		Unicode: "\U0001F430",
		Alias:   ":rabbit:",
	},
	"\U0001F431  :cat:": {
		Name:    "Cat Face",
		Unicode: "\U0001F431",
		Alias:   ":cat:",
	},
	"\U0001F432  :dragon_face:": {
		Name:    "Dragon Face",
		Unicode: "\U0001F432",
		Alias:   ":dragon_face:",
	},
	"\U0001F433  :whale:": {
		Name:    "Spouting Whale",
		Unicode: "\U0001F433",
		Alias:   ":whale:",
	},
	"\U0001F434  :horse:": {
		Name:    "Horse Face",
		Unicode: "\U0001F434",
		Alias:   ":horse:",
	},
	"\U0001F435  :monkey_face:": {
		Name:    "Monkey Face",
		Unicode: "\U0001F435",
		Alias:   ":monkey_face:",
	},
	"\U0001F436  :dog:": {
		Name:    "Dog Face",
		Unicode: "\U0001F436",
		Alias:   ":dog:",
	},
	"\U0001F437  :pig:": {
		Name:    "Pig Face",
		Unicode: "\U0001F437",
		Alias:   ":pig:",
	},
	"\U0001F438  :frog:": {
		Name:    "Frog Face",
		Unicode: "\U0001F438",
		Alias:   ":frog:",
	},
	"\U0001F439  :hamster:": {
		Name:    "Hamster Face",
		Unicode: "\U0001F439",
		Alias:   ":hamster:",
	},
	"\U0001F43A  :wolf:": {
		Name:    "Wolf Face",
		Unicode: "\U0001F43A",
		Alias:   ":wolf:",
	},
	"\U0001F43B\u200D\u2744\uFE0F  :polar_bear:": {
		Name:    "Polar Bear",
		Unicode: "\U0001F43B\u200D\u2744\uFE0F",
		Alias:   ":polar_bear:",
	},
	"\U0001F43B  :bear:": {
		Name:    "Bear Face",
		Unicode: "\U0001F43B",
		Alias:   ":bear:",
	},
	"\U0001F43C  :panda_face:": {
		Name:    "Panda Face",
		Unicode: "\U0001F43C",
		Alias:   ":panda_face:",
	},
	"\U0001F43D  :pig_nose:": {
		Name:    "Pig Nose",
		Unicode: "\U0001F43D",
		Alias:   ":pig_nose:",
	},
	"\U0001F43E  :feet:": {
		Name:    "Paw Prints",
		Unicode: "\U0001F43E",
		Alias:   ":feet:",
	},
	"\U0001F43F\uFE0F  :chipmunk:": {
		Name:    "Chipmunk",
		Unicode: "\U0001F43F\uFE0F",
		Alias:   ":chipmunk:",
	},
	"\U0001F440  :eyes:": {
		Name:    "Eyes",
		Unicode: "\U0001F440",
		Alias:   ":eyes:",
	},
	"\U0001F441\uFE0F\u200D\u1F5E8\uFE0F :eye-in-speech-bubble:": {
		Name:    "Eye In Speech Bubble",
		Unicode: "\U0001F441\uFE0F\u200D\u1F5E8\uFE0F",
		Alias:   ":eye-in-speech-bubble:",
	},
	"\U0001F441\uFE0F  :eye:": {
		Name:    "Eye",
		Unicode: "\U0001F441\uFE0F",
		Alias:   ":eye:",
	},
	"\U0001F442  :ear:": {
		Name:    "Ear",
		Unicode: "\U0001F442",
		Alias:   ":ear:",
	},
	"\U0001F443  :nose:": {
		Name:    "Nose",
		Unicode: "\U0001F443",
		Alias:   ":nose:",
	},
	"\U0001F444  :lips:": {
		Name:    "Mouth",
		Unicode: "\U0001F444",
		Alias:   ":lips:",
	},
	"\U0001F445  :tongue:": {
		Name:    "Tongue",
		Unicode: "\U0001F445",
		Alias:   ":tongue:",
	},
	"\U0001F446  :point_up_2:": {
		Name:    "White Up Pointing Backhand Index",
		Unicode: "\U0001F446",
		Alias:   ":point_up_2:",
	},
	"\U0001F447  :point_down:": {
		Name:    "White Down Pointing Backhand Index",
		Unicode: "\U0001F447",
		Alias:   ":point_down:",
	},
	"\U0001F448  :point_left:": {
		Name:    "White Left Pointing Backhand Index",
		Unicode: "\U0001F448",
		Alias:   ":point_left:",
	},
	"\U0001F449  :point_right:": {
		Name:    "White Right Pointing Backhand Index",
		Unicode: "\U0001F449",
		Alias:   ":point_right:",
	},
	"\U0001F44A  :facepunch:": {
		Name:    "Fisted Hand Sign",
		Unicode: "\U0001F44A",
		Alias:   ":facepunch:",
	},
	"\U0001F44B  :wave:": {
		Name:    "Waving Hand Sign",
		Unicode: "\U0001F44B",
		Alias:   ":wave:",
	},
	"\U0001F44C  :ok_hand:": {
		Name:    "Ok Hand Sign",
		Unicode: "\U0001F44C",
		Alias:   ":ok_hand:",
	},
	"\U0001F44D :+1:": {
		Name:    "Thumbs Up Sign",
		Unicode: "\U0001F44D",
		Alias:   ":+1:",
	},
	"\U0001F44E :-1:": {
		Name:    "Thumbs Down Sign",
		Unicode: "\U0001F44E",
		Alias:   ":-1:",
	},
	"\U0001F44F  :clap:": {
		Name:    "Clapping Hands Sign",
		Unicode: "\U0001F44F",
		Alias:   ":clap:",
	},
	"\U0001F450  :open_hands:": {
		Name:    "Open Hands Sign",
		Unicode: "\U0001F450",
		Alias:   ":open_hands:",
	},
	"\U0001F451  :crown:": {
		Name:    "Crown",
		Unicode: "\U0001F451",
		Alias:   ":crown:",
	},
	"\U0001F452  :womans_hat:": {
		Name:    "Womans Hat",
		Unicode: "\U0001F452",
		Alias:   ":womans_hat:",
	},
	"\U0001F453  :eyeglasses:": {
		Name:    "Eyeglasses",
		Unicode: "\U0001F453",
		Alias:   ":eyeglasses:",
	},
	"\U0001F454  :necktie:": {
		Name:    "Necktie",
		Unicode: "\U0001F454",
		Alias:   ":necktie:",
	},
	"\U0001F455  :shirt:": {
		Name:    "T-Shirt",
		Unicode: "\U0001F455",
		Alias:   ":shirt:",
	},
	"\U0001F456  :jeans:": {
		Name:    "Jeans",
		Unicode: "\U0001F456",
		Alias:   ":jeans:",
	},
	"\U0001F457  :dress:": {
		Name:    "Dress",
		Unicode: "\U0001F457",
		Alias:   ":dress:",
	},
	"\U0001F458  :kimono:": {
		Name:    "Kimono",
		Unicode: "\U0001F458",
		Alias:   ":kimono:",
	},
	"\U0001F459  :bikini:": {
		Name:    "Bikini",
		Unicode: "\U0001F459",
		Alias:   ":bikini:",
	},
	"\U0001F45A  :womans_clothes:": {
		Name:    "Womans Clothes",
		Unicode: "\U0001F45A",
		Alias:   ":womans_clothes:",
	},
	"\U0001F45B  :purse:": {
		Name:    "Purse",
		Unicode: "\U0001F45B",
		Alias:   ":purse:",
	},
	"\U0001F45C  :handbag:": {
		Name:    "Handbag",
		Unicode: "\U0001F45C",
		Alias:   ":handbag:",
	},
	"\U0001F45D  :pouch:": {
		Name:    "Pouch",
		Unicode: "\U0001F45D",
		Alias:   ":pouch:",
	},
	"\U0001F45E  :mans_shoe:": {
		Name:    "Mans Shoe",
		Unicode: "\U0001F45E",
		Alias:   ":mans_shoe:",
	},
	"\U0001F45F  :athletic_shoe:": {
		Name:    "Athletic Shoe",
		Unicode: "\U0001F45F",
		Alias:   ":athletic_shoe:",
	},
	"\U0001F460  :high_heel:": {
		Name:    "High-Heeled Shoe",
		Unicode: "\U0001F460",
		Alias:   ":high_heel:",
	},
	"\U0001F461  :sandal:": {
		Name:    "Womans Sandal",
		Unicode: "\U0001F461",
		Alias:   ":sandal:",
	},
	"\U0001F462  :boot:": {
		Name:    "Womans Boots",
		Unicode: "\U0001F462",
		Alias:   ":boot:",
	},
	"\U0001F463  :footprints:": {
		Name:    "Footprints",
		Unicode: "\U0001F463",
		Alias:   ":footprints:",
	},
	"\U0001F464  :bust_in_silhouette:": {
		Name:    "Bust In Silhouette",
		Unicode: "\U0001F464",
		Alias:   ":bust_in_silhouette:",
	},
	"\U0001F465  :busts_in_silhouette:": {
		Name:    "Busts In Silhouette",
		Unicode: "\U0001F465",
		Alias:   ":busts_in_silhouette:",
	},
	"\U0001F466  :boy:": {
		Name:    "Boy",
		Unicode: "\U0001F466",
		Alias:   ":boy:",
	},
	"\U0001F467  :girl:": {
		Name:    "Girl",
		Unicode: "\U0001F467",
		Alias:   ":girl:",
	},
	"\U0001F468\u200D\u1F33E :male-farmer:": {
		Name:    "Man Farmer",
		Unicode: "\U0001F468\u200D\u1F33E",
		Alias:   ":male-farmer:",
	},
	"\U0001F468\u200D\u1F373 :male-cook:": {
		Name:    "Man Cook",
		Unicode: "\U0001F468\u200D\u1F373",
		Alias:   ":male-cook:",
	},
	"\U0001F468\u200D\u1F37C  :man_feeding_baby:": {
		Name:    "Man Feeding Baby",
		Unicode: "\U0001F468\u200D\u1F37C",
		Alias:   ":man_feeding_baby:",
	},
	"\U0001F468\u200D\u1F393 :male-student:": {
		Name:    "Man Student",
		Unicode: "\U0001F468\u200D\u1F393",
		Alias:   ":male-student:",
	},
	"\U0001F468\u200D\u1F3A4 :male-singer:": {
		Name:    "Man Singer",
		Unicode: "\U0001F468\u200D\u1F3A4",
		Alias:   ":male-singer:",
	},
	"\U0001F468\u200D\u1F3A8 :male-artist:": {
		Name:    "Man Artist",
		Unicode: "\U0001F468\u200D\u1F3A8",
		Alias:   ":male-artist:",
	},
	"\U0001F468\u200D\u1F3EB :male-teacher:": {
		Name:    "Man Teacher",
		Unicode: "\U0001F468\u200D\u1F3EB",
		Alias:   ":male-teacher:",
	},
	"\U0001F468\u200D\u1F3ED :male-factory-worker:": {
		Name:    "Man Factory Worker",
		Unicode: "\U0001F468\u200D\u1F3ED",
		Alias:   ":male-factory-worker:",
	},
	"\U0001F468\u200D\u1F466\u200D\u1F466 :man-boy-boy:": {
		Name:    "Family: Man, Boy, Boy",
		Unicode: "\U0001F468\u200D\u1F466\u200D\u1F466",
		Alias:   ":man-boy-boy:",
	},
	"\U0001F468\u200D\u1F466 :man-boy:": {
		Name:    "Family: Man, Boy",
		Unicode: "\U0001F468\u200D\u1F466",
		Alias:   ":man-boy:",
	},
	"\U0001F468\u200D\u1F467\u200D\u1F466 :man-girl-boy:": {
		Name:    "Family: Man, Girl, Boy",
		Unicode: "\U0001F468\u200D\u1F467\u200D\u1F466",
		Alias:   ":man-girl-boy:",
	},
	"\U0001F468\u200D\u1F467\u200D\u1F467 :man-girl-girl:": {
		Name:    "Family: Man, Girl, Girl",
		Unicode: "\U0001F468\u200D\u1F467\u200D\u1F467",
		Alias:   ":man-girl-girl:",
	},
	"\U0001F468\u200D\u1F467 :man-girl:": {
		Name:    "Family: Man, Girl",
		Unicode: "\U0001F468\u200D\u1F467",
		Alias:   ":man-girl:",
	},
	"\U0001F468\u200D\u1F468\u200D\u1F466 :man-man-boy:": {
		Name:    "Family: Man, Man, Boy",
		Unicode: "\U0001F468\u200D\u1F468\u200D\u1F466",
		Alias:   ":man-man-boy:",
	},
	"\U0001F468\u200D\u1F468\u200D\u1F466\u200D\u1F466 :man-man-boy-boy:": {
		Name:    "Family: Man, Man, Boy, Boy",
		Unicode: "\U0001F468\u200D\u1F468\u200D\u1F466\u200D\u1F466",
		Alias:   ":man-man-boy-boy:",
	},
	"\U0001F468\u200D\u1F468\u200D\u1F467 :man-man-girl:": {
		Name:    "Family: Man, Man, Girl",
		Unicode: "\U0001F468\u200D\u1F468\u200D\u1F467",
		Alias:   ":man-man-girl:",
	},
	"\U0001F468\u200D\u1F468\u200D\u1F467\u200D\u1F466 :man-man-girl-boy:": {
		Name:    "Family: Man, Man, Girl, Boy",
		Unicode: "\U0001F468\u200D\u1F468\u200D\u1F467\u200D\u1F466",
		Alias:   ":man-man-girl-boy:",
	},
	"\U0001F468\u200D\u1F468\u200D\u1F467\u200D\u1F467 :man-man-girl-girl:": {
		Name:    "Family: Man, Man, Girl, Girl",
		Unicode: "\U0001F468\u200D\u1F468\u200D\u1F467\u200D\u1F467",
		Alias:   ":man-man-girl-girl:",
	},
	"\U0001F468\u200D\u1F469\u200D\u1F466 :man-woman-boy:": {
		Name:    "Family: Man, Woman, Boy",
		Unicode: "\U0001F468\u200D\u1F469\u200D\u1F466",
		Alias:   ":man-woman-boy:",
	},
	"\U0001F468\u200D\u1F469\u200D\u1F466\u200D\u1F466 :man-woman-boy-boy:": {
		Name:    "Family: Man, Woman, Boy, Boy",
		Unicode: "\U0001F468\u200D\u1F469\u200D\u1F466\u200D\u1F466",
		Alias:   ":man-woman-boy-boy:",
	},
	"\U0001F468\u200D\u1F469\u200D\u1F467 :man-woman-girl:": {
		Name:    "Family: Man, Woman, Girl",
		Unicode: "\U0001F468\u200D\u1F469\u200D\u1F467",
		Alias:   ":man-woman-girl:",
	},
	"\U0001F468\u200D\u1F469\u200D\u1F467\u200D\u1F466 :man-woman-girl-boy:": {
		Name:    "Family: Man, Woman, Girl, Boy",
		Unicode: "\U0001F468\u200D\u1F469\u200D\u1F467\u200D\u1F466",
		Alias:   ":man-woman-girl-boy:",
	},
	"\U0001F468\u200D\u1F469\u200D\u1F467\u200D\u1F467 :man-woman-girl-girl:": {
		Name:    "Family: Man, Woman, Girl, Girl",
		Unicode: "\U0001F468\u200D\u1F469\u200D\u1F467\u200D\u1F467",
		Alias:   ":man-woman-girl-girl:",
	},
	"\U0001F468\u200D\u1F4BB :male-technologist:": {
		Name:    "Man Technologist",
		Unicode: "\U0001F468\u200D\u1F4BB",
		Alias:   ":male-technologist:",
	},
	"\U0001F468\u200D\u1F4BC :male-office-worker:": {
		Name:    "Man Office Worker",
		Unicode: "\U0001F468\u200D\u1F4BC",
		Alias:   ":male-office-worker:",
	},
	"\U0001F468\u200D\u1F527 :male-mechanic:": {
		Name:    "Man Mechanic",
		Unicode: "\U0001F468\u200D\u1F527",
		Alias:   ":male-mechanic:",
	},
	"\U0001F468\u200D\u1F52C :male-scientist:": {
		Name:    "Man Scientist",
		Unicode: "\U0001F468\u200D\u1F52C",
		Alias:   ":male-scientist:",
	},
	"\U0001F468\u200D\u1F680 :male-astronaut:": {
		Name:    "Man Astronaut",
		Unicode: "\U0001F468\u200D\u1F680",
		Alias:   ":male-astronaut:",
	},
	"\U0001F468\u200D\u1F692 :male-firefighter:": {
		Name:    "Man Firefighter",
		Unicode: "\U0001F468\u200D\u1F692",
		Alias:   ":male-firefighter:",
	},
	"\U0001F468\u200D\u1F9AF  :man_with_probing_cane:": {
		Name:    "Man With White Cane",
		Unicode: "\U0001F468\u200D\u1F9AF",
		Alias:   ":man_with_probing_cane:",
	},
	"\U0001F468\u200D\u1F9B0  :red_haired_man:": {
		Name:    "Man: Red Hair",
		Unicode: "\U0001F468\u200D\u1F9B0",
		Alias:   ":red_haired_man:",
	},
	"\U0001F468\u200D\u1F9B1  :curly_haired_man:": {
		Name:    "Man: Curly Hair",
		Unicode: "\U0001F468\u200D\u1F9B1",
		Alias:   ":curly_haired_man:",
	},
	"\U0001F468\u200D\u1F9B2  :bald_man:": {
		Name:    "Man: Bald",
		Unicode: "\U0001F468\u200D\u1F9B2",
		Alias:   ":bald_man:",
	},
	"\U0001F468\u200D\u1F9B3  :white_haired_man:": {
		Name:    "Man: White Hair",
		Unicode: "\U0001F468\u200D\u1F9B3",
		Alias:   ":white_haired_man:",
	},
	"\U0001F468\u200D\u1F9BC  :man_in_motorized_wheelchair:": {
		Name:    "Man In Motorized Wheelchair",
		Unicode: "\U0001F468\u200D\u1F9BC",
		Alias:   ":man_in_motorized_wheelchair:",
	},
	"\U0001F468\u200D\u1F9BD  :man_in_manual_wheelchair:": {
		Name:    "Man In Manual Wheelchair",
		Unicode: "\U0001F468\u200D\u1F9BD",
		Alias:   ":man_in_manual_wheelchair:",
	},
	"\U0001F468\u200D\u2695\uFE0F :male-doctor:": {
		Name:    "Man Health Worker",
		Unicode: "\U0001F468\u200D\u2695\uFE0F",
		Alias:   ":male-doctor:",
	},
	"\U0001F468\u200D\u2696\uFE0F :male-judge:": {
		Name:    "Man Judge",
		Unicode: "\U0001F468\u200D\u2696\uFE0F",
		Alias:   ":male-judge:",
	},
	"\U0001F468\u200D\u2708\uFE0F :male-pilot:": {
		Name:    "Man Pilot",
		Unicode: "\U0001F468\u200D\u2708\uFE0F",
		Alias:   ":male-pilot:",
	},
	"\U0001F468\u200D\u2764\uFE0F\u200D\u1F468 :man-heart-man:": {
		Name:    "Couple With Heart: Man, Man",
		Unicode: "\U0001F468\u200D\u2764\uFE0F\u200D\u1F468",
		Alias:   ":man-heart-man:",
	},
	"\U0001F468\u200D\u2764\uFE0F\u200D\u1F48B\u200D\u1F468 :man-kiss-man:": {
		Name:    "Kiss: Man, Man",
		Unicode: "\U0001F468\u200D\u2764\uFE0F\u200D\u1F48B\u200D\u1F468",
		Alias:   ":man-kiss-man:",
	},
	"\U0001F468  :man:": {
		Name:    "Man",
		Unicode: "\U0001F468",
		Alias:   ":man:",
	},
	"\U0001F469\u200D\u1F33E :female-farmer:": {
		Name:    "Woman Farmer",
		Unicode: "\U0001F469\u200D\u1F33E",
		Alias:   ":female-farmer:",
	},
	"\U0001F469\u200D\u1F373 :female-cook:": {
		Name:    "Woman Cook",
		Unicode: "\U0001F469\u200D\u1F373",
		Alias:   ":female-cook:",
	},
	"\U0001F469\u200D\u1F37C  :woman_feeding_baby:": {
		Name:    "Woman Feeding Baby",
		Unicode: "\U0001F469\u200D\u1F37C",
		Alias:   ":woman_feeding_baby:",
	},
	"\U0001F469\u200D\u1F393 :female-student:": {
		Name:    "Woman Student",
		Unicode: "\U0001F469\u200D\u1F393",
		Alias:   ":female-student:",
	},
	"\U0001F469\u200D\u1F3A4 :female-singer:": {
		Name:    "Woman Singer",
		Unicode: "\U0001F469\u200D\u1F3A4",
		Alias:   ":female-singer:",
	},
	"\U0001F469\u200D\u1F3A8 :female-artist:": {
		Name:    "Woman Artist",
		Unicode: "\U0001F469\u200D\u1F3A8",
		Alias:   ":female-artist:",
	},
	"\U0001F469\u200D\u1F3EB :female-teacher:": {
		Name:    "Woman Teacher",
		Unicode: "\U0001F469\u200D\u1F3EB",
		Alias:   ":female-teacher:",
	},
	"\U0001F469\u200D\u1F3ED :female-factory-worker:": {
		Name:    "Woman Factory Worker",
		Unicode: "\U0001F469\u200D\u1F3ED",
		Alias:   ":female-factory-worker:",
	},
	"\U0001F469\u200D\u1F466\u200D\u1F466 :woman-boy-boy:": {
		Name:    "Family: Woman, Boy, Boy",
		Unicode: "\U0001F469\u200D\u1F466\u200D\u1F466",
		Alias:   ":woman-boy-boy:",
	},
	"\U0001F469\u200D\u1F466 :woman-boy:": {
		Name:    "Family: Woman, Boy",
		Unicode: "\U0001F469\u200D\u1F466",
		Alias:   ":woman-boy:",
	},
	"\U0001F469\u200D\u1F467\u200D\u1F466 :woman-girl-boy:": {
		Name:    "Family: Woman, Girl, Boy",
		Unicode: "\U0001F469\u200D\u1F467\u200D\u1F466",
		Alias:   ":woman-girl-boy:",
	},
	"\U0001F469\u200D\u1F467\u200D\u1F467 :woman-girl-girl:": {
		Name:    "Family: Woman, Girl, Girl",
		Unicode: "\U0001F469\u200D\u1F467\u200D\u1F467",
		Alias:   ":woman-girl-girl:",
	},
	"\U0001F469\u200D\u1F467 :woman-girl:": {
		Name:    "Family: Woman, Girl",
		Unicode: "\U0001F469\u200D\u1F467",
		Alias:   ":woman-girl:",
	},
	"\U0001F469\u200D\u1F469\u200D\u1F466 :woman-woman-boy:": {
		Name:    "Family: Woman, Woman, Boy",
		Unicode: "\U0001F469\u200D\u1F469\u200D\u1F466",
		Alias:   ":woman-woman-boy:",
	},
	"\U0001F469\u200D\u1F469\u200D\u1F466\u200D\u1F466 :woman-woman-boy-boy:": {
		Name:    "Family: Woman, Woman, Boy, Boy",
		Unicode: "\U0001F469\u200D\u1F469\u200D\u1F466\u200D\u1F466",
		Alias:   ":woman-woman-boy-boy:",
	},
	"\U0001F469\u200D\u1F469\u200D\u1F467 :woman-woman-girl:": {
		Name:    "Family: Woman, Woman, Girl",
		Unicode: "\U0001F469\u200D\u1F469\u200D\u1F467",
		Alias:   ":woman-woman-girl:",
	},
	"\U0001F469\u200D\u1F469\u200D\u1F467\u200D\u1F466 :woman-woman-girl-boy:": {
		Name:    "Family: Woman, Woman, Girl, Boy",
		Unicode: "\U0001F469\u200D\u1F469\u200D\u1F467\u200D\u1F466",
		Alias:   ":woman-woman-girl-boy:",
	},
	"\U0001F469\u200D\u1F469\u200D\u1F467\u200D\u1F467 :woman-woman-girl-girl:": {
		Name:    "Family: Woman, Woman, Girl, Girl",
		Unicode: "\U0001F469\u200D\u1F469\u200D\u1F467\u200D\u1F467",
		Alias:   ":woman-woman-girl-girl:",
	},
	"\U0001F469\u200D\u1F4BB :female-technologist:": {
		Name:    "Woman Technologist",
		Unicode: "\U0001F469\u200D\u1F4BB",
		Alias:   ":female-technologist:",
	},
	"\U0001F469\u200D\u1F4BC :female-office-worker:": {
		Name:    "Woman Office Worker",
		Unicode: "\U0001F469\u200D\u1F4BC",
		Alias:   ":female-office-worker:",
	},
	"\U0001F469\u200D\u1F527 :female-mechanic:": {
		Name:    "Woman Mechanic",
		Unicode: "\U0001F469\u200D\u1F527",
		Alias:   ":female-mechanic:",
	},
	"\U0001F469\u200D\u1F52C :female-scientist:": {
		Name:    "Woman Scientist",
		Unicode: "\U0001F469\u200D\u1F52C",
		Alias:   ":female-scientist:",
	},
	"\U0001F469\u200D\u1F680 :female-astronaut:": {
		Name:    "Woman Astronaut",
		Unicode: "\U0001F469\u200D\u1F680",
		Alias:   ":female-astronaut:",
	},
	"\U0001F469\u200D\u1F692 :female-firefighter:": {
		Name:    "Woman Firefighter",
		Unicode: "\U0001F469\u200D\u1F692",
		Alias:   ":female-firefighter:",
	},
	"\U0001F469\u200D\u1F9AF  :woman_with_probing_cane:": {
		Name:    "Woman With White Cane",
		Unicode: "\U0001F469\u200D\u1F9AF",
		Alias:   ":woman_with_probing_cane:",
	},
	"\U0001F469\u200D\u1F9B0  :red_haired_woman:": {
		Name:    "Woman: Red Hair",
		Unicode: "\U0001F469\u200D\u1F9B0",
		Alias:   ":red_haired_woman:",
	},
	"\U0001F469\u200D\u1F9B1  :curly_haired_woman:": {
		Name:    "Woman: Curly Hair",
		Unicode: "\U0001F469\u200D\u1F9B1",
		Alias:   ":curly_haired_woman:",
	},
	"\U0001F469\u200D\u1F9B2  :bald_woman:": {
		Name:    "Woman: Bald",
		Unicode: "\U0001F469\u200D\u1F9B2",
		Alias:   ":bald_woman:",
	},
	"\U0001F469\u200D\u1F9B3  :white_haired_woman:": {
		Name:    "Woman: White Hair",
		Unicode: "\U0001F469\u200D\u1F9B3",
		Alias:   ":white_haired_woman:",
	},
	"\U0001F469\u200D\u1F9BC  :woman_in_motorized_wheelchair:": {
		Name:    "Woman In Motorized Wheelchair",
		Unicode: "\U0001F469\u200D\u1F9BC",
		Alias:   ":woman_in_motorized_wheelchair:",
	},
	"\U0001F469\u200D\u1F9BD  :woman_in_manual_wheelchair:": {
		Name:    "Woman In Manual Wheelchair",
		Unicode: "\U0001F469\u200D\u1F9BD",
		Alias:   ":woman_in_manual_wheelchair:",
	},
	"\U0001F469\u200D\u2695\uFE0F :female-doctor:": {
		Name:    "Woman Health Worker",
		Unicode: "\U0001F469\u200D\u2695\uFE0F",
		Alias:   ":female-doctor:",
	},
	"\U0001F469\u200D\u2696\uFE0F :female-judge:": {
		Name:    "Woman Judge",
		Unicode: "\U0001F469\u200D\u2696\uFE0F",
		Alias:   ":female-judge:",
	},
	"\U0001F469\u200D\u2708\uFE0F :female-pilot:": {
		Name:    "Woman Pilot",
		Unicode: "\U0001F469\u200D\u2708\uFE0F",
		Alias:   ":female-pilot:",
	},
	"\U0001F469\u200D\u2764\uFE0F\u200D\u1F468 :woman-heart-man:": {
		Name:    "Couple With Heart: Woman, Man",
		Unicode: "\U0001F469\u200D\u2764\uFE0F\u200D\u1F468",
		Alias:   ":woman-heart-man:",
	},
	"\U0001F469\u200D\u2764\uFE0F\u200D\u1F469 :woman-heart-woman:": {
		Name:    "Couple With Heart: Woman, Woman",
		Unicode: "\U0001F469\u200D\u2764\uFE0F\u200D\u1F469",
		Alias:   ":woman-heart-woman:",
	},
	"\U0001F469\u200D\u2764\uFE0F\u200D\u1F48B\u200D\u1F468 :woman-kiss-man:": {
		Name:    "Kiss: Woman, Man",
		Unicode: "\U0001F469\u200D\u2764\uFE0F\u200D\u1F48B\u200D\u1F468",
		Alias:   ":woman-kiss-man:",
	},
	"\U0001F469\u200D\u2764\uFE0F\u200D\u1F48B\u200D\u1F469 :woman-kiss-woman:": {
		Name:    "Kiss: Woman, Woman",
		Unicode: "\U0001F469\u200D\u2764\uFE0F\u200D\u1F48B\u200D\u1F469",
		Alias:   ":woman-kiss-woman:",
	},
	"\U0001F469  :woman:": {
		Name:    "Woman",
		Unicode: "\U0001F469",
		Alias:   ":woman:",
	},
	"\U0001F46A  :family:": {
		Name:    "Family",
		Unicode: "\U0001F46A",
		Alias:   ":family:",
	},
	"\U0001F46B  :man_and_woman_holding_hands:": {
		Name:    "Man And Woman Holding Hands",
		Unicode: "\U0001F46B",
		Alias:   ":man_and_woman_holding_hands:",
	},
	"\U0001F46C  :two_men_holding_hands:": {
		Name:    "Two Men Holding Hands",
		Unicode: "\U0001F46C",
		Alias:   ":two_men_holding_hands:",
	},
	"\U0001F46D  :two_women_holding_hands:": {
		Name:    "Two Women Holding Hands",
		Unicode: "\U0001F46D",
		Alias:   ":two_women_holding_hands:",
	},
	"\U0001F46E\u200D\u2640\uFE0F :female-police-officer:": {
		Name:    "Woman Police Officer",
		Unicode: "\U0001F46E\u200D\u2640\uFE0F",
		Alias:   ":female-police-officer:",
	},
	"\U0001F46E\u200D\u2642\uFE0F :male-police-officer:": {
		Name:    "Man Police Officer",
		Unicode: "\U0001F46E\u200D\u2642\uFE0F",
		Alias:   ":male-police-officer:",
	},
	"\U0001F46E  :cop:": {
		Name:    "Police Officer",
		Unicode: "\U0001F46E",
		Alias:   ":cop:",
	},
	"\U0001F46F\u200D\u2640\uFE0F :women-with-bunny-ears-partying:": {
		Name:    "Women With Bunny Ears",
		Unicode: "\U0001F46F\u200D\u2640\uFE0F",
		Alias:   ":women-with-bunny-ears-partying:",
	},
	"\U0001F46F\u200D\u2642\uFE0F :men-with-bunny-ears-partying:": {
		Name:    "Men With Bunny Ears",
		Unicode: "\U0001F46F\u200D\u2642\uFE0F",
		Alias:   ":men-with-bunny-ears-partying:",
	},
	"\U0001F46F  :dancers:": {
		Name:    "Woman With Bunny Ears",
		Unicode: "\U0001F46F",
		Alias:   ":dancers:",
	},
	"\U0001F470\u200D\u2640\uFE0F  :woman_with_veil:": {
		Name:    "Woman With Veil",
		Unicode: "\U0001F470\u200D\u2640\uFE0F",
		Alias:   ":woman_with_veil:",
	},
	"\U0001F470\u200D\u2642\uFE0F  :man_with_veil:": {
		Name:    "Man With Veil",
		Unicode: "\U0001F470\u200D\u2642\uFE0F",
		Alias:   ":man_with_veil:",
	},
	"\U0001F470  :bride_with_veil:": {
		Name:    "Bride With Veil",
		Unicode: "\U0001F470",
		Alias:   ":bride_with_veil:",
	},
	"\U0001F471\u200D\u2640\uFE0F :blond-haired-woman:": {
		Name:    "Woman: Blond Hair",
		Unicode: "\U0001F471\u200D\u2640\uFE0F",
		Alias:   ":blond-haired-woman:",
	},
	"\U0001F471\u200D\u2642\uFE0F :blond-haired-man:": {
		Name:    "Man: Blond Hair",
		Unicode: "\U0001F471\u200D\u2642\uFE0F",
		Alias:   ":blond-haired-man:",
	},
	"\U0001F471  :person_with_blond_hair:": {
		Name:    "Person With Blond Hair",
		Unicode: "\U0001F471",
		Alias:   ":person_with_blond_hair:",
	},
	"\U0001F472  :man_with_gua_pi_mao:": {
		Name:    "Man With Gua Pi Mao",
		Unicode: "\U0001F472",
		Alias:   ":man_with_gua_pi_mao:",
	},
	"\U0001F473\u200D\u2640\uFE0F :woman-wearing-turban:": {
		Name:    "Woman Wearing Turban",
		Unicode: "\U0001F473\u200D\u2640\uFE0F",
		Alias:   ":woman-wearing-turban:",
	},
	"\U0001F473\u200D\u2642\uFE0F :man-wearing-turban:": {
		Name:    "Man Wearing Turban",
		Unicode: "\U0001F473\u200D\u2642\uFE0F",
		Alias:   ":man-wearing-turban:",
	},
	"\U0001F473  :man_with_turban:": {
		Name:    "Man With Turban",
		Unicode: "\U0001F473",
		Alias:   ":man_with_turban:",
	},
	"\U0001F474  :older_man:": {
		Name:    "Older Man",
		Unicode: "\U0001F474",
		Alias:   ":older_man:",
	},
	"\U0001F475  :older_woman:": {
		Name:    "Older Woman",
		Unicode: "\U0001F475",
		Alias:   ":older_woman:",
	},
	"\U0001F476  :baby:": {
		Name:    "Baby",
		Unicode: "\U0001F476",
		Alias:   ":baby:",
	},
	"\U0001F477\u200D\u2640\uFE0F :female-construction-worker:": {
		Name:    "Woman Construction Worker",
		Unicode: "\U0001F477\u200D\u2640\uFE0F",
		Alias:   ":female-construction-worker:",
	},
	"\U0001F477\u200D\u2642\uFE0F :male-construction-worker:": {
		Name:    "Man Construction Worker",
		Unicode: "\U0001F477\u200D\u2642\uFE0F",
		Alias:   ":male-construction-worker:",
	},
	"\U0001F477  :construction_worker:": {
		Name:    "Construction Worker",
		Unicode: "\U0001F477",
		Alias:   ":construction_worker:",
	},
	"\U0001F478  :princess:": {
		Name:    "Princess",
		Unicode: "\U0001F478",
		Alias:   ":princess:",
	},
	"\U0001F479  :japanese_ogre:": {
		Name:    "Japanese Ogre",
		Unicode: "\U0001F479",
		Alias:   ":japanese_ogre:",
	},
	"\U0001F47A  :japanese_goblin:": {
		Name:    "Japanese Goblin",
		Unicode: "\U0001F47A",
		Alias:   ":japanese_goblin:",
	},
	"\U0001F47B  :ghost:": {
		Name:    "Ghost",
		Unicode: "\U0001F47B",
		Alias:   ":ghost:",
	},
	"\U0001F47C  :angel:": {
		Name:    "Baby Angel",
		Unicode: "\U0001F47C",
		Alias:   ":angel:",
	},
	"\U0001F47D  :alien:": {
		Name:    "Extraterrestrial Alien",
		Unicode: "\U0001F47D",
		Alias:   ":alien:",
	},
	"\U0001F47E  :space_invader:": {
		Name:    "Alien Monster",
		Unicode: "\U0001F47E",
		Alias:   ":space_invader:",
	},
	"\U0001F47F  :imp:": {
		Name:    "Imp",
		Unicode: "\U0001F47F",
		Alias:   ":imp:",
	},
	"\U0001F480  :skull:": {
		Name:    "Skull",
		Unicode: "\U0001F480",
		Alias:   ":skull:",
	},
	"\U0001F481\u200D\u2640\uFE0F :woman-tipping-hand:": {
		Name:    "Woman Tipping Hand",
		Unicode: "\U0001F481\u200D\u2640\uFE0F",
		Alias:   ":woman-tipping-hand:",
	},
	"\U0001F481\u200D\u2642\uFE0F :man-tipping-hand:": {
		Name:    "Man Tipping Hand",
		Unicode: "\U0001F481\u200D\u2642\uFE0F",
		Alias:   ":man-tipping-hand:",
	},
	"\U0001F481  :information_desk_person:": {
		Name:    "Information Desk Person",
		Unicode: "\U0001F481",
		Alias:   ":information_desk_person:",
	},
	"\U0001F482\u200D\u2640\uFE0F :female-guard:": {
		Name:    "Woman Guard",
		Unicode: "\U0001F482\u200D\u2640\uFE0F",
		Alias:   ":female-guard:",
	},
	"\U0001F482\u200D\u2642\uFE0F :male-guard:": {
		Name:    "Man Guard",
		Unicode: "\U0001F482\u200D\u2642\uFE0F",
		Alias:   ":male-guard:",
	},
	"\U0001F482  :guardsman:": {
		Name:    "Guardsman",
		Unicode: "\U0001F482",
		Alias:   ":guardsman:",
	},
	"\U0001F483  :dancer:": {
		Name:    "Dancer",
		Unicode: "\U0001F483",
		Alias:   ":dancer:",
	},
	"\U0001F484  :lipstick:": {
		Name:    "Lipstick",
		Unicode: "\U0001F484",
		Alias:   ":lipstick:",
	},
	"\U0001F485  :nail_care:": {
		Name:    "Nail Polish",
		Unicode: "\U0001F485",
		Alias:   ":nail_care:",
	},
	"\U0001F486\u200D\u2640\uFE0F :woman-getting-massage:": {
		Name:    "Woman Getting Massage",
		Unicode: "\U0001F486\u200D\u2640\uFE0F",
		Alias:   ":woman-getting-massage:",
	},
	"\U0001F486\u200D\u2642\uFE0F :man-getting-massage:": {
		Name:    "Man Getting Massage",
		Unicode: "\U0001F486\u200D\u2642\uFE0F",
		Alias:   ":man-getting-massage:",
	},
	"\U0001F486  :massage:": {
		Name:    "Face Massage",
		Unicode: "\U0001F486",
		Alias:   ":massage:",
	},
	"\U0001F487\u200D\u2640\uFE0F :woman-getting-haircut:": {
		Name:    "Woman Getting Haircut",
		Unicode: "\U0001F487\u200D\u2640\uFE0F",
		Alias:   ":woman-getting-haircut:",
	},
	"\U0001F487\u200D\u2642\uFE0F :man-getting-haircut:": {
		Name:    "Man Getting Haircut",
		Unicode: "\U0001F487\u200D\u2642\uFE0F",
		Alias:   ":man-getting-haircut:",
	},
	"\U0001F487  :haircut:": {
		Name:    "Haircut",
		Unicode: "\U0001F487",
		Alias:   ":haircut:",
	},
	"\U0001F488  :barber:": {
		Name:    "Barber Pole",
		Unicode: "\U0001F488",
		Alias:   ":barber:",
	},
	"\U0001F489  :syringe:": {
		Name:    "Syringe",
		Unicode: "\U0001F489",
		Alias:   ":syringe:",
	},
	"\U0001F48A  :pill:": {
		Name:    "Pill",
		Unicode: "\U0001F48A",
		Alias:   ":pill:",
	},
	"\U0001F48B  :kiss:": {
		Name:    "Kiss Mark",
		Unicode: "\U0001F48B",
		Alias:   ":kiss:",
	},
	"\U0001F48C  :love_letter:": {
		Name:    "Love Letter",
		Unicode: "\U0001F48C",
		Alias:   ":love_letter:",
	},
	"\U0001F48D  :ring:": {
		Name:    "Ring",
		Unicode: "\U0001F48D",
		Alias:   ":ring:",
	},
	"\U0001F48E  :gem:": {
		Name:    "Gem Stone",
		Unicode: "\U0001F48E",
		Alias:   ":gem:",
	},
	"\U0001F48F  :couplekiss:": {
		Name:    "Kiss",
		Unicode: "\U0001F48F",
		Alias:   ":couplekiss:",
	},
	"\U0001F490  :bouquet:": {
		Name:    "Bouquet",
		Unicode: "\U0001F490",
		Alias:   ":bouquet:",
	},
	"\U0001F491  :couple_with_heart:": {
		Name:    "Couple With Heart",
		Unicode: "\U0001F491",
		Alias:   ":couple_with_heart:",
	},
	"\U0001F492  :wedding:": {
		Name:    "Wedding",
		Unicode: "\U0001F492",
		Alias:   ":wedding:",
	},
	"\U0001F493  :heartbeat:": {
		Name:    "Beating Heart",
		Unicode: "\U0001F493",
		Alias:   ":heartbeat:",
	},
	"\U0001F494  :broken_heart:": {
		Name:    "Broken Heart",
		Unicode: "\U0001F494",
		Alias:   ":broken_heart:",
	},
	"\U0001F495  :two_hearts:": {
		Name:    "Two Hearts",
		Unicode: "\U0001F495",
		Alias:   ":two_hearts:",
	},
	"\U0001F496  :sparkling_heart:": {
		Name:    "Sparkling Heart",
		Unicode: "\U0001F496",
		Alias:   ":sparkling_heart:",
	},
	"\U0001F497  :heartpulse:": {
		Name:    "Growing Heart",
		Unicode: "\U0001F497",
		Alias:   ":heartpulse:",
	},
	"\U0001F498  :cupid:": {
		Name:    "Heart With Arrow",
		Unicode: "\U0001F498",
		Alias:   ":cupid:",
	},
	"\U0001F499  :blue_heart:": {
		Name:    "Blue Heart",
		Unicode: "\U0001F499",
		Alias:   ":blue_heart:",
	},
	"\U0001F49A  :green_heart:": {
		Name:    "Green Heart",
		Unicode: "\U0001F49A",
		Alias:   ":green_heart:",
	},
	"\U0001F49B  :yellow_heart:": {
		Name:    "Yellow Heart",
		Unicode: "\U0001F49B",
		Alias:   ":yellow_heart:",
	},
	"\U0001F49C  :purple_heart:": {
		Name:    "Purple Heart",
		Unicode: "\U0001F49C",
		Alias:   ":purple_heart:",
	},
	"\U0001F49D  :gift_heart:": {
		Name:    "Heart With Ribbon",
		Unicode: "\U0001F49D",
		Alias:   ":gift_heart:",
	},
	"\U0001F49E  :revolving_hearts:": {
		Name:    "Revolving Hearts",
		Unicode: "\U0001F49E",
		Alias:   ":revolving_hearts:",
	},
	"\U0001F49F  :heart_decoration:": {
		Name:    "Heart Decoration",
		Unicode: "\U0001F49F",
		Alias:   ":heart_decoration:",
	},
	"\U0001F4A0  :diamond_shape_with_a_dot_inside:": {
		Name:    "Diamond Shape With A Dot Inside",
		Unicode: "\U0001F4A0",
		Alias:   ":diamond_shape_with_a_dot_inside:",
	},
	"\U0001F4A1  :bulb:": {
		Name:    "Electric Light Bulb",
		Unicode: "\U0001F4A1",
		Alias:   ":bulb:",
	},
	"\U0001F4A2  :anger:": {
		Name:    "Anger Symbol",
		Unicode: "\U0001F4A2",
		Alias:   ":anger:",
	},
	"\U0001F4A3  :bomb:": {
		Name:    "Bomb",
		Unicode: "\U0001F4A3",
		Alias:   ":bomb:",
	},
	"\U0001F4A4  :zzz:": {
		Name:    "Sleeping Symbol",
		Unicode: "\U0001F4A4",
		Alias:   ":zzz:",
	},
	"\U0001F4A5  :boom:": {
		Name:    "Collision Symbol",
		Unicode: "\U0001F4A5",
		Alias:   ":boom:",
	},
	"\U0001F4A6  :sweat_drops:": {
		Name:    "Splashing Sweat Symbol",
		Unicode: "\U0001F4A6",
		Alias:   ":sweat_drops:",
	},
	"\U0001F4A7  :droplet:": {
		Name:    "Droplet",
		Unicode: "\U0001F4A7",
		Alias:   ":droplet:",
	},
	"\U0001F4A8  :dash:": {
		Name:    "Dash Symbol",
		Unicode: "\U0001F4A8",
		Alias:   ":dash:",
	},
	"\U0001F4A9  :hankey:": {
		Name:    "Pile Of Poo",
		Unicode: "\U0001F4A9",
		Alias:   ":hankey:",
	},
	"\U0001F4AA  :muscle:": {
		Name:    "Flexed Biceps",
		Unicode: "\U0001F4AA",
		Alias:   ":muscle:",
	},
	"\U0001F4AB  :dizzy:": {
		Name:    "Dizzy Symbol",
		Unicode: "\U0001F4AB",
		Alias:   ":dizzy:",
	},
	"\U0001F4AC  :speech_balloon:": {
		Name:    "Speech Balloon",
		Unicode: "\U0001F4AC",
		Alias:   ":speech_balloon:",
	},
	"\U0001F4AD  :thought_balloon:": {
		Name:    "Thought Balloon",
		Unicode: "\U0001F4AD",
		Alias:   ":thought_balloon:",
	},
	"\U0001F4AE  :white_flower:": {
		Name:    "White Flower",
		Unicode: "\U0001F4AE",
		Alias:   ":white_flower:",
	},
	"\U0001F4AF  :100:": {
		Name:    "Hundred Points Symbol",
		Unicode: "\U0001F4AF",
		Alias:   ":100:",
	},
	"\U0001F4B0  :moneybag:": {
		Name:    "Money Bag",
		Unicode: "\U0001F4B0",
		Alias:   ":moneybag:",
	},
	"\U0001F4B1  :currency_exchange:": {
		Name:    "Currency Exchange",
		Unicode: "\U0001F4B1",
		Alias:   ":currency_exchange:",
	},
	"\U0001F4B2  :heavy_dollar_sign:": {
		Name:    "Heavy Dollar Sign",
		Unicode: "\U0001F4B2",
		Alias:   ":heavy_dollar_sign:",
	},
	"\U0001F4B3  :credit_card:": {
		Name:    "Credit Card",
		Unicode: "\U0001F4B3",
		Alias:   ":credit_card:",
	},
	"\U0001F4B4  :yen:": {
		Name:    "Banknote With Yen Sign",
		Unicode: "\U0001F4B4",
		Alias:   ":yen:",
	},
	"\U0001F4B5  :dollar:": {
		Name:    "Banknote With Dollar Sign",
		Unicode: "\U0001F4B5",
		Alias:   ":dollar:",
	},
	"\U0001F4B6  :euro:": {
		Name:    "Banknote With Euro Sign",
		Unicode: "\U0001F4B6",
		Alias:   ":euro:",
	},
	"\U0001F4B7  :pound:": {
		Name:    "Banknote With Pound Sign",
		Unicode: "\U0001F4B7",
		Alias:   ":pound:",
	},
	"\U0001F4B8  :money_with_wings:": {
		Name:    "Money With Wings",
		Unicode: "\U0001F4B8",
		Alias:   ":money_with_wings:",
	},
	"\U0001F4B9  :chart:": {
		Name:    "Chart With Upwards Trend And Yen Sign",
		Unicode: "\U0001F4B9",
		Alias:   ":chart:",
	},
	"\U0001F4BA  :seat:": {
		Name:    "Seat",
		Unicode: "\U0001F4BA",
		Alias:   ":seat:",
	},
	"\U0001F4BB  :computer:": {
		Name:    "Personal Computer",
		Unicode: "\U0001F4BB",
		Alias:   ":computer:",
	},
	"\U0001F4BC  :briefcase:": {
		Name:    "Briefcase",
		Unicode: "\U0001F4BC",
		Alias:   ":briefcase:",
	},
	"\U0001F4BD  :minidisc:": {
		Name:    "Minidisc",
		Unicode: "\U0001F4BD",
		Alias:   ":minidisc:",
	},
	"\U0001F4BE  :floppy_disk:": {
		Name:    "Floppy Disk",
		Unicode: "\U0001F4BE",
		Alias:   ":floppy_disk:",
	},
	"\U0001F4BF  :cd:": {
		Name:    "Optical Disc",
		Unicode: "\U0001F4BF",
		Alias:   ":cd:",
	},
	"\U0001F4C0  :dvd:": {
		Name:    "Dvd",
		Unicode: "\U0001F4C0",
		Alias:   ":dvd:",
	},
	"\U0001F4C1  :file_folder:": {
		Name:    "File Folder",
		Unicode: "\U0001F4C1",
		Alias:   ":file_folder:",
	},
	"\U0001F4C2  :open_file_folder:": {
		Name:    "Open File Folder",
		Unicode: "\U0001F4C2",
		Alias:   ":open_file_folder:",
	},
	"\U0001F4C3  :page_with_curl:": {
		Name:    "Page With Curl",
		Unicode: "\U0001F4C3",
		Alias:   ":page_with_curl:",
	},
	"\U0001F4C4  :page_facing_up:": {
		Name:    "Page Facing Up",
		Unicode: "\U0001F4C4",
		Alias:   ":page_facing_up:",
	},
	"\U0001F4C5  :date:": {
		Name:    "Calendar",
		Unicode: "\U0001F4C5",
		Alias:   ":date:",
	},
	"\U0001F4C6  :calendar:": {
		Name:    "Tear-Off Calendar",
		Unicode: "\U0001F4C6",
		Alias:   ":calendar:",
	},
	"\U0001F4C7  :card_index:": {
		Name:    "Card Index",
		Unicode: "\U0001F4C7",
		Alias:   ":card_index:",
	},
	"\U0001F4C8  :chart_with_upwards_trend:": {
		Name:    "Chart With Upwards Trend",
		Unicode: "\U0001F4C8",
		Alias:   ":chart_with_upwards_trend:",
	},
	"\U0001F4C9  :chart_with_downwards_trend:": {
		Name:    "Chart With Downwards Trend",
		Unicode: "\U0001F4C9",
		Alias:   ":chart_with_downwards_trend:",
	},
	"\U0001F4CA  :bar_chart:": {
		Name:    "Bar Chart",
		Unicode: "\U0001F4CA",
		Alias:   ":bar_chart:",
	},
	"\U0001F4CB  :clipboard:": {
		Name:    "Clipboard",
		Unicode: "\U0001F4CB",
		Alias:   ":clipboard:",
	},
	"\U0001F4CC  :pushpin:": {
		Name:    "Pushpin",
		Unicode: "\U0001F4CC",
		Alias:   ":pushpin:",
	},
	"\U0001F4CD  :round_pushpin:": {
		Name:    "Round Pushpin",
		Unicode: "\U0001F4CD",
		Alias:   ":round_pushpin:",
	},
	"\U0001F4CE  :paperclip:": {
		Name:    "Paperclip",
		Unicode: "\U0001F4CE",
		Alias:   ":paperclip:",
	},
	"\U0001F4CF  :straight_ruler:": {
		Name:    "Straight Ruler",
		Unicode: "\U0001F4CF",
		Alias:   ":straight_ruler:",
	},
	"\U0001F4D0  :triangular_ruler:": {
		Name:    "Triangular Ruler",
		Unicode: "\U0001F4D0",
		Alias:   ":triangular_ruler:",
	},
	"\U0001F4D1  :bookmark_tabs:": {
		Name:    "Bookmark Tabs",
		Unicode: "\U0001F4D1",
		Alias:   ":bookmark_tabs:",
	},
	"\U0001F4D2  :ledger:": {
		Name:    "Ledger",
		Unicode: "\U0001F4D2",
		Alias:   ":ledger:",
	},
	"\U0001F4D3  :notebook:": {
		Name:    "Notebook",
		Unicode: "\U0001F4D3",
		Alias:   ":notebook:",
	},
	"\U0001F4D4  :notebook_with_decorative_cover:": {
		Name:    "Notebook With Decorative Cover",
		Unicode: "\U0001F4D4",
		Alias:   ":notebook_with_decorative_cover:",
	},
	"\U0001F4D5  :closed_book:": {
		Name:    "Closed Book",
		Unicode: "\U0001F4D5",
		Alias:   ":closed_book:",
	},
	"\U0001F4D6  :book:": {
		Name:    "Open Book",
		Unicode: "\U0001F4D6",
		Alias:   ":book:",
	},
	"\U0001F4D7  :green_book:": {
		Name:    "Green Book",
		Unicode: "\U0001F4D7",
		Alias:   ":green_book:",
	},
	"\U0001F4D8  :blue_book:": {
		Name:    "Blue Book",
		Unicode: "\U0001F4D8",
		Alias:   ":blue_book:",
	},
	"\U0001F4D9  :orange_book:": {
		Name:    "Orange Book",
		Unicode: "\U0001F4D9",
		Alias:   ":orange_book:",
	},
	"\U0001F4DA  :books:": {
		Name:    "Books",
		Unicode: "\U0001F4DA",
		Alias:   ":books:",
	},
	"\U0001F4DB  :name_badge:": {
		Name:    "Name Badge",
		Unicode: "\U0001F4DB",
		Alias:   ":name_badge:",
	},
	"\U0001F4DC  :scroll:": {
		Name:    "Scroll",
		Unicode: "\U0001F4DC",
		Alias:   ":scroll:",
	},
	"\U0001F4DD  :memo:": {
		Name:    "Memo",
		Unicode: "\U0001F4DD",
		Alias:   ":memo:",
	},
	"\U0001F4DE  :telephone_receiver:": {
		Name:    "Telephone Receiver",
		Unicode: "\U0001F4DE",
		Alias:   ":telephone_receiver:",
	},
	"\U0001F4DF  :pager:": {
		Name:    "Pager",
		Unicode: "\U0001F4DF",
		Alias:   ":pager:",
	},
	"\U0001F4E0  :fax:": {
		Name:    "Fax Machine",
		Unicode: "\U0001F4E0",
		Alias:   ":fax:",
	},
	"\U0001F4E1  :satellite_antenna:": {
		Name:    "Satellite Antenna",
		Unicode: "\U0001F4E1",
		Alias:   ":satellite_antenna:",
	},
	"\U0001F4E2  :loudspeaker:": {
		Name:    "Public Address Loudspeaker",
		Unicode: "\U0001F4E2",
		Alias:   ":loudspeaker:",
	},
	"\U0001F4E3  :mega:": {
		Name:    "Cheering Megaphone",
		Unicode: "\U0001F4E3",
		Alias:   ":mega:",
	},
	"\U0001F4E4  :outbox_tray:": {
		Name:    "Outbox Tray",
		Unicode: "\U0001F4E4",
		Alias:   ":outbox_tray:",
	},
	"\U0001F4E5  :inbox_tray:": {
		Name:    "Inbox Tray",
		Unicode: "\U0001F4E5",
		Alias:   ":inbox_tray:",
	},
	"\U0001F4E6  :package:": {
		Name:    "Package",
		Unicode: "\U0001F4E6",
		Alias:   ":package:",
	},
	"\U0001F4E7 :e-mail:": {
		Name:    "E-Mail Symbol",
		Unicode: "\U0001F4E7",
		Alias:   ":e-mail:",
	},
	"\U0001F4E8  :incoming_envelope:": {
		Name:    "Incoming Envelope",
		Unicode: "\U0001F4E8",
		Alias:   ":incoming_envelope:",
	},
	"\U0001F4E9  :envelope_with_arrow:": {
		Name:    "Envelope With Downwards Arrow Above",
		Unicode: "\U0001F4E9",
		Alias:   ":envelope_with_arrow:",
	},
	"\U0001F4EA  :mailbox_closed:": {
		Name:    "Closed Mailbox With Lowered Flag",
		Unicode: "\U0001F4EA",
		Alias:   ":mailbox_closed:",
	},
	"\U0001F4EB  :mailbox:": {
		Name:    "Closed Mailbox With Raised Flag",
		Unicode: "\U0001F4EB",
		Alias:   ":mailbox:",
	},
	"\U0001F4EC  :mailbox_with_mail:": {
		Name:    "Open Mailbox With Raised Flag",
		Unicode: "\U0001F4EC",
		Alias:   ":mailbox_with_mail:",
	},
	"\U0001F4ED  :mailbox_with_no_mail:": {
		Name:    "Open Mailbox With Lowered Flag",
		Unicode: "\U0001F4ED",
		Alias:   ":mailbox_with_no_mail:",
	},
	"\U0001F4EE  :postbox:": {
		Name:    "Postbox",
		Unicode: "\U0001F4EE",
		Alias:   ":postbox:",
	},
	"\U0001F4EF  :postal_horn:": {
		Name:    "Postal Horn",
		Unicode: "\U0001F4EF",
		Alias:   ":postal_horn:",
	},
	"\U0001F4F0  :newspaper:": {
		Name:    "Newspaper",
		Unicode: "\U0001F4F0",
		Alias:   ":newspaper:",
	},
	"\U0001F4F1  :iphone:": {
		Name:    "Mobile Phone",
		Unicode: "\U0001F4F1",
		Alias:   ":iphone:",
	},
	"\U0001F4F2  :calling:": {
		Name:    "Mobile Phone With Rightwards Arrow At Left",
		Unicode: "\U0001F4F2",
		Alias:   ":calling:",
	},
	"\U0001F4F3  :vibration_mode:": {
		Name:    "Vibration Mode",
		Unicode: "\U0001F4F3",
		Alias:   ":vibration_mode:",
	},
	"\U0001F4F4  :mobile_phone_off:": {
		Name:    "Mobile Phone Off",
		Unicode: "\U0001F4F4",
		Alias:   ":mobile_phone_off:",
	},
	"\U0001F4F5  :no_mobile_phones:": {
		Name:    "No Mobile Phones",
		Unicode: "\U0001F4F5",
		Alias:   ":no_mobile_phones:",
	},
	"\U0001F4F6  :signal_strength:": {
		Name:    "Antenna With Bars",
		Unicode: "\U0001F4F6",
		Alias:   ":signal_strength:",
	},
	"\U0001F4F7  :camera:": {
		Name:    "Camera",
		Unicode: "\U0001F4F7",
		Alias:   ":camera:",
	},
	"\U0001F4F8  :camera_with_flash:": {
		Name:    "Camera With Flash",
		Unicode: "\U0001F4F8",
		Alias:   ":camera_with_flash:",
	},
	"\U0001F4F9  :video_camera:": {
		Name:    "Video Camera",
		Unicode: "\U0001F4F9",
		Alias:   ":video_camera:",
	},
	"\U0001F4FA  :tv:": {
		Name:    "Television",
		Unicode: "\U0001F4FA",
		Alias:   ":tv:",
	},
	"\U0001F4FB  :radio:": {
		Name:    "Radio",
		Unicode: "\U0001F4FB",
		Alias:   ":radio:",
	},
	"\U0001F4FC  :vhs:": {
		Name:    "Videocassette",
		Unicode: "\U0001F4FC",
		Alias:   ":vhs:",
	},
	"\U0001F4FD\uFE0F  :film_projector:": {
		Name:    "Film Projector",
		Unicode: "\U0001F4FD\uFE0F",
		Alias:   ":film_projector:",
	},
	"\U0001F4FF  :prayer_beads:": {
		Name:    "Prayer Beads",
		Unicode: "\U0001F4FF",
		Alias:   ":prayer_beads:",
	},
	"\U0001F500  :twisted_rightwards_arrows:": {
		Name:    "Twisted Rightwards Arrows",
		Unicode: "\U0001F500",
		Alias:   ":twisted_rightwards_arrows:",
	},
	"\U0001F501  :repeat:": {
		Name:    "Clockwise Rightwards And Leftwards Open Circle Arrows",
		Unicode: "\U0001F501",
		Alias:   ":repeat:",
	},
	"\U0001F502  :repeat_one:": {
		Name:    "Clockwise Rightwards And Leftwards Open Circle Arrows With Circled One Overlay",
		Unicode: "\U0001F502",
		Alias:   ":repeat_one:",
	},
	"\U0001F503  :arrows_clockwise:": {
		Name:    "Clockwise Downwards And Upwards Open Circle Arrows",
		Unicode: "\U0001F503",
		Alias:   ":arrows_clockwise:",
	},
	"\U0001F504  :arrows_counterclockwise:": {
		Name:    "Anticlockwise Downwards And Upwards Open Circle Arrows",
		Unicode: "\U0001F504",
		Alias:   ":arrows_counterclockwise:",
	},
	"\U0001F505  :low_brightness:": {
		Name:    "Low Brightness Symbol",
		Unicode: "\U0001F505",
		Alias:   ":low_brightness:",
	},
	"\U0001F506  :high_brightness:": {
		Name:    "High Brightness Symbol",
		Unicode: "\U0001F506",
		Alias:   ":high_brightness:",
	},
	"\U0001F507  :mute:": {
		Name:    "Speaker With Cancellation Stroke",
		Unicode: "\U0001F507",
		Alias:   ":mute:",
	},
	"\U0001F508  :speaker:": {
		Name:    "Speaker",
		Unicode: "\U0001F508",
		Alias:   ":speaker:",
	},
	"\U0001F509  :sound:": {
		Name:    "Speaker With One Sound Wave",
		Unicode: "\U0001F509",
		Alias:   ":sound:",
	},
	"\U0001F50A  :loud_sound:": {
		Name:    "Speaker With Three Sound Waves",
		Unicode: "\U0001F50A",
		Alias:   ":loud_sound:",
	},
	"\U0001F50B  :battery:": {
		Name:    "Battery",
		Unicode: "\U0001F50B",
		Alias:   ":battery:",
	},
	"\U0001F50C  :electric_plug:": {
		Name:    "Electric Plug",
		Unicode: "\U0001F50C",
		Alias:   ":electric_plug:",
	},
	"\U0001F50D  :mag:": {
		Name:    "Left-Pointing Magnifying Glass",
		Unicode: "\U0001F50D",
		Alias:   ":mag:",
	},
	"\U0001F50E  :mag_right:": {
		Name:    "Right-Pointing Magnifying Glass",
		Unicode: "\U0001F50E",
		Alias:   ":mag_right:",
	},
	"\U0001F50F  :lock_with_ink_pen:": {
		Name:    "Lock With Ink Pen",
		Unicode: "\U0001F50F",
		Alias:   ":lock_with_ink_pen:",
	},
	"\U0001F510  :closed_lock_with_key:": {
		Name:    "Closed Lock With Key",
		Unicode: "\U0001F510",
		Alias:   ":closed_lock_with_key:",
	},
	"\U0001F511  :key:": {
		Name:    "Key",
		Unicode: "\U0001F511",
		Alias:   ":key:",
	},
	"\U0001F512  :lock:": {
		Name:    "Lock",
		Unicode: "\U0001F512",
		Alias:   ":lock:",
	},
	"\U0001F513  :unlock:": {
		Name:    "Open Lock",
		Unicode: "\U0001F513",
		Alias:   ":unlock:",
	},
	"\U0001F514  :bell:": {
		Name:    "Bell",
		Unicode: "\U0001F514",
		Alias:   ":bell:",
	},
	"\U0001F515  :no_bell:": {
		Name:    "Bell With Cancellation Stroke",
		Unicode: "\U0001F515",
		Alias:   ":no_bell:",
	},
	"\U0001F516  :bookmark:": {
		Name:    "Bookmark",
		Unicode: "\U0001F516",
		Alias:   ":bookmark:",
	},
	"\U0001F517  :link:": {
		Name:    "Link Symbol",
		Unicode: "\U0001F517",
		Alias:   ":link:",
	},
	"\U0001F518  :radio_button:": {
		Name:    "Radio Button",
		Unicode: "\U0001F518",
		Alias:   ":radio_button:",
	},
	"\U0001F519  :back:": {
		Name:    "Back With Leftwards Arrow Above",
		Unicode: "\U0001F519",
		Alias:   ":back:",
	},
	"\U0001F51A  :end:": {
		Name:    "End With Leftwards Arrow Above",
		Unicode: "\U0001F51A",
		Alias:   ":end:",
	},
	"\U0001F51B  :on:": {
		Name:    "On With Exclamation Mark With Left Right Arrow Above",
		Unicode: "\U0001F51B",
		Alias:   ":on:",
	},
	"\U0001F51C  :soon:": {
		Name:    "Soon With Rightwards Arrow Above",
		Unicode: "\U0001F51C",
		Alias:   ":soon:",
	},
	"\U0001F51D  :top:": {
		Name:    "Top With Upwards Arrow Above",
		Unicode: "\U0001F51D",
		Alias:   ":top:",
	},
	"\U0001F51E  :underage:": {
		Name:    "No One Under Eighteen Symbol",
		Unicode: "\U0001F51E",
		Alias:   ":underage:",
	},
	"\U0001F51F  :keycap_ten:": {
		Name:    "Keycap Ten",
		Unicode: "\U0001F51F",
		Alias:   ":keycap_ten:",
	},
	"\U0001F520  :capital_abcd:": {
		Name:    "Input Symbol For Latin Capital Letters",
		Unicode: "\U0001F520",
		Alias:   ":capital_abcd:",
	},
	"\U0001F521  :abcd:": {
		Name:    "Input Symbol For Latin Small Letters",
		Unicode: "\U0001F521",
		Alias:   ":abcd:",
	},
	"\U0001F522  :1234:": {
		Name:    "Input Symbol For Numbers",
		Unicode: "\U0001F522",
		Alias:   ":1234:",
	},
	"\U0001F523  :symbols:": {
		Name:    "Input Symbol For Symbols",
		Unicode: "\U0001F523",
		Alias:   ":symbols:",
	},
	"\U0001F524  :abc:": {
		Name:    "Input Symbol For Latin Letters",
		Unicode: "\U0001F524",
		Alias:   ":abc:",
	},
	"\U0001F525  :fire:": {
		Name:    "Fire",
		Unicode: "\U0001F525",
		Alias:   ":fire:",
	},
	"\U0001F526  :flashlight:": {
		Name:    "Electric Torch",
		Unicode: "\U0001F526",
		Alias:   ":flashlight:",
	},
	"\U0001F527  :wrench:": {
		Name:    "Wrench",
		Unicode: "\U0001F527",
		Alias:   ":wrench:",
	},
	"\U0001F528  :hammer:": {
		Name:    "Hammer",
		Unicode: "\U0001F528",
		Alias:   ":hammer:",
	},
	"\U0001F529  :nut_and_bolt:": {
		Name:    "Nut And Bolt",
		Unicode: "\U0001F529",
		Alias:   ":nut_and_bolt:",
	},
	"\U0001F52A  :hocho:": {
		Name:    "Hocho",
		Unicode: "\U0001F52A",
		Alias:   ":hocho:",
	},
	"\U0001F52B  :gun:": {
		Name:    "Pistol",
		Unicode: "\U0001F52B",
		Alias:   ":gun:",
	},
	"\U0001F52C  :microscope:": {
		Name:    "Microscope",
		Unicode: "\U0001F52C",
		Alias:   ":microscope:",
	},
	"\U0001F52D  :telescope:": {
		Name:    "Telescope",
		Unicode: "\U0001F52D",
		Alias:   ":telescope:",
	},
	"\U0001F52E  :crystal_ball:": {
		Name:    "Crystal Ball",
		Unicode: "\U0001F52E",
		Alias:   ":crystal_ball:",
	},
	"\U0001F52F  :six_pointed_star:": {
		Name:    "Six Pointed Star With Middle Dot",
		Unicode: "\U0001F52F",
		Alias:   ":six_pointed_star:",
	},
	"\U0001F530  :beginner:": {
		Name:    "Japanese Symbol For Beginner",
		Unicode: "\U0001F530",
		Alias:   ":beginner:",
	},
	"\U0001F531  :trident:": {
		Name:    "Trident Emblem",
		Unicode: "\U0001F531",
		Alias:   ":trident:",
	},
	"\U0001F532  :black_square_button:": {
		Name:    "Black Square Button",
		Unicode: "\U0001F532",
		Alias:   ":black_square_button:",
	},
	"\U0001F533  :white_square_button:": {
		Name:    "White Square Button",
		Unicode: "\U0001F533",
		Alias:   ":white_square_button:",
	},
	"\U0001F534  :red_circle:": {
		Name:    "Large Red Circle",
		Unicode: "\U0001F534",
		Alias:   ":red_circle:",
	},
	"\U0001F535  :large_blue_circle:": {
		Name:    "Large Blue Circle",
		Unicode: "\U0001F535",
		Alias:   ":large_blue_circle:",
	},
	"\U0001F536  :large_orange_diamond:": {
		Name:    "Large Orange Diamond",
		Unicode: "\U0001F536",
		Alias:   ":large_orange_diamond:",
	},
	"\U0001F537  :large_blue_diamond:": {
		Name:    "Large Blue Diamond",
		Unicode: "\U0001F537",
		Alias:   ":large_blue_diamond:",
	},
	"\U0001F538  :small_orange_diamond:": {
		Name:    "Small Orange Diamond",
		Unicode: "\U0001F538",
		Alias:   ":small_orange_diamond:",
	},
	"\U0001F539  :small_blue_diamond:": {
		Name:    "Small Blue Diamond",
		Unicode: "\U0001F539",
		Alias:   ":small_blue_diamond:",
	},
	"\U0001F53A  :small_red_triangle:": {
		Name:    "Up-Pointing Red Triangle",
		Unicode: "\U0001F53A",
		Alias:   ":small_red_triangle:",
	},
	"\U0001F53B  :small_red_triangle_down:": {
		Name:    "Down-Pointing Red Triangle",
		Unicode: "\U0001F53B",
		Alias:   ":small_red_triangle_down:",
	},
	"\U0001F53C  :arrow_up_small:": {
		Name:    "Up-Pointing Small Red Triangle",
		Unicode: "\U0001F53C",
		Alias:   ":arrow_up_small:",
	},
	"\U0001F53D  :arrow_down_small:": {
		Name:    "Down-Pointing Small Red Triangle",
		Unicode: "\U0001F53D",
		Alias:   ":arrow_down_small:",
	},
	"\U0001F549\uFE0F  :om_symbol:": {
		Name:    "Om",
		Unicode: "\U0001F549\uFE0F",
		Alias:   ":om_symbol:",
	},
	"\U0001F54A\uFE0F  :dove_of_peace:": {
		Name:    "Dove",
		Unicode: "\U0001F54A\uFE0F",
		Alias:   ":dove_of_peace:",
	},
	"\U0001F54B  :kaaba:": {
		Name:    "Kaaba",
		Unicode: "\U0001F54B",
		Alias:   ":kaaba:",
	},
	"\U0001F54C  :mosque:": {
		Name:    "Mosque",
		Unicode: "\U0001F54C",
		Alias:   ":mosque:",
	},
	"\U0001F54D  :synagogue:": {
		Name:    "Synagogue",
		Unicode: "\U0001F54D",
		Alias:   ":synagogue:",
	},
	"\U0001F54E  :menorah_with_nine_branches:": {
		Name:    "Menorah With Nine Branches",
		Unicode: "\U0001F54E",
		Alias:   ":menorah_with_nine_branches:",
	},
	"\U0001F550  :clock1:": {
		Name:    "Clock Face One Oclock",
		Unicode: "\U0001F550",
		Alias:   ":clock1:",
	},
	"\U0001F551  :clock2:": {
		Name:    "Clock Face Two Oclock",
		Unicode: "\U0001F551",
		Alias:   ":clock2:",
	},
	"\U0001F552  :clock3:": {
		Name:    "Clock Face Three Oclock",
		Unicode: "\U0001F552",
		Alias:   ":clock3:",
	},
	"\U0001F553  :clock4:": {
		Name:    "Clock Face Four Oclock",
		Unicode: "\U0001F553",
		Alias:   ":clock4:",
	},
	"\U0001F554  :clock5:": {
		Name:    "Clock Face Five Oclock",
		Unicode: "\U0001F554",
		Alias:   ":clock5:",
	},
	"\U0001F555  :clock6:": {
		Name:    "Clock Face Six Oclock",
		Unicode: "\U0001F555",
		Alias:   ":clock6:",
	},
	"\U0001F556  :clock7:": {
		Name:    "Clock Face Seven Oclock",
		Unicode: "\U0001F556",
		Alias:   ":clock7:",
	},
	"\U0001F557  :clock8:": {
		Name:    "Clock Face Eight Oclock",
		Unicode: "\U0001F557",
		Alias:   ":clock8:",
	},
	"\U0001F558  :clock9:": {
		Name:    "Clock Face Nine Oclock",
		Unicode: "\U0001F558",
		Alias:   ":clock9:",
	},
	"\U0001F559  :clock10:": {
		Name:    "Clock Face Ten Oclock",
		Unicode: "\U0001F559",
		Alias:   ":clock10:",
	},
	"\U0001F55A  :clock11:": {
		Name:    "Clock Face Eleven Oclock",
		Unicode: "\U0001F55A",
		Alias:   ":clock11:",
	},
	"\U0001F55B  :clock12:": {
		Name:    "Clock Face Twelve Oclock",
		Unicode: "\U0001F55B",
		Alias:   ":clock12:",
	},
	"\U0001F55C  :clock130:": {
		Name:    "Clock Face One-Thirty",
		Unicode: "\U0001F55C",
		Alias:   ":clock130:",
	},
	"\U0001F55D  :clock230:": {
		Name:    "Clock Face Two-Thirty",
		Unicode: "\U0001F55D",
		Alias:   ":clock230:",
	},
	"\U0001F55E  :clock330:": {
		Name:    "Clock Face Three-Thirty",
		Unicode: "\U0001F55E",
		Alias:   ":clock330:",
	},
	"\U0001F55F  :clock430:": {
		Name:    "Clock Face Four-Thirty",
		Unicode: "\U0001F55F",
		Alias:   ":clock430:",
	},
	"\U0001F560  :clock530:": {
		Name:    "Clock Face Five-Thirty",
		Unicode: "\U0001F560",
		Alias:   ":clock530:",
	},
	"\U0001F561  :clock630:": {
		Name:    "Clock Face Six-Thirty",
		Unicode: "\U0001F561",
		Alias:   ":clock630:",
	},
	"\U0001F562  :clock730:": {
		Name:    "Clock Face Seven-Thirty",
		Unicode: "\U0001F562",
		Alias:   ":clock730:",
	},
	"\U0001F563  :clock830:": {
		Name:    "Clock Face Eight-Thirty",
		Unicode: "\U0001F563",
		Alias:   ":clock830:",
	},
	"\U0001F564  :clock930:": {
		Name:    "Clock Face Nine-Thirty",
		Unicode: "\U0001F564",
		Alias:   ":clock930:",
	},
	"\U0001F565  :clock1030:": {
		Name:    "Clock Face Ten-Thirty",
		Unicode: "\U0001F565",
		Alias:   ":clock1030:",
	},
	"\U0001F566  :clock1130:": {
		Name:    "Clock Face Eleven-Thirty",
		Unicode: "\U0001F566",
		Alias:   ":clock1130:",
	},
	"\U0001F567  :clock1230:": {
		Name:    "Clock Face Twelve-Thirty",
		Unicode: "\U0001F567",
		Alias:   ":clock1230:",
	},
	"\U0001F56F\uFE0F  :candle:": {
		Name:    "Candle",
		Unicode: "\U0001F56F\uFE0F",
		Alias:   ":candle:",
	},
	"\U0001F570\uFE0F  :mantelpiece_clock:": {
		Name:    "Mantelpiece Clock",
		Unicode: "\U0001F570\uFE0F",
		Alias:   ":mantelpiece_clock:",
	},
	"\U0001F573\uFE0F  :hole:": {
		Name:    "Hole",
		Unicode: "\U0001F573\uFE0F",
		Alias:   ":hole:",
	},
	"\U0001F574\uFE0F  :man_in_business_suit_levitating:": {
		Name:    "Person In Suit Levitating",
		Unicode: "\U0001F574\uFE0F",
		Alias:   ":man_in_business_suit_levitating:",
	},
	"\U0001F575\uFE0F\u200D\u2640\uFE0F :female-detective:": {
		Name:    "Woman Detective",
		Unicode: "\U0001F575\uFE0F\u200D\u2640\uFE0F",
		Alias:   ":female-detective:",
	},
	"\U0001F575\uFE0F\u200D\u2642\uFE0F :male-detective:": {
		Name:    "Man Detective",
		Unicode: "\U0001F575\uFE0F\u200D\u2642\uFE0F",
		Alias:   ":male-detective:",
	},
	"\U0001F575\uFE0F  :sleuth_or_spy:": {
		Name:    "Detective",
		Unicode: "\U0001F575\uFE0F",
		Alias:   ":sleuth_or_spy:",
	},
	"\U0001F576\uFE0F  :dark_sunglasses:": {
		Name:    "Sunglasses",
		Unicode: "\U0001F576\uFE0F",
		Alias:   ":dark_sunglasses:",
	},
	"\U0001F577\uFE0F  :spider:": {
		Name:    "Spider",
		Unicode: "\U0001F577\uFE0F",
		Alias:   ":spider:",
	},
	"\U0001F578\uFE0F  :spider_web:": {
		Name:    "Spider Web",
		Unicode: "\U0001F578\uFE0F",
		Alias:   ":spider_web:",
	},
	"\U0001F579\uFE0F  :joystick:": {
		Name:    "Joystick",
		Unicode: "\U0001F579\uFE0F",
		Alias:   ":joystick:",
	},
	"\U0001F57A  :man_dancing:": {
		Name:    "Man Dancing",
		Unicode: "\U0001F57A",
		Alias:   ":man_dancing:",
	},
	"\U0001F587\uFE0F  :linked_paperclips:": {
		Name:    "Linked Paperclips",
		Unicode: "\U0001F587\uFE0F",
		Alias:   ":linked_paperclips:",
	},
	"\U0001F58A\uFE0F  :lower_left_ballpoint_pen:": {
		Name:    "Pen",
		Unicode: "\U0001F58A\uFE0F",
		Alias:   ":lower_left_ballpoint_pen:",
	},
	"\U0001F58B\uFE0F  :lower_left_fountain_pen:": {
		Name:    "Fountain Pen",
		Unicode: "\U0001F58B\uFE0F",
		Alias:   ":lower_left_fountain_pen:",
	},
	"\U0001F58C\uFE0F  :lower_left_paintbrush:": {
		Name:    "Paintbrush",
		Unicode: "\U0001F58C\uFE0F",
		Alias:   ":lower_left_paintbrush:",
	},
	"\U0001F58D\uFE0F  :lower_left_crayon:": {
		Name:    "Crayon",
		Unicode: "\U0001F58D\uFE0F",
		Alias:   ":lower_left_crayon:",
	},
	"\U0001F590\uFE0F  :raised_hand_with_fingers_splayed:": {
		Name:    "Hand With Fingers Splayed",
		Unicode: "\U0001F590\uFE0F",
		Alias:   ":raised_hand_with_fingers_splayed:",
	},
	"\U0001F595  :middle_finger:": {
		Name:    "Reversed Hand With Middle Finger Extended",
		Unicode: "\U0001F595",
		Alias:   ":middle_finger:",
	},
	"\U0001F596 :spock-hand:": {
		Name:    "Raised Hand With Part Between Middle And Ring Fingers",
		Unicode: "\U0001F596",
		Alias:   ":spock-hand:",
	},
	"\U0001F5A4  :black_heart:": {
		Name:    "Black Heart",
		Unicode: "\U0001F5A4",
		Alias:   ":black_heart:",
	},
	"\U0001F5A5\uFE0F  :desktop_computer:": {
		Name:    "Desktop Computer",
		Unicode: "\U0001F5A5\uFE0F",
		Alias:   ":desktop_computer:",
	},
	"\U0001F5A8\uFE0F  :printer:": {
		Name:    "Printer",
		Unicode: "\U0001F5A8\uFE0F",
		Alias:   ":printer:",
	},
	"\U0001F5B1\uFE0F  :three_button_mouse:": {
		Name:    "Computer Mouse",
		Unicode: "\U0001F5B1\uFE0F",
		Alias:   ":three_button_mouse:",
	},
	"\U0001F5B2\uFE0F  :trackball:": {
		Name:    "Trackball",
		Unicode: "\U0001F5B2\uFE0F",
		Alias:   ":trackball:",
	},
	"\U0001F5BC\uFE0F  :frame_with_picture:": {
		Name:    "Framed Picture",
		Unicode: "\U0001F5BC\uFE0F",
		Alias:   ":frame_with_picture:",
	},
	"\U0001F5C2\uFE0F  :card_index_dividers:": {
		Name:    "Card Index Dividers",
		Unicode: "\U0001F5C2\uFE0F",
		Alias:   ":card_index_dividers:",
	},
	"\U0001F5C3\uFE0F  :card_file_box:": {
		Name:    "Card File Box",
		Unicode: "\U0001F5C3\uFE0F",
		Alias:   ":card_file_box:",
	},
	"\U0001F5C4\uFE0F  :file_cabinet:": {
		Name:    "File Cabinet",
		Unicode: "\U0001F5C4\uFE0F",
		Alias:   ":file_cabinet:",
	},
	"\U0001F5D1\uFE0F  :wastebasket:": {
		Name:    "Wastebasket",
		Unicode: "\U0001F5D1\uFE0F",
		Alias:   ":wastebasket:",
	},
	"\U0001F5D2\uFE0F  :spiral_note_pad:": {
		Name:    "Spiral Notepad",
		Unicode: "\U0001F5D2\uFE0F",
		Alias:   ":spiral_note_pad:",
	},
	"\U0001F5D3\uFE0F  :spiral_calendar_pad:": {
		Name:    "Spiral Calendar",
		Unicode: "\U0001F5D3\uFE0F",
		Alias:   ":spiral_calendar_pad:",
	},
	"\U0001F5DC\uFE0F  :compression:": {
		Name:    "Clamp",
		Unicode: "\U0001F5DC\uFE0F",
		Alias:   ":compression:",
	},
	"\U0001F5DD\uFE0F  :old_key:": {
		Name:    "Old Key",
		Unicode: "\U0001F5DD\uFE0F",
		Alias:   ":old_key:",
	},
	"\U0001F5DE\uFE0F  :rolled_up_newspaper:": {
		Name:    "Rolled-Up Newspaper",
		Unicode: "\U0001F5DE\uFE0F",
		Alias:   ":rolled_up_newspaper:",
	},
	"\U0001F5E1\uFE0F  :dagger_knife:": {
		Name:    "Dagger",
		Unicode: "\U0001F5E1\uFE0F",
		Alias:   ":dagger_knife:",
	},
	"\U0001F5E3\uFE0F  :speaking_head_in_silhouette:": {
		Name:    "Speaking Head",
		Unicode: "\U0001F5E3\uFE0F",
		Alias:   ":speaking_head_in_silhouette:",
	},
	"\U0001F5E8\uFE0F  :left_speech_bubble:": {
		Name:    "Left Speech Bubble",
		Unicode: "\U0001F5E8\uFE0F",
		Alias:   ":left_speech_bubble:",
	},
	"\U0001F5EF\uFE0F  :right_anger_bubble:": {
		Name:    "Right Anger Bubble",
		Unicode: "\U0001F5EF\uFE0F",
		Alias:   ":right_anger_bubble:",
	},
	"\U0001F5F3\uFE0F  :ballot_box_with_ballot:": {
		Name:    "Ballot Box With Ballot",
		Unicode: "\U0001F5F3\uFE0F",
		Alias:   ":ballot_box_with_ballot:",
	},
	"\U0001F5FA\uFE0F  :world_map:": {
		Name:    "World Map",
		Unicode: "\U0001F5FA\uFE0F",
		Alias:   ":world_map:",
	},
	"\U0001F5FB  :mount_fuji:": {
		Name:    "Mount Fuji",
		Unicode: "\U0001F5FB",
		Alias:   ":mount_fuji:",
	},
	"\U0001F5FC  :tokyo_tower:": {
		Name:    "Tokyo Tower",
		Unicode: "\U0001F5FC",
		Alias:   ":tokyo_tower:",
	},
	"\U0001F5FD  :statue_of_liberty:": {
		Name:    "Statue Of Liberty",
		Unicode: "\U0001F5FD",
		Alias:   ":statue_of_liberty:",
	},
	"\U0001F5FE  :japan:": {
		Name:    "Silhouette Of Japan",
		Unicode: "\U0001F5FE",
		Alias:   ":japan:",
	},
	"\U0001F5FF  :moyai:": {
		Name:    "Moyai",
		Unicode: "\U0001F5FF",
		Alias:   ":moyai:",
	},
	"\U0001F600  :grinning:": {
		Name:    "Grinning Face",
		Unicode: "\U0001F600",
		Alias:   ":grinning:",
	},
	"\U0001F601  :grin:": {
		Name:    "Grinning Face With Smiling Eyes",
		Unicode: "\U0001F601",
		Alias:   ":grin:",
	},
	"\U0001F602  :joy:": {
		Name:    "Face With Tears Of Joy",
		Unicode: "\U0001F602",
		Alias:   ":joy:",
	},
	"\U0001F603  :smiley:": {
		Name:    "Smiling Face With Open Mouth",
		Unicode: "\U0001F603",
		Alias:   ":smiley:",
	},
	"\U0001F604  :smile:": {
		Name:    "Smiling Face With Open Mouth And Smiling Eyes",
		Unicode: "\U0001F604",
		Alias:   ":smile:",
	},
	"\U0001F605  :sweat_smile:": {
		Name:    "Smiling Face With Open Mouth And Cold Sweat",
		Unicode: "\U0001F605",
		Alias:   ":sweat_smile:",
	},
	"\U0001F606  :laughing:": {
		Name:    "Smiling Face With Open Mouth And Tightly-Closed Eyes",
		Unicode: "\U0001F606",
		Alias:   ":laughing:",
	},
	"\U0001F607  :innocent:": {
		Name:    "Smiling Face With Halo",
		Unicode: "\U0001F607",
		Alias:   ":innocent:",
	},
	"\U0001F608  :smiling_imp:": {
		Name:    "Smiling Face With Horns",
		Unicode: "\U0001F608",
		Alias:   ":smiling_imp:",
	},
	"\U0001F609  :wink:": {
		Name:    "Winking Face",
		Unicode: "\U0001F609",
		Alias:   ":wink:",
	},
	"\U0001F60A  :blush:": {
		Name:    "Smiling Face With Smiling Eyes",
		Unicode: "\U0001F60A",
		Alias:   ":blush:",
	},
	"\U0001F60B  :yum:": {
		Name:    "Face Savouring Delicious Food",
		Unicode: "\U0001F60B",
		Alias:   ":yum:",
	},
	"\U0001F60C  :relieved:": {
		Name:    "Relieved Face",
		Unicode: "\U0001F60C",
		Alias:   ":relieved:",
	},
	"\U0001F60D  :heart_eyes:": {
		Name:    "Smiling Face With Heart-Shaped Eyes",
		Unicode: "\U0001F60D",
		Alias:   ":heart_eyes:",
	},
	"\U0001F60E  :sunglasses:": {
		Name:    "Smiling Face With Sunglasses",
		Unicode: "\U0001F60E",
		Alias:   ":sunglasses:",
	},
	"\U0001F60F  :smirk:": {
		Name:    "Smirking Face",
		Unicode: "\U0001F60F",
		Alias:   ":smirk:",
	},
	"\U0001F610  :neutral_face:": {
		Name:    "Neutral Face",
		Unicode: "\U0001F610",
		Alias:   ":neutral_face:",
	},
	"\U0001F611  :expressionless:": {
		Name:    "Expressionless Face",
		Unicode: "\U0001F611",
		Alias:   ":expressionless:",
	},
	"\U0001F612  :unamused:": {
		Name:    "Unamused Face",
		Unicode: "\U0001F612",
		Alias:   ":unamused:",
	},
	"\U0001F613  :sweat:": {
		Name:    "Face With Cold Sweat",
		Unicode: "\U0001F613",
		Alias:   ":sweat:",
	},
	"\U0001F614  :pensive:": {
		Name:    "Pensive Face",
		Unicode: "\U0001F614",
		Alias:   ":pensive:",
	},
	"\U0001F615  :confused:": {
		Name:    "Confused Face",
		Unicode: "\U0001F615",
		Alias:   ":confused:",
	},
	"\U0001F616  :confounded:": {
		Name:    "Confounded Face",
		Unicode: "\U0001F616",
		Alias:   ":confounded:",
	},
	"\U0001F617  :kissing:": {
		Name:    "Kissing Face",
		Unicode: "\U0001F617",
		Alias:   ":kissing:",
	},
	"\U0001F618  :kissing_heart:": {
		Name:    "Face Throwing A Kiss",
		Unicode: "\U0001F618",
		Alias:   ":kissing_heart:",
	},
	"\U0001F619  :kissing_smiling_eyes:": {
		Name:    "Kissing Face With Smiling Eyes",
		Unicode: "\U0001F619",
		Alias:   ":kissing_smiling_eyes:",
	},
	"\U0001F61A  :kissing_closed_eyes:": {
		Name:    "Kissing Face With Closed Eyes",
		Unicode: "\U0001F61A",
		Alias:   ":kissing_closed_eyes:",
	},
	"\U0001F61B  :stuck_out_tongue:": {
		Name:    "Face With Stuck-Out Tongue",
		Unicode: "\U0001F61B",
		Alias:   ":stuck_out_tongue:",
	},
	"\U0001F61C  :stuck_out_tongue_winking_eye:": {
		Name:    "Face With Stuck-Out Tongue And Winking Eye",
		Unicode: "\U0001F61C",
		Alias:   ":stuck_out_tongue_winking_eye:",
	},
	"\U0001F61D  :stuck_out_tongue_closed_eyes:": {
		Name:    "Face With Stuck-Out Tongue And Tightly-Closed Eyes",
		Unicode: "\U0001F61D",
		Alias:   ":stuck_out_tongue_closed_eyes:",
	},
	"\U0001F61E  :disappointed:": {
		Name:    "Disappointed Face",
		Unicode: "\U0001F61E",
		Alias:   ":disappointed:",
	},
	"\U0001F61F  :worried:": {
		Name:    "Worried Face",
		Unicode: "\U0001F61F",
		Alias:   ":worried:",
	},
	"\U0001F620  :angry:": {
		Name:    "Angry Face",
		Unicode: "\U0001F620",
		Alias:   ":angry:",
	},
	"\U0001F621  :rage:": {
		Name:    "Pouting Face",
		Unicode: "\U0001F621",
		Alias:   ":rage:",
	},
	"\U0001F622  :cry:": {
		Name:    "Crying Face",
		Unicode: "\U0001F622",
		Alias:   ":cry:",
	},
	"\U0001F623  :persevere:": {
		Name:    "Persevering Face",
		Unicode: "\U0001F623",
		Alias:   ":persevere:",
	},
	"\U0001F624  :triumph:": {
		Name:    "Face With Look Of Triumph",
		Unicode: "\U0001F624",
		Alias:   ":triumph:",
	},
	"\U0001F625  :disappointed_relieved:": {
		Name:    "Disappointed But Relieved Face",
		Unicode: "\U0001F625",
		Alias:   ":disappointed_relieved:",
	},
	"\U0001F626  :frowning:": {
		Name:    "Frowning Face With Open Mouth",
		Unicode: "\U0001F626",
		Alias:   ":frowning:",
	},
	"\U0001F627  :anguished:": {
		Name:    "Anguished Face",
		Unicode: "\U0001F627",
		Alias:   ":anguished:",
	},
	"\U0001F628  :fearful:": {
		Name:    "Fearful Face",
		Unicode: "\U0001F628",
		Alias:   ":fearful:",
	},
	"\U0001F629  :weary:": {
		Name:    "Weary Face",
		Unicode: "\U0001F629",
		Alias:   ":weary:",
	},
	"\U0001F62A  :sleepy:": {
		Name:    "Sleepy Face",
		Unicode: "\U0001F62A",
		Alias:   ":sleepy:",
	},
	"\U0001F62B  :tired_face:": {
		Name:    "Tired Face",
		Unicode: "\U0001F62B",
		Alias:   ":tired_face:",
	},
	"\U0001F62C  :grimacing:": {
		Name:    "Grimacing Face",
		Unicode: "\U0001F62C",
		Alias:   ":grimacing:",
	},
	"\U0001F62D  :sob:": {
		Name:    "Loudly Crying Face",
		Unicode: "\U0001F62D",
		Alias:   ":sob:",
	},
	"\U0001F62E\u200D\u1F4A8  :face_exhaling:": {
		Name:    "Face Exhaling",
		Unicode: "\U0001F62E\u200D\u1F4A8",
		Alias:   ":face_exhaling:",
	},
	"\U0001F62E  :open_mouth:": {
		Name:    "Face With Open Mouth",
		Unicode: "\U0001F62E",
		Alias:   ":open_mouth:",
	},
	"\U0001F62F  :hushed:": {
		Name:    "Hushed Face",
		Unicode: "\U0001F62F",
		Alias:   ":hushed:",
	},
	"\U0001F630  :cold_sweat:": {
		Name:    "Face With Open Mouth And Cold Sweat",
		Unicode: "\U0001F630",
		Alias:   ":cold_sweat:",
	},
	"\U0001F631  :scream:": {
		Name:    "Face Screaming In Fear",
		Unicode: "\U0001F631",
		Alias:   ":scream:",
	},
	"\U0001F632  :astonished:": {
		Name:    "Astonished Face",
		Unicode: "\U0001F632",
		Alias:   ":astonished:",
	},
	"\U0001F633  :flushed:": {
		Name:    "Flushed Face",
		Unicode: "\U0001F633",
		Alias:   ":flushed:",
	},
	"\U0001F634  :sleeping:": {
		Name:    "Sleeping Face",
		Unicode: "\U0001F634",
		Alias:   ":sleeping:",
	},
	"\U0001F635\u200D\u1F4AB  :face_with_spiral_eyes:": {
		Name:    "Face With Spiral Eyes",
		Unicode: "\U0001F635\u200D\u1F4AB",
		Alias:   ":face_with_spiral_eyes:",
	},
	"\U0001F635  :dizzy_face:": {
		Name:    "Dizzy Face",
		Unicode: "\U0001F635",
		Alias:   ":dizzy_face:",
	},
	"\U0001F636\u200D\u1F32B\uFE0F  :face_in_clouds:": {
		Name:    "Face In Clouds",
		Unicode: "\U0001F636\u200D\u1F32B\uFE0F",
		Alias:   ":face_in_clouds:",
	},
	"\U0001F636  :no_mouth:": {
		Name:    "Face Without Mouth",
		Unicode: "\U0001F636",
		Alias:   ":no_mouth:",
	},
	"\U0001F637  :mask:": {
		Name:    "Face With Medical Mask",
		Unicode: "\U0001F637",
		Alias:   ":mask:",
	},
	"\U0001F638  :smile_cat:": {
		Name:    "Grinning Cat Face With Smiling Eyes",
		Unicode: "\U0001F638",
		Alias:   ":smile_cat:",
	},
	"\U0001F639  :joy_cat:": {
		Name:    "Cat Face With Tears Of Joy",
		Unicode: "\U0001F639",
		Alias:   ":joy_cat:",
	},
	"\U0001F63A  :smiley_cat:": {
		Name:    "Smiling Cat Face With Open Mouth",
		Unicode: "\U0001F63A",
		Alias:   ":smiley_cat:",
	},
	"\U0001F63B  :heart_eyes_cat:": {
		Name:    "Smiling Cat Face With Heart-Shaped Eyes",
		Unicode: "\U0001F63B",
		Alias:   ":heart_eyes_cat:",
	},
	"\U0001F63C  :smirk_cat:": {
		Name:    "Cat Face With Wry Smile",
		Unicode: "\U0001F63C",
		Alias:   ":smirk_cat:",
	},
	"\U0001F63D  :kissing_cat:": {
		Name:    "Kissing Cat Face With Closed Eyes",
		Unicode: "\U0001F63D",
		Alias:   ":kissing_cat:",
	},
	"\U0001F63E  :pouting_cat:": {
		Name:    "Pouting Cat Face",
		Unicode: "\U0001F63E",
		Alias:   ":pouting_cat:",
	},
	"\U0001F63F  :crying_cat_face:": {
		Name:    "Crying Cat Face",
		Unicode: "\U0001F63F",
		Alias:   ":crying_cat_face:",
	},
	"\U0001F640  :scream_cat:": {
		Name:    "Weary Cat Face",
		Unicode: "\U0001F640",
		Alias:   ":scream_cat:",
	},
	"\U0001F641  :slightly_frowning_face:": {
		Name:    "Slightly Frowning Face",
		Unicode: "\U0001F641",
		Alias:   ":slightly_frowning_face:",
	},
	"\U0001F642  :slightly_smiling_face:": {
		Name:    "Slightly Smiling Face",
		Unicode: "\U0001F642",
		Alias:   ":slightly_smiling_face:",
	},
	"\U0001F643  :upside_down_face:": {
		Name:    "Upside-Down Face",
		Unicode: "\U0001F643",
		Alias:   ":upside_down_face:",
	},
	"\U0001F644  :face_with_rolling_eyes:": {
		Name:    "Face With Rolling Eyes",
		Unicode: "\U0001F644",
		Alias:   ":face_with_rolling_eyes:",
	},
	"\U0001F645\u200D\u2640\uFE0F :woman-gesturing-no:": {
		Name:    "Woman Gesturing No",
		Unicode: "\U0001F645\u200D\u2640\uFE0F",
		Alias:   ":woman-gesturing-no:",
	},
	"\U0001F645\u200D\u2642\uFE0F :man-gesturing-no:": {
		Name:    "Man Gesturing No",
		Unicode: "\U0001F645\u200D\u2642\uFE0F",
		Alias:   ":man-gesturing-no:",
	},
	"\U0001F645  :no_good:": {
		Name:    "Face With No Good Gesture",
		Unicode: "\U0001F645",
		Alias:   ":no_good:",
	},
	"\U0001F646\u200D\u2640\uFE0F :woman-gesturing-ok:": {
		Name:    "Woman Gesturing Ok",
		Unicode: "\U0001F646\u200D\u2640\uFE0F",
		Alias:   ":woman-gesturing-ok:",
	},
	"\U0001F646\u200D\u2642\uFE0F :man-gesturing-ok:": {
		Name:    "Man Gesturing Ok",
		Unicode: "\U0001F646\u200D\u2642\uFE0F",
		Alias:   ":man-gesturing-ok:",
	},
	"\U0001F646  :ok_woman:": {
		Name:    "Face With Ok Gesture",
		Unicode: "\U0001F646",
		Alias:   ":ok_woman:",
	},
	"\U0001F647\u200D\u2640\uFE0F :woman-bowing:": {
		Name:    "Woman Bowing",
		Unicode: "\U0001F647\u200D\u2640\uFE0F",
		Alias:   ":woman-bowing:",
	},
	"\U0001F647\u200D\u2642\uFE0F :man-bowing:": {
		Name:    "Man Bowing",
		Unicode: "\U0001F647\u200D\u2642\uFE0F",
		Alias:   ":man-bowing:",
	},
	"\U0001F647  :bow:": {
		Name:    "Person Bowing Deeply",
		Unicode: "\U0001F647",
		Alias:   ":bow:",
	},
	"\U0001F648  :see_no_evil:": {
		Name:    "See-No-Evil Monkey",
		Unicode: "\U0001F648",
		Alias:   ":see_no_evil:",
	},
	"\U0001F649  :hear_no_evil:": {
		Name:    "Hear-No-Evil Monkey",
		Unicode: "\U0001F649",
		Alias:   ":hear_no_evil:",
	},
	"\U0001F64A  :speak_no_evil:": {
		Name:    "Speak-No-Evil Monkey",
		Unicode: "\U0001F64A",
		Alias:   ":speak_no_evil:",
	},
	"\U0001F64B\u200D\u2640\uFE0F :woman-raising-hand:": {
		Name:    "Woman Raising Hand",
		Unicode: "\U0001F64B\u200D\u2640\uFE0F",
		Alias:   ":woman-raising-hand:",
	},
	"\U0001F64B\u200D\u2642\uFE0F :man-raising-hand:": {
		Name:    "Man Raising Hand",
		Unicode: "\U0001F64B\u200D\u2642\uFE0F",
		Alias:   ":man-raising-hand:",
	},
	"\U0001F64B  :raising_hand:": {
		Name:    "Happy Person Raising One Hand",
		Unicode: "\U0001F64B",
		Alias:   ":raising_hand:",
	},
	"\U0001F64C  :raised_hands:": {
		Name:    "Person Raising Both Hands In Celebration",
		Unicode: "\U0001F64C",
		Alias:   ":raised_hands:",
	},
	"\U0001F64D\u200D\u2640\uFE0F :woman-frowning:": {
		Name:    "Woman Frowning",
		Unicode: "\U0001F64D\u200D\u2640\uFE0F",
		Alias:   ":woman-frowning:",
	},
	"\U0001F64D\u200D\u2642\uFE0F :man-frowning:": {
		Name:    "Man Frowning",
		Unicode: "\U0001F64D\u200D\u2642\uFE0F",
		Alias:   ":man-frowning:",
	},
	"\U0001F64D  :person_frowning:": {
		Name:    "Person Frowning",
		Unicode: "\U0001F64D",
		Alias:   ":person_frowning:",
	},
	"\U0001F64E\u200D\u2640\uFE0F :woman-pouting:": {
		Name:    "Woman Pouting",
		Unicode: "\U0001F64E\u200D\u2640\uFE0F",
		Alias:   ":woman-pouting:",
	},
	"\U0001F64E\u200D\u2642\uFE0F :man-pouting:": {
		Name:    "Man Pouting",
		Unicode: "\U0001F64E\u200D\u2642\uFE0F",
		Alias:   ":man-pouting:",
	},
	"\U0001F64E  :person_with_pouting_face:": {
		Name:    "Person With Pouting Face",
		Unicode: "\U0001F64E",
		Alias:   ":person_with_pouting_face:",
	},
	"\U0001F64F  :pray:": {
		Name:    "Person With Folded Hands",
		Unicode: "\U0001F64F",
		Alias:   ":pray:",
	},
	"\U0001F680  :rocket:": {
		Name:    "Rocket",
		Unicode: "\U0001F680",
		Alias:   ":rocket:",
	},
	"\U0001F681  :helicopter:": {
		Name:    "Helicopter",
		Unicode: "\U0001F681",
		Alias:   ":helicopter:",
	},
	"\U0001F682  :steam_locomotive:": {
		Name:    "Steam Locomotive",
		Unicode: "\U0001F682",
		Alias:   ":steam_locomotive:",
	},
	"\U0001F683  :railway_car:": {
		Name:    "Railway Car",
		Unicode: "\U0001F683",
		Alias:   ":railway_car:",
	},
	"\U0001F684  :bullettrain_side:": {
		Name:    "High-Speed Train",
		Unicode: "\U0001F684",
		Alias:   ":bullettrain_side:",
	},
	"\U0001F685  :bullettrain_front:": {
		Name:    "High-Speed Train With Bullet Nose",
		Unicode: "\U0001F685",
		Alias:   ":bullettrain_front:",
	},
	"\U0001F686  :train2:": {
		Name:    "Train",
		Unicode: "\U0001F686",
		Alias:   ":train2:",
	},
	"\U0001F687  :metro:": {
		Name:    "Metro",
		Unicode: "\U0001F687",
		Alias:   ":metro:",
	},
	"\U0001F688  :light_rail:": {
		Name:    "Light Rail",
		Unicode: "\U0001F688",
		Alias:   ":light_rail:",
	},
	"\U0001F689  :station:": {
		Name:    "Station",
		Unicode: "\U0001F689",
		Alias:   ":station:",
	},
	"\U0001F68A  :tram:": {
		Name:    "Tram",
		Unicode: "\U0001F68A",
		Alias:   ":tram:",
	},
	"\U0001F68B  :train:": {
		Name:    "Tram Car",
		Unicode: "\U0001F68B",
		Alias:   ":train:",
	},
	"\U0001F68C  :bus:": {
		Name:    "Bus",
		Unicode: "\U0001F68C",
		Alias:   ":bus:",
	},
	"\U0001F68D  :oncoming_bus:": {
		Name:    "Oncoming Bus",
		Unicode: "\U0001F68D",
		Alias:   ":oncoming_bus:",
	},
	"\U0001F68E  :trolleybus:": {
		Name:    "Trolleybus",
		Unicode: "\U0001F68E",
		Alias:   ":trolleybus:",
	},
	"\U0001F68F  :busstop:": {
		Name:    "Bus Stop",
		Unicode: "\U0001F68F",
		Alias:   ":busstop:",
	},
	"\U0001F690  :minibus:": {
		Name:    "Minibus",
		Unicode: "\U0001F690",
		Alias:   ":minibus:",
	},
	"\U0001F691  :ambulance:": {
		Name:    "Ambulance",
		Unicode: "\U0001F691",
		Alias:   ":ambulance:",
	},
	"\U0001F692  :fire_engine:": {
		Name:    "Fire Engine",
		Unicode: "\U0001F692",
		Alias:   ":fire_engine:",
	},
	"\U0001F693  :police_car:": {
		Name:    "Police Car",
		Unicode: "\U0001F693",
		Alias:   ":police_car:",
	},
	"\U0001F694  :oncoming_police_car:": {
		Name:    "Oncoming Police Car",
		Unicode: "\U0001F694",
		Alias:   ":oncoming_police_car:",
	},
	"\U0001F695  :taxi:": {
		Name:    "Taxi",
		Unicode: "\U0001F695",
		Alias:   ":taxi:",
	},
	"\U0001F696  :oncoming_taxi:": {
		Name:    "Oncoming Taxi",
		Unicode: "\U0001F696",
		Alias:   ":oncoming_taxi:",
	},
	"\U0001F697  :car:": {
		Name:    "Automobile",
		Unicode: "\U0001F697",
		Alias:   ":car:",
	},
	"\U0001F698  :oncoming_automobile:": {
		Name:    "Oncoming Automobile",
		Unicode: "\U0001F698",
		Alias:   ":oncoming_automobile:",
	},
	"\U0001F699  :blue_car:": {
		Name:    "Recreational Vehicle",
		Unicode: "\U0001F699",
		Alias:   ":blue_car:",
	},
	"\U0001F69A  :truck:": {
		Name:    "Delivery Truck",
		Unicode: "\U0001F69A",
		Alias:   ":truck:",
	},
	"\U0001F69B  :articulated_lorry:": {
		Name:    "Articulated Lorry",
		Unicode: "\U0001F69B",
		Alias:   ":articulated_lorry:",
	},
	"\U0001F69C  :tractor:": {
		Name:    "Tractor",
		Unicode: "\U0001F69C",
		Alias:   ":tractor:",
	},
	"\U0001F69D  :monorail:": {
		Name:    "Monorail",
		Unicode: "\U0001F69D",
		Alias:   ":monorail:",
	},
	"\U0001F69E  :mountain_railway:": {
		Name:    "Mountain Railway",
		Unicode: "\U0001F69E",
		Alias:   ":mountain_railway:",
	},
	"\U0001F69F  :suspension_railway:": {
		Name:    "Suspension Railway",
		Unicode: "\U0001F69F",
		Alias:   ":suspension_railway:",
	},
	"\U0001F6A0  :mountain_cableway:": {
		Name:    "Mountain Cableway",
		Unicode: "\U0001F6A0",
		Alias:   ":mountain_cableway:",
	},
	"\U0001F6A1  :aerial_tramway:": {
		Name:    "Aerial Tramway",
		Unicode: "\U0001F6A1",
		Alias:   ":aerial_tramway:",
	},
	"\U0001F6A2  :ship:": {
		Name:    "Ship",
		Unicode: "\U0001F6A2",
		Alias:   ":ship:",
	},
	"\U0001F6A3\u200D\u2640\uFE0F :woman-rowing-boat:": {
		Name:    "Woman Rowing Boat",
		Unicode: "\U0001F6A3\u200D\u2640\uFE0F",
		Alias:   ":woman-rowing-boat:",
	},
	"\U0001F6A3\u200D\u2642\uFE0F :man-rowing-boat:": {
		Name:    "Man Rowing Boat",
		Unicode: "\U0001F6A3\u200D\u2642\uFE0F",
		Alias:   ":man-rowing-boat:",
	},
	"\U0001F6A3  :rowboat:": {
		Name:    "Rowboat",
		Unicode: "\U0001F6A3",
		Alias:   ":rowboat:",
	},
	"\U0001F6A4  :speedboat:": {
		Name:    "Speedboat",
		Unicode: "\U0001F6A4",
		Alias:   ":speedboat:",
	},
	"\U0001F6A5  :traffic_light:": {
		Name:    "Horizontal Traffic Light",
		Unicode: "\U0001F6A5",
		Alias:   ":traffic_light:",
	},
	"\U0001F6A6  :vertical_traffic_light:": {
		Name:    "Vertical Traffic Light",
		Unicode: "\U0001F6A6",
		Alias:   ":vertical_traffic_light:",
	},
	"\U0001F6A7  :construction:": {
		Name:    "Construction Sign",
		Unicode: "\U0001F6A7",
		Alias:   ":construction:",
	},
	"\U0001F6A8  :rotating_light:": {
		Name:    "Police Cars Revolving Light",
		Unicode: "\U0001F6A8",
		Alias:   ":rotating_light:",
	},
	"\U0001F6A9  :triangular_flag_on_post:": {
		Name:    "Triangular Flag On Post",
		Unicode: "\U0001F6A9",
		Alias:   ":triangular_flag_on_post:",
	},
	"\U0001F6AA  :door:": {
		Name:    "Door",
		Unicode: "\U0001F6AA",
		Alias:   ":door:",
	},
	"\U0001F6AB  :no_entry_sign:": {
		Name:    "No Entry Sign",
		Unicode: "\U0001F6AB",
		Alias:   ":no_entry_sign:",
	},
	"\U0001F6AC  :smoking:": {
		Name:    "Smoking Symbol",
		Unicode: "\U0001F6AC",
		Alias:   ":smoking:",
	},
	"\U0001F6AD  :no_smoking:": {
		Name:    "No Smoking Symbol",
		Unicode: "\U0001F6AD",
		Alias:   ":no_smoking:",
	},
	"\U0001F6AE  :put_litter_in_its_place:": {
		Name:    "Put Litter In Its Place Symbol",
		Unicode: "\U0001F6AE",
		Alias:   ":put_litter_in_its_place:",
	},
	"\U0001F6AF  :do_not_litter:": {
		Name:    "Do Not Litter Symbol",
		Unicode: "\U0001F6AF",
		Alias:   ":do_not_litter:",
	},
	"\U0001F6B0  :potable_water:": {
		Name:    "Potable Water Symbol",
		Unicode: "\U0001F6B0",
		Alias:   ":potable_water:",
	},
	"\U0001F6B1 :non-potable_water:": {
		Name:    "Non-Potable Water Symbol",
		Unicode: "\U0001F6B1",
		Alias:   ":non-potable_water:",
	},
	"\U0001F6B2  :bike:": {
		Name:    "Bicycle",
		Unicode: "\U0001F6B2",
		Alias:   ":bike:",
	},
	"\U0001F6B3  :no_bicycles:": {
		Name:    "No Bicycles",
		Unicode: "\U0001F6B3",
		Alias:   ":no_bicycles:",
	},
	"\U0001F6B4\u200D\u2640\uFE0F :woman-biking:": {
		Name:    "Woman Biking",
		Unicode: "\U0001F6B4\u200D\u2640\uFE0F",
		Alias:   ":woman-biking:",
	},
	"\U0001F6B4\u200D\u2642\uFE0F :man-biking:": {
		Name:    "Man Biking",
		Unicode: "\U0001F6B4\u200D\u2642\uFE0F",
		Alias:   ":man-biking:",
	},
	"\U0001F6B4  :bicyclist:": {
		Name:    "Bicyclist",
		Unicode: "\U0001F6B4",
		Alias:   ":bicyclist:",
	},
	"\U0001F6B5\u200D\u2640\uFE0F :woman-mountain-biking:": {
		Name:    "Woman Mountain Biking",
		Unicode: "\U0001F6B5\u200D\u2640\uFE0F",
		Alias:   ":woman-mountain-biking:",
	},
	"\U0001F6B5\u200D\u2642\uFE0F :man-mountain-biking:": {
		Name:    "Man Mountain Biking",
		Unicode: "\U0001F6B5\u200D\u2642\uFE0F",
		Alias:   ":man-mountain-biking:",
	},
	"\U0001F6B5  :mountain_bicyclist:": {
		Name:    "Mountain Bicyclist",
		Unicode: "\U0001F6B5",
		Alias:   ":mountain_bicyclist:",
	},
	"\U0001F6B6\u200D\u2640\uFE0F :woman-walking:": {
		Name:    "Woman Walking",
		Unicode: "\U0001F6B6\u200D\u2640\uFE0F",
		Alias:   ":woman-walking:",
	},
	"\U0001F6B6\u200D\u2642\uFE0F :man-walking:": {
		Name:    "Man Walking",
		Unicode: "\U0001F6B6\u200D\u2642\uFE0F",
		Alias:   ":man-walking:",
	},
	"\U0001F6B6  :walking:": {
		Name:    "Pedestrian",
		Unicode: "\U0001F6B6",
		Alias:   ":walking:",
	},
	"\U0001F6B7  :no_pedestrians:": {
		Name:    "No Pedestrians",
		Unicode: "\U0001F6B7",
		Alias:   ":no_pedestrians:",
	},
	"\U0001F6B8  :children_crossing:": {
		Name:    "Children Crossing",
		Unicode: "\U0001F6B8",
		Alias:   ":children_crossing:",
	},
	"\U0001F6B9  :mens:": {
		Name:    "Mens Symbol",
		Unicode: "\U0001F6B9",
		Alias:   ":mens:",
	},
	"\U0001F6BA  :womens:": {
		Name:    "Womens Symbol",
		Unicode: "\U0001F6BA",
		Alias:   ":womens:",
	},
	"\U0001F6BB  :restroom:": {
		Name:    "Restroom",
		Unicode: "\U0001F6BB",
		Alias:   ":restroom:",
	},
	"\U0001F6BC  :baby_symbol:": {
		Name:    "Baby Symbol",
		Unicode: "\U0001F6BC",
		Alias:   ":baby_symbol:",
	},
	"\U0001F6BD  :toilet:": {
		Name:    "Toilet",
		Unicode: "\U0001F6BD",
		Alias:   ":toilet:",
	},
	"\U0001F6BE  :wc:": {
		Name:    "Water Closet",
		Unicode: "\U0001F6BE",
		Alias:   ":wc:",
	},
	"\U0001F6BF  :shower:": {
		Name:    "Shower",
		Unicode: "\U0001F6BF",
		Alias:   ":shower:",
	},
	"\U0001F6C0  :bath:": {
		Name:    "Bath",
		Unicode: "\U0001F6C0",
		Alias:   ":bath:",
	},
	"\U0001F6C1  :bathtub:": {
		Name:    "Bathtub",
		Unicode: "\U0001F6C1",
		Alias:   ":bathtub:",
	},
	"\U0001F6C2  :passport_control:": {
		Name:    "Passport Control",
		Unicode: "\U0001F6C2",
		Alias:   ":passport_control:",
	},
	"\U0001F6C3  :customs:": {
		Name:    "Customs",
		Unicode: "\U0001F6C3",
		Alias:   ":customs:",
	},
	"\U0001F6C4  :baggage_claim:": {
		Name:    "Baggage Claim",
		Unicode: "\U0001F6C4",
		Alias:   ":baggage_claim:",
	},
	"\U0001F6C5  :left_luggage:": {
		Name:    "Left Luggage",
		Unicode: "\U0001F6C5",
		Alias:   ":left_luggage:",
	},
	"\U0001F6CB\uFE0F  :couch_and_lamp:": {
		Name:    "Couch And Lamp",
		Unicode: "\U0001F6CB\uFE0F",
		Alias:   ":couch_and_lamp:",
	},
	"\U0001F6CC  :sleeping_accommodation:": {
		Name:    "Sleeping Accommodation",
		Unicode: "\U0001F6CC",
		Alias:   ":sleeping_accommodation:",
	},
	"\U0001F6CD\uFE0F  :shopping_bags:": {
		Name:    "Shopping Bags",
		Unicode: "\U0001F6CD\uFE0F",
		Alias:   ":shopping_bags:",
	},
	"\U0001F6CE\uFE0F  :bellhop_bell:": {
		Name:    "Bellhop Bell",
		Unicode: "\U0001F6CE\uFE0F",
		Alias:   ":bellhop_bell:",
	},
	"\U0001F6CF\uFE0F  :bed:": {
		Name:    "Bed",
		Unicode: "\U0001F6CF\uFE0F",
		Alias:   ":bed:",
	},
	"\U0001F6D0  :place_of_worship:": {
		Name:    "Place Of Worship",
		Unicode: "\U0001F6D0",
		Alias:   ":place_of_worship:",
	},
	"\U0001F6D1  :octagonal_sign:": {
		Name:    "Octagonal Sign",
		Unicode: "\U0001F6D1",
		Alias:   ":octagonal_sign:",
	},
	"\U0001F6D2  :shopping_trolley:": {
		Name:    "Shopping Trolley",
		Unicode: "\U0001F6D2",
		Alias:   ":shopping_trolley:",
	},
	"\U0001F6D5  :hindu_temple:": {
		Name:    "Hindu Temple",
		Unicode: "\U0001F6D5",
		Alias:   ":hindu_temple:",
	},
	"\U0001F6D6  :hut:": {
		Name:    "Hut",
		Unicode: "\U0001F6D6",
		Alias:   ":hut:",
	},
	"\U0001F6D7  :elevator:": {
		Name:    "Elevator",
		Unicode: "\U0001F6D7",
		Alias:   ":elevator:",
	},
	"\U0001F6DD  :playground_slide:": {
		Name:    "Playground Slide",
		Unicode: "\U0001F6DD",
		Alias:   ":playground_slide:",
	},
	"\U0001F6DE  :wheel:": {
		Name:    "Wheel",
		Unicode: "\U0001F6DE",
		Alias:   ":wheel:",
	},
	"\U0001F6DF  :ring_buoy:": {
		Name:    "Ring Buoy",
		Unicode: "\U0001F6DF",
		Alias:   ":ring_buoy:",
	},
	"\U0001F6E0\uFE0F  :hammer_and_wrench:": {
		Name:    "Hammer And Wrench",
		Unicode: "\U0001F6E0\uFE0F",
		Alias:   ":hammer_and_wrench:",
	},
	"\U0001F6E1\uFE0F  :shield:": {
		Name:    "Shield",
		Unicode: "\U0001F6E1\uFE0F",
		Alias:   ":shield:",
	},
	"\U0001F6E2\uFE0F  :oil_drum:": {
		Name:    "Oil Drum",
		Unicode: "\U0001F6E2\uFE0F",
		Alias:   ":oil_drum:",
	},
	"\U0001F6E3\uFE0F  :motorway:": {
		Name:    "Motorway",
		Unicode: "\U0001F6E3\uFE0F",
		Alias:   ":motorway:",
	},
	"\U0001F6E4\uFE0F  :railway_track:": {
		Name:    "Railway Track",
		Unicode: "\U0001F6E4\uFE0F",
		Alias:   ":railway_track:",
	},
	"\U0001F6E5\uFE0F  :motor_boat:": {
		Name:    "Motor Boat",
		Unicode: "\U0001F6E5\uFE0F",
		Alias:   ":motor_boat:",
	},
	"\U0001F6E9\uFE0F  :small_airplane:": {
		Name:    "Small Airplane",
		Unicode: "\U0001F6E9\uFE0F",
		Alias:   ":small_airplane:",
	},
	"\U0001F6EB  :airplane_departure:": {
		Name:    "Airplane Departure",
		Unicode: "\U0001F6EB",
		Alias:   ":airplane_departure:",
	},
	"\U0001F6EC  :airplane_arriving:": {
		Name:    "Airplane Arriving",
		Unicode: "\U0001F6EC",
		Alias:   ":airplane_arriving:",
	},
	"\U0001F6F0\uFE0F  :satellite:": {
		Name:    "Satellite",
		Unicode: "\U0001F6F0\uFE0F",
		Alias:   ":satellite:",
	},
	"\U0001F6F3\uFE0F  :passenger_ship:": {
		Name:    "Passenger Ship",
		Unicode: "\U0001F6F3\uFE0F",
		Alias:   ":passenger_ship:",
	},
	"\U0001F6F4  :scooter:": {
		Name:    "Scooter",
		Unicode: "\U0001F6F4",
		Alias:   ":scooter:",
	},
	"\U0001F6F5  :motor_scooter:": {
		Name:    "Motor Scooter",
		Unicode: "\U0001F6F5",
		Alias:   ":motor_scooter:",
	},
	"\U0001F6F6  :canoe:": {
		Name:    "Canoe",
		Unicode: "\U0001F6F6",
		Alias:   ":canoe:",
	},
	"\U0001F6F7  :sled:": {
		Name:    "Sled",
		Unicode: "\U0001F6F7",
		Alias:   ":sled:",
	},
	"\U0001F6F8  :flying_saucer:": {
		Name:    "Flying Saucer",
		Unicode: "\U0001F6F8",
		Alias:   ":flying_saucer:",
	},
	"\U0001F6F9  :skateboard:": {
		Name:    "Skateboard",
		Unicode: "\U0001F6F9",
		Alias:   ":skateboard:",
	},
	"\U0001F6FA  :auto_rickshaw:": {
		Name:    "Auto Rickshaw",
		Unicode: "\U0001F6FA",
		Alias:   ":auto_rickshaw:",
	},
	"\U0001F6FB  :pickup_truck:": {
		Name:    "Pickup Truck",
		Unicode: "\U0001F6FB",
		Alias:   ":pickup_truck:",
	},
	"\U0001F6FC  :roller_skate:": {
		Name:    "Roller Skate",
		Unicode: "\U0001F6FC",
		Alias:   ":roller_skate:",
	},
	"\U0001F7E0  :large_orange_circle:": {
		Name:    "Large Orange Circle",
		Unicode: "\U0001F7E0",
		Alias:   ":large_orange_circle:",
	},
	"\U0001F7E1  :large_yellow_circle:": {
		Name:    "Large Yellow Circle",
		Unicode: "\U0001F7E1",
		Alias:   ":large_yellow_circle:",
	},
	"\U0001F7E2  :large_green_circle:": {
		Name:    "Large Green Circle",
		Unicode: "\U0001F7E2",
		Alias:   ":large_green_circle:",
	},
	"\U0001F7E3  :large_purple_circle:": {
		Name:    "Large Purple Circle",
		Unicode: "\U0001F7E3",
		Alias:   ":large_purple_circle:",
	},
	"\U0001F7E4  :large_brown_circle:": {
		Name:    "Large Brown Circle",
		Unicode: "\U0001F7E4",
		Alias:   ":large_brown_circle:",
	},
	"\U0001F7E5  :large_red_square:": {
		Name:    "Large Red Square",
		Unicode: "\U0001F7E5",
		Alias:   ":large_red_square:",
	},
	"\U0001F7E6  :large_blue_square:": {
		Name:    "Large Blue Square",
		Unicode: "\U0001F7E6",
		Alias:   ":large_blue_square:",
	},
	"\U0001F7E7  :large_orange_square:": {
		Name:    "Large Orange Square",
		Unicode: "\U0001F7E7",
		Alias:   ":large_orange_square:",
	},
	"\U0001F7E8  :large_yellow_square:": {
		Name:    "Large Yellow Square",
		Unicode: "\U0001F7E8",
		Alias:   ":large_yellow_square:",
	},
	"\U0001F7E9  :large_green_square:": {
		Name:    "Large Green Square",
		Unicode: "\U0001F7E9",
		Alias:   ":large_green_square:",
	},
	"\U0001F7EA  :large_purple_square:": {
		Name:    "Large Purple Square",
		Unicode: "\U0001F7EA",
		Alias:   ":large_purple_square:",
	},
	"\U0001F7EB  :large_brown_square:": {
		Name:    "Large Brown Square",
		Unicode: "\U0001F7EB",
		Alias:   ":large_brown_square:",
	},
	"\U0001F7F0  :heavy_equals_sign:": {
		Name:    "Heavy Equals Sign",
		Unicode: "\U0001F7F0",
		Alias:   ":heavy_equals_sign:",
	},
	"\U0001F90C  :pinched_fingers:": {
		Name:    "Pinched Fingers",
		Unicode: "\U0001F90C",
		Alias:   ":pinched_fingers:",
	},
	"\U0001F90D  :white_heart:": {
		Name:    "White Heart",
		Unicode: "\U0001F90D",
		Alias:   ":white_heart:",
	},
	"\U0001F90E  :brown_heart:": {
		Name:    "Brown Heart",
		Unicode: "\U0001F90E",
		Alias:   ":brown_heart:",
	},
	"\U0001F90F  :pinching_hand:": {
		Name:    "Pinching Hand",
		Unicode: "\U0001F90F",
		Alias:   ":pinching_hand:",
	},
	"\U0001F910  :zipper_mouth_face:": {
		Name:    "Zipper-Mouth Face",
		Unicode: "\U0001F910",
		Alias:   ":zipper_mouth_face:",
	},
	"\U0001F911  :money_mouth_face:": {
		Name:    "Money-Mouth Face",
		Unicode: "\U0001F911",
		Alias:   ":money_mouth_face:",
	},
	"\U0001F912  :face_with_thermometer:": {
		Name:    "Face With Thermometer",
		Unicode: "\U0001F912",
		Alias:   ":face_with_thermometer:",
	},
	"\U0001F913  :nerd_face:": {
		Name:    "Nerd Face",
		Unicode: "\U0001F913",
		Alias:   ":nerd_face:",
	},
	"\U0001F914  :thinking_face:": {
		Name:    "Thinking Face",
		Unicode: "\U0001F914",
		Alias:   ":thinking_face:",
	},
	"\U0001F915  :face_with_head_bandage:": {
		Name:    "Face With Head-Bandage",
		Unicode: "\U0001F915",
		Alias:   ":face_with_head_bandage:",
	},
	"\U0001F916  :robot_face:": {
		Name:    "Robot Face",
		Unicode: "\U0001F916",
		Alias:   ":robot_face:",
	},
	"\U0001F917  :hugging_face:": {
		Name:    "Hugging Face",
		Unicode: "\U0001F917",
		Alias:   ":hugging_face:",
	},
	"\U0001F918  :the_horns:": {
		Name:    "Sign Of The Horns",
		Unicode: "\U0001F918",
		Alias:   ":the_horns:",
	},
	"\U0001F919  :call_me_hand:": {
		Name:    "Call Me Hand",
		Unicode: "\U0001F919",
		Alias:   ":call_me_hand:",
	},
	"\U0001F91A  :raised_back_of_hand:": {
		Name:    "Raised Back Of Hand",
		Unicode: "\U0001F91A",
		Alias:   ":raised_back_of_hand:",
	},
	"\U0001F91B :left-facing_fist:": {
		Name:    "Left-Facing Fist",
		Unicode: "\U0001F91B",
		Alias:   ":left-facing_fist:",
	},
	"\U0001F91C :right-facing_fist:": {
		Name:    "Right-Facing Fist",
		Unicode: "\U0001F91C",
		Alias:   ":right-facing_fist:",
	},
	"\U0001F91D  :handshake:": {
		Name:    "Handshake",
		Unicode: "\U0001F91D",
		Alias:   ":handshake:",
	},
	"\U0001F91E  :crossed_fingers:": {
		Name:    "Hand With Index And Middle Fingers Crossed",
		Unicode: "\U0001F91E",
		Alias:   ":crossed_fingers:",
	},
	"\U0001F91F  :i_love_you_hand_sign:": {
		Name:    "I Love You Hand Sign",
		Unicode: "\U0001F91F",
		Alias:   ":i_love_you_hand_sign:",
	},
	"\U0001F920  :face_with_cowboy_hat:": {
		Name:    "Face With Cowboy Hat",
		Unicode: "\U0001F920",
		Alias:   ":face_with_cowboy_hat:",
	},
	"\U0001F921  :clown_face:": {
		Name:    "Clown Face",
		Unicode: "\U0001F921",
		Alias:   ":clown_face:",
	},
	"\U0001F922  :nauseated_face:": {
		Name:    "Nauseated Face",
		Unicode: "\U0001F922",
		Alias:   ":nauseated_face:",
	},
	"\U0001F923  :rolling_on_the_floor_laughing:": {
		Name:    "Rolling On The Floor Laughing",
		Unicode: "\U0001F923",
		Alias:   ":rolling_on_the_floor_laughing:",
	},
	"\U0001F924  :drooling_face:": {
		Name:    "Drooling Face",
		Unicode: "\U0001F924",
		Alias:   ":drooling_face:",
	},
	"\U0001F925  :lying_face:": {
		Name:    "Lying Face",
		Unicode: "\U0001F925",
		Alias:   ":lying_face:",
	},
	"\U0001F926\u200D\u2640\uFE0F :woman-facepalming:": {
		Name:    "Woman Facepalming",
		Unicode: "\U0001F926\u200D\u2640\uFE0F",
		Alias:   ":woman-facepalming:",
	},
	"\U0001F926\u200D\u2642\uFE0F :man-facepalming:": {
		Name:    "Man Facepalming",
		Unicode: "\U0001F926\u200D\u2642\uFE0F",
		Alias:   ":man-facepalming:",
	},
	"\U0001F926  :face_palm:": {
		Name:    "Face Palm",
		Unicode: "\U0001F926",
		Alias:   ":face_palm:",
	},
	"\U0001F927  :sneezing_face:": {
		Name:    "Sneezing Face",
		Unicode: "\U0001F927",
		Alias:   ":sneezing_face:",
	},
	"\U0001F928  :face_with_raised_eyebrow:": {
		Name:    "Face With One Eyebrow Raised",
		Unicode: "\U0001F928",
		Alias:   ":face_with_raised_eyebrow:",
	},
	"\U0001F929 :star-struck:": {
		Name:    "Grinning Face With Star Eyes",
		Unicode: "\U0001F929",
		Alias:   ":star-struck:",
	},
	"\U0001F92A  :zany_face:": {
		Name:    "Grinning Face With One Large And One Small Eye",
		Unicode: "\U0001F92A",
		Alias:   ":zany_face:",
	},
	"\U0001F92B  :shushing_face:": {
		Name:    "Face With Finger Covering Closed Lips",
		Unicode: "\U0001F92B",
		Alias:   ":shushing_face:",
	},
	"\U0001F92C  :face_with_symbols_on_mouth:": {
		Name:    "Serious Face With Symbols Covering Mouth",
		Unicode: "\U0001F92C",
		Alias:   ":face_with_symbols_on_mouth:",
	},
	"\U0001F92D  :face_with_hand_over_mouth:": {
		Name:    "Smiling Face With Smiling Eyes And Hand Covering Mouth",
		Unicode: "\U0001F92D",
		Alias:   ":face_with_hand_over_mouth:",
	},
	"\U0001F92E  :face_vomiting:": {
		Name:    "Face With Open Mouth Vomiting",
		Unicode: "\U0001F92E",
		Alias:   ":face_vomiting:",
	},
	"\U0001F92F  :exploding_head:": {
		Name:    "Shocked Face With Exploding Head",
		Unicode: "\U0001F92F",
		Alias:   ":exploding_head:",
	},
	"\U0001F930  :pregnant_woman:": {
		Name:    "Pregnant Woman",
		Unicode: "\U0001F930",
		Alias:   ":pregnant_woman:",
	},
	"\U0001F931 :breast-feeding:": {
		Name:    "Breast-Feeding",
		Unicode: "\U0001F931",
		Alias:   ":breast-feeding:",
	},
	"\U0001F932  :palms_up_together:": {
		Name:    "Palms Up Together",
		Unicode: "\U0001F932",
		Alias:   ":palms_up_together:",
	},
	"\U0001F933  :selfie:": {
		Name:    "Selfie",
		Unicode: "\U0001F933",
		Alias:   ":selfie:",
	},
	"\U0001F934  :prince:": {
		Name:    "Prince",
		Unicode: "\U0001F934",
		Alias:   ":prince:",
	},
	"\U0001F935\u200D\u2640\uFE0F  :woman_in_tuxedo:": {
		Name:    "Woman In Tuxedo",
		Unicode: "\U0001F935\u200D\u2640\uFE0F",
		Alias:   ":woman_in_tuxedo:",
	},
	"\U0001F935\u200D\u2642\uFE0F  :man_in_tuxedo:": {
		Name:    "Man In Tuxedo",
		Unicode: "\U0001F935\u200D\u2642\uFE0F",
		Alias:   ":man_in_tuxedo:",
	},
	"\U0001F935  :person_in_tuxedo:": {
		Name:    "Man In Tuxedo",
		Unicode: "\U0001F935",
		Alias:   ":person_in_tuxedo:",
	},
	"\U0001F936  :mrs_claus:": {
		Name:    "Mother Christmas",
		Unicode: "\U0001F936",
		Alias:   ":mrs_claus:",
	},
	"\U0001F937\u200D\u2640\uFE0F :woman-shrugging:": {
		Name:    "Woman Shrugging",
		Unicode: "\U0001F937\u200D\u2640\uFE0F",
		Alias:   ":woman-shrugging:",
	},
	"\U0001F937\u200D\u2642\uFE0F :man-shrugging:": {
		Name:    "Man Shrugging",
		Unicode: "\U0001F937\u200D\u2642\uFE0F",
		Alias:   ":man-shrugging:",
	},
	"\U0001F937  :shrug:": {
		Name:    "Shrug",
		Unicode: "\U0001F937",
		Alias:   ":shrug:",
	},
	"\U0001F938\u200D\u2640\uFE0F :woman-cartwheeling:": {
		Name:    "Woman Cartwheeling",
		Unicode: "\U0001F938\u200D\u2640\uFE0F",
		Alias:   ":woman-cartwheeling:",
	},
	"\U0001F938\u200D\u2642\uFE0F :man-cartwheeling:": {
		Name:    "Man Cartwheeling",
		Unicode: "\U0001F938\u200D\u2642\uFE0F",
		Alias:   ":man-cartwheeling:",
	},
	"\U0001F938  :person_doing_cartwheel:": {
		Name:    "Person Doing Cartwheel",
		Unicode: "\U0001F938",
		Alias:   ":person_doing_cartwheel:",
	},
	"\U0001F939\u200D\u2640\uFE0F :woman-juggling:": {
		Name:    "Woman Juggling",
		Unicode: "\U0001F939\u200D\u2640\uFE0F",
		Alias:   ":woman-juggling:",
	},
	"\U0001F939\u200D\u2642\uFE0F :man-juggling:": {
		Name:    "Man Juggling",
		Unicode: "\U0001F939\u200D\u2642\uFE0F",
		Alias:   ":man-juggling:",
	},
	"\U0001F939  :juggling:": {
		Name:    "Juggling",
		Unicode: "\U0001F939",
		Alias:   ":juggling:",
	},
	"\U0001F93A  :fencer:": {
		Name:    "Fencer",
		Unicode: "\U0001F93A",
		Alias:   ":fencer:",
	},
	"\U0001F93C\u200D\u2640\uFE0F :woman-wrestling:": {
		Name:    "Women Wrestling",
		Unicode: "\U0001F93C\u200D\u2640\uFE0F",
		Alias:   ":woman-wrestling:",
	},
	"\U0001F93C\u200D\u2642\uFE0F :man-wrestling:": {
		Name:    "Men Wrestling",
		Unicode: "\U0001F93C\u200D\u2642\uFE0F",
		Alias:   ":man-wrestling:",
	},
	"\U0001F93C  :wrestlers:": {
		Name:    "Wrestlers",
		Unicode: "\U0001F93C",
		Alias:   ":wrestlers:",
	},
	"\U0001F93D\u200D\u2640\uFE0F :woman-playing-water-polo:": {
		Name:    "Woman Playing Water Polo",
		Unicode: "\U0001F93D\u200D\u2640\uFE0F",
		Alias:   ":woman-playing-water-polo:",
	},
	"\U0001F93D\u200D\u2642\uFE0F :man-playing-water-polo:": {
		Name:    "Man Playing Water Polo",
		Unicode: "\U0001F93D\u200D\u2642\uFE0F",
		Alias:   ":man-playing-water-polo:",
	},
	"\U0001F93D  :water_polo:": {
		Name:    "Water Polo",
		Unicode: "\U0001F93D",
		Alias:   ":water_polo:",
	},
	"\U0001F93E\u200D\u2640\uFE0F :woman-playing-handball:": {
		Name:    "Woman Playing Handball",
		Unicode: "\U0001F93E\u200D\u2640\uFE0F",
		Alias:   ":woman-playing-handball:",
	},
	"\U0001F93E\u200D\u2642\uFE0F :man-playing-handball:": {
		Name:    "Man Playing Handball",
		Unicode: "\U0001F93E\u200D\u2642\uFE0F",
		Alias:   ":man-playing-handball:",
	},
	"\U0001F93E  :handball:": {
		Name:    "Handball",
		Unicode: "\U0001F93E",
		Alias:   ":handball:",
	},
	"\U0001F93F  :diving_mask:": {
		Name:    "Diving Mask",
		Unicode: "\U0001F93F",
		Alias:   ":diving_mask:",
	},
	"\U0001F940  :wilted_flower:": {
		Name:    "Wilted Flower",
		Unicode: "\U0001F940",
		Alias:   ":wilted_flower:",
	},
	"\U0001F941  :drum_with_drumsticks:": {
		Name:    "Drum With Drumsticks",
		Unicode: "\U0001F941",
		Alias:   ":drum_with_drumsticks:",
	},
	"\U0001F942  :clinking_glasses:": {
		Name:    "Clinking Glasses",
		Unicode: "\U0001F942",
		Alias:   ":clinking_glasses:",
	},
	"\U0001F943  :tumbler_glass:": {
		Name:    "Tumbler Glass",
		Unicode: "\U0001F943",
		Alias:   ":tumbler_glass:",
	},
	"\U0001F944  :spoon:": {
		Name:    "Spoon",
		Unicode: "\U0001F944",
		Alias:   ":spoon:",
	},
	"\U0001F945  :goal_net:": {
		Name:    "Goal Net",
		Unicode: "\U0001F945",
		Alias:   ":goal_net:",
	},
	"\U0001F947  :first_place_medal:": {
		Name:    "First Place Medal",
		Unicode: "\U0001F947",
		Alias:   ":first_place_medal:",
	},
	"\U0001F948  :second_place_medal:": {
		Name:    "Second Place Medal",
		Unicode: "\U0001F948",
		Alias:   ":second_place_medal:",
	},
	"\U0001F949  :third_place_medal:": {
		Name:    "Third Place Medal",
		Unicode: "\U0001F949",
		Alias:   ":third_place_medal:",
	},
	"\U0001F94A  :boxing_glove:": {
		Name:    "Boxing Glove",
		Unicode: "\U0001F94A",
		Alias:   ":boxing_glove:",
	},
	"\U0001F94B  :martial_arts_uniform:": {
		Name:    "Martial Arts Uniform",
		Unicode: "\U0001F94B",
		Alias:   ":martial_arts_uniform:",
	},
	"\U0001F94C  :curling_stone:": {
		Name:    "Curling Stone",
		Unicode: "\U0001F94C",
		Alias:   ":curling_stone:",
	},
	"\U0001F94D  :lacrosse:": {
		Name:    "Lacrosse Stick And Ball",
		Unicode: "\U0001F94D",
		Alias:   ":lacrosse:",
	},
	"\U0001F94E  :softball:": {
		Name:    "Softball",
		Unicode: "\U0001F94E",
		Alias:   ":softball:",
	},
	"\U0001F94F  :flying_disc:": {
		Name:    "Flying Disc",
		Unicode: "\U0001F94F",
		Alias:   ":flying_disc:",
	},
	"\U0001F950  :croissant:": {
		Name:    "Croissant",
		Unicode: "\U0001F950",
		Alias:   ":croissant:",
	},
	"\U0001F951  :avocado:": {
		Name:    "Avocado",
		Unicode: "\U0001F951",
		Alias:   ":avocado:",
	},
	"\U0001F952  :cucumber:": {
		Name:    "Cucumber",
		Unicode: "\U0001F952",
		Alias:   ":cucumber:",
	},
	"\U0001F953  :bacon:": {
		Name:    "Bacon",
		Unicode: "\U0001F953",
		Alias:   ":bacon:",
	},
	"\U0001F954  :potato:": {
		Name:    "Potato",
		Unicode: "\U0001F954",
		Alias:   ":potato:",
	},
	"\U0001F955  :carrot:": {
		Name:    "Carrot",
		Unicode: "\U0001F955",
		Alias:   ":carrot:",
	},
	"\U0001F956  :baguette_bread:": {
		Name:    "Baguette Bread",
		Unicode: "\U0001F956",
		Alias:   ":baguette_bread:",
	},
	"\U0001F957  :green_salad:": {
		Name:    "Green Salad",
		Unicode: "\U0001F957",
		Alias:   ":green_salad:",
	},
	"\U0001F958  :shallow_pan_of_food:": {
		Name:    "Shallow Pan Of Food",
		Unicode: "\U0001F958",
		Alias:   ":shallow_pan_of_food:",
	},
	"\U0001F959  :stuffed_flatbread:": {
		Name:    "Stuffed Flatbread",
		Unicode: "\U0001F959",
		Alias:   ":stuffed_flatbread:",
	},
	"\U0001F95A  :egg:": {
		Name:    "Egg",
		Unicode: "\U0001F95A",
		Alias:   ":egg:",
	},
	"\U0001F95B  :glass_of_milk:": {
		Name:    "Glass Of Milk",
		Unicode: "\U0001F95B",
		Alias:   ":glass_of_milk:",
	},
	"\U0001F95C  :peanuts:": {
		Name:    "Peanuts",
		Unicode: "\U0001F95C",
		Alias:   ":peanuts:",
	},
	"\U0001F95D  :kiwifruit:": {
		Name:    "Kiwifruit",
		Unicode: "\U0001F95D",
		Alias:   ":kiwifruit:",
	},
	"\U0001F95E  :pancakes:": {
		Name:    "Pancakes",
		Unicode: "\U0001F95E",
		Alias:   ":pancakes:",
	},
	"\U0001F95F  :dumpling:": {
		Name:    "Dumpling",
		Unicode: "\U0001F95F",
		Alias:   ":dumpling:",
	},
	"\U0001F960  :fortune_cookie:": {
		Name:    "Fortune Cookie",
		Unicode: "\U0001F960",
		Alias:   ":fortune_cookie:",
	},
	"\U0001F961  :takeout_box:": {
		Name:    "Takeout Box",
		Unicode: "\U0001F961",
		Alias:   ":takeout_box:",
	},
	"\U0001F962  :chopsticks:": {
		Name:    "Chopsticks",
		Unicode: "\U0001F962",
		Alias:   ":chopsticks:",
	},
	"\U0001F963  :bowl_with_spoon:": {
		Name:    "Bowl With Spoon",
		Unicode: "\U0001F963",
		Alias:   ":bowl_with_spoon:",
	},
	"\U0001F964  :cup_with_straw:": {
		Name:    "Cup With Straw",
		Unicode: "\U0001F964",
		Alias:   ":cup_with_straw:",
	},
	"\U0001F965  :coconut:": {
		Name:    "Coconut",
		Unicode: "\U0001F965",
		Alias:   ":coconut:",
	},
	"\U0001F966  :broccoli:": {
		Name:    "Broccoli",
		Unicode: "\U0001F966",
		Alias:   ":broccoli:",
	},
	"\U0001F967  :pie:": {
		Name:    "Pie",
		Unicode: "\U0001F967",
		Alias:   ":pie:",
	},
	"\U0001F968  :pretzel:": {
		Name:    "Pretzel",
		Unicode: "\U0001F968",
		Alias:   ":pretzel:",
	},
	"\U0001F969  :cut_of_meat:": {
		Name:    "Cut Of Meat",
		Unicode: "\U0001F969",
		Alias:   ":cut_of_meat:",
	},
	"\U0001F96A  :sandwich:": {
		Name:    "Sandwich",
		Unicode: "\U0001F96A",
		Alias:   ":sandwich:",
	},
	"\U0001F96B  :canned_food:": {
		Name:    "Canned Food",
		Unicode: "\U0001F96B",
		Alias:   ":canned_food:",
	},
	"\U0001F96C  :leafy_green:": {
		Name:    "Leafy Green",
		Unicode: "\U0001F96C",
		Alias:   ":leafy_green:",
	},
	"\U0001F96D  :mango:": {
		Name:    "Mango",
		Unicode: "\U0001F96D",
		Alias:   ":mango:",
	},
	"\U0001F96E  :moon_cake:": {
		Name:    "Moon Cake",
		Unicode: "\U0001F96E",
		Alias:   ":moon_cake:",
	},
	"\U0001F96F  :bagel:": {
		Name:    "Bagel",
		Unicode: "\U0001F96F",
		Alias:   ":bagel:",
	},
	"\U0001F970  :smiling_face_with_3_hearts:": {
		Name:    "Smiling Face With Smiling Eyes And Three Hearts",
		Unicode: "\U0001F970",
		Alias:   ":smiling_face_with_3_hearts:",
	},
	"\U0001F971  :yawning_face:": {
		Name:    "Yawning Face",
		Unicode: "\U0001F971",
		Alias:   ":yawning_face:",
	},
	"\U0001F972  :smiling_face_with_tear:": {
		Name:    "Smiling Face With Tear",
		Unicode: "\U0001F972",
		Alias:   ":smiling_face_with_tear:",
	},
	"\U0001F973  :partying_face:": {
		Name:    "Face With Party Horn And Party Hat",
		Unicode: "\U0001F973",
		Alias:   ":partying_face:",
	},
	"\U0001F974  :woozy_face:": {
		Name:    "Face With Uneven Eyes And Wavy Mouth",
		Unicode: "\U0001F974",
		Alias:   ":woozy_face:",
	},
	"\U0001F975  :hot_face:": {
		Name:    "Overheated Face",
		Unicode: "\U0001F975",
		Alias:   ":hot_face:",
	},
	"\U0001F976  :cold_face:": {
		Name:    "Freezing Face",
		Unicode: "\U0001F976",
		Alias:   ":cold_face:",
	},
	"\U0001F977  :ninja:": {
		Name:    "Ninja",
		Unicode: "\U0001F977",
		Alias:   ":ninja:",
	},
	"\U0001F978  :disguised_face:": {
		Name:    "Disguised Face",
		Unicode: "\U0001F978",
		Alias:   ":disguised_face:",
	},
	"\U0001F979  :face_holding_back_tears:": {
		Name:    "Face Holding Back Tears",
		Unicode: "\U0001F979",
		Alias:   ":face_holding_back_tears:",
	},
	"\U0001F97A  :pleading_face:": {
		Name:    "Face With Pleading Eyes",
		Unicode: "\U0001F97A",
		Alias:   ":pleading_face:",
	},
	"\U0001F97B  :sari:": {
		Name:    "Sari",
		Unicode: "\U0001F97B",
		Alias:   ":sari:",
	},
	"\U0001F97C  :lab_coat:": {
		Name:    "Lab Coat",
		Unicode: "\U0001F97C",
		Alias:   ":lab_coat:",
	},
	"\U0001F97D  :goggles:": {
		Name:    "Goggles",
		Unicode: "\U0001F97D",
		Alias:   ":goggles:",
	},
	"\U0001F97E  :hiking_boot:": {
		Name:    "Hiking Boot",
		Unicode: "\U0001F97E",
		Alias:   ":hiking_boot:",
	},
	"\U0001F97F  :womans_flat_shoe:": {
		Name:    "Flat Shoe",
		Unicode: "\U0001F97F",
		Alias:   ":womans_flat_shoe:",
	},
	"\U0001F980  :crab:": {
		Name:    "Crab",
		Unicode: "\U0001F980",
		Alias:   ":crab:",
	},
	"\U0001F981  :lion_face:": {
		Name:    "Lion Face",
		Unicode: "\U0001F981",
		Alias:   ":lion_face:",
	},
	"\U0001F982  :scorpion:": {
		Name:    "Scorpion",
		Unicode: "\U0001F982",
		Alias:   ":scorpion:",
	},
	"\U0001F983  :turkey:": {
		Name:    "Turkey",
		Unicode: "\U0001F983",
		Alias:   ":turkey:",
	},
	"\U0001F984  :unicorn_face:": {
		Name:    "Unicorn Face",
		Unicode: "\U0001F984",
		Alias:   ":unicorn_face:",
	},
	"\U0001F985  :eagle:": {
		Name:    "Eagle",
		Unicode: "\U0001F985",
		Alias:   ":eagle:",
	},
	"\U0001F986  :duck:": {
		Name:    "Duck",
		Unicode: "\U0001F986",
		Alias:   ":duck:",
	},
	"\U0001F987  :bat:": {
		Name:    "Bat",
		Unicode: "\U0001F987",
		Alias:   ":bat:",
	},
	"\U0001F988  :shark:": {
		Name:    "Shark",
		Unicode: "\U0001F988",
		Alias:   ":shark:",
	},
	"\U0001F989  :owl:": {
		Name:    "Owl",
		Unicode: "\U0001F989",
		Alias:   ":owl:",
	},
	"\U0001F98A  :fox_face:": {
		Name:    "Fox Face",
		Unicode: "\U0001F98A",
		Alias:   ":fox_face:",
	},
	"\U0001F98B  :butterfly:": {
		Name:    "Butterfly",
		Unicode: "\U0001F98B",
		Alias:   ":butterfly:",
	},
	"\U0001F98C  :deer:": {
		Name:    "Deer",
		Unicode: "\U0001F98C",
		Alias:   ":deer:",
	},
	"\U0001F98D  :gorilla:": {
		Name:    "Gorilla",
		Unicode: "\U0001F98D",
		Alias:   ":gorilla:",
	},
	"\U0001F98E  :lizard:": {
		Name:    "Lizard",
		Unicode: "\U0001F98E",
		Alias:   ":lizard:",
	},
	"\U0001F98F  :rhinoceros:": {
		Name:    "Rhinoceros",
		Unicode: "\U0001F98F",
		Alias:   ":rhinoceros:",
	},
	"\U0001F990  :shrimp:": {
		Name:    "Shrimp",
		Unicode: "\U0001F990",
		Alias:   ":shrimp:",
	},
	"\U0001F991  :squid:": {
		Name:    "Squid",
		Unicode: "\U0001F991",
		Alias:   ":squid:",
	},
	"\U0001F992  :giraffe_face:": {
		Name:    "Giraffe Face",
		Unicode: "\U0001F992",
		Alias:   ":giraffe_face:",
	},
	"\U0001F993  :zebra_face:": {
		Name:    "Zebra Face",
		Unicode: "\U0001F993",
		Alias:   ":zebra_face:",
	},
	"\U0001F994  :hedgehog:": {
		Name:    "Hedgehog",
		Unicode: "\U0001F994",
		Alias:   ":hedgehog:",
	},
	"\U0001F995  :sauropod:": {
		Name:    "Sauropod",
		Unicode: "\U0001F995",
		Alias:   ":sauropod:",
	},
	"\U0001F996 :t-rex:": {
		Name:    "T-Rex",
		Unicode: "\U0001F996",
		Alias:   ":t-rex:",
	},
	"\U0001F997  :cricket:": {
		Name:    "Cricket",
		Unicode: "\U0001F997",
		Alias:   ":cricket:",
	},
	"\U0001F998  :kangaroo:": {
		Name:    "Kangaroo",
		Unicode: "\U0001F998",
		Alias:   ":kangaroo:",
	},
	"\U0001F999  :llama:": {
		Name:    "Llama",
		Unicode: "\U0001F999",
		Alias:   ":llama:",
	},
	"\U0001F99A  :peacock:": {
		Name:    "Peacock",
		Unicode: "\U0001F99A",
		Alias:   ":peacock:",
	},
	"\U0001F99B  :hippopotamus:": {
		Name:    "Hippopotamus",
		Unicode: "\U0001F99B",
		Alias:   ":hippopotamus:",
	},
	"\U0001F99C  :parrot:": {
		Name:    "Parrot",
		Unicode: "\U0001F99C",
		Alias:   ":parrot:",
	},
	"\U0001F99D  :raccoon:": {
		Name:    "Raccoon",
		Unicode: "\U0001F99D",
		Alias:   ":raccoon:",
	},
	"\U0001F99E  :lobster:": {
		Name:    "Lobster",
		Unicode: "\U0001F99E",
		Alias:   ":lobster:",
	},
	"\U0001F99F  :mosquito:": {
		Name:    "Mosquito",
		Unicode: "\U0001F99F",
		Alias:   ":mosquito:",
	},
	"\U0001F9A0  :microbe:": {
		Name:    "Microbe",
		Unicode: "\U0001F9A0",
		Alias:   ":microbe:",
	},
	"\U0001F9A1  :badger:": {
		Name:    "Badger",
		Unicode: "\U0001F9A1",
		Alias:   ":badger:",
	},
	"\U0001F9A2  :swan:": {
		Name:    "Swan",
		Unicode: "\U0001F9A2",
		Alias:   ":swan:",
	},
	"\U0001F9A3  :mammoth:": {
		Name:    "Mammoth",
		Unicode: "\U0001F9A3",
		Alias:   ":mammoth:",
	},
	"\U0001F9A4  :dodo:": {
		Name:    "Dodo",
		Unicode: "\U0001F9A4",
		Alias:   ":dodo:",
	},
	"\U0001F9A5  :sloth:": {
		Name:    "Sloth",
		Unicode: "\U0001F9A5",
		Alias:   ":sloth:",
	},
	"\U0001F9A6  :otter:": {
		Name:    "Otter",
		Unicode: "\U0001F9A6",
		Alias:   ":otter:",
	},
	"\U0001F9A7  :orangutan:": {
		Name:    "Orangutan",
		Unicode: "\U0001F9A7",
		Alias:   ":orangutan:",
	},
	"\U0001F9A8  :skunk:": {
		Name:    "Skunk",
		Unicode: "\U0001F9A8",
		Alias:   ":skunk:",
	},
	"\U0001F9A9  :flamingo:": {
		Name:    "Flamingo",
		Unicode: "\U0001F9A9",
		Alias:   ":flamingo:",
	},
	"\U0001F9AA  :oyster:": {
		Name:    "Oyster",
		Unicode: "\U0001F9AA",
		Alias:   ":oyster:",
	},
	"\U0001F9AB  :beaver:": {
		Name:    "Beaver",
		Unicode: "\U0001F9AB",
		Alias:   ":beaver:",
	},
	"\U0001F9AC  :bison:": {
		Name:    "Bison",
		Unicode: "\U0001F9AC",
		Alias:   ":bison:",
	},
	"\U0001F9AD  :seal:": {
		Name:    "Seal",
		Unicode: "\U0001F9AD",
		Alias:   ":seal:",
	},
	"\U0001F9AE  :guide_dog:": {
		Name:    "Guide Dog",
		Unicode: "\U0001F9AE",
		Alias:   ":guide_dog:",
	},
	"\U0001F9AF  :probing_cane:": {
		Name:    "Probing Cane",
		Unicode: "\U0001F9AF",
		Alias:   ":probing_cane:",
	},
	"\U0001F9B4  :bone:": {
		Name:    "Bone",
		Unicode: "\U0001F9B4",
		Alias:   ":bone:",
	},
	"\U0001F9B5  :leg:": {
		Name:    "Leg",
		Unicode: "\U0001F9B5",
		Alias:   ":leg:",
	},
	"\U0001F9B6  :foot:": {
		Name:    "Foot",
		Unicode: "\U0001F9B6",
		Alias:   ":foot:",
	},
	"\U0001F9B7  :tooth:": {
		Name:    "Tooth",
		Unicode: "\U0001F9B7",
		Alias:   ":tooth:",
	},
	"\U0001F9B8\u200D\u2640\uFE0F  :female_superhero:": {
		Name:    "Woman Superhero",
		Unicode: "\U0001F9B8\u200D\u2640\uFE0F",
		Alias:   ":female_superhero:",
	},
	"\U0001F9B8\u200D\u2642\uFE0F  :male_superhero:": {
		Name:    "Man Superhero",
		Unicode: "\U0001F9B8\u200D\u2642\uFE0F",
		Alias:   ":male_superhero:",
	},
	"\U0001F9B8  :superhero:": {
		Name:    "Superhero",
		Unicode: "\U0001F9B8",
		Alias:   ":superhero:",
	},
	"\U0001F9B9\u200D\u2640\uFE0F  :female_supervillain:": {
		Name:    "Woman Supervillain",
		Unicode: "\U0001F9B9\u200D\u2640\uFE0F",
		Alias:   ":female_supervillain:",
	},
	"\U0001F9B9\u200D\u2642\uFE0F  :male_supervillain:": {
		Name:    "Man Supervillain",
		Unicode: "\U0001F9B9\u200D\u2642\uFE0F",
		Alias:   ":male_supervillain:",
	},
	"\U0001F9B9  :supervillain:": {
		Name:    "Supervillain",
		Unicode: "\U0001F9B9",
		Alias:   ":supervillain:",
	},
	"\U0001F9BA  :safety_vest:": {
		Name:    "Safety Vest",
		Unicode: "\U0001F9BA",
		Alias:   ":safety_vest:",
	},
	"\U0001F9BB  :ear_with_hearing_aid:": {
		Name:    "Ear With Hearing Aid",
		Unicode: "\U0001F9BB",
		Alias:   ":ear_with_hearing_aid:",
	},
	"\U0001F9BC  :motorized_wheelchair:": {
		Name:    "Motorized Wheelchair",
		Unicode: "\U0001F9BC",
		Alias:   ":motorized_wheelchair:",
	},
	"\U0001F9BD  :manual_wheelchair:": {
		Name:    "Manual Wheelchair",
		Unicode: "\U0001F9BD",
		Alias:   ":manual_wheelchair:",
	},
	"\U0001F9BE  :mechanical_arm:": {
		Name:    "Mechanical Arm",
		Unicode: "\U0001F9BE",
		Alias:   ":mechanical_arm:",
	},
	"\U0001F9BF  :mechanical_leg:": {
		Name:    "Mechanical Leg",
		Unicode: "\U0001F9BF",
		Alias:   ":mechanical_leg:",
	},
	"\U0001F9C0  :cheese_wedge:": {
		Name:    "Cheese Wedge",
		Unicode: "\U0001F9C0",
		Alias:   ":cheese_wedge:",
	},
	"\U0001F9C1  :cupcake:": {
		Name:    "Cupcake",
		Unicode: "\U0001F9C1",
		Alias:   ":cupcake:",
	},
	"\U0001F9C2  :salt:": {
		Name:    "Salt Shaker",
		Unicode: "\U0001F9C2",
		Alias:   ":salt:",
	},
	"\U0001F9C3  :beverage_box:": {
		Name:    "Beverage Box",
		Unicode: "\U0001F9C3",
		Alias:   ":beverage_box:",
	},
	"\U0001F9C4  :garlic:": {
		Name:    "Garlic",
		Unicode: "\U0001F9C4",
		Alias:   ":garlic:",
	},
	"\U0001F9C5  :onion:": {
		Name:    "Onion",
		Unicode: "\U0001F9C5",
		Alias:   ":onion:",
	},
	"\U0001F9C6  :falafel:": {
		Name:    "Falafel",
		Unicode: "\U0001F9C6",
		Alias:   ":falafel:",
	},
	"\U0001F9C7  :waffle:": {
		Name:    "Waffle",
		Unicode: "\U0001F9C7",
		Alias:   ":waffle:",
	},
	"\U0001F9C8  :butter:": {
		Name:    "Butter",
		Unicode: "\U0001F9C8",
		Alias:   ":butter:",
	},
	"\U0001F9C9  :mate_drink:": {
		Name:    "Mate Drink",
		Unicode: "\U0001F9C9",
		Alias:   ":mate_drink:",
	},
	"\U0001F9CA  :ice_cube:": {
		Name:    "Ice Cube",
		Unicode: "\U0001F9CA",
		Alias:   ":ice_cube:",
	},
	"\U0001F9CB  :bubble_tea:": {
		Name:    "Bubble Tea",
		Unicode: "\U0001F9CB",
		Alias:   ":bubble_tea:",
	},
	"\U0001F9CC  :troll:": {
		Name:    "Troll",
		Unicode: "\U0001F9CC",
		Alias:   ":troll:",
	},
	"\U0001F9CD\u200D\u2640\uFE0F  :woman_standing:": {
		Name:    "Woman Standing",
		Unicode: "\U0001F9CD\u200D\u2640\uFE0F",
		Alias:   ":woman_standing:",
	},
	"\U0001F9CD\u200D\u2642\uFE0F  :man_standing:": {
		Name:    "Man Standing",
		Unicode: "\U0001F9CD\u200D\u2642\uFE0F",
		Alias:   ":man_standing:",
	},
	"\U0001F9CD  :standing_person:": {
		Name:    "Standing Person",
		Unicode: "\U0001F9CD",
		Alias:   ":standing_person:",
	},
	"\U0001F9CE\u200D\u2640\uFE0F  :woman_kneeling:": {
		Name:    "Woman Kneeling",
		Unicode: "\U0001F9CE\u200D\u2640\uFE0F",
		Alias:   ":woman_kneeling:",
	},
	"\U0001F9CE\u200D\u2642\uFE0F  :man_kneeling:": {
		Name:    "Man Kneeling",
		Unicode: "\U0001F9CE\u200D\u2642\uFE0F",
		Alias:   ":man_kneeling:",
	},
	"\U0001F9CE  :kneeling_person:": {
		Name:    "Kneeling Person",
		Unicode: "\U0001F9CE",
		Alias:   ":kneeling_person:",
	},
	"\U0001F9CF\u200D\u2640\uFE0F  :deaf_woman:": {
		Name:    "Deaf Woman",
		Unicode: "\U0001F9CF\u200D\u2640\uFE0F",
		Alias:   ":deaf_woman:",
	},
	"\U0001F9CF\u200D\u2642\uFE0F  :deaf_man:": {
		Name:    "Deaf Man",
		Unicode: "\U0001F9CF\u200D\u2642\uFE0F",
		Alias:   ":deaf_man:",
	},
	"\U0001F9CF  :deaf_person:": {
		Name:    "Deaf Person",
		Unicode: "\U0001F9CF",
		Alias:   ":deaf_person:",
	},
	"\U0001F9D0  :face_with_monocle:": {
		Name:    "Face With Monocle",
		Unicode: "\U0001F9D0",
		Alias:   ":face_with_monocle:",
	},
	"\U0001F9D1\u200D\u1F33E  :farmer:": {
		Name:    "Farmer",
		Unicode: "\U0001F9D1\u200D\u1F33E",
		Alias:   ":farmer:",
	},
	"\U0001F9D1\u200D\u1F373  :cook:": {
		Name:    "Cook",
		Unicode: "\U0001F9D1\u200D\u1F373",
		Alias:   ":cook:",
	},
	"\U0001F9D1\u200D\u1F37C  :person_feeding_baby:": {
		Name:    "Person Feeding Baby",
		Unicode: "\U0001F9D1\u200D\u1F37C",
		Alias:   ":person_feeding_baby:",
	},
	"\U0001F9D1\u200D\u1F384  :mx_claus:": {
		Name:    "Mx Claus",
		Unicode: "\U0001F9D1\u200D\u1F384",
		Alias:   ":mx_claus:",
	},
	"\U0001F9D1\u200D\u1F393  :student:": {
		Name:    "Student",
		Unicode: "\U0001F9D1\u200D\u1F393",
		Alias:   ":student:",
	},
	"\U0001F9D1\u200D\u1F3A4  :singer:": {
		Name:    "Singer",
		Unicode: "\U0001F9D1\u200D\u1F3A4",
		Alias:   ":singer:",
	},
	"\U0001F9D1\u200D\u1F3A8  :artist:": {
		Name:    "Artist",
		Unicode: "\U0001F9D1\u200D\u1F3A8",
		Alias:   ":artist:",
	},
	"\U0001F9D1\u200D\u1F3EB  :teacher:": {
		Name:    "Teacher",
		Unicode: "\U0001F9D1\u200D\u1F3EB",
		Alias:   ":teacher:",
	},
	"\U0001F9D1\u200D\u1F3ED  :factory_worker:": {
		Name:    "Factory Worker",
		Unicode: "\U0001F9D1\u200D\u1F3ED",
		Alias:   ":factory_worker:",
	},
	"\U0001F9D1\u200D\u1F4BB  :technologist:": {
		Name:    "Technologist",
		Unicode: "\U0001F9D1\u200D\u1F4BB",
		Alias:   ":technologist:",
	},
	"\U0001F9D1\u200D\u1F4BC  :office_worker:": {
		Name:    "Office Worker",
		Unicode: "\U0001F9D1\u200D\u1F4BC",
		Alias:   ":office_worker:",
	},
	"\U0001F9D1\u200D\u1F527  :mechanic:": {
		Name:    "Mechanic",
		Unicode: "\U0001F9D1\u200D\u1F527",
		Alias:   ":mechanic:",
	},
	"\U0001F9D1\u200D\u1F52C  :scientist:": {
		Name:    "Scientist",
		Unicode: "\U0001F9D1\u200D\u1F52C",
		Alias:   ":scientist:",
	},
	"\U0001F9D1\u200D\u1F680  :astronaut:": {
		Name:    "Astronaut",
		Unicode: "\U0001F9D1\u200D\u1F680",
		Alias:   ":astronaut:",
	},
	"\U0001F9D1\u200D\u1F692  :firefighter:": {
		Name:    "Firefighter",
		Unicode: "\U0001F9D1\u200D\u1F692",
		Alias:   ":firefighter:",
	},
	"\U0001F9D1\u200D\u1F91D\u200D\u1F9D1  :people_holding_hands:": {
		Name:    "People Holding Hands",
		Unicode: "\U0001F9D1\u200D\u1F91D\u200D\u1F9D1",
		Alias:   ":people_holding_hands:",
	},
	"\U0001F9D1\u200D\u1F9AF  :person_with_probing_cane:": {
		Name:    "Person With White Cane",
		Unicode: "\U0001F9D1\u200D\u1F9AF",
		Alias:   ":person_with_probing_cane:",
	},
	"\U0001F9D1\u200D\u1F9B0  :red_haired_person:": {
		Name:    "Person: Red Hair",
		Unicode: "\U0001F9D1\u200D\u1F9B0",
		Alias:   ":red_haired_person:",
	},
	"\U0001F9D1\u200D\u1F9B1  :curly_haired_person:": {
		Name:    "Person: Curly Hair",
		Unicode: "\U0001F9D1\u200D\u1F9B1",
		Alias:   ":curly_haired_person:",
	},
	"\U0001F9D1\u200D\u1F9B2  :bald_person:": {
		Name:    "Person: Bald",
		Unicode: "\U0001F9D1\u200D\u1F9B2",
		Alias:   ":bald_person:",
	},
	"\U0001F9D1\u200D\u1F9B3  :white_haired_person:": {
		Name:    "Person: White Hair",
		Unicode: "\U0001F9D1\u200D\u1F9B3",
		Alias:   ":white_haired_person:",
	},
	"\U0001F9D1\u200D\u1F9BC  :person_in_motorized_wheelchair:": {
		Name:    "Person In Motorized Wheelchair",
		Unicode: "\U0001F9D1\u200D\u1F9BC",
		Alias:   ":person_in_motorized_wheelchair:",
	},
	"\U0001F9D1\u200D\u1F9BD  :person_in_manual_wheelchair:": {
		Name:    "Person In Manual Wheelchair",
		Unicode: "\U0001F9D1\u200D\u1F9BD",
		Alias:   ":person_in_manual_wheelchair:",
	},
	"\U0001F9D1\u200D\u2695\uFE0F  :health_worker:": {
		Name:    "Health Worker",
		Unicode: "\U0001F9D1\u200D\u2695\uFE0F",
		Alias:   ":health_worker:",
	},
	"\U0001F9D1\u200D\u2696\uFE0F  :judge:": {
		Name:    "Judge",
		Unicode: "\U0001F9D1\u200D\u2696\uFE0F",
		Alias:   ":judge:",
	},
	"\U0001F9D1\u200D\u2708\uFE0F  :pilot:": {
		Name:    "Pilot",
		Unicode: "\U0001F9D1\u200D\u2708\uFE0F",
		Alias:   ":pilot:",
	},
	"\U0001F9D1  :adult:": {
		Name:    "Adult",
		Unicode: "\U0001F9D1",
		Alias:   ":adult:",
	},
	"\U0001F9D2  :child:": {
		Name:    "Child",
		Unicode: "\U0001F9D2",
		Alias:   ":child:",
	},
	"\U0001F9D3  :older_adult:": {
		Name:    "Older Adult",
		Unicode: "\U0001F9D3",
		Alias:   ":older_adult:",
	},
	"\U0001F9D4\u200D\u2640\uFE0F  :woman_with_beard:": {
		Name:    "Woman: Beard",
		Unicode: "\U0001F9D4\u200D\u2640\uFE0F",
		Alias:   ":woman_with_beard:",
	},
	"\U0001F9D4\u200D\u2642\uFE0F  :man_with_beard:": {
		Name:    "Man: Beard",
		Unicode: "\U0001F9D4\u200D\u2642\uFE0F",
		Alias:   ":man_with_beard:",
	},
	"\U0001F9D4  :bearded_person:": {
		Name:    "Bearded Person",
		Unicode: "\U0001F9D4",
		Alias:   ":bearded_person:",
	},
	"\U0001F9D5  :person_with_headscarf:": {
		Name:    "Person With Headscarf",
		Unicode: "\U0001F9D5",
		Alias:   ":person_with_headscarf:",
	},
	"\U0001F9D6\u200D\u2640\uFE0F  :woman_in_steamy_room:": {
		Name:    "Woman In Steamy Room",
		Unicode: "\U0001F9D6\u200D\u2640\uFE0F",
		Alias:   ":woman_in_steamy_room:",
	},
	"\U0001F9D6\u200D\u2642\uFE0F  :man_in_steamy_room:": {
		Name:    "Man In Steamy Room",
		Unicode: "\U0001F9D6\u200D\u2642\uFE0F",
		Alias:   ":man_in_steamy_room:",
	},
	"\U0001F9D6  :person_in_steamy_room:": {
		Name:    "Person In Steamy Room",
		Unicode: "\U0001F9D6",
		Alias:   ":person_in_steamy_room:",
	},
	"\U0001F9D7\u200D\u2640\uFE0F  :woman_climbing:": {
		Name:    "Woman Climbing",
		Unicode: "\U0001F9D7\u200D\u2640\uFE0F",
		Alias:   ":woman_climbing:",
	},
	"\U0001F9D7\u200D\u2642\uFE0F  :man_climbing:": {
		Name:    "Man Climbing",
		Unicode: "\U0001F9D7\u200D\u2642\uFE0F",
		Alias:   ":man_climbing:",
	},
	"\U0001F9D7  :person_climbing:": {
		Name:    "Person Climbing",
		Unicode: "\U0001F9D7",
		Alias:   ":person_climbing:",
	},
	"\U0001F9D8\u200D\u2640\uFE0F  :woman_in_lotus_position:": {
		Name:    "Woman In Lotus Position",
		Unicode: "\U0001F9D8\u200D\u2640\uFE0F",
		Alias:   ":woman_in_lotus_position:",
	},
	"\U0001F9D8\u200D\u2642\uFE0F  :man_in_lotus_position:": {
		Name:    "Man In Lotus Position",
		Unicode: "\U0001F9D8\u200D\u2642\uFE0F",
		Alias:   ":man_in_lotus_position:",
	},
	"\U0001F9D8  :person_in_lotus_position:": {
		Name:    "Person In Lotus Position",
		Unicode: "\U0001F9D8",
		Alias:   ":person_in_lotus_position:",
	},
	"\U0001F9D9\u200D\u2640\uFE0F  :female_mage:": {
		Name:    "Woman Mage",
		Unicode: "\U0001F9D9\u200D\u2640\uFE0F",
		Alias:   ":female_mage:",
	},
	"\U0001F9D9\u200D\u2642\uFE0F  :male_mage:": {
		Name:    "Man Mage",
		Unicode: "\U0001F9D9\u200D\u2642\uFE0F",
		Alias:   ":male_mage:",
	},
	"\U0001F9D9  :mage:": {
		Name:    "Mage",
		Unicode: "\U0001F9D9",
		Alias:   ":mage:",
	},
	"\U0001F9DA\u200D\u2640\uFE0F  :female_fairy:": {
		Name:    "Woman Fairy",
		Unicode: "\U0001F9DA\u200D\u2640\uFE0F",
		Alias:   ":female_fairy:",
	},
	"\U0001F9DA\u200D\u2642\uFE0F  :male_fairy:": {
		Name:    "Man Fairy",
		Unicode: "\U0001F9DA\u200D\u2642\uFE0F",
		Alias:   ":male_fairy:",
	},
	"\U0001F9DA  :fairy:": {
		Name:    "Fairy",
		Unicode: "\U0001F9DA",
		Alias:   ":fairy:",
	},
	"\U0001F9DB\u200D\u2640\uFE0F  :female_vampire:": {
		Name:    "Woman Vampire",
		Unicode: "\U0001F9DB\u200D\u2640\uFE0F",
		Alias:   ":female_vampire:",
	},
	"\U0001F9DB\u200D\u2642\uFE0F  :male_vampire:": {
		Name:    "Man Vampire",
		Unicode: "\U0001F9DB\u200D\u2642\uFE0F",
		Alias:   ":male_vampire:",
	},
	"\U0001F9DB  :vampire:": {
		Name:    "Vampire",
		Unicode: "\U0001F9DB",
		Alias:   ":vampire:",
	},
	"\U0001F9DC\u200D\u2640\uFE0F  :mermaid:": {
		Name:    "Mermaid",
		Unicode: "\U0001F9DC\u200D\u2640\uFE0F",
		Alias:   ":mermaid:",
	},
	"\U0001F9DC\u200D\u2642\uFE0F  :merman:": {
		Name:    "Merman",
		Unicode: "\U0001F9DC\u200D\u2642\uFE0F",
		Alias:   ":merman:",
	},
	"\U0001F9DC  :merperson:": {
		Name:    "Merperson",
		Unicode: "\U0001F9DC",
		Alias:   ":merperson:",
	},
	"\U0001F9DD\u200D\u2640\uFE0F  :female_elf:": {
		Name:    "Woman Elf",
		Unicode: "\U0001F9DD\u200D\u2640\uFE0F",
		Alias:   ":female_elf:",
	},
	"\U0001F9DD\u200D\u2642\uFE0F  :male_elf:": {
		Name:    "Man Elf",
		Unicode: "\U0001F9DD\u200D\u2642\uFE0F",
		Alias:   ":male_elf:",
	},
	"\U0001F9DD  :elf:": {
		Name:    "Elf",
		Unicode: "\U0001F9DD",
		Alias:   ":elf:",
	},
	"\U0001F9DE\u200D\u2640\uFE0F  :female_genie:": {
		Name:    "Woman Genie",
		Unicode: "\U0001F9DE\u200D\u2640\uFE0F",
		Alias:   ":female_genie:",
	},
	"\U0001F9DE\u200D\u2642\uFE0F  :male_genie:": {
		Name:    "Man Genie",
		Unicode: "\U0001F9DE\u200D\u2642\uFE0F",
		Alias:   ":male_genie:",
	},
	"\U0001F9DE  :genie:": {
		Name:    "Genie",
		Unicode: "\U0001F9DE",
		Alias:   ":genie:",
	},
	"\U0001F9DF\u200D\u2640\uFE0F  :female_zombie:": {
		Name:    "Woman Zombie",
		Unicode: "\U0001F9DF\u200D\u2640\uFE0F",
		Alias:   ":female_zombie:",
	},
	"\U0001F9DF\u200D\u2642\uFE0F  :male_zombie:": {
		Name:    "Man Zombie",
		Unicode: "\U0001F9DF\u200D\u2642\uFE0F",
		Alias:   ":male_zombie:",
	},
	"\U0001F9DF  :zombie:": {
		Name:    "Zombie",
		Unicode: "\U0001F9DF",
		Alias:   ":zombie:",
	},
	"\U0001F9E0  :brain:": {
		Name:    "Brain",
		Unicode: "\U0001F9E0",
		Alias:   ":brain:",
	},
	"\U0001F9E1  :orange_heart:": {
		Name:    "Orange Heart",
		Unicode: "\U0001F9E1",
		Alias:   ":orange_heart:",
	},
	"\U0001F9E2  :billed_cap:": {
		Name:    "Billed Cap",
		Unicode: "\U0001F9E2",
		Alias:   ":billed_cap:",
	},
	"\U0001F9E3  :scarf:": {
		Name:    "Scarf",
		Unicode: "\U0001F9E3",
		Alias:   ":scarf:",
	},
	"\U0001F9E4  :gloves:": {
		Name:    "Gloves",
		Unicode: "\U0001F9E4",
		Alias:   ":gloves:",
	},
	"\U0001F9E5  :coat:": {
		Name:    "Coat",
		Unicode: "\U0001F9E5",
		Alias:   ":coat:",
	},
	"\U0001F9E6  :socks:": {
		Name:    "Socks",
		Unicode: "\U0001F9E6",
		Alias:   ":socks:",
	},
	"\U0001F9E7  :red_envelope:": {
		Name:    "Red Gift Envelope",
		Unicode: "\U0001F9E7",
		Alias:   ":red_envelope:",
	},
	"\U0001F9E8  :firecracker:": {
		Name:    "Firecracker",
		Unicode: "\U0001F9E8",
		Alias:   ":firecracker:",
	},
	"\U0001F9E9  :jigsaw:": {
		Name:    "Jigsaw Puzzle Piece",
		Unicode: "\U0001F9E9",
		Alias:   ":jigsaw:",
	},
	"\U0001F9EA  :test_tube:": {
		Name:    "Test Tube",
		Unicode: "\U0001F9EA",
		Alias:   ":test_tube:",
	},
	"\U0001F9EB  :petri_dish:": {
		Name:    "Petri Dish",
		Unicode: "\U0001F9EB",
		Alias:   ":petri_dish:",
	},
	"\U0001F9EC  :dna:": {
		Name:    "Dna Double Helix",
		Unicode: "\U0001F9EC",
		Alias:   ":dna:",
	},
	"\U0001F9ED  :compass:": {
		Name:    "Compass",
		Unicode: "\U0001F9ED",
		Alias:   ":compass:",
	},
	"\U0001F9EE  :abacus:": {
		Name:    "Abacus",
		Unicode: "\U0001F9EE",
		Alias:   ":abacus:",
	},
	"\U0001F9EF  :fire_extinguisher:": {
		Name:    "Fire Extinguisher",
		Unicode: "\U0001F9EF",
		Alias:   ":fire_extinguisher:",
	},
	"\U0001F9F0  :toolbox:": {
		Name:    "Toolbox",
		Unicode: "\U0001F9F0",
		Alias:   ":toolbox:",
	},
	"\U0001F9F1  :bricks:": {
		Name:    "Brick",
		Unicode: "\U0001F9F1",
		Alias:   ":bricks:",
	},
	"\U0001F9F2  :magnet:": {
		Name:    "Magnet",
		Unicode: "\U0001F9F2",
		Alias:   ":magnet:",
	},
	"\U0001F9F3  :luggage:": {
		Name:    "Luggage",
		Unicode: "\U0001F9F3",
		Alias:   ":luggage:",
	},
	"\U0001F9F4  :lotion_bottle:": {
		Name:    "Lotion Bottle",
		Unicode: "\U0001F9F4",
		Alias:   ":lotion_bottle:",
	},
	"\U0001F9F5  :thread:": {
		Name:    "Spool Of Thread",
		Unicode: "\U0001F9F5",
		Alias:   ":thread:",
	},
	"\U0001F9F6  :yarn:": {
		Name:    "Ball Of Yarn",
		Unicode: "\U0001F9F6",
		Alias:   ":yarn:",
	},
	"\U0001F9F7  :safety_pin:": {
		Name:    "Safety Pin",
		Unicode: "\U0001F9F7",
		Alias:   ":safety_pin:",
	},
	"\U0001F9F8  :teddy_bear:": {
		Name:    "Teddy Bear",
		Unicode: "\U0001F9F8",
		Alias:   ":teddy_bear:",
	},
	"\U0001F9F9  :broom:": {
		Name:    "Broom",
		Unicode: "\U0001F9F9",
		Alias:   ":broom:",
	},
	"\U0001F9FA  :basket:": {
		Name:    "Basket",
		Unicode: "\U0001F9FA",
		Alias:   ":basket:",
	},
	"\U0001F9FB  :roll_of_paper:": {
		Name:    "Roll Of Paper",
		Unicode: "\U0001F9FB",
		Alias:   ":roll_of_paper:",
	},
	"\U0001F9FC  :soap:": {
		Name:    "Bar Of Soap",
		Unicode: "\U0001F9FC",
		Alias:   ":soap:",
	},
	"\U0001F9FD  :sponge:": {
		Name:    "Sponge",
		Unicode: "\U0001F9FD",
		Alias:   ":sponge:",
	},
	"\U0001F9FE  :receipt:": {
		Name:    "Receipt",
		Unicode: "\U0001F9FE",
		Alias:   ":receipt:",
	},
	"\U0001F9FF  :nazar_amulet:": {
		Name:    "Nazar Amulet",
		Unicode: "\U0001F9FF",
		Alias:   ":nazar_amulet:",
	},
	"\U0001FA70  :ballet_shoes:": {
		Name:    "Ballet Shoes",
		Unicode: "\U0001FA70",
		Alias:   ":ballet_shoes:",
	},
	"\U0001FA71 :one-piece_swimsuit:": {
		Name:    "One-Piece Swimsuit",
		Unicode: "\U0001FA71",
		Alias:   ":one-piece_swimsuit:",
	},
	"\U0001FA72  :briefs:": {
		Name:    "Briefs",
		Unicode: "\U0001FA72",
		Alias:   ":briefs:",
	},
	"\U0001FA73  :shorts:": {
		Name:    "Shorts",
		Unicode: "\U0001FA73",
		Alias:   ":shorts:",
	},
	"\U0001FA74  :thong_sandal:": {
		Name:    "Thong Sandal",
		Unicode: "\U0001FA74",
		Alias:   ":thong_sandal:",
	},
	"\U0001FA78  :drop_of_blood:": {
		Name:    "Drop Of Blood",
		Unicode: "\U0001FA78",
		Alias:   ":drop_of_blood:",
	},
	"\U0001FA79  :adhesive_bandage:": {
		Name:    "Adhesive Bandage",
		Unicode: "\U0001FA79",
		Alias:   ":adhesive_bandage:",
	},
	"\U0001FA7A  :stethoscope:": {
		Name:    "Stethoscope",
		Unicode: "\U0001FA7A",
		Alias:   ":stethoscope:",
	},
	"\U0001FA7B :x-ray:": {
		Name:    "X-Ray",
		Unicode: "\U0001FA7B",
		Alias:   ":x-ray:",
	},
	"\U0001FA7C  :crutch:": {
		Name:    "Crutch",
		Unicode: "\U0001FA7C",
		Alias:   ":crutch:",
	},
	"\U0001FA80 :yo-yo:": {
		Name:    "Yo-Yo",
		Unicode: "\U0001FA80",
		Alias:   ":yo-yo:",
	},
	"\U0001FA81  :kite:": {
		Name:    "Kite",
		Unicode: "\U0001FA81",
		Alias:   ":kite:",
	},
	"\U0001FA82  :parachute:": {
		Name:    "Parachute",
		Unicode: "\U0001FA82",
		Alias:   ":parachute:",
	},
	"\U0001FA83  :boomerang:": {
		Name:    "Boomerang",
		Unicode: "\U0001FA83",
		Alias:   ":boomerang:",
	},
	"\U0001FA84  :magic_wand:": {
		Name:    "Magic Wand",
		Unicode: "\U0001FA84",
		Alias:   ":magic_wand:",
	},
	"\U0001FA85  :pinata:": {
		Name:    "Pinata",
		Unicode: "\U0001FA85",
		Alias:   ":pinata:",
	},
	"\U0001FA86  :nesting_dolls:": {
		Name:    "Nesting Dolls",
		Unicode: "\U0001FA86",
		Alias:   ":nesting_dolls:",
	},
	"\U0001FA90  :ringed_planet:": {
		Name:    "Ringed Planet",
		Unicode: "\U0001FA90",
		Alias:   ":ringed_planet:",
	},
	"\U0001FA91  :chair:": {
		Name:    "Chair",
		Unicode: "\U0001FA91",
		Alias:   ":chair:",
	},
	"\U0001FA92  :razor:": {
		Name:    "Razor",
		Unicode: "\U0001FA92",
		Alias:   ":razor:",
	},
	"\U0001FA93  :axe:": {
		Name:    "Axe",
		Unicode: "\U0001FA93",
		Alias:   ":axe:",
	},
	"\U0001FA94  :diya_lamp:": {
		Name:    "Diya Lamp",
		Unicode: "\U0001FA94",
		Alias:   ":diya_lamp:",
	},
	"\U0001FA95  :banjo:": {
		Name:    "Banjo",
		Unicode: "\U0001FA95",
		Alias:   ":banjo:",
	},
	"\U0001FA96  :military_helmet:": {
		Name:    "Military Helmet",
		Unicode: "\U0001FA96",
		Alias:   ":military_helmet:",
	},
	"\U0001FA97  :accordion:": {
		Name:    "Accordion",
		Unicode: "\U0001FA97",
		Alias:   ":accordion:",
	},
	"\U0001FA98  :long_drum:": {
		Name:    "Long Drum",
		Unicode: "\U0001FA98",
		Alias:   ":long_drum:",
	},
	"\U0001FA99  :coin:": {
		Name:    "Coin",
		Unicode: "\U0001FA99",
		Alias:   ":coin:",
	},
	"\U0001FA9A  :carpentry_saw:": {
		Name:    "Carpentry Saw",
		Unicode: "\U0001FA9A",
		Alias:   ":carpentry_saw:",
	},
	"\U0001FA9B  :screwdriver:": {
		Name:    "Screwdriver",
		Unicode: "\U0001FA9B",
		Alias:   ":screwdriver:",
	},
	"\U0001FA9C  :ladder:": {
		Name:    "Ladder",
		Unicode: "\U0001FA9C",
		Alias:   ":ladder:",
	},
	"\U0001FA9D  :hook:": {
		Name:    "Hook",
		Unicode: "\U0001FA9D",
		Alias:   ":hook:",
	},
	"\U0001FA9E  :mirror:": {
		Name:    "Mirror",
		Unicode: "\U0001FA9E",
		Alias:   ":mirror:",
	},
	"\U0001FA9F  :window:": {
		Name:    "Window",
		Unicode: "\U0001FA9F",
		Alias:   ":window:",
	},
	"\U0001FAA0  :plunger:": {
		Name:    "Plunger",
		Unicode: "\U0001FAA0",
		Alias:   ":plunger:",
	},
	"\U0001FAA1  :sewing_needle:": {
		Name:    "Sewing Needle",
		Unicode: "\U0001FAA1",
		Alias:   ":sewing_needle:",
	},
	"\U0001FAA2  :knot:": {
		Name:    "Knot",
		Unicode: "\U0001FAA2",
		Alias:   ":knot:",
	},
	"\U0001FAA3  :bucket:": {
		Name:    "Bucket",
		Unicode: "\U0001FAA3",
		Alias:   ":bucket:",
	},
	"\U0001FAA4  :mouse_trap:": {
		Name:    "Mouse Trap",
		Unicode: "\U0001FAA4",
		Alias:   ":mouse_trap:",
	},
	"\U0001FAA5  :toothbrush:": {
		Name:    "Toothbrush",
		Unicode: "\U0001FAA5",
		Alias:   ":toothbrush:",
	},
	"\U0001FAA6  :headstone:": {
		Name:    "Headstone",
		Unicode: "\U0001FAA6",
		Alias:   ":headstone:",
	},
	"\U0001FAA7  :placard:": {
		Name:    "Placard",
		Unicode: "\U0001FAA7",
		Alias:   ":placard:",
	},
	"\U0001FAA8  :rock:": {
		Name:    "Rock",
		Unicode: "\U0001FAA8",
		Alias:   ":rock:",
	},
	"\U0001FAA9  :mirror_ball:": {
		Name:    "Mirror Ball",
		Unicode: "\U0001FAA9",
		Alias:   ":mirror_ball:",
	},
	"\U0001FAAA  :identification_card:": {
		Name:    "Identification Card",
		Unicode: "\U0001FAAA",
		Alias:   ":identification_card:",
	},
	"\U0001FAAB  :low_battery:": {
		Name:    "Low Battery",
		Unicode: "\U0001FAAB",
		Alias:   ":low_battery:",
	},
	"\U0001FAAC  :hamsa:": {
		Name:    "Hamsa",
		Unicode: "\U0001FAAC",
		Alias:   ":hamsa:",
	},
	"\U0001FAB0  :fly:": {
		Name:    "Fly",
		Unicode: "\U0001FAB0",
		Alias:   ":fly:",
	},
	"\U0001FAB1  :worm:": {
		Name:    "Worm",
		Unicode: "\U0001FAB1",
		Alias:   ":worm:",
	},
	"\U0001FAB2  :beetle:": {
		Name:    "Beetle",
		Unicode: "\U0001FAB2",
		Alias:   ":beetle:",
	},
	"\U0001FAB3  :cockroach:": {
		Name:    "Cockroach",
		Unicode: "\U0001FAB3",
		Alias:   ":cockroach:",
	},
	"\U0001FAB4  :potted_plant:": {
		Name:    "Potted Plant",
		Unicode: "\U0001FAB4",
		Alias:   ":potted_plant:",
	},
	"\U0001FAB5  :wood:": {
		Name:    "Wood",
		Unicode: "\U0001FAB5",
		Alias:   ":wood:",
	},
	"\U0001FAB6  :feather:": {
		Name:    "Feather",
		Unicode: "\U0001FAB6",
		Alias:   ":feather:",
	},
	"\U0001FAB7  :lotus:": {
		Name:    "Lotus",
		Unicode: "\U0001FAB7",
		Alias:   ":lotus:",
	},
	"\U0001FAB8  :coral:": {
		Name:    "Coral",
		Unicode: "\U0001FAB8",
		Alias:   ":coral:",
	},
	"\U0001FAB9  :empty_nest:": {
		Name:    "Empty Nest",
		Unicode: "\U0001FAB9",
		Alias:   ":empty_nest:",
	},
	"\U0001FABA  :nest_with_eggs:": {
		Name:    "Nest With Eggs",
		Unicode: "\U0001FABA",
		Alias:   ":nest_with_eggs:",
	},
	"\U0001FAC0  :anatomical_heart:": {
		Name:    "Anatomical Heart",
		Unicode: "\U0001FAC0",
		Alias:   ":anatomical_heart:",
	},
	"\U0001FAC1  :lungs:": {
		Name:    "Lungs",
		Unicode: "\U0001FAC1",
		Alias:   ":lungs:",
	},
	"\U0001FAC2  :people_hugging:": {
		Name:    "People Hugging",
		Unicode: "\U0001FAC2",
		Alias:   ":people_hugging:",
	},
	"\U0001FAC3  :pregnant_man:": {
		Name:    "Pregnant Man",
		Unicode: "\U0001FAC3",
		Alias:   ":pregnant_man:",
	},
	"\U0001FAC4  :pregnant_person:": {
		Name:    "Pregnant Person",
		Unicode: "\U0001FAC4",
		Alias:   ":pregnant_person:",
	},
	"\U0001FAC5  :person_with_crown:": {
		Name:    "Person With Crown",
		Unicode: "\U0001FAC5",
		Alias:   ":person_with_crown:",
	},
	"\U0001FAD0  :blueberries:": {
		Name:    "Blueberries",
		Unicode: "\U0001FAD0",
		Alias:   ":blueberries:",
	},
	"\U0001FAD1  :bell_pepper:": {
		Name:    "Bell Pepper",
		Unicode: "\U0001FAD1",
		Alias:   ":bell_pepper:",
	},
	"\U0001FAD2  :olive:": {
		Name:    "Olive",
		Unicode: "\U0001FAD2",
		Alias:   ":olive:",
	},
	"\U0001FAD3  :flatbread:": {
		Name:    "Flatbread",
		Unicode: "\U0001FAD3",
		Alias:   ":flatbread:",
	},
	"\U0001FAD4  :tamale:": {
		Name:    "Tamale",
		Unicode: "\U0001FAD4",
		Alias:   ":tamale:",
	},
	"\U0001FAD5  :fondue:": {
		Name:    "Fondue",
		Unicode: "\U0001FAD5",
		Alias:   ":fondue:",
	},
	"\U0001FAD6  :teapot:": {
		Name:    "Teapot",
		Unicode: "\U0001FAD6",
		Alias:   ":teapot:",
	},
	"\U0001FAD7  :pouring_liquid:": {
		Name:    "Pouring Liquid",
		Unicode: "\U0001FAD7",
		Alias:   ":pouring_liquid:",
	},
	"\U0001FAD8  :beans:": {
		Name:    "Beans",
		Unicode: "\U0001FAD8",
		Alias:   ":beans:",
	},
	"\U0001FAD9  :jar:": {
		Name:    "Jar",
		Unicode: "\U0001FAD9",
		Alias:   ":jar:",
	},
	"\U0001FAE0  :melting_face:": {
		Name:    "Melting Face",
		Unicode: "\U0001FAE0",
		Alias:   ":melting_face:",
	},
	"\U0001FAE1  :saluting_face:": {
		Name:    "Saluting Face",
		Unicode: "\U0001FAE1",
		Alias:   ":saluting_face:",
	},
	"\U0001FAE2  :face_with_open_eyes_and_hand_over_mouth:": {
		Name:    "Face With Open Eyes And Hand Over Mouth",
		Unicode: "\U0001FAE2",
		Alias:   ":face_with_open_eyes_and_hand_over_mouth:",
	},
	"\U0001FAE3  :face_with_peeking_eye:": {
		Name:    "Face With Peeking Eye",
		Unicode: "\U0001FAE3",
		Alias:   ":face_with_peeking_eye:",
	},
	"\U0001FAE4  :face_with_diagonal_mouth:": {
		Name:    "Face With Diagonal Mouth",
		Unicode: "\U0001FAE4",
		Alias:   ":face_with_diagonal_mouth:",
	},
	"\U0001FAE5  :dotted_line_face:": {
		Name:    "Dotted Line Face",
		Unicode: "\U0001FAE5",
		Alias:   ":dotted_line_face:",
	},
	"\U0001FAE6  :biting_lip:": {
		Name:    "Biting Lip",
		Unicode: "\U0001FAE6",
		Alias:   ":biting_lip:",
	},
	"\U0001FAE7  :bubbles:": {
		Name:    "Bubbles",
		Unicode: "\U0001FAE7",
		Alias:   ":bubbles:",
	},
	"\U0001FAF0  :hand_with_index_finger_and_thumb_crossed:": {
		Name:    "Hand With Index Finger And Thumb Crossed",
		Unicode: "\U0001FAF0",
		Alias:   ":hand_with_index_finger_and_thumb_crossed:",
	},
	"\U0001FAF1  :rightwards_hand:": {
		Name:    "Rightwards Hand",
		Unicode: "\U0001FAF1",
		Alias:   ":rightwards_hand:",
	},
	"\U0001FAF2  :leftwards_hand:": {
		Name:    "Leftwards Hand",
		Unicode: "\U0001FAF2",
		Alias:   ":leftwards_hand:",
	},
	"\U0001FAF3  :palm_down_hand:": {
		Name:    "Palm Down Hand",
		Unicode: "\U0001FAF3",
		Alias:   ":palm_down_hand:",
	},
	"\U0001FAF4  :palm_up_hand:": {
		Name:    "Palm Up Hand",
		Unicode: "\U0001FAF4",
		Alias:   ":palm_up_hand:",
	},
	"\U0001FAF5  :index_pointing_at_the_viewer:": {
		Name:    "Index Pointing At The Viewer",
		Unicode: "\U0001FAF5",
		Alias:   ":index_pointing_at_the_viewer:",
	},
	"\U0001FAF6  :heart_hands:": {
		Name:    "Heart Hands",
		Unicode: "\U0001FAF6",
		Alias:   ":heart_hands:",
	},
	"\u203C\uFE0F  :bangbang:": {
		Name:    "Double Exclamation Mark",
		Unicode: "\u203C\uFE0F",
		Alias:   ":bangbang:",
	},
	"\u2049\uFE0F  :interrobang:": {
		Name:    "Exclamation Question Mark",
		Unicode: "\u2049\uFE0F",
		Alias:   ":interrobang:",
	},
	"\u2122\uFE0F  :tm:": {
		Name:    "Trade Mark Sign",
		Unicode: "\u2122\uFE0F",
		Alias:   ":tm:",
	},
	"\u2139\uFE0F  :information_source:": {
		Name:    "Information Source",
		Unicode: "\u2139\uFE0F",
		Alias:   ":information_source:",
	},
	"\u2194\uFE0F  :left_right_arrow:": {
		Name:    "Left Right Arrow",
		Unicode: "\u2194\uFE0F",
		Alias:   ":left_right_arrow:",
	},
	"\u2195\uFE0F  :arrow_up_down:": {
		Name:    "Up Down Arrow",
		Unicode: "\u2195\uFE0F",
		Alias:   ":arrow_up_down:",
	},
	"\u2196\uFE0F  :arrow_upper_left:": {
		Name:    "North West Arrow",
		Unicode: "\u2196\uFE0F",
		Alias:   ":arrow_upper_left:",
	},
	"\u2197\uFE0F  :arrow_upper_right:": {
		Name:    "North East Arrow",
		Unicode: "\u2197\uFE0F",
		Alias:   ":arrow_upper_right:",
	},
	"\u2198\uFE0F  :arrow_lower_right:": {
		Name:    "South East Arrow",
		Unicode: "\u2198\uFE0F",
		Alias:   ":arrow_lower_right:",
	},
	"\u2199\uFE0F  :arrow_lower_left:": {
		Name:    "South West Arrow",
		Unicode: "\u2199\uFE0F",
		Alias:   ":arrow_lower_left:",
	},
	"\u21A9\uFE0F  :leftwards_arrow_with_hook:": {
		Name:    "Leftwards Arrow With Hook",
		Unicode: "\u21A9\uFE0F",
		Alias:   ":leftwards_arrow_with_hook:",
	},
	"\u21AA\uFE0F  :arrow_right_hook:": {
		Name:    "Rightwards Arrow With Hook",
		Unicode: "\u21AA\uFE0F",
		Alias:   ":arrow_right_hook:",
	},
	"\u231A  :watch:": {
		Name:    "Watch",
		Unicode: "\u231A",
		Alias:   ":watch:",
	},
	"\u231B  :hourglass:": {
		Name:    "Hourglass",
		Unicode: "\u231B",
		Alias:   ":hourglass:",
	},
	"\u2328\uFE0F  :keyboard:": {
		Name:    "Keyboard",
		Unicode: "\u2328\uFE0F",
		Alias:   ":keyboard:",
	},
	"\u23CF\uFE0F  :eject:": {
		Name:    "Eject Button",
		Unicode: "\u23CF\uFE0F",
		Alias:   ":eject:",
	},
	"\u23E9  :fast_forward:": {
		Name:    "Black Right-Pointing Double Triangle",
		Unicode: "\u23E9",
		Alias:   ":fast_forward:",
	},
	"\u23EA  :rewind:": {
		Name:    "Black Left-Pointing Double Triangle",
		Unicode: "\u23EA",
		Alias:   ":rewind:",
	},
	"\u23EB  :arrow_double_up:": {
		Name:    "Black Up-Pointing Double Triangle",
		Unicode: "\u23EB",
		Alias:   ":arrow_double_up:",
	},
	"\u23EC  :arrow_double_down:": {
		Name:    "Black Down-Pointing Double Triangle",
		Unicode: "\u23EC",
		Alias:   ":arrow_double_down:",
	},
	"\u23ED\uFE0F  :black_right_pointing_double_triangle_with_vertical_bar:": {
		Name:    "Next Track Button",
		Unicode: "\u23ED\uFE0F",
		Alias:   ":black_right_pointing_double_triangle_with_vertical_bar:",
	},
	"\u23EE\uFE0F  :black_left_pointing_double_triangle_with_vertical_bar:": {
		Name:    "Last Track Button",
		Unicode: "\u23EE\uFE0F",
		Alias:   ":black_left_pointing_double_triangle_with_vertical_bar:",
	},
	"\u23EF\uFE0F  :black_right_pointing_triangle_with_double_vertical_bar:": {
		Name:    "Play Or Pause Button",
		Unicode: "\u23EF\uFE0F",
		Alias:   ":black_right_pointing_triangle_with_double_vertical_bar:",
	},
	"\u23F0  :alarm_clock:": {
		Name:    "Alarm Clock",
		Unicode: "\u23F0",
		Alias:   ":alarm_clock:",
	},
	"\u23F1\uFE0F  :stopwatch:": {
		Name:    "Stopwatch",
		Unicode: "\u23F1\uFE0F",
		Alias:   ":stopwatch:",
	},
	"\u23F2\uFE0F  :timer_clock:": {
		Name:    "Timer Clock",
		Unicode: "\u23F2\uFE0F",
		Alias:   ":timer_clock:",
	},
	"\u23F3  :hourglass_flowing_sand:": {
		Name:    "Hourglass With Flowing Sand",
		Unicode: "\u23F3",
		Alias:   ":hourglass_flowing_sand:",
	},
	"\u23F8\uFE0F  :double_vertical_bar:": {
		Name:    "Pause Button",
		Unicode: "\u23F8\uFE0F",
		Alias:   ":double_vertical_bar:",
	},
	"\u23F9\uFE0F  :black_square_for_stop:": {
		Name:    "Stop Button",
		Unicode: "\u23F9\uFE0F",
		Alias:   ":black_square_for_stop:",
	},
	"\u23FA\uFE0F  :black_circle_for_record:": {
		Name:    "Record Button",
		Unicode: "\u23FA\uFE0F",
		Alias:   ":black_circle_for_record:",
	},
	"\u24C2\uFE0F  :m:": {
		Name:    "Circled Latin Capital Letter M",
		Unicode: "\u24C2\uFE0F",
		Alias:   ":m:",
	},
	"\u25AA\uFE0F  :black_small_square:": {
		Name:    "Black Small Square",
		Unicode: "\u25AA\uFE0F",
		Alias:   ":black_small_square:",
	},
	"\u25AB\uFE0F  :white_small_square:": {
		Name:    "White Small Square",
		Unicode: "\u25AB\uFE0F",
		Alias:   ":white_small_square:",
	},
	"\u25B6\uFE0F  :arrow_forward:": {
		Name:    "Black Right-Pointing Triangle",
		Unicode: "\u25B6\uFE0F",
		Alias:   ":arrow_forward:",
	},
	"\u25C0\uFE0F  :arrow_backward:": {
		Name:    "Black Left-Pointing Triangle",
		Unicode: "\u25C0\uFE0F",
		Alias:   ":arrow_backward:",
	},
	"\u25FB\uFE0F  :white_medium_square:": {
		Name:    "White Medium Square",
		Unicode: "\u25FB\uFE0F",
		Alias:   ":white_medium_square:",
	},
	"\u25FC\uFE0F  :black_medium_square:": {
		Name:    "Black Medium Square",
		Unicode: "\u25FC\uFE0F",
		Alias:   ":black_medium_square:",
	},
	"\u25FD  :white_medium_small_square:": {
		Name:    "White Medium Small Square",
		Unicode: "\u25FD",
		Alias:   ":white_medium_small_square:",
	},
	"\u25FE  :black_medium_small_square:": {
		Name:    "Black Medium Small Square",
		Unicode: "\u25FE",
		Alias:   ":black_medium_small_square:",
	},
	"\u2600\uFE0F  :sunny:": {
		Name:    "Black Sun With Rays",
		Unicode: "\u2600\uFE0F",
		Alias:   ":sunny:",
	},
	"\u2601\uFE0F  :cloud:": {
		Name:    "Cloud",
		Unicode: "\u2601\uFE0F",
		Alias:   ":cloud:",
	},
	"\u2602\uFE0F  :umbrella:": {
		Name:    "Umbrella",
		Unicode: "\u2602\uFE0F",
		Alias:   ":umbrella:",
	},
	"\u2603\uFE0F  :snowman:": {
		Name:    "Snowman",
		Unicode: "\u2603\uFE0F",
		Alias:   ":snowman:",
	},
	"\u2604\uFE0F  :comet:": {
		Name:    "Comet",
		Unicode: "\u2604\uFE0F",
		Alias:   ":comet:",
	},
	"\u260E\uFE0F  :phone:": {
		Name:    "Black Telephone",
		Unicode: "\u260E\uFE0F",
		Alias:   ":phone:",
	},
	"\u2611\uFE0F  :ballot_box_with_check:": {
		Name:    "Ballot Box With Check",
		Unicode: "\u2611\uFE0F",
		Alias:   ":ballot_box_with_check:",
	},
	"\u2614  :umbrella_with_rain_drops:": {
		Name:    "Umbrella With Rain Drops",
		Unicode: "\u2614",
		Alias:   ":umbrella_with_rain_drops:",
	},
	"\u2615  :coffee:": {
		Name:    "Hot Beverage",
		Unicode: "\u2615",
		Alias:   ":coffee:",
	},
	"\u2618\uFE0F  :shamrock:": {
		Name:    "Shamrock",
		Unicode: "\u2618\uFE0F",
		Alias:   ":shamrock:",
	},
	"\u261D\uFE0F  :point_up:": {
		Name:    "White Up Pointing Index",
		Unicode: "\u261D\uFE0F",
		Alias:   ":point_up:",
	},
	"\u2620\uFE0F  :skull_and_crossbones:": {
		Name:    "Skull And Crossbones",
		Unicode: "\u2620\uFE0F",
		Alias:   ":skull_and_crossbones:",
	},
	"\u2622\uFE0F  :radioactive_sign:": {
		Name:    "Radioactive",
		Unicode: "\u2622\uFE0F",
		Alias:   ":radioactive_sign:",
	},
	"\u2623\uFE0F  :biohazard_sign:": {
		Name:    "Biohazard",
		Unicode: "\u2623\uFE0F",
		Alias:   ":biohazard_sign:",
	},
	"\u2626\uFE0F  :orthodox_cross:": {
		Name:    "Orthodox Cross",
		Unicode: "\u2626\uFE0F",
		Alias:   ":orthodox_cross:",
	},
	"\u262A\uFE0F  :star_and_crescent:": {
		Name:    "Star And Crescent",
		Unicode: "\u262A\uFE0F",
		Alias:   ":star_and_crescent:",
	},
	"\u262E\uFE0F  :peace_symbol:": {
		Name:    "Peace Symbol",
		Unicode: "\u262E\uFE0F",
		Alias:   ":peace_symbol:",
	},
	"\u262F\uFE0F  :yin_yang:": {
		Name:    "Yin Yang",
		Unicode: "\u262F\uFE0F",
		Alias:   ":yin_yang:",
	},
	"\u2638\uFE0F  :wheel_of_dharma:": {
		Name:    "Wheel Of Dharma",
		Unicode: "\u2638\uFE0F",
		Alias:   ":wheel_of_dharma:",
	},
	"\u2639\uFE0F  :white_frowning_face:": {
		Name:    "Frowning Face",
		Unicode: "\u2639\uFE0F",
		Alias:   ":white_frowning_face:",
	},
	"\u263A\uFE0F  :relaxed:": {
		Name:    "White Smiling Face",
		Unicode: "\u263A\uFE0F",
		Alias:   ":relaxed:",
	},
	"\u2640\uFE0F  :female_sign:": {
		Name:    "Female Sign",
		Unicode: "\u2640\uFE0F",
		Alias:   ":female_sign:",
	},
	"\u2642\uFE0F  :male_sign:": {
		Name:    "Male Sign",
		Unicode: "\u2642\uFE0F",
		Alias:   ":male_sign:",
	},
	"\u2648  :aries:": {
		Name:    "Aries",
		Unicode: "\u2648",
		Alias:   ":aries:",
	},
	"\u2649  :taurus:": {
		Name:    "Taurus",
		Unicode: "\u2649",
		Alias:   ":taurus:",
	},
	"\u264A  :gemini:": {
		Name:    "Gemini",
		Unicode: "\u264A",
		Alias:   ":gemini:",
	},
	"\u264B  :cancer:": {
		Name:    "Cancer",
		Unicode: "\u264B",
		Alias:   ":cancer:",
	},
	"\u264C  :leo:": {
		Name:    "Leo",
		Unicode: "\u264C",
		Alias:   ":leo:",
	},
	"\u264D  :virgo:": {
		Name:    "Virgo",
		Unicode: "\u264D",
		Alias:   ":virgo:",
	},
	"\u264E  :libra:": {
		Name:    "Libra",
		Unicode: "\u264E",
		Alias:   ":libra:",
	},
	"\u264F  :scorpius:": {
		Name:    "Scorpius",
		Unicode: "\u264F",
		Alias:   ":scorpius:",
	},
	"\u2650  :sagittarius:": {
		Name:    "Sagittarius",
		Unicode: "\u2650",
		Alias:   ":sagittarius:",
	},
	"\u2651  :capricorn:": {
		Name:    "Capricorn",
		Unicode: "\u2651",
		Alias:   ":capricorn:",
	},
	"\u2652  :aquarius:": {
		Name:    "Aquarius",
		Unicode: "\u2652",
		Alias:   ":aquarius:",
	},
	"\u2653  :pisces:": {
		Name:    "Pisces",
		Unicode: "\u2653",
		Alias:   ":pisces:",
	},
	"\u265F\uFE0F  :chess_pawn:": {
		Name:    "Chess Pawn",
		Unicode: "\u265F\uFE0F",
		Alias:   ":chess_pawn:",
	},
	"\u2660\uFE0F  :spades:": {
		Name:    "Black Spade Suit",
		Unicode: "\u2660\uFE0F",
		Alias:   ":spades:",
	},
	"\u2663\uFE0F  :clubs:": {
		Name:    "Black Club Suit",
		Unicode: "\u2663\uFE0F",
		Alias:   ":clubs:",
	},
	"\u2665\uFE0F  :hearts:": {
		Name:    "Black Heart Suit",
		Unicode: "\u2665\uFE0F",
		Alias:   ":hearts:",
	},
	"\u2666\uFE0F  :diamonds:": {
		Name:    "Black Diamond Suit",
		Unicode: "\u2666\uFE0F",
		Alias:   ":diamonds:",
	},
	"\u2668\uFE0F  :hotsprings:": {
		Name:    "Hot Springs",
		Unicode: "\u2668\uFE0F",
		Alias:   ":hotsprings:",
	},
	"\u267B\uFE0F  :recycle:": {
		Name:    "Black Universal Recycling Symbol",
		Unicode: "\u267B\uFE0F",
		Alias:   ":recycle:",
	},
	"\u267E\uFE0F  :infinity:": {
		Name:    "Infinity",
		Unicode: "\u267E\uFE0F",
		Alias:   ":infinity:",
	},
	"\u267F  :wheelchair:": {
		Name:    "Wheelchair Symbol",
		Unicode: "\u267F",
		Alias:   ":wheelchair:",
	},
	"\u2692\uFE0F  :hammer_and_pick:": {
		Name:    "Hammer And Pick",
		Unicode: "\u2692\uFE0F",
		Alias:   ":hammer_and_pick:",
	},
	"\u2693  :anchor:": {
		Name:    "Anchor",
		Unicode: "\u2693",
		Alias:   ":anchor:",
	},
	"\u2694\uFE0F  :crossed_swords:": {
		Name:    "Crossed Swords",
		Unicode: "\u2694\uFE0F",
		Alias:   ":crossed_swords:",
	},
	"\u2695\uFE0F  :medical_symbol:": {
		Name:    "Medical Symbol",
		Unicode: "\u2695\uFE0F",
		Alias:   ":medical_symbol:",
	},
	"\u2696\uFE0F  :scales:": {
		Name:    "Balance Scale",
		Unicode: "\u2696\uFE0F",
		Alias:   ":scales:",
	},
	"\u2697\uFE0F  :alembic:": {
		Name:    "Alembic",
		Unicode: "\u2697\uFE0F",
		Alias:   ":alembic:",
	},
	"\u2699\uFE0F  :gear:": {
		Name:    "Gear",
		Unicode: "\u2699\uFE0F",
		Alias:   ":gear:",
	},
	"\u269B\uFE0F  :atom_symbol:": {
		Name:    "Atom Symbol",
		Unicode: "\u269B\uFE0F",
		Alias:   ":atom_symbol:",
	},
	"\u269C\uFE0F  :fleur_de_lis:": {
		Name:    "Fleur-De-Lis",
		Unicode: "\u269C\uFE0F",
		Alias:   ":fleur_de_lis:",
	},
	"\u26A0\uFE0F  :warning:": {
		Name:    "Warning Sign",
		Unicode: "\u26A0\uFE0F",
		Alias:   ":warning:",
	},
	"\u26A1  :zap:": {
		Name:    "High Voltage Sign",
		Unicode: "\u26A1",
		Alias:   ":zap:",
	},
	"\u26A7\uFE0F  :transgender_symbol:": {
		Name:    "Transgender Symbol",
		Unicode: "\u26A7\uFE0F",
		Alias:   ":transgender_symbol:",
	},
	"\u26AA  :white_circle:": {
		Name:    "Medium White Circle",
		Unicode: "\u26AA",
		Alias:   ":white_circle:",
	},
	"\u26AB  :black_circle:": {
		Name:    "Medium Black Circle",
		Unicode: "\u26AB",
		Alias:   ":black_circle:",
	},
	"\u26B0\uFE0F  :coffin:": {
		Name:    "Coffin",
		Unicode: "\u26B0\uFE0F",
		Alias:   ":coffin:",
	},
	"\u26B1\uFE0F  :funeral_urn:": {
		Name:    "Funeral Urn",
		Unicode: "\u26B1\uFE0F",
		Alias:   ":funeral_urn:",
	},
	"\u26BD  :soccer:": {
		Name:    "Soccer Ball",
		Unicode: "\u26BD",
		Alias:   ":soccer:",
	},
	"\u26BE  :baseball:": {
		Name:    "Baseball",
		Unicode: "\u26BE",
		Alias:   ":baseball:",
	},
	"\u26C4  :snowman_without_snow:": {
		Name:    "Snowman Without Snow",
		Unicode: "\u26C4",
		Alias:   ":snowman_without_snow:",
	},
	"\u26C5  :partly_sunny:": {
		Name:    "Sun Behind Cloud",
		Unicode: "\u26C5",
		Alias:   ":partly_sunny:",
	},
	"\u26C8\uFE0F  :thunder_cloud_and_rain:": {
		Name:    "Cloud With Lightning And Rain",
		Unicode: "\u26C8\uFE0F",
		Alias:   ":thunder_cloud_and_rain:",
	},
	"\u26CE  :ophiuchus:": {
		Name:    "Ophiuchus",
		Unicode: "\u26CE",
		Alias:   ":ophiuchus:",
	},
	"\u26CF\uFE0F  :pick:": {
		Name:    "Pick",
		Unicode: "\u26CF\uFE0F",
		Alias:   ":pick:",
	},
	"\u26D1\uFE0F  :helmet_with_white_cross:": {
		Name:    "Rescue Worker\u2019S Helmet",
		Unicode: "\u26D1\uFE0F",
		Alias:   ":helmet_with_white_cross:",
	},
	"\u26D3\uFE0F  :chains:": {
		Name:    "Chains",
		Unicode: "\u26D3\uFE0F",
		Alias:   ":chains:",
	},
	"\u26D4  :no_entry:": {
		Name:    "No Entry",
		Unicode: "\u26D4",
		Alias:   ":no_entry:",
	},
	"\u26E9\uFE0F  :shinto_shrine:": {
		Name:    "Shinto Shrine",
		Unicode: "\u26E9\uFE0F",
		Alias:   ":shinto_shrine:",
	},
	"\u26EA  :church:": {
		Name:    "Church",
		Unicode: "\u26EA",
		Alias:   ":church:",
	},
	"\u26F0\uFE0F  :mountain:": {
		Name:    "Mountain",
		Unicode: "\u26F0\uFE0F",
		Alias:   ":mountain:",
	},
	"\u26F1\uFE0F  :umbrella_on_ground:": {
		Name:    "Umbrella On Ground",
		Unicode: "\u26F1\uFE0F",
		Alias:   ":umbrella_on_ground:",
	},
	"\u26F2  :fountain:": {
		Name:    "Fountain",
		Unicode: "\u26F2",
		Alias:   ":fountain:",
	},
	"\u26F3  :golf:": {
		Name:    "Flag In Hole",
		Unicode: "\u26F3",
		Alias:   ":golf:",
	},
	"\u26F4\uFE0F  :ferry:": {
		Name:    "Ferry",
		Unicode: "\u26F4\uFE0F",
		Alias:   ":ferry:",
	},
	"\u26F5  :boat:": {
		Name:    "Sailboat",
		Unicode: "\u26F5",
		Alias:   ":boat:",
	},
	"\u26F7\uFE0F  :skier:": {
		Name:    "Skier",
		Unicode: "\u26F7\uFE0F",
		Alias:   ":skier:",
	},
	"\u26F8\uFE0F  :ice_skate:": {
		Name:    "Ice Skate",
		Unicode: "\u26F8\uFE0F",
		Alias:   ":ice_skate:",
	},
	"\u26F9\uFE0F\u200D\u2640\uFE0F :woman-bouncing-ball:": {
		Name:    "Woman Bouncing Ball",
		Unicode: "\u26F9\uFE0F\u200D\u2640\uFE0F",
		Alias:   ":woman-bouncing-ball:",
	},
	"\u26F9\uFE0F\u200D\u2642\uFE0F :man-bouncing-ball:": {
		Name:    "Man Bouncing Ball",
		Unicode: "\u26F9\uFE0F\u200D\u2642\uFE0F",
		Alias:   ":man-bouncing-ball:",
	},
	"\u26F9\uFE0F  :person_with_ball:": {
		Name:    "Person Bouncing Ball",
		Unicode: "\u26F9\uFE0F",
		Alias:   ":person_with_ball:",
	},
	"\u26FA  :tent:": {
		Name:    "Tent",
		Unicode: "\u26FA",
		Alias:   ":tent:",
	},
	"\u26FD  :fuelpump:": {
		Name:    "Fuel Pump",
		Unicode: "\u26FD",
		Alias:   ":fuelpump:",
	},
	"\u2702\uFE0F  :scissors:": {
		Name:    "Black Scissors",
		Unicode: "\u2702\uFE0F",
		Alias:   ":scissors:",
	},
	"\u2705  :white_check_mark:": {
		Name:    "White Heavy Check Mark",
		Unicode: "\u2705",
		Alias:   ":white_check_mark:",
	},
	"\u2708\uFE0F  :airplane:": {
		Name:    "Airplane",
		Unicode: "\u2708\uFE0F",
		Alias:   ":airplane:",
	},
	"\u2709\uFE0F  :email:": {
		Name:    "Envelope",
		Unicode: "\u2709\uFE0F",
		Alias:   ":email:",
	},
	"\u270A  :fist:": {
		Name:    "Raised Fist",
		Unicode: "\u270A",
		Alias:   ":fist:",
	},
	"\u270B  :hand:": {
		Name:    "Raised Hand",
		Unicode: "\u270B",
		Alias:   ":hand:",
	},
	"\u270C\uFE0F  :v:": {
		Name:    "Victory Hand",
		Unicode: "\u270C\uFE0F",
		Alias:   ":v:",
	},
	"\u270D\uFE0F  :writing_hand:": {
		Name:    "Writing Hand",
		Unicode: "\u270D\uFE0F",
		Alias:   ":writing_hand:",
	},
	"\u270F\uFE0F  :pencil2:": {
		Name:    "Pencil",
		Unicode: "\u270F\uFE0F",
		Alias:   ":pencil2:",
	},
	"\u2712\uFE0F  :black_nib:": {
		Name:    "Black Nib",
		Unicode: "\u2712\uFE0F",
		Alias:   ":black_nib:",
	},
	"\u2714\uFE0F  :heavy_check_mark:": {
		Name:    "Heavy Check Mark",
		Unicode: "\u2714\uFE0F",
		Alias:   ":heavy_check_mark:",
	},
	"\u2716\uFE0F  :heavy_multiplication_x:": {
		Name:    "Heavy Multiplication X",
		Unicode: "\u2716\uFE0F",
		Alias:   ":heavy_multiplication_x:",
	},
	"\u271D\uFE0F  :latin_cross:": {
		Name:    "Latin Cross",
		Unicode: "\u271D\uFE0F",
		Alias:   ":latin_cross:",
	},
	"\u2721\uFE0F  :star_of_david:": {
		Name:    "Star Of David",
		Unicode: "\u2721\uFE0F",
		Alias:   ":star_of_david:",
	},
	"\u2728  :sparkles:": {
		Name:    "Sparkles",
		Unicode: "\u2728",
		Alias:   ":sparkles:",
	},
	"\u2733\uFE0F  :eight_spoked_asterisk:": {
		Name:    "Eight Spoked Asterisk",
		Unicode: "\u2733\uFE0F",
		Alias:   ":eight_spoked_asterisk:",
	},
	"\u2734\uFE0F  :eight_pointed_black_star:": {
		Name:    "Eight Pointed Black Star",
		Unicode: "\u2734\uFE0F",
		Alias:   ":eight_pointed_black_star:",
	},
	"\u2744\uFE0F  :snowflake:": {
		Name:    "Snowflake",
		Unicode: "\u2744\uFE0F",
		Alias:   ":snowflake:",
	},
	"\u2747\uFE0F  :sparkle:": {
		Name:    "Sparkle",
		Unicode: "\u2747\uFE0F",
		Alias:   ":sparkle:",
	},
	"\u274C  :x:": {
		Name:    "Cross Mark",
		Unicode: "\u274C",
		Alias:   ":x:",
	},
	"\u274E  :negative_squared_cross_mark:": {
		Name:    "Negative Squared Cross Mark",
		Unicode: "\u274E",
		Alias:   ":negative_squared_cross_mark:",
	},
	"\u2753  :question:": {
		Name:    "Black Question Mark Ornament",
		Unicode: "\u2753",
		Alias:   ":question:",
	},
	"\u2754  :grey_question:": {
		Name:    "White Question Mark Ornament",
		Unicode: "\u2754",
		Alias:   ":grey_question:",
	},
	"\u2755  :grey_exclamation:": {
		Name:    "White Exclamation Mark Ornament",
		Unicode: "\u2755",
		Alias:   ":grey_exclamation:",
	},
	"\u2757  :exclamation:": {
		Name:    "Heavy Exclamation Mark Symbol",
		Unicode: "\u2757",
		Alias:   ":exclamation:",
	},
	"\u2763\uFE0F  :heavy_heart_exclamation_mark_ornament:": {
		Name:    "Heart Exclamation",
		Unicode: "\u2763\uFE0F",
		Alias:   ":heavy_heart_exclamation_mark_ornament:",
	},
	"\u2764\uFE0F\u200D\u1F525  :heart_on_fire:": {
		Name:    "Heart On Fire",
		Unicode: "\u2764\uFE0F\u200D\u1F525",
		Alias:   ":heart_on_fire:",
	},
	"\u2764\uFE0F\u200D\u1FA79  :mending_heart:": {
		Name:    "Mending Heart",
		Unicode: "\u2764\uFE0F\u200D\u1FA79",
		Alias:   ":mending_heart:",
	},
	"\u2764\uFE0F  :heart:": {
		Name:    "Heavy Black Heart",
		Unicode: "\u2764\uFE0F",
		Alias:   ":heart:",
	},
	"\u2795  :heavy_plus_sign:": {
		Name:    "Heavy Plus Sign",
		Unicode: "\u2795",
		Alias:   ":heavy_plus_sign:",
	},
	"\u2796  :heavy_minus_sign:": {
		Name:    "Heavy Minus Sign",
		Unicode: "\u2796",
		Alias:   ":heavy_minus_sign:",
	},
	"\u2797  :heavy_division_sign:": {
		Name:    "Heavy Division Sign",
		Unicode: "\u2797",
		Alias:   ":heavy_division_sign:",
	},
	"\u27A1\uFE0F  :arrow_right:": {
		Name:    "Black Rightwards Arrow",
		Unicode: "\u27A1\uFE0F",
		Alias:   ":arrow_right:",
	},
	"\u27B0  :curly_loop:": {
		Name:    "Curly Loop",
		Unicode: "\u27B0",
		Alias:   ":curly_loop:",
	},
	"\u27BF  :loop:": {
		Name:    "Double Curly Loop",
		Unicode: "\u27BF",
		Alias:   ":loop:",
	},
	"\u2934\uFE0F  :arrow_heading_up:": {
		Name:    "Arrow Pointing Rightwards Then Curving Upwards",
		Unicode: "\u2934\uFE0F",
		Alias:   ":arrow_heading_up:",
	},
	"\u2935\uFE0F  :arrow_heading_down:": {
		Name:    "Arrow Pointing Rightwards Then Curving Downwards",
		Unicode: "\u2935\uFE0F",
		Alias:   ":arrow_heading_down:",
	},
	"\u2B05\uFE0F  :arrow_left:": {
		Name:    "Leftwards Black Arrow",
		Unicode: "\u2B05\uFE0F",
		Alias:   ":arrow_left:",
	},
	"\u2B06\uFE0F  :arrow_up:": {
		Name:    "Upwards Black Arrow",
		Unicode: "\u2B06\uFE0F",
		Alias:   ":arrow_up:",
	},
	"\u2B07\uFE0F  :arrow_down:": {
		Name:    "Downwards Black Arrow",
		Unicode: "\u2B07\uFE0F",
		Alias:   ":arrow_down:",
	},
	"\u2B1B  :black_large_square:": {
		Name:    "Black Large Square",
		Unicode: "\u2B1B",
		Alias:   ":black_large_square:",
	},
	"\u2B1C  :white_large_square:": {
		Name:    "White Large Square",
		Unicode: "\u2B1C",
		Alias:   ":white_large_square:",
	},
	"\u2B50  :star:": {
		Name:    "White Medium Star",
		Unicode: "\u2B50",
		Alias:   ":star:",
	},
	"\u2B55  :o:": {
		Name:    "Heavy Large Circle",
		Unicode: "\u2B55",
		Alias:   ":o:",
	},
	"\u3030\uFE0F  :wavy_dash:": {
		Name:    "Wavy Dash",
		Unicode: "\u3030\uFE0F",
		Alias:   ":wavy_dash:",
	},
	"\u303D\uFE0F  :part_alternation_mark:": {
		Name:    "Part Alternation Mark",
		Unicode: "\u303D\uFE0F",
		Alias:   ":part_alternation_mark:",
	},
	"\u3297\uFE0F  :congratulations:": {
		Name:    "Circled Ideograph Congratulation",
		Unicode: "\u3297\uFE0F",
		Alias:   ":congratulations:",
	},
	"\u3299\uFE0F  :secret:": {
		Name:    "Circled Ideograph Secret",
		Unicode: "\u3299\uFE0F",
		Alias:   ":secret:",
	},
}
