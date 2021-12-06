package pinyin

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type pinyinFunc func(string, Args) [][]string
type testCase struct {
	args   Args
	result [][]string
}

func testPinyin(t *testing.T, s string, d []testCase, f pinyinFunc) {
	for _, tc := range d {
		v := f(s, tc.args)
		if !reflect.DeepEqual(v, tc.result) {
			t.Errorf("Expected %s, got %s", tc.result, v)
		}
	}
}

var finals = []string{
	// a
	"a1", "ā", "a2", "á", "a3", "ǎ", "a4", "à",
	// o
	"o1", "ō", "o2", "ó", "o3", "ǒ", "o4", "ò",
	// e
	"e1", "ē", "e2", "é", "e3", "ě", "e4", "è",
	// i
	"i1", "ī", "i2", "í", "i3", "ǐ", "i4", "ì",
	// u
	"u1", "ū", "u2", "ú", "u3", "ǔ", "u4", "ù",
	// v
	"v1", "ǖ", "v2", "ǘ", "v3", "ǚ", "v4", "ǜ",

	// ai
	"ai1", "āi", "ai2", "ái", "ai3", "ǎi", "ai4", "ài",
	// ei
	"ei1", "ēi", "ei2", "éi", "ei3", "ěi", "ei4", "èi",
	// ui
	"ui1", "uī", "ui2", "uí", "ui3", "uǐ", "ui4", "uì",
	// ao
	"ao1", "āo", "ao2", "áo", "ao3", "ǎo", "ao4", "ào",
	// ou
	"ou1", "ōu", "ou2", "óu", "ou3", "ǒu", "ou4", "òu",
	// iu
	"iu1", "īu", "iu2", "íu", "iu3", "ǐu", "iu4", "ìu",

	// ie
	"ie1", "iē", "ie2", "ié", "ie3", "iě", "ie4", "iè",
	// ve
	"ue1", "üē", "ue2", "üé", "ue3", "üě", "ue4", "üè",
	// er
	"er1", "ēr", "er2", "ér", "er3", "ěr", "er4", "èr",

	// an
	"an1", "ān", "an2", "án", "an3", "ǎn", "an4", "àn",
	// en
	"en1", "ēn", "en2", "én", "en3", "ěn", "en4", "èn",
	// in
	"in1", "īn", "in2", "ín", "in3", "ǐn", "in4", "ìn",
	// un/vn
	"un1", "ūn", "un2", "ún", "un3", "ǔn", "un4", "ùn",

	// ang
	"ang1", "āng", "ang2", "áng", "ang3", "ǎng", "ang4", "àng",
	// eng
	"eng1", "ēng", "eng2", "éng", "eng3", "ěng", "eng4", "èng",
	// ing
	"ing1", "īng", "ing2", "íng", "ing3", "ǐng", "ing4", "ìng",
	// ong
	"ong1", "ōng", "ong2", "óng", "ong3", "ǒng", "ong4", "òng",
}

func TestPinyin(t *testing.T) {
	cc := strings.Join(finals, " ")
	for i, v := range PinyinDict {
		l := strings.Split(v, ",")
		if len(l) > 1 {
			for _, vv := range l {

				jn := strings.Index(cc, vv)
				if jn < 0 {
					fmt.Printf("0x%x; %v\n", i, v)
				}

			}
		}
	}
	py := NewArgs()
	py.Heteronym = true
	for _, v := range "单系黄河清周志华龟抬头望明月" {
		fmt.Printf("%x; %v\n", v, string(v))
	}
	s := Pinyin("单系黄河清周志华龟抬头望明月", py)
	fmt.Println("s:", s)
	t.Error("aaaa")
}

func TestNoneHans(t *testing.T) {
	s := "abc"
	v := Pinyin(s, NewArgs())
	value := [][]string{}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestLazyPinyin(t *testing.T) {
	s := "中国人"
	v := LazyPinyin(s, Args{})
	value := []string{"zhong", "guo", "ren"}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}

	s = "中国人abc"
	v = LazyPinyin(s, Args{})
	value = []string{"zhong", "guo", "ren"}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestSlug(t *testing.T) {
	s := "中国人"
	v := Slug(s, Args{})
	value := "zhongguoren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	v = Slug(s, Args{Separator: ","})
	value = "zhong,guo,ren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	a := NewArgs()
	v = Slug(s, a)
	value = "zhong-guo-ren"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}

	s = "中国人abc，,中"
	v = Slug(s, a)
	value = "zhong-guo-ren-zhong"
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestFinal(t *testing.T) {
	value := "an"
	v := final("an")
	if v != value {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestFallback(t *testing.T) {
	hans := "中国人abc"
	testData := []testCase{
		// default
		{
			NewArgs(),
			[][]string{
				{"zhong"},
				{"guo"},
				{"ren"},
			},
		},
		// custom
		{
			Args{
				Fallback: func(r rune, a Args) []string {
					return []string{"la"}
				},
			},
			[][]string{
				{"zhong"},
				{"guo"},
				{"ren"},
				{"la"},
				{"la"},
				{"la"},
			},
		},
		// custom
		{
			Args{
				Heteronym: true,
				Fallback: func(r rune, a Args) []string {
					return []string{"la", "wo"}
				},
			},
			[][]string{
				{"zhong", "zhong"},
				{"guo"},
				{"ren"},
				{"la", "wo"},
				{"la", "wo"},
				{"la", "wo"},
			},
		},
	}
	testPinyin(t, hans, testData, Pinyin)
}

type testItem struct {
	hans   string
	args   Args
	result [][]string
}

func testPinyinUpdate(t *testing.T, d []testItem, f pinyinFunc) {
	for _, tc := range d {
		v := f(tc.hans, tc.args)
		if !reflect.DeepEqual(v, tc.result) {
			t.Errorf("Expected %s, got %s", tc.result, v)
		}
	}
}

func TestUpdated(t *testing.T) {
	testData := []testItem{
		// 误把 yu 放到声母列表了
		{"鱼", Args{Style: Tone2}, [][]string{{"yu2"}}},
		{"鱼", Args{Style: Tone3}, [][]string{{"yu2"}}},
		{"鱼", Args{Style: Finals}, [][]string{{"v"}}},
		{"雨", Args{Style: Tone2}, [][]string{{"yu3"}}},
		{"雨", Args{Style: Tone3}, [][]string{{"yu3"}}},
		{"雨", Args{Style: Finals}, [][]string{{"v"}}},
		{"元", Args{Style: Tone2}, [][]string{{"yua2n"}}},
		{"元", Args{Style: Tone3}, [][]string{{"yuan2"}}},
		{"元", Args{Style: Finals}, [][]string{{"van"}}},
		// y, w 也不是拼音, yu的韵母是v, yi的韵母是i, wu的韵母是u
		{"呀", Args{Style: Initials}, [][]string{{""}}},
		{"呀", Args{Style: Tone2}, [][]string{{"ya"}}},
		{"呀", Args{Style: Tone3}, [][]string{{"ya"}}},
		{"呀", Args{Style: Finals}, [][]string{{"ia"}}},
		{"无", Args{Style: Initials}, [][]string{{""}}},
		{"无", Args{Style: Tone2}, [][]string{{"wu2"}}},
		{"无", Args{Style: Tone3}, [][]string{{"wu2"}}},
		{"无", Args{Style: Finals}, [][]string{{"u"}}},
		{"衣", Args{Style: Tone2}, [][]string{{"yi1"}}},
		{"衣", Args{Style: Tone3}, [][]string{{"yi1"}}},
		{"衣", Args{Style: Finals}, [][]string{{"i"}}},
		{"万", Args{Style: Tone2}, [][]string{{"wa4n"}}},
		{"万", Args{Style: Tone3}, [][]string{{"wan4"}}},
		{"万", Args{Style: Finals}, [][]string{{"uan"}}},
		// ju, qu, xu 的韵母应该是 v
		{"具", Args{Style: FinalsTone}, [][]string{{"ǜ"}}},
		{"具", Args{Style: FinalsTone2}, [][]string{{"v4"}}},
		{"具", Args{Style: FinalsTone3}, [][]string{{"v4"}}},
		{"具", Args{Style: Finals}, [][]string{{"v"}}},
		{"取", Args{Style: FinalsTone}, [][]string{{"ǚ"}}},
		{"取", Args{Style: FinalsTone2}, [][]string{{"v3"}}},
		{"取", Args{Style: FinalsTone3}, [][]string{{"v3"}}},
		{"取", Args{Style: Finals}, [][]string{{"v"}}},
		{"徐", Args{Style: FinalsTone}, [][]string{{"ǘ"}}},
		{"徐", Args{Style: FinalsTone2}, [][]string{{"v2"}}},
		{"徐", Args{Style: FinalsTone3}, [][]string{{"v2"}}},
		{"徐", Args{Style: Finals}, [][]string{{"v"}}},
		// # ń
		{"嗯", Args{Style: Normal}, [][]string{{"n"}}},
		{"嗯", Args{Style: Tone}, [][]string{{"ń"}}},
		{"嗯", Args{Style: Tone2}, [][]string{{"n2"}}},
		{"嗯", Args{Style: Tone3}, [][]string{{"n2"}}},
		{"嗯", Args{Style: Initials}, [][]string{{""}}},
		{"嗯", Args{Style: FirstLetter}, [][]string{{"n"}}},
		{"嗯", Args{Style: Finals}, [][]string{{"n"}}},
		{"嗯", Args{Style: FinalsTone}, [][]string{{"ń"}}},
		{"嗯", Args{Style: FinalsTone2}, [][]string{{"n2"}}},
		{"嗯", Args{Style: FinalsTone3}, [][]string{{"n2"}}},
		// # ḿ  \u1e3f  U+1E3F
		{"呣", Args{Style: Normal}, [][]string{{"m"}}},
		{"呣", Args{Style: Tone}, [][]string{{"ḿ"}}},
		{"呣", Args{Style: Tone2}, [][]string{{"m2"}}},
		{"呣", Args{Style: Tone3}, [][]string{{"m2"}}},
		{"呣", Args{Style: Initials}, [][]string{{""}}},
		{"呣", Args{Style: FirstLetter}, [][]string{{"m"}}},
		{"呣", Args{Style: Finals}, [][]string{{"m"}}},
		{"呣", Args{Style: FinalsTone}, [][]string{{"ḿ"}}},
		{"呣", Args{Style: FinalsTone2}, [][]string{{"m2"}}},
		{"呣", Args{Style: FinalsTone3}, [][]string{{"m2"}}},
		// 去除 0
		{"啊", Args{Style: Tone2}, [][]string{{"a"}}},
		{"啊", Args{Style: Tone3}, [][]string{{"a"}}},
		{"侵略", Args{Style: Tone2}, [][]string{{"qi1n"}, {"lve4"}}},
		{"侵略", Args{Style: FinalsTone2}, [][]string{{"i1n"}, {"ve4"}}},
		{"侵略", Args{Style: FinalsTone3}, [][]string{{"in1"}, {"ve4"}}},
	}
	testPinyinUpdate(t, testData, Pinyin)
}

func TestConvert(t *testing.T) {
	s := "中国人"
	v := Convert(s, nil)
	value := [][]string{{"zhong"}, {"guo"}, {"ren"}}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}

	a := NewArgs()
	v = Convert(s, &a)
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestLazyConvert(t *testing.T) {
	s := "中国人"
	v := LazyConvert(s, nil)
	value := []string{"zhong", "guo", "ren"}
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}

	a := NewArgs()
	v = LazyConvert(s, &a)
	if !reflect.DeepEqual(v, value) {
		t.Errorf("Expected %s, got %s", value, v)
	}
}

func TestPinyin_fallback_issue_35(t *testing.T) {
	a := NewArgs()
	a.Separator = ""
	a.Style = FirstLetter
	a.Fallback = func(r rune, a Args) []string {
		return []string{string(r)}
	}
	var s = "重。,a庆"
	v := Pinyin(s, a)
	expect := [][]string{{"z"}, {"。"}, {","}, {"a"}, {"q"}}
	if !reflect.DeepEqual(v, expect) {
		t.Errorf("Expected %s, got %s", expect, v)
	}
}
