package jzodiac_test

import (
	"fmt"
	"time"

	"github.com/spiegel-im-spiegel/jzodiac"
)

func ExampleZodiacDayNumber() {
	t := time.Date(2021, time.July, 28, 0, 0, 0, 0, jzodiac.JST)
	kan, shi := jzodiac.ZodiacDayNumber(t)
	fmt.Printf("%v is %v%v\n", t.Format("2006-01-02"), kan, shi)
	// Output:
	// 2021-07-28 is 丁丑
}

func ExampleZodiacYearNumber() {
	kan, shi := jzodiac.ZodiacYearNumber(2021)
	fmt.Printf("Year %v is %v%v\n", 2021, kan, shi)
	// Output:
	// Year 2021 is 辛丑
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
