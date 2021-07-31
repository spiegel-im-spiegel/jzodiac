package jzodiac_test

import (
	"testing"
	"time"

	"github.com/spiegel-im-spiegel/jzodiac"
)

func TestKan10(t *testing.T) {
	testCases := []struct {
		kan  jzodiac.Kan10
		name string
	}{
		{kan: jzodiac.Kinoe, name: "甲"},
		{kan: jzodiac.Kinoto, name: "乙"},
		{kan: jzodiac.Hinoe, name: "丙"},
		{kan: jzodiac.Hinoto, name: "丁"},
		{kan: jzodiac.Tsutinoe, name: "戊"},
		{kan: jzodiac.Tsutinoto, name: "己"},
		{kan: jzodiac.Kanoe, name: "庚"},
		{kan: jzodiac.Kanoto, name: "辛"},
		{kan: jzodiac.Mizunoe, name: "壬"},
		{kan: jzodiac.Mizunoto, name: "癸"},
		{kan: jzodiac.Kan10(10), name: "甲"},
	}

	for _, tc := range testCases {
		str := tc.kan.String()
		if str != tc.name {
			t.Errorf("jzodiac.Kan10(%v) is \"%v\", want %v", uint(tc.kan), str, tc.name)
		}
	}
}

func TestShi12(t *testing.T) {
	testCases := []struct {
		shi  jzodiac.Shi12
		name string
	}{
		{shi: jzodiac.Rat, name: "子"},
		{shi: jzodiac.Ox, name: "丑"},
		{shi: jzodiac.Tiger, name: "寅"},
		{shi: jzodiac.Rabbit, name: "卯"},
		{shi: jzodiac.Dragon, name: "辰"},
		{shi: jzodiac.Snake, name: "巳"},
		{shi: jzodiac.Horse, name: "午"},
		{shi: jzodiac.Sheep, name: "未"},
		{shi: jzodiac.Monkey, name: "申"},
		{shi: jzodiac.Rooster, name: "酉"},
		{shi: jzodiac.Dog, name: "戌"},
		{shi: jzodiac.Boar, name: "亥"},
		{shi: jzodiac.Shi12(12), name: "子"},
	}

	for _, tc := range testCases {
		str := tc.shi.String()
		if str != tc.name {
			t.Errorf("jzodiac.Shi12(%v) is \"%v\", want %v", uint(tc.shi), str, tc.name)
		}
	}
}

func TestZodiac(t *testing.T) {
	testCases := []struct {
		t       time.Time
		kanYear jzodiac.Kan10
		shiYear jzodiac.Shi12
		kanDay  jzodiac.Kan10
		shiDay  jzodiac.Shi12
	}{
		{t: time.Date(1983, time.January, 1, 0, 0, 0, 0, jzodiac.JST), kanYear: jzodiac.Mizunoto, shiYear: jzodiac.Boar, kanDay: jzodiac.Tsutinoto, shiDay: jzodiac.Ox},
		{t: time.Date(1984, time.January, 1, 0, 0, 0, 0, jzodiac.JST), kanYear: jzodiac.Kinoe, shiYear: jzodiac.Rat, kanDay: jzodiac.Kinoe, shiDay: jzodiac.Horse},
		{t: time.Date(1985, time.January, 1, 0, 0, 0, 0, jzodiac.JST), kanYear: jzodiac.Kinoto, shiYear: jzodiac.Ox, kanDay: jzodiac.Kanoe, shiDay: jzodiac.Rat},
		{t: time.Date(2000, time.January, 1, 0, 0, 0, 0, jzodiac.JST), kanYear: jzodiac.Kanoe, shiYear: jzodiac.Dragon, kanDay: jzodiac.Tsutinoe, shiDay: jzodiac.Horse},
		{t: time.Date(2001, time.January, 1, 0, 0, 0, 0, jzodiac.JST), kanYear: jzodiac.Kanoto, shiYear: jzodiac.Snake, kanDay: jzodiac.Kinoe, shiDay: jzodiac.Rat},
		{t: time.Date(2002, time.January, 1, 0, 0, 0, 0, jzodiac.JST), kanYear: jzodiac.Mizunoe, shiYear: jzodiac.Horse, kanDay: jzodiac.Tsutinoto, shiDay: jzodiac.Snake},
	}

	for _, tc := range testCases {
		kanYear, shiYear := jzodiac.ZodiacYearNumber(tc.t.Year())
		if kanYear != tc.kanYear {
			t.Errorf("result of ZodiacYearNumber(\"%v\") is \"%v\", want %v", tc.t, kanYear, tc.kanYear)
		}
		if shiYear != tc.shiYear {
			t.Errorf("result of ZodiacYearNumber(\"%v\") is \"%v\", want %v", tc.t, shiYear, tc.shiYear)
		}
		kanDay, shiDay := jzodiac.ZodiacDayNumber(tc.t)
		if kanDay != tc.kanDay {
			t.Errorf("result of ZodiacDayNumber(\"%v\") is \"%v\", want %v", tc.t, kanDay, tc.kanDay)
		}
		if shiYear != tc.shiYear {
			t.Errorf("result of ZodiacDayNumber(\"%v\") is \"%v\", want %v", tc.t, shiDay, tc.shiDay)
		}
	}
}

/* Copyright 2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
