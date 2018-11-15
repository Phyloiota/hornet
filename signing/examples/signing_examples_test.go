package examples

import (
	"fmt"
	"github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/signing"
	"github.com/iotaledger/iota.go/trinary"
)

// i req: seed, The seed from which to derive the subseed from.
// i req: index, The index of the subseed.
// o: Trits, The Trits representation of the subseed.
// o: error, Returned for invalid seeds and internal errors.
func ExampleSubseed() {
	seed := "ZLNM9UHJWKTTDEZOTH9CXDEIFUJQCIACDPJIXPOWBDW9LTBHC9AQRIXTIHYLIIURLZCXNSTGNIVC9ISVB"
	subseed, err := signing.Subseed(seed, 0)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(subseed))
	// output: CEFLDDLMF9TO9ZLLTYXIPVFIJKAOFRIQLGNYIDZCTDYSWMNXPYNGFAKHQDY9ABGGQZHEFTXKWKWZXEIUD
}

// i req: subseed, The subseed from which to derive the private key from.
// i req: securityLevel, The used security level.
// o: Trits, The Trits representation of the private key.
// o: error, Returned for internal errors.
func ExampleKey() {
	subseed := "CEFLDDLMF9TO9ZLLTYXIPVFIJKAOFRIQLGNYIDZCTDYSWMNXPYNGFAKHQDY9ABGGQZHEFTXKWKWZXEIUD"
	subseedTrits := trinary.MustTrytesToTrits(subseed)
	key, err := signing.Key(subseedTrits, consts.SecurityLevelMedium)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(key))
	// output: D9DWVXXXMGBR9BKHQMMRTQIQROKTLOZNNYHHETDLHVE9FUIBGLEVSTHMJHNHXRRYWHBLNUICBOQHVGBRDSAJZKOL9EQNXGHCETSGW9LTGJPMKD9DEFITRUDVUUSKR9BMMAABVJIKICTLEGBSDMHGGOEALGWBZCF9YCWUANXAZOZJFQARYCKNOJTCXBYEWENKABGSUQXLOM9ZNEDVZQMUPWROQVQXTHLDUKMRULEZDOW9MEXSELWIKHYTCXQNWYIRLNQWUEOBPBQZTHNMRSMNIRQ9BKDHPDGN9VRZEBXKTFCCXOETJKX9YYJDB9HGLCOHMPOXRXBBV9YDSLYPBVABTNOWKXLSTGZAA9EXHLLGWIISYOSQQKYYOXSJDXVSLHDASEXDYJTFRRHLRMKLHZESDGTDGLUNLOWIETHXWEWPDCJRGMKFMXVNPLBGRYMAXTTAILUGAFUIXGTDP9JIVZFMEJKZRIYPLQNPSPQGZWZJZQXUVTJGRMCCJLNOLFZUCSEPNNNSYKLHJZH9YXTTTEPMDBWGXYWAMAMIDXYYSUBJKFAIH9XXBGBXMSBPJZGIDSELVW9PCTDYNKRCCZQIUIRIBNVWY9EFTZPJFLOF9GRJJQEPNOXHJHZE9KJXFX9TZZHNIVAQUSHAWTDQVGNMRMNRCKRZWCKWBBWKERTTBYKINHRYVTLVHSEDJTXVQUPBYJEQ9FFEEZOPG9HJCLPLOI9RKDDTCS9QWOZMYQJBKHRLHLCMUJODPLZFVFBINPLBLDI9EHNECSMTRWIIPVNRAWARRJSXMXCDSIENRC9MFIY9CYVLUOIWCONKRBVDRENXLXQKOGGLXWPSMQJUKUDRHBWVDQEWWSJKREESCIBIQNIWWPPWNLCBD9WYXOXCAHWLCIHTFSSNMAFPUUKZQEKWZKOYVIYSCYTOKNOSEPWPTVQBUUDUVDECYJWJQRYOPIKDLTHEAAJYQWFNMCGXRHBJTIPCPVGBVYOQTAWPHQBZEEYVNBNMHVRJDFBT9OFYRPUEIHBQIYB9TUZUMLLZAYVZKEBTQGKJPXPACTOKIG9DMUABKEFROB9FRM99JSPXZQDXTTGPPKUYMLLA9LEUFZYMWHDXTVOAUTJLC9SMEBTJMZPWNKGZTZZOYLOC9JGOQUOYYOSUSS9GNYNLROXMPFYIXPA9UOKATDFVJIGDGLMNGBCORUQMTFAQI9PDUUDRPBCRUJW9ZDHXDVGIIZLAKZQBQXZJZFZZDKKMTBMPYFEKJB9RHRZYJJMQTFLIYRHGUWEJWDD9K9INENNZPET9AJPCBPIBOZPHQTVDSCUGHFWZIDAOHXDCOUTANGUNBUHWGHXODGZCMANRSNUKFXXJLUGOZVZMFZTVSUBCGVMJAEQJFKC9HYMVNOULDFUCCXYBAGGXXDVE9ZETYAGJFXNE9DKVKDLGLYXHTFJDWKJZXUNUUVHKYRHANXJVFWVNFYWLEPMFFMMVIYRDUWSKOHDH9VFLMAKVPOVBH9GOS9TOUXDTMZPMTSCLPRUEPSNWAMVIONKWMUJYQYDZKSNCURUCARCGQUPTTVDBHQFPXMDYODFFAFTQVCJGXPZUZEZVCVVONHH9ZHLVYKUORZBGVBODHWMDMZCCQYZKQCZHHMQNPEVKTFAXFGNKCKJXQWSYKIFPVQFCPBTSCNQVHDTXBWWBGQZ9PAQCQFNSKQCAHIHOGYNVBENKJNWCBRFAU9H9ZDORYB9HDVIZOCUBP9PLCMXPEIOIOFPYWWLYWHVUJVMQCOGIMU99X9QRNURAAMROSXVYRJBK9LQCGREWFNLSPOXRGAINPAMOMZJMEAELKHJRQFLKYTEDEFGK9CNQUXPXHO9LUMZITY9MSESJJN9ZSZLNMIKLFABBKFMLGNRT9YLHUHBYMPPFXMWDPTM9T9GSBOYXSAGCRTPB9WVBOBA9QGTRSSJJUAKPZAOXNGSENLZEWOLRRFYVEGWOQJBGMXQRMDOQMLVIJKFUW9KCQHNZ9G9OBMCJSQIFPXDNQTOLFBHDHDQVBAZPBDW9SKV9TUGHW9JLPBCGLOWCFFYUCCDJARXMMZVXPIJKWTNKTXXRPFGCJUQVOASZYHDPAXFWTEQDCCZRQZCNHUXBEMLAENNPZHTID9PUCM9AQMJSJBXVH9TWUEUVRWZRKHAMPBKYYMRPXDNUWZHBCO9NQJ99MZAIBTXOFU9RDCKKIDOEACSPMIYPUFXXGBYCIVV9AHTLPPNGWJHOQBPNFKYKGCTFFEQLRPYEPMIOXLBYTKZ9YZPWXNBZVUBDAWHBDPIFCTGICEHQNTMT9AAUTI9KXAVKEZEWGFZSPVXHOVKNDXKINGSMSWGMUXHGTOAKHASIAIMEDPTPFYVE9BBSIB9OL9YBCEDHEVPCPJHTUCXFCFASRFZIDKZRSQJVEHSTKHVQXFYJFCOJDPDSQFAPEWVURE9FBSORQEXYBANRQGAWLNVPBXVEBHNULWKJJKPCXCP9QXNWVATDNRBHVWXRIQCNQMDJHMTIIHSHEYXJQVWTDASLSSISNTYMTGJUVHRPVXOZFOMYEIPK9DSLZNGEPPZDCCFGJVBJEHVSOPZANIJYYFZMSAAHNBEXJ9ONTQWMHAFKYTGUNLR9GPHWODRSIDYSJGCEOFKETUKMOBFWBXYXUVD9OAQPDDGAGXICQAVUGXGKPPZYZFNGLIYNEKTUCCWXZQLWZXDKIMQCBMCZIPN9JW9AXFNKMJMEXEJPRDJXVGLJSYUZXUPWK9E9UYBXIZXCGBRCLLWGI9RHVVUTSLYMIREQEPSMIKSG9FTOJWVECSXYERNKMSV9JZXKSTRFVYRENPUOALHGYGIRNMYBEHCZLVMPDEWEFACCJUHRJFPIX9GJZWNZLOLRIL9QJUVEENRRQEBF9OWMXEBU9APHLCTYOKBWFHIWARHPWSZWRWDPGFOGHEBHXHGIODOAG9FDMXCUNVPBPETXQSWHKLEKRHIPDZ9OYENTVRQTMNKGGLVDDIPMLFPFAKQGXQJYHRUXQNWH9STRJE9DEEPMRVBUWAZWWJUNUGBKX9ULPLYAOLKTPNLIAVDWWHK9NPFFL9QCZMXGVSWJWWC9JIJRZWJSDXSJEUYESNLINGBFBTEYMURQP9RRBD9XSNILEFSDXK9ZPRPLYJHYCWAKVOSVEOUFHBSVQOWBEKNVWKMICTUXNPROPPWGYRAPGLHFCMVWCLWEVKCASBGEGYOMJJLMDXDZPNADTSFPOFKRQKQSEULQTTJARXHKTDMNZEIVSZPIKFGTOMDFQEQEBOVGKGNC9ZMVF9AEAHEXZ9CXUAWTSGNPPUJQRREXAFJCLISFHQJUWFDNETQUCMXTEWNJUSAKRNXHA9RFYUDFYSNOFTLLYJH9YTOPWPITUBTFFGWAFXQ9CYORFC9XSUQRWEJXPWWF9IGDIYGTPLVOCFLLCHOZADUBFXQSHMXCURSCVHIJSVBBDYCYCLEEFRLIKMYHZJABDKFMWXOAVLYHKTYTWEXFUPKCPDCMTZYQRVFLGJEYFUXRDYNAJE9ZPA9XPDZCTOWZBANDDTSLI9OWYOYSGZGXWHNBQDMDGPZCMBERFPYQSBHQFFXQCK9YPBPBIOATPYJZPLJTDVYFKZFJCGSNLXJKSOYTHLJIAGLPBNLQVPBPTDSDVFNXXQRQQLNCBPAHOKWKRZFOCA9MAARTPFNYNOESMDTML99HHOIODAHDOYCULH9YNFUVWCLYECPHGVBXYVFMLYWXJU9ZTQXWBEPDWAOTMPADNUESWPQRXLJVLAKKMBRYQYLSLLBPFMZJPTKABQHMYBQAGQIYHIM9XSOEAFAYBHJLCVZIDFMNIQYXTFQCKXWWDKKKFKJWNT9DTTWLPNSOXBRVPPJMQUP9NEYJRNUWUODWZWTYHVTPSKLBPVRCERSMHFVVWJEUSHNXSYXFIBKHYRBSPAJPJ9V9UATWMVRXSCP9UVNM9ANXIXEKSXMXOXPOBHFLKIMHCEVJGAFWYFSMVNLYUSVBUHNGQNZMXIXBPUPDWAHDSIDNTTMHGEG9JFVDXGQVZMOWYRXUJWRIGTRAEVZGQDAXGILD9IQQUWAM9LHNHHRYXIYOZBDNEVORRDDOKICQCDMTDJUPIWZLEYCOOFSYOJLVKXLSBOTGAHXMLQYQEBEP9PZ9KKUOYIGECFTXXNBJIOJXMHCPJ9HOTLW9PSOHFAL9UTNUXVKF9IWMNJ9FITOC9FIYWQFSXPWAKYMIUHBK99DKLCTZUANNGVEEWEVVKMNQPQVVLG9ICNTACUTUPZYMYZKZTVEEVUGFWJGGWITRLJA9NHFYDVHGQLOBOUJHHCAHFRCFZIEXNZRLZFBVQUXOIKYKQR9JAQAZLAEYTBSREGMVEILTOWFXPHMNGGCTGLECZB
}

// i req: key, The private key from which to derive the digests from.
// o: Trits, The Trits representation of the digests.
// o: error, Returned for internal errors.
func ExampleDigests() {
	key := "D9DWVXXXMGBR9BKHQMMRTQIQROKTLOZNNYHHETDLHVE9FUIBGLEVSTHMJHNHXRRYWHBLNUICBOQHVGBRDSAJZKOL9EQNXGHCETSGW9LTGJPMKD9DEFITRUDVUUSKR9BMMAABVJIKICTLEGBSDMHGGOEALGWBZCF9YCWUANXAZOZJFQARYCKNOJTCXBYEWENKABGSUQXLOM9ZNEDVZQMUPWROQVQXTHLDUKMRULEZDOW9MEXSELWIKHYTCXQNWYIRLNQWUEOBPBQZTHNMRSMNIRQ9BKDHPDGN9VRZEBXKTFCCXOETJKX9YYJDB9HGLCOHMPOXRXBBV9YDSLYPBVABTNOWKXLSTGZAA9EXHLLGWIISYOSQQKYYOXSJDXVSLHDASEXDYJTFRRHLRMKLHZESDGTDGLUNLOWIETHXWEWPDCJRGMKFMXVNPLBGRYMAXTTAILUGAFUIXGTDP9JIVZFMEJKZRIYPLQNPSPQGZWZJZQXUVTJGRMCCJLNOLFZUCSEPNNNSYKLHJZH9YXTTTEPMDBWGXYWAMAMIDXYYSUBJKFAIH9XXBGBXMSBPJZGIDSELVW9PCTDYNKRCCZQIUIRIBNVWY9EFTZPJFLOF9GRJJQEPNOXHJHZE9KJXFX9TZZHNIVAQUSHAWTDQVGNMRMNRCKRZWCKWBBWKERTTBYKINHRYVTLVHSEDJTXVQUPBYJEQ9FFEEZOPG9HJCLPLOI9RKDDTCS9QWOZMYQJBKHRLHLCMUJODPLZFVFBINPLBLDI9EHNECSMTRWIIPVNRAWARRJSXMXCDSIENRC9MFIY9CYVLUOIWCONKRBVDRENXLXQKOGGLXWPSMQJUKUDRHBWVDQEWWSJKREESCIBIQNIWWPPWNLCBD9WYXOXCAHWLCIHTFSSNMAFPUUKZQEKWZKOYVIYSCYTOKNOSEPWPTVQBUUDUVDECYJWJQRYOPIKDLTHEAAJYQWFNMCGXRHBJTIPCPVGBVYOQTAWPHQBZEEYVNBNMHVRJDFBT9OFYRPUEIHBQIYB9TUZUMLLZAYVZKEBTQGKJPXPACTOKIG9DMUABKEFROB9FRM99JSPXZQDXTTGPPKUYMLLA9LEUFZYMWHDXTVOAUTJLC9SMEBTJMZPWNKGZTZZOYLOC9JGOQUOYYOSUSS9GNYNLROXMPFYIXPA9UOKATDFVJIGDGLMNGBCORUQMTFAQI9PDUUDRPBCRUJW9ZDHXDVGIIZLAKZQBQXZJZFZZDKKMTBMPYFEKJB9RHRZYJJMQTFLIYRHGUWEJWDD9K9INENNZPET9AJPCBPIBOZPHQTVDSCUGHFWZIDAOHXDCOUTANGUNBUHWGHXODGZCMANRSNUKFXXJLUGOZVZMFZTVSUBCGVMJAEQJFKC9HYMVNOULDFUCCXYBAGGXXDVE9ZETYAGJFXNE9DKVKDLGLYXHTFJDWKJZXUNUUVHKYRHANXJVFWVNFYWLEPMFFMMVIYRDUWSKOHDH9VFLMAKVPOVBH9GOS9TOUXDTMZPMTSCLPRUEPSNWAMVIONKWMUJYQYDZKSNCURUCARCGQUPTTVDBHQFPXMDYODFFAFTQVCJGXPZUZEZVCVVONHH9ZHLVYKUORZBGVBODHWMDMZCCQYZKQCZHHMQNPEVKTFAXFGNKCKJXQWSYKIFPVQFCPBTSCNQVHDTXBWWBGQZ9PAQCQFNSKQCAHIHOGYNVBENKJNWCBRFAU9H9ZDORYB9HDVIZOCUBP9PLCMXPEIOIOFPYWWLYWHVUJVMQCOGIMU99X9QRNURAAMROSXVYRJBK9LQCGREWFNLSPOXRGAINPAMOMZJMEAELKHJRQFLKYTEDEFGK9CNQUXPXHO9LUMZITY9MSESJJN9ZSZLNMIKLFABBKFMLGNRT9YLHUHBYMPPFXMWDPTM9T9GSBOYXSAGCRTPB9WVBOBA9QGTRSSJJUAKPZAOXNGSENLZEWOLRRFYVEGWOQJBGMXQRMDOQMLVIJKFUW9KCQHNZ9G9OBMCJSQIFPXDNQTOLFBHDHDQVBAZPBDW9SKV9TUGHW9JLPBCGLOWCFFYUCCDJARXMMZVXPIJKWTNKTXXRPFGCJUQVOASZYHDPAXFWTEQDCCZRQZCNHUXBEMLAENNPZHTID9PUCM9AQMJSJBXVH9TWUEUVRWZRKHAMPBKYYMRPXDNUWZHBCO9NQJ99MZAIBTXOFU9RDCKKIDOEACSPMIYPUFXXGBYCIVV9AHTLPPNGWJHOQBPNFKYKGCTFFEQLRPYEPMIOXLBYTKZ9YZPWXNBZVUBDAWHBDPIFCTGICEHQNTMT9AAUTI9KXAVKEZEWGFZSPVXHOVKNDXKINGSMSWGMUXHGTOAKHASIAIMEDPTPFYVE9BBSIB9OL9YBCEDHEVPCPJHTUCXFCFASRFZIDKZRSQJVEHSTKHVQXFYJFCOJDPDSQFAPEWVURE9FBSORQEXYBANRQGAWLNVPBXVEBHNULWKJJKPCXCP9QXNWVATDNRBHVWXRIQCNQMDJHMTIIHSHEYXJQVWTDASLSSISNTYMTGJUVHRPVXOZFOMYEIPK9DSLZNGEPPZDCCFGJVBJEHVSOPZANIJYYFZMSAAHNBEXJ9ONTQWMHAFKYTGUNLR9GPHWODRSIDYSJGCEOFKETUKMOBFWBXYXUVD9OAQPDDGAGXICQAVUGXGKPPZYZFNGLIYNEKTUCCWXZQLWZXDKIMQCBMCZIPN9JW9AXFNKMJMEXEJPRDJXVGLJSYUZXUPWK9E9UYBXIZXCGBRCLLWGI9RHVVUTSLYMIREQEPSMIKSG9FTOJWVECSXYERNKMSV9JZXKSTRFVYRENPUOALHGYGIRNMYBEHCZLVMPDEWEFACCJUHRJFPIX9GJZWNZLOLRIL9QJUVEENRRQEBF9OWMXEBU9APHLCTYOKBWFHIWARHPWSZWRWDPGFOGHEBHXHGIODOAG9FDMXCUNVPBPETXQSWHKLEKRHIPDZ9OYENTVRQTMNKGGLVDDIPMLFPFAKQGXQJYHRUXQNWH9STRJE9DEEPMRVBUWAZWWJUNUGBKX9ULPLYAOLKTPNLIAVDWWHK9NPFFL9QCZMXGVSWJWWC9JIJRZWJSDXSJEUYESNLINGBFBTEYMURQP9RRBD9XSNILEFSDXK9ZPRPLYJHYCWAKVOSVEOUFHBSVQOWBEKNVWKMICTUXNPROPPWGYRAPGLHFCMVWCLWEVKCASBGEGYOMJJLMDXDZPNADTSFPOFKRQKQSEULQTTJARXHKTDMNZEIVSZPIKFGTOMDFQEQEBOVGKGNC9ZMVF9AEAHEXZ9CXUAWTSGNPPUJQRREXAFJCLISFHQJUWFDNETQUCMXTEWNJUSAKRNXHA9RFYUDFYSNOFTLLYJH9YTOPWPITUBTFFGWAFXQ9CYORFC9XSUQRWEJXPWWF9IGDIYGTPLVOCFLLCHOZADUBFXQSHMXCURSCVHIJSVBBDYCYCLEEFRLIKMYHZJABDKFMWXOAVLYHKTYTWEXFUPKCPDCMTZYQRVFLGJEYFUXRDYNAJE9ZPA9XPDZCTOWZBANDDTSLI9OWYOYSGZGXWHNBQDMDGPZCMBERFPYQSBHQFFXQCK9YPBPBIOATPYJZPLJTDVYFKZFJCGSNLXJKSOYTHLJIAGLPBNLQVPBPTDSDVFNXXQRQQLNCBPAHOKWKRZFOCA9MAARTPFNYNOESMDTML99HHOIODAHDOYCULH9YNFUVWCLYECPHGVBXYVFMLYWXJU9ZTQXWBEPDWAOTMPADNUESWPQRXLJVLAKKMBRYQYLSLLBPFMZJPTKABQHMYBQAGQIYHIM9XSOEAFAYBHJLCVZIDFMNIQYXTFQCKXWWDKKKFKJWNT9DTTWLPNSOXBRVPPJMQUP9NEYJRNUWUODWZWTYHVTPSKLBPVRCERSMHFVVWJEUSHNXSYXFIBKHYRBSPAJPJ9V9UATWMVRXSCP9UVNM9ANXIXEKSXMXOXPOBHFLKIMHCEVJGAFWYFSMVNLYUSVBUHNGQNZMXIXBPUPDWAHDSIDNTTMHGEG9JFVDXGQVZMOWYRXUJWRIGTRAEVZGQDAXGILD9IQQUWAM9LHNHHRYXIYOZBDNEVORRDDOKICQCDMTDJUPIWZLEYCOOFSYOJLVKXLSBOTGAHXMLQYQEBEP9PZ9KKUOYIGECFTXXNBJIOJXMHCPJ9HOTLW9PSOHFAL9UTNUXVKF9IWMNJ9FITOC9FIYWQFSXPWAKYMIUHBK99DKLCTZUANNGVEEWEVVKMNQPQVVLG9ICNTACUTUPZYMYZKZTVEEVUGFWJGGWITRLJA9NHFYDVHGQLOBOUJHHCAHFRCFZIEXNZRLZFBVQUXOIKYKQR9JAQAZLAEYTBSREGMVEILTOWFXPHMNGGCTGLECZB"
	keyTrits := trinary.MustTrytesToTrits(key)
	digests, err := signing.Digests(keyTrits)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(trinary.MustTritsToTrytes(digests))
	// output: MUVADERKIZGMEYJVHGVWBKMQMMXOPWYVOXYPNAGDNKBLHWIBUALWLWSSNDXLYAIIWX9NQRRAOQIVIHWLAIRTWWSF9TGEIKFGMCDWNIXPIYKRTSBHJIONSTSSVUCBYHS9SOZB9PSAOSJUIYQYTUV9NXLZCZWHUALYWW
}

// i req: digests, The digests from which to derive the address from.
// o: Trits, The Trits representation of the address.
// o: error, Returned for internal errors.
func ExampleAddress() {
	digests := "MUVADERKIZGMEYJVHGVWBKMQMMXOPWYVOXYPNAGDNKBLHWIBUALWLWSSNDXLYAIIWX9NQRRAOQIVIHWLAIRTWWSF9TGEIKFGMCDWNIXPIYKRTSBHJIONSTSSVUCBYHS9SOZB9PSAOSJUIYQYTUV9NXLZCZWHUALYWW"
	digestsTrits := trinary.MustTrytesToTrits(digests)
	address, err := signing.Address(digestsTrits)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(address)
	// output: CLAAFXEY9AHHCSZCXNKDRZEJHIAFVKYORWNOZAGFPAZYNTSLCXUAG9WBSXBRXYEDPVPLXYVDCBCEKRUBD
}

// i req: bundleHash, The bundle hash to normalize.
// o: []int8, The normalized int8 slice without any 13/M values.
func ExampleNormalizedBundleHash() {
	bundleHash := "VAJOHANFEOTRSIPCLG9MIPENDFPLQQUGSBLBHMKZ9XVCUSWIKJOOHSPWJAXVLPTAKMPURYAYD9ONODVOW"
	normBundleHash := signing.NormalizedBundleHash(bundleHash)
	fmt.Println(normBundleHash)
	// output: []int8{8, 1, 10, -12, 8, 1, -13, 6, 5, -12, -7, -9, -8, 9, -11, 3, 12, 7, 0, 13, 9, -11, 5, -13, 4, 6, -11, -3, -10, -10, -6, 7, -8, 2, 12, 2, 8, 13, 11, -1, 0, -3, -5, 3, -6, -8, -4, 9, 11, 10, -12, -12, 8, -8, 13, 13, 13, 13, 13, -5, 12, -11, -7, 1, 11, 13, -11, -6, -9, -2, 1, -2, 4, 0, -12, -13, -12, 4, -5, -12, -4}
}

// i req: normalizedBundleHashFragment, The normalized bundle hash.
// i req: keyFragment, The fragment of the private key.
// o: Trits, The Trits representation of the signed message signature fragment.
// o: error, Returned for internal errors.
func ExampleSignatureFragment() {
	// the private key
	key := "D9DWVXXXMGBR9BKHQMMRTQIQROKTLOZNNYHHETDLHVE9FUIBGLEVSTHMJHNHXRRYWHBLNUICBOQHVGBRDSAJZKOL9EQNXGHCETSGW9LTGJPMKD9DEFITRUDVUUSKR9BMMAABVJIKICTLEGBSDMHGGOEALGWBZCF9YCWUANXAZOZJFQARYCKNOJTCXBYEWENKABGSUQXLOM9ZNEDVZQMUPWROQVQXTHLDUKMRULEZDOW9MEXSELWIKHYTCXQNWYIRLNQWUEOBPBQZTHNMRSMNIRQ9BKDHPDGN9VRZEBXKTFCCXOETJKX9YYJDB9HGLCOHMPOXRXBBV9YDSLYPBVABTNOWKXLSTGZAA9EXHLLGWIISYOSQQKYYOXSJDXVSLHDASEXDYJTFRRHLRMKLHZESDGTDGLUNLOWIETHXWEWPDCJRGMKFMXVNPLBGRYMAXTTAILUGAFUIXGTDP9JIVZFMEJKZRIYPLQNPSPQGZWZJZQXUVTJGRMCCJLNOLFZUCSEPNNNSYKLHJZH9YXTTTEPMDBWGXYWAMAMIDXYYSUBJKFAIH9XXBGBXMSBPJZGIDSELVW9PCTDYNKRCCZQIUIRIBNVWY9EFTZPJFLOF9GRJJQEPNOXHJHZE9KJXFX9TZZHNIVAQUSHAWTDQVGNMRMNRCKRZWCKWBBWKERTTBYKINHRYVTLVHSEDJTXVQUPBYJEQ9FFEEZOPG9HJCLPLOI9RKDDTCS9QWOZMYQJBKHRLHLCMUJODPLZFVFBINPLBLDI9EHNECSMTRWIIPVNRAWARRJSXMXCDSIENRC9MFIY9CYVLUOIWCONKRBVDRENXLXQKOGGLXWPSMQJUKUDRHBWVDQEWWSJKREESCIBIQNIWWPPWNLCBD9WYXOXCAHWLCIHTFSSNMAFPUUKZQEKWZKOYVIYSCYTOKNOSEPWPTVQBUUDUVDECYJWJQRYOPIKDLTHEAAJYQWFNMCGXRHBJTIPCPVGBVYOQTAWPHQBZEEYVNBNMHVRJDFBT9OFYRPUEIHBQIYB9TUZUMLLZAYVZKEBTQGKJPXPACTOKIG9DMUABKEFROB9FRM99JSPXZQDXTTGPPKUYMLLA9LEUFZYMWHDXTVOAUTJLC9SMEBTJMZPWNKGZTZZOYLOC9JGOQUOYYOSUSS9GNYNLROXMPFYIXPA9UOKATDFVJIGDGLMNGBCORUQMTFAQI9PDUUDRPBCRUJW9ZDHXDVGIIZLAKZQBQXZJZFZZDKKMTBMPYFEKJB9RHRZYJJMQTFLIYRHGUWEJWDD9K9INENNZPET9AJPCBPIBOZPHQTVDSCUGHFWZIDAOHXDCOUTANGUNBUHWGHXODGZCMANRSNUKFXXJLUGOZVZMFZTVSUBCGVMJAEQJFKC9HYMVNOULDFUCCXYBAGGXXDVE9ZETYAGJFXNE9DKVKDLGLYXHTFJDWKJZXUNUUVHKYRHANXJVFWVNFYWLEPMFFMMVIYRDUWSKOHDH9VFLMAKVPOVBH9GOS9TOUXDTMZPMTSCLPRUEPSNWAMVIONKWMUJYQYDZKSNCURUCARCGQUPTTVDBHQFPXMDYODFFAFTQVCJGXPZUZEZVCVVONHH9ZHLVYKUORZBGVBODHWMDMZCCQYZKQCZHHMQNPEVKTFAXFGNKCKJXQWSYKIFPVQFCPBTSCNQVHDTXBWWBGQZ9PAQCQFNSKQCAHIHOGYNVBENKJNWCBRFAU9H9ZDORYB9HDVIZOCUBP9PLCMXPEIOIOFPYWWLYWHVUJVMQCOGIMU99X9QRNURAAMROSXVYRJBK9LQCGREWFNLSPOXRGAINPAMOMZJMEAELKHJRQFLKYTEDEFGK9CNQUXPXHO9LUMZITY9MSESJJN9ZSZLNMIKLFABBKFMLGNRT9YLHUHBYMPPFXMWDPTM9T9GSBOYXSAGCRTPB9WVBOBA9QGTRSSJJUAKPZAOXNGSENLZEWOLRRFYVEGWOQJBGMXQRMDOQMLVIJKFUW9KCQHNZ9G9OBMCJSQIFPXDNQTOLFBHDHDQVBAZPBDW9SKV9TUGHW9JLPBCGLOWCFFYUCCDJARXMMZVXPIJKWTNKTXXRPFGCJUQVOASZYHDPAXFWTEQDCCZRQZCNHUXBEMLAENNPZHTID9PUCM9AQMJSJBXVH9TWUEUVRWZRKHAMPBKYYMRPXDNUWZHBCO9NQJ99MZAIBTXOFU9RDCKKIDOEACSPMIYPUFXXGBYCIVV9AHTLPPNGWJHOQBPNFKYKGCTFFEQLRPYEPMIOXLBYTKZ9YZPWXNBZVUBDAWHBDPIFCTGICEHQNTMT9AAUTI9KXAVKEZEWGFZSPVXHOVKNDXKINGSMSWGMUXHGTOAKHASIAIMEDPTPFYVE9BBSIB9OL9YBCEDHEVPCPJHTUCXFCFASRFZIDKZRSQJVEHSTKHVQXFYJFCOJDPDSQFAPEWVURE9FBSORQEXYBANRQGAWLNVPBXVEBHNULWKJJKPCXCP9QXNWVATDNRBHVWXRIQCNQMDJHMTIIHSHEYXJQVWTDASLSSISNTYMTGJUVHRPVXOZFOMYEIPK9DSLZNGEPPZDCCFGJVBJEHVSOPZANIJYYFZMSAAHNBEXJ9ONTQWMHAFKYTGUNLR9GPHWODRSIDYSJGCEOFKETUKMOBFWBXYXUVD9OAQPDDGAGXICQAVUGXGKPPZYZFNGLIYNEKTUCCWXZQLWZXDKIMQCBMCZIPN9JW9AXFNKMJMEXEJPRDJXVGLJSYUZXUPWK9E9UYBXIZXCGBRCLLWGI9RHVVUTSLYMIREQEPSMIKSG9FTOJWVECSXYERNKMSV9JZXKSTRFVYRENPUOALHGYGIRNMYBEHCZLVMPDEWEFACCJUHRJFPIX9GJZWNZLOLRIL9QJUVEENRRQEBF9OWMXEBU9APHLCTYOKBWFHIWARHPWSZWRWDPGFOGHEBHXHGIODOAG9FDMXCUNVPBPETXQSWHKLEKRHIPDZ9OYENTVRQTMNKGGLVDDIPMLFPFAKQGXQJYHRUXQNWH9STRJE9DEEPMRVBUWAZWWJUNUGBKX9ULPLYAOLKTPNLIAVDWWHK9NPFFL9QCZMXGVSWJWWC9JIJRZWJSDXSJEUYESNLINGBFBTEYMURQP9RRBD9XSNILEFSDXK9ZPRPLYJHYCWAKVOSVEOUFHBSVQOWBEKNVWKMICTUXNPROPPWGYRAPGLHFCMVWCLWEVKCASBGEGYOMJJLMDXDZPNADTSFPOFKRQKQSEULQTTJARXHKTDMNZEIVSZPIKFGTOMDFQEQEBOVGKGNC9ZMVF9AEAHEXZ9CXUAWTSGNPPUJQRREXAFJCLISFHQJUWFDNETQUCMXTEWNJUSAKRNXHA9RFYUDFYSNOFTLLYJH9YTOPWPITUBTFFGWAFXQ9CYORFC9XSUQRWEJXPWWF9IGDIYGTPLVOCFLLCHOZADUBFXQSHMXCURSCVHIJSVBBDYCYCLEEFRLIKMYHZJABDKFMWXOAVLYHKTYTWEXFUPKCPDCMTZYQRVFLGJEYFUXRDYNAJE9ZPA9XPDZCTOWZBANDDTSLI9OWYOYSGZGXWHNBQDMDGPZCMBERFPYQSBHQFFXQCK9YPBPBIOATPYJZPLJTDVYFKZFJCGSNLXJKSOYTHLJIAGLPBNLQVPBPTDSDVFNXXQRQQLNCBPAHOKWKRZFOCA9MAARTPFNYNOESMDTML99HHOIODAHDOYCULH9YNFUVWCLYECPHGVBXYVFMLYWXJU9ZTQXWBEPDWAOTMPADNUESWPQRXLJVLAKKMBRYQYLSLLBPFMZJPTKABQHMYBQAGQIYHIM9XSOEAFAYBHJLCVZIDFMNIQYXTFQCKXWWDKKKFKJWNT9DTTWLPNSOXBRVPPJMQUP9NEYJRNUWUODWZWTYHVTPSKLBPVRCERSMHFVVWJEUSHNXSYXFIBKHYRBSPAJPJ9V9UATWMVRXSCP9UVNM9ANXIXEKSXMXOXPOBHFLKIMHCEVJGAFWYFSMVNLYUSVBUHNGQNZMXIXBPUPDWAHDSIDNTTMHGEG9JFVDXGQVZMOWYRXUJWRIGTRAEVZGQDAXGILD9IQQUWAM9LHNHHRYXIYOZBDNEVORRDDOKICQCDMTDJUPIWZLEYCOOFSYOJLVKXLSBOTGAHXMLQYQEBEP9PZ9KKUOYIGECFTXXNBJIOJXMHCPJ9HOTLW9PSOHFAL9UTNUXVKF9IWMNJ9FITOC9FIYWQFSXPWAKYMIUHBK99DKLCTZUANNGVEEWEVVKMNQPQVVLG9ICNTACUTUPZYMYZKZTVEEVUGFWJGGWITRLJA9NHFYDVHGQLOBOUJHHCAHFRCFZIEXNZRLZFBVQUXOIKYKQR9JAQAZLAEYTBSREGMVEILTOWFXPHMNGGCTGLECZB"

	// the normalized bundle hash
	normalized := signing.NormalizedBundleHash("VAJOHANFEOTRSIPCLG9MIPENDFPLQQUGSBLBHMKZ9XVCUSWIKJOOHSPWJAXVLPTAKMPURYAYD9ONODVOW")

	// the amount of signature fragments
	fragments := make([]trinary.Trytes, consts.SecurityLevelMedium)
	keyTrits := trinary.MustTrytesToTrits(key)

	// we are generating 2 signature fragments since we are using SecurityLevelMedium/2
	for i := 0; i < int(consts.SecurityLevelMedium); i++ {
		// generate the signed signature fragment by supplying the correct
		// parts of the normalized bundle hash and private key
		frag, err := signing.SignatureFragment(
			normalized[i*consts.HashTrytesSize/3:(i+1)*consts.HashTrytesSize/3],
			keyTrits[i*consts.KeyFragmentLength:(i+1)*consts.KeyFragmentLength],
		)
		if err != nil {
			// handle error
			return
		}
		fragments[i] = trinary.MustTritsToTrytes(frag)
	}

	fmt.Println(fragments)
	// output:
	// "PDXDZUYANYKSVRVGNUIHLZWJZJMFQNZSDF9IZXVQVNVSMJB9KYURDMJAULFNWQFZJGDQWISOKTHFPMRJD99GMYTVB9W9NUBWMZFCOLUNAQULNXDYZRAIYLE99PNFWVFUFLTFGBBEWHXVGWYLKOKWEH9ROCIHSTYVLWRVUZH9UPRFIYEFNTOAPPCLRKCVINPWZHZFDSRLOLOHHDYXLFOECERGDCHDHQEQHD9HRQZK9X9KXWD9ABWCDELTQ9HURJQ99KR9RIMZSROMLSLJWJKWYDIMIXJBWXFSBSILGQOJORXHECMWQSAIX9UAPNWMHNNWRDRCMYUUTQOX9E9WXROHNAAEQAEKOGEPZOCSHZGDMGQGXXBDXSVDAJFIABJPYXZUUB9YGJQQKEUXGCRPAHUPCJCFS99HOQXMTGTLTAFAHLAXETHXXDV9YEGPSDOOCYEXUKKULS9TRTQCYBLL9XDLZKRXOFKDZJNTUKP9XCJQSYAVIEWAYLTG9ALVMPOZMRKK9JUDVOFRICJHZX9AWGXQRHGORJJMRGHANXUCQAILZQBESXNXKBVXZTBYCJRPMDJWROL9ZGSNZEZYWZHTZE9TJZXDAUKBHDSJWCLDFLDNLIKKWBIBQVF9RKIERKYQFYNAEXKZDTCBOHXQINIDLLHKGVMSNYCXCVGJOVRBLQDPMBHOWARWRWMLCGMMAASHFWBZGTCNGAHWUNJJBKRHFCYVFISKCLSO9EXWUWWAHCZVBJUPASHRGA9ARFLFRTABZMVDOHPORJYWWFDVXZCJWCAIBAGGQZTIRMSNARYYYBPLAXZPYIX9NULWHETDUCIVTRJAPLQVEUPSYPWBHGEPNCPBSPFOUIBHEUUJRWLSDPSBTOMITFREWHANYQWHRZDLPFLKDC9UADBIIVISENQZSPLYLWIXKINPVDJHYFZJIDXVUJUKPAOHURNCOBREXVNLOVSRHSUABOSDRHIYPMXPWVGFKSRDUQBUIYDTKGWOYEV9IQTBNUEFYCNVVPUOYJATSVWLMOISRRENCKY9MWHZAHI9QTJDJGIMBGRWZTD9HIDJPDAZHPZ9BQRRVWQWWBVAZPZJTKCBEFYBYNQAOULRALPQVWIEWWZFCUXVBB9TFXCGQFZUE9TUAOYITCHBLMG9XLVRHAMAOGBKMWNRPEFBDIPCMRFZUQNHBXOVGVVSIBHOCOAMEVMQNITSCKXDQOYWYRXCRCUKQROZIPPNBNOFRBWEREDCRPNQJLZEDCODIGJWLLQSDTVL9TQLXMCDZCBXHMSTDQNJJIOLQFZEEOCBWZUQMBHXKNOHGLEY9RUPPBQPUMROEKEMBNFSINZIRRPWMFJHANFJSUC9ZYHULTZDLSKOPSCLKOUVW9KBYOPXDDNAZBJZGFWMZKFECDC9RVGNWDESFPHBPBLFAOWTRBKQBKQNQJV9ETJNMOLLDZQQAIJF9MIMFLVB9XVGYGAMYRCSVVFEEUPRLJSZTAPEYJLOGAPEPUXRATWHMMJZRUJHZWQFSMFEMQQWJBJXRFPEJVUU9RL9TBDZKSNCURUCARCGQUPTTVDBHQFPXMDYODFFAFTQVCJGXPZUZEZVCVVONHH9ZHLVYKUORZBGVBODHWMDMZCCCHPVFMARSWEGYWRAFBAGAISLXNWCCBXHPOMBKRXWVZFAVLOZKHKCJLUPELPQIYBSHOZRUWGWHFLJVJSUZFSLYGI9OHESGTRMD9OPPNPMKKOLUIO9XNIPPHLM99NGVBYDCWJXEOZO9AIIHVTTLRSEVMCSQNHBUGXIZBGXTHDKCJGXBWPKOFYOCDZLDAN9PFWGORYIYIYUSI9NXRFQSLKNQY9VYQLZDMBJ9INAALKOMGMJSGGZPSXMIJZYWOGI9YJQVYRWHLERDIATQNTQJFHSWCGWFCHWTQOXBJFKEUSYVQJY9ZBJXHXCFFOSYMFONGVHUWIZICQLOIETCXGIOAECKUKXJI9NQZKUDKJ9LMZVAC9SBMPEKHWDDEWPYFTCZNKPMRVAOEHXPOVWZBVIMJEVXUVINIIHKDXWVLWJ9NDUUSCTTYAEVCXFXXZQIHAHNVOTKEAUTJVIB9T9D9WTJAANUOJFDXBQBTOSUCUK9CVRSIL9XDNLMPMGOMMITSGCGDWBZPRDOCOOOBUPKYAZHGEOBBWQWXHETDHYRXUULMPNVDWWABGQIPCTIA9",
	// "KQFGJD9NZPAJGHGQSQLCYLNXZKIFRRFDKSBFFYMITXAQIMUDVTQV9GINJQKZAJJLKMVXVTJMCBGLNXQUYNHZCOI9TPADSBZMDASJPLVNTWN9FMKFSVKFPNGQWBFYMLE9SNQARTFUQPMDHMGKOINWXKPAWEKGDPLSQYTVPWEWRWPRUDDDTNRHKPLKQWZBQFSHLPYABTJPEDLEQ9JOUNUALMYZFLNVIELETFJRRZXTYDULMQRZBPCV9QEDNNNUCESOYVGQHX9HQAVHBPLMJJHLHZTZM9ZJSKODVMUCSJHXZAQESJLTVRIWWTHNDRHTRNTEQLKCNAXANVNPAZILDSDDVXCOMRBJLDKWXFFHZFICZFAUOT9RDKMQZPHMXXOHLKPE9LFMPUZEQVSHMTDHTUERWFMDWINQQOGBKMNJWAWAGVYVWFMUBUEVVGOVAKUSUIDPDSVDCGK99GNFYMADPNBSCININUVCGHULYRIFWZIJVZF9MERWKMLZJYUIGWXUEONCAIMNVLDFMNQ9MNUJTDVWKLUXBVOXFVZMQZWECTQQLJZHRQIVAXWRZFYMNZNQRYNDOWQIMTGAETINJQZYIFOZZYZPSYWGCJEMUVBWLLHASHWESJXPATRYCUYKPYOIMEI9YJDDIXRDHQTMXNWNPVFWWKLCWHOHIIKSUDMAUCSBTZFCSPBOHRIUYOFZSAYEUYUUV9GTFGOWNSFCALKKABRBXIEJ9BSVUEPHGVP9PEZABBY9WUFDJXLFRRPGPDPVRAYHONWAEUFCQNJFPOESGPUXJCQZRWYDWDJIECIYION9W9QJYHRUXQNWH9STRJE9DEEPMRVBUWAZWWJUNUGBKX9ULPLYAOLKTPNLIAVDWWHK9NPFFL9QCZMXGVSWJWWLXJKFJPGRUGYHCCXEGHMHYPSZ9BYJGAAO9IKMUBEWEXKYDIHHCMUFOTZTRCCWYLUTDGNDYWFMTBXQQNPZBWCZIEWQGRB99LSPNEPPYLYAOATFBFDYKFJVWAIGKOWMHVVNSCWTQINKCNFATXWYDNWVFKYVJOGTRTKJXGBGMSVFSBIHRAV9LIKKGXFVVZEBSQNDMUDKXKPWDAKBZ9LGAGWFVRGMQHYLPIHIFIHWVJYQTFFESWZSWXWSGHKFITTECNEOSND9KZVENVTEUKJGVPWRVTOECNFYKLQSWLXYTRLQNUMU99MMRTDTPCFCT9JWUVRKOBCYUGJINRHAXWIJETZUZVQAFZKXNDYJIYEVJVJKKIRNDUAXPPYLWQHHCTCCUCNEIHNJZUS9SKSC9BCNBEFDFB9LAAPXPPNKEHVDJYLIKJSHMUFQYWYPTXRUXWBGENPHQWVUQJAICGOSDHCVTUERBAKOOVQEVPOMZRSJBHBPSMBHRTVIFYC9H9SJXLBEIJACFQP9OLWUWCL9FLDGJFCJBRXYNDARGFOUJHFTVWUTMBVKIHHDLIRXNAIIGMDW9AH9XCWMFSJWHKAEBYYNDZNJV9WIBNDFSCKMWOKSQSBKSLMIHSWICNNPNMBYMDKEBBSWRSHBZBXOOQEQONNBHFH9AF9NSYHZASPRVFURAOQCQGNMPOZUXKRYSGKQXNGUDBLRZMSD9ITWZWG9OSAAZCCOGPAXI9LCMIKKEDHCUQXSYC9MXASXEJIGIBS9CGELPOMZODPHTNUPKS9VRO9IPBAYJOAZNBUYJFJTZWO9IRDLDQ9TWKZEZHHBJZGV9EWXZYYOVSDJTOQOAZ9AUAAAGGJXJFIESPDAXGVHVAMWEYK9ICKOCSIV9YATX9HPTDA9RUAGU9ZXEMJMPSYFFHDAVFUXFPRUKKFPITCUAUSXP9MJXAQTEUNIIABWDVFGUJBIE9DVVMYKFVBZHACKZYEMKPEAJEZSBMTRJLRWYEIJWPVMAJRDSVGXKETEDEUYXPNILWXXNK9LNW9B9ZHWDOOZPQXES9VXNK9AGOCKISNOOOEDQMMDMUMRUOIVYEUBCJRGZEEUXGUK9EJEBECSYODYRLXRLFLOFZONTCQLXLQXXRJNUJY9DBXPHMLIRIKAWELNTVKVPBGGDEELPCH9BTEQYVLFMFTY99ENBTGWJNBZWRFHQ9WYQBZRYUXMHUFYGXNYNWDRKODSBPRBOYCZLGDWJUVXMFVFZXILWIZEQEPSFSCXILLVQBKSDHHSPWIGCMEUGXDD9WUIOWQUWAQLYPA",
}

// i req: normalizedBundleHashFragment, The fragment of the normalized bundle hash.
// i req: signatureFragment, The signature fragment corresponding to the bundle hash fragment.
// o: Trits, The Trits representation of the digest.
// o: error, Returned for internal errors.
func ExampleDigest() {
	normalized := signing.NormalizedBundleHash("VAJOHANFEOTRSIPCLG9MIPENDFPLQQUGSBLBHMKZ9XVCUSWIKJOOHSPWJAXVLPTAKMPURYAYD9ONODVOW")
	normalizedFragments := make([][]int8, 3)
	for i := 0; i < 3; i++ {
		normalizedFragments[i] = normalized[i*27 : (i+1)*27]
	}

	// the signed messages extracted from transactions
	signedFragments := []trinary.Trytes{
		"PDXDZUYANYKSVRVGNUIHLZWJZJMFQNZSDF9IZXVQVNVSMJB9KYURDMJAULFNWQFZJGDQWISOKTHFPMRJD99GMYTVB9W9NUBWMZFCOLUNAQULNXDYZRAIYLE99PNFWVFUFLTFGBBEWHXVGWYLKOKWEH9ROCIHSTYVLWRVUZH9UPRFIYEFNTOAPPCLRKCVINPWZHZFDSRLOLOHHDYXLFOECERGDCHDHQEQHD9HRQZK9X9KXWD9ABWCDELTQ9HURJQ99KR9RIMZSROMLSLJWJKWYDIMIXJBWXFSBSILGQOJORXHECMWQSAIX9UAPNWMHNNWRDRCMYUUTQOX9E9WXROHNAAEQAEKOGEPZOCSHZGDMGQGXXBDXSVDAJFIABJPYXZUUB9YGJQQKEUXGCRPAHUPCJCFS99HOQXMTGTLTAFAHLAXETHXXDV9YEGPSDOOCYEXUKKULS9TRTQCYBLL9XDLZKRXOFKDZJNTUKP9XCJQSYAVIEWAYLTG9ALVMPOZMRKK9JUDVOFRICJHZX9AWGXQRHGORJJMRGHANXUCQAILZQBESXNXKBVXZTBYCJRPMDJWROL9ZGSNZEZYWZHTZE9TJZXDAUKBHDSJWCLDFLDNLIKKWBIBQVF9RKIERKYQFYNAEXKZDTCBOHXQINIDLLHKGVMSNYCXCVGJOVRBLQDPMBHOWARWRWMLCGMMAASHFWBZGTCNGAHWUNJJBKRHFCYVFISKCLSO9EXWUWWAHCZVBJUPASHRGA9ARFLFRTABZMVDOHPORJYWWFDVXZCJWCAIBAGGQZTIRMSNARYYYBPLAXZPYIX9NULWHETDUCIVTRJAPLQVEUPSYPWBHGEPNCPBSPFOUIBHEUUJRWLSDPSBTOMITFREWHANYQWHRZDLPFLKDC9UADBIIVISENQZSPLYLWIXKINPVDJHYFZJIDXVUJUKPAOHURNCOBREXVNLOVSRHSUABOSDRHIYPMXPWVGFKSRDUQBUIYDTKGWOYEV9IQTBNUEFYCNVVPUOYJATSVWLMOISRRENCKY9MWHZAHI9QTJDJGIMBGRWZTD9HIDJPDAZHPZ9BQRRVWQWWBVAZPZJTKCBEFYBYNQAOULRALPQVWIEWWZFCUXVBB9TFXCGQFZUE9TUAOYITCHBLMG9XLVRHAMAOGBKMWNRPEFBDIPCMRFZUQNHBXOVGVVSIBHOCOAMEVMQNITSCKXDQOYWYRXCRCUKQROZIPPNBNOFRBWEREDCRPNQJLZEDCODIGJWLLQSDTVL9TQLXMCDZCBXHMSTDQNJJIOLQFZEEOCBWZUQMBHXKNOHGLEY9RUPPBQPUMROEKEMBNFSINZIRRPWMFJHANFJSUC9ZYHULTZDLSKOPSCLKOUVW9KBYOPXDDNAZBJZGFWMZKFECDC9RVGNWDESFPHBPBLFAOWTRBKQBKQNQJV9ETJNMOLLDZQQAIJF9MIMFLVB9XVGYGAMYRCSVVFEEUPRLJSZTAPEYJLOGAPEPUXRATWHMMJZRUJHZWQFSMFEMQQWJBJXRFPEJVUU9RL9TBDZKSNCURUCARCGQUPTTVDBHQFPXMDYODFFAFTQVCJGXPZUZEZVCVVONHH9ZHLVYKUORZBGVBODHWMDMZCCCHPVFMARSWEGYWRAFBAGAISLXNWCCBXHPOMBKRXWVZFAVLOZKHKCJLUPELPQIYBSHOZRUWGWHFLJVJSUZFSLYGI9OHESGTRMD9OPPNPMKKOLUIO9XNIPPHLM99NGVBYDCWJXEOZO9AIIHVTTLRSEVMCSQNHBUGXIZBGXTHDKCJGXBWPKOFYOCDZLDAN9PFWGORYIYIYUSI9NXRFQSLKNQY9VYQLZDMBJ9INAALKOMGMJSGGZPSXMIJZYWOGI9YJQVYRWHLERDIATQNTQJFHSWCGWFCHWTQOXBJFKEUSYVQJY9ZBJXHXCFFOSYMFONGVHUWIZICQLOIETCXGIOAECKUKXJI9NQZKUDKJ9LMZVAC9SBMPEKHWDDEWPYFTCZNKPMRVAOEHXPOVWZBVIMJEVXUVINIIHKDXWVLWJ9NDUUSCTTYAEVCXFXXZQIHAHNVOTKEAUTJVIB9T9D9WTJAANUOJFDXBQBTOSUCUK9CVRSIL9XDNLMPMGOMMITSGCGDWBZPRDOCOOOBUPKYAZHGEOBBWQWXHETDHYRXUULMPNVDWWABGQIPCTIA9",
		"KQFGJD9NZPAJGHGQSQLCYLNXZKIFRRFDKSBFFYMITXAQIMUDVTQV9GINJQKZAJJLKMVXVTJMCBGLNXQUYNHZCOI9TPADSBZMDASJPLVNTWN9FMKFSVKFPNGQWBFYMLE9SNQARTFUQPMDHMGKOINWXKPAWEKGDPLSQYTVPWEWRWPRUDDDTNRHKPLKQWZBQFSHLPYABTJPEDLEQ9JOUNUALMYZFLNVIELETFJRRZXTYDULMQRZBPCV9QEDNNNUCESOYVGQHX9HQAVHBPLMJJHLHZTZM9ZJSKODVMUCSJHXZAQESJLTVRIWWTHNDRHTRNTEQLKCNAXANVNPAZILDSDDVXCOMRBJLDKWXFFHZFICZFAUOT9RDKMQZPHMXXOHLKPE9LFMPUZEQVSHMTDHTUERWFMDWINQQOGBKMNJWAWAGVYVWFMUBUEVVGOVAKUSUIDPDSVDCGK99GNFYMADPNBSCININUVCGHULYRIFWZIJVZF9MERWKMLZJYUIGWXUEONCAIMNVLDFMNQ9MNUJTDVWKLUXBVOXFVZMQZWECTQQLJZHRQIVAXWRZFYMNZNQRYNDOWQIMTGAETINJQZYIFOZZYZPSYWGCJEMUVBWLLHASHWESJXPATRYCUYKPYOIMEI9YJDDIXRDHQTMXNWNPVFWWKLCWHOHIIKSUDMAUCSBTZFCSPBOHRIUYOFZSAYEUYUUV9GTFGOWNSFCALKKABRBXIEJ9BSVUEPHGVP9PEZABBY9WUFDJXLFRRPGPDPVRAYHONWAEUFCQNJFPOESGPUXJCQZRWYDWDJIECIYION9W9QJYHRUXQNWH9STRJE9DEEPMRVBUWAZWWJUNUGBKX9ULPLYAOLKTPNLIAVDWWHK9NPFFL9QCZMXGVSWJWWLXJKFJPGRUGYHCCXEGHMHYPSZ9BYJGAAO9IKMUBEWEXKYDIHHCMUFOTZTRCCWYLUTDGNDYWFMTBXQQNPZBWCZIEWQGRB99LSPNEPPYLYAOATFBFDYKFJVWAIGKOWMHVVNSCWTQINKCNFATXWYDNWVFKYVJOGTRTKJXGBGMSVFSBIHRAV9LIKKGXFVVZEBSQNDMUDKXKPWDAKBZ9LGAGWFVRGMQHYLPIHIFIHWVJYQTFFESWZSWXWSGHKFITTECNEOSND9KZVENVTEUKJGVPWRVTOECNFYKLQSWLXYTRLQNUMU99MMRTDTPCFCT9JWUVRKOBCYUGJINRHAXWIJETZUZVQAFZKXNDYJIYEVJVJKKIRNDUAXPPYLWQHHCTCCUCNEIHNJZUS9SKSC9BCNBEFDFB9LAAPXPPNKEHVDJYLIKJSHMUFQYWYPTXRUXWBGENPHQWVUQJAICGOSDHCVTUERBAKOOVQEVPOMZRSJBHBPSMBHRTVIFYC9H9SJXLBEIJACFQP9OLWUWCL9FLDGJFCJBRXYNDARGFOUJHFTVWUTMBVKIHHDLIRXNAIIGMDW9AH9XCWMFSJWHKAEBYYNDZNJV9WIBNDFSCKMWOKSQSBKSLMIHSWICNNPNMBYMDKEBBSWRSHBZBXOOQEQONNBHFH9AF9NSYHZASPRVFURAOQCQGNMPOZUXKRYSGKQXNGUDBLRZMSD9ITWZWG9OSAAZCCOGPAXI9LCMIKKEDHCUQXSYC9MXASXEJIGIBS9CGELPOMZODPHTNUPKS9VRO9IPBAYJOAZNBUYJFJTZWO9IRDLDQ9TWKZEZHHBJZGV9EWXZYYOVSDJTOQOAZ9AUAAAGGJXJFIESPDAXGVHVAMWEYK9ICKOCSIV9YATX9HPTDA9RUAGU9ZXEMJMPSYFFHDAVFUXFPRUKKFPITCUAUSXP9MJXAQTEUNIIABWDVFGUJBIE9DVVMYKFVBZHACKZYEMKPEAJEZSBMTRJLRWYEIJWPVMAJRDSVGXKETEDEUYXPNILWXXNK9LNW9B9ZHWDOOZPQXES9VXNK9AGOCKISNOOOEDQMMDMUMRUOIVYEUBCJRGZEEUXGUK9EJEBECSYODYRLXRLFLOFZONTCQLXLQXXRJNUJY9DBXPHMLIRIKAWELNTVKVPBGGDEELPCH9BTEQYVLFMFTY99ENBTGWJNBZWRFHQ9WYQBZRYUXMHUFYGXNYNWDRKODSBPRBOYCZLGDWJUVXMFVFZXILWIZEQEPSFSCXILLVQBKSDHHSPWIGCMEUGXDD9WUIOWQUWAQLYPA",
	}

	digests := make([]trinary.Trits, 2)
	for i := 0; i < len(signedFragments); i++ {
		digest, err := signing.Digest(
			normalizedFragments[i%3],
			trinary.MustTrytesToTrits(signedFragments[i]),
		)
		if err != nil {
			// handle error
			return
		}
		digests[i] = digest
	}

	fmt.Println(trinary.MustTritsToTrytes(digests[0]))
	fmt.Println(trinary.MustTritsToTrytes(digests[1]))
	// output:
	// MUVADERKIZGMEYJVHGVWBKMQMMXOPWYVOXYPNAGDNKBLHWIBUALWLWSSNDXLYAIIWX9NQRRAOQIVIHWLA
	// IRTWWSF9TGEIKFGMCDWNIXPIYKRTSBHJIONSTSSVUCBYHS9SOZB9PSAOSJUIYQYTUV9NXLZCZWHUALYWW
}

// i req: expectedAddress, The address to validate against to check whether the signatures are valid.
// i req: fragments, The signed signature fragments.
// i req: bundleHash, The hash of the bundle.
// o: bool, Whether the signatures are valid.
// o: error, Returned for internal errors.
func ExampleValidateSignatures() {
	address := "CLAAFXEY9AHHCSZCXNKDRZEJHIAFVKYORWNOZAGFPAZYNTSLCXUAG9WBSXBRXYEDPVPLXYVDCBCEKRUBD"
	fragments := []trinary.Trytes{
		"PDXDZUYANYKSVRVGNUIHLZWJZJMFQNZSDF9IZXVQVNVSMJB9KYURDMJAULFNWQFZJGDQWISOKTHFPMRJD99GMYTVB9W9NUBWMZFCOLUNAQULNXDYZRAIYLE99PNFWVFUFLTFGBBEWHXVGWYLKOKWEH9ROCIHSTYVLWRVUZH9UPRFIYEFNTOAPPCLRKCVINPWZHZFDSRLOLOHHDYXLFOECERGDCHDHQEQHD9HRQZK9X9KXWD9ABWCDELTQ9HURJQ99KR9RIMZSROMLSLJWJKWYDIMIXJBWXFSBSILGQOJORXHECMWQSAIX9UAPNWMHNNWRDRCMYUUTQOX9E9WXROHNAAEQAEKOGEPZOCSHZGDMGQGXXBDXSVDAJFIABJPYXZUUB9YGJQQKEUXGCRPAHUPCJCFS99HOQXMTGTLTAFAHLAXETHXXDV9YEGPSDOOCYEXUKKULS9TRTQCYBLL9XDLZKRXOFKDZJNTUKP9XCJQSYAVIEWAYLTG9ALVMPOZMRKK9JUDVOFRICJHZX9AWGXQRHGORJJMRGHANXUCQAILZQBESXNXKBVXZTBYCJRPMDJWROL9ZGSNZEZYWZHTZE9TJZXDAUKBHDSJWCLDFLDNLIKKWBIBQVF9RKIERKYQFYNAEXKZDTCBOHXQINIDLLHKGVMSNYCXCVGJOVRBLQDPMBHOWARWRWMLCGMMAASHFWBZGTCNGAHWUNJJBKRHFCYVFISKCLSO9EXWUWWAHCZVBJUPASHRGA9ARFLFRTABZMVDOHPORJYWWFDVXZCJWCAIBAGGQZTIRMSNARYYYBPLAXZPYIX9NULWHETDUCIVTRJAPLQVEUPSYPWBHGEPNCPBSPFOUIBHEUUJRWLSDPSBTOMITFREWHANYQWHRZDLPFLKDC9UADBIIVISENQZSPLYLWIXKINPVDJHYFZJIDXVUJUKPAOHURNCOBREXVNLOVSRHSUABOSDRHIYPMXPWVGFKSRDUQBUIYDTKGWOYEV9IQTBNUEFYCNVVPUOYJATSVWLMOISRRENCKY9MWHZAHI9QTJDJGIMBGRWZTD9HIDJPDAZHPZ9BQRRVWQWWBVAZPZJTKCBEFYBYNQAOULRALPQVWIEWWZFCUXVBB9TFXCGQFZUE9TUAOYITCHBLMG9XLVRHAMAOGBKMWNRPEFBDIPCMRFZUQNHBXOVGVVSIBHOCOAMEVMQNITSCKXDQOYWYRXCRCUKQROZIPPNBNOFRBWEREDCRPNQJLZEDCODIGJWLLQSDTVL9TQLXMCDZCBXHMSTDQNJJIOLQFZEEOCBWZUQMBHXKNOHGLEY9RUPPBQPUMROEKEMBNFSINZIRRPWMFJHANFJSUC9ZYHULTZDLSKOPSCLKOUVW9KBYOPXDDNAZBJZGFWMZKFECDC9RVGNWDESFPHBPBLFAOWTRBKQBKQNQJV9ETJNMOLLDZQQAIJF9MIMFLVB9XVGYGAMYRCSVVFEEUPRLJSZTAPEYJLOGAPEPUXRATWHMMJZRUJHZWQFSMFEMQQWJBJXRFPEJVUU9RL9TBDZKSNCURUCARCGQUPTTVDBHQFPXMDYODFFAFTQVCJGXPZUZEZVCVVONHH9ZHLVYKUORZBGVBODHWMDMZCCCHPVFMARSWEGYWRAFBAGAISLXNWCCBXHPOMBKRXWVZFAVLOZKHKCJLUPELPQIYBSHOZRUWGWHFLJVJSUZFSLYGI9OHESGTRMD9OPPNPMKKOLUIO9XNIPPHLM99NGVBYDCWJXEOZO9AIIHVTTLRSEVMCSQNHBUGXIZBGXTHDKCJGXBWPKOFYOCDZLDAN9PFWGORYIYIYUSI9NXRFQSLKNQY9VYQLZDMBJ9INAALKOMGMJSGGZPSXMIJZYWOGI9YJQVYRWHLERDIATQNTQJFHSWCGWFCHWTQOXBJFKEUSYVQJY9ZBJXHXCFFOSYMFONGVHUWIZICQLOIETCXGIOAECKUKXJI9NQZKUDKJ9LMZVAC9SBMPEKHWDDEWPYFTCZNKPMRVAOEHXPOVWZBVIMJEVXUVINIIHKDXWVLWJ9NDUUSCTTYAEVCXFXXZQIHAHNVOTKEAUTJVIB9T9D9WTJAANUOJFDXBQBTOSUCUK9CVRSIL9XDNLMPMGOMMITSGCGDWBZPRDOCOOOBUPKYAZHGEOBBWQWXHETDHYRXUULMPNVDWWABGQIPCTIA9",
		"KQFGJD9NZPAJGHGQSQLCYLNXZKIFRRFDKSBFFYMITXAQIMUDVTQV9GINJQKZAJJLKMVXVTJMCBGLNXQUYNHZCOI9TPADSBZMDASJPLVNTWN9FMKFSVKFPNGQWBFYMLE9SNQARTFUQPMDHMGKOINWXKPAWEKGDPLSQYTVPWEWRWPRUDDDTNRHKPLKQWZBQFSHLPYABTJPEDLEQ9JOUNUALMYZFLNVIELETFJRRZXTYDULMQRZBPCV9QEDNNNUCESOYVGQHX9HQAVHBPLMJJHLHZTZM9ZJSKODVMUCSJHXZAQESJLTVRIWWTHNDRHTRNTEQLKCNAXANVNPAZILDSDDVXCOMRBJLDKWXFFHZFICZFAUOT9RDKMQZPHMXXOHLKPE9LFMPUZEQVSHMTDHTUERWFMDWINQQOGBKMNJWAWAGVYVWFMUBUEVVGOVAKUSUIDPDSVDCGK99GNFYMADPNBSCININUVCGHULYRIFWZIJVZF9MERWKMLZJYUIGWXUEONCAIMNVLDFMNQ9MNUJTDVWKLUXBVOXFVZMQZWECTQQLJZHRQIVAXWRZFYMNZNQRYNDOWQIMTGAETINJQZYIFOZZYZPSYWGCJEMUVBWLLHASHWESJXPATRYCUYKPYOIMEI9YJDDIXRDHQTMXNWNPVFWWKLCWHOHIIKSUDMAUCSBTZFCSPBOHRIUYOFZSAYEUYUUV9GTFGOWNSFCALKKABRBXIEJ9BSVUEPHGVP9PEZABBY9WUFDJXLFRRPGPDPVRAYHONWAEUFCQNJFPOESGPUXJCQZRWYDWDJIECIYION9W9QJYHRUXQNWH9STRJE9DEEPMRVBUWAZWWJUNUGBKX9ULPLYAOLKTPNLIAVDWWHK9NPFFL9QCZMXGVSWJWWLXJKFJPGRUGYHCCXEGHMHYPSZ9BYJGAAO9IKMUBEWEXKYDIHHCMUFOTZTRCCWYLUTDGNDYWFMTBXQQNPZBWCZIEWQGRB99LSPNEPPYLYAOATFBFDYKFJVWAIGKOWMHVVNSCWTQINKCNFATXWYDNWVFKYVJOGTRTKJXGBGMSVFSBIHRAV9LIKKGXFVVZEBSQNDMUDKXKPWDAKBZ9LGAGWFVRGMQHYLPIHIFIHWVJYQTFFESWZSWXWSGHKFITTECNEOSND9KZVENVTEUKJGVPWRVTOECNFYKLQSWLXYTRLQNUMU99MMRTDTPCFCT9JWUVRKOBCYUGJINRHAXWIJETZUZVQAFZKXNDYJIYEVJVJKKIRNDUAXPPYLWQHHCTCCUCNEIHNJZUS9SKSC9BCNBEFDFB9LAAPXPPNKEHVDJYLIKJSHMUFQYWYPTXRUXWBGENPHQWVUQJAICGOSDHCVTUERBAKOOVQEVPOMZRSJBHBPSMBHRTVIFYC9H9SJXLBEIJACFQP9OLWUWCL9FLDGJFCJBRXYNDARGFOUJHFTVWUTMBVKIHHDLIRXNAIIGMDW9AH9XCWMFSJWHKAEBYYNDZNJV9WIBNDFSCKMWOKSQSBKSLMIHSWICNNPNMBYMDKEBBSWRSHBZBXOOQEQONNBHFH9AF9NSYHZASPRVFURAOQCQGNMPOZUXKRYSGKQXNGUDBLRZMSD9ITWZWG9OSAAZCCOGPAXI9LCMIKKEDHCUQXSYC9MXASXEJIGIBS9CGELPOMZODPHTNUPKS9VRO9IPBAYJOAZNBUYJFJTZWO9IRDLDQ9TWKZEZHHBJZGV9EWXZYYOVSDJTOQOAZ9AUAAAGGJXJFIESPDAXGVHVAMWEYK9ICKOCSIV9YATX9HPTDA9RUAGU9ZXEMJMPSYFFHDAVFUXFPRUKKFPITCUAUSXP9MJXAQTEUNIIABWDVFGUJBIE9DVVMYKFVBZHACKZYEMKPEAJEZSBMTRJLRWYEIJWPVMAJRDSVGXKETEDEUYXPNILWXXNK9LNW9B9ZHWDOOZPQXES9VXNK9AGOCKISNOOOEDQMMDMUMRUOIVYEUBCJRGZEEUXGUK9EJEBECSYODYRLXRLFLOFZONTCQLXLQXXRJNUJY9DBXPHMLIRIKAWELNTVKVPBGGDEELPCH9BTEQYVLFMFTY99ENBTGWJNBZWRFHQ9WYQBZRYUXMHUFYGXNYNWDRKODSBPRBOYCZLGDWJUVXMFVFZXILWIZEQEPSFSCXILLVQBKSDHHSPWIGCMEUGXDD9WUIOWQUWAQLYPA",
	}
	bundleHash := "VAJOHANFEOTRSIPCLG9MIPENDFPLQQUGSBLBHMKZ9XVCUSWIKJOOHSPWJAXVLPTAKMPURYAYD9ONODVOW"

	valid, err := signing.ValidateSignatures(address, fragments, bundleHash)
	if err != nil {
		// handle error
		return
	}
	fmt.Println(valid)
	// output: true
}
