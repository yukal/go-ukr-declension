package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/franela/goblin"
)

// The possibility of conducting tests in a specific group:
// For more details see testing_helper.go
//
// go test -v -count=1 ./...
// go test -v -count=1 -run Basic ./...
// go test -v -count=1 -run BasicFeminine ./...

func TestBasicFeminine(t *testing.T) {
	g := goblin.Goblin(t)

	getTitleOf := func(cases []string, caseNum int) string {
		return fmt.Sprintf("%s: %-10s => %s",
			getLastLetters(string(cases[C_NOMINATIVE]), 2),
			string(cases[C_NOMINATIVE]),
			string(cases[caseNum]),
		)
	}

	g.Describe("Basic: Feminine", func() {
		tdtItems := [][]string{
			{"", "Анастасія", "Анастасії", "Анастасії", "Анастасію", "Анастасією", "Анастасії", "Анастасіє"},
			{"", "Ія", "Ії", "Ії", "Ію", "Ією", "Ії", "Іє"},
			{"", "Рая", "Раї", "Раї", "Раю", "Раєю", "Раї", "Рає"},
			{"", "Геля", "Гелі", "Гелі", "Гелю", "Гелею", "Гелі", "Гелю"},
			{"", "Леся", "Лесі", "Лесі", "Лесю", "Лесею", "Лесі", "Лесю"},
			{"", "Фрея", "Фреї", "Фреї", "Фрею", "Фреєю", "Фреї", "Фреє"},
			{"", "Проня", "Проні", "Проні", "Проню", "Пронею", "Проні", "Проню"},
			{"", "Арʼя", "Арʼї", "Арʼї", "Арʼю", "Арʼєю", "Арʼї", "Арʼє"},
			{"", "Гандзя", "Гандзі", "Гандзі", "Гандзю", "Гандзею", "Гандзі", "Гандзю"},
			{"", "Катря", "Катрі", "Катрі", "Катрю", "Катрею", "Катрі", "Катрю"},
			{"", "Вівдя", "Вівді", "Вівді", "Вівдю", "Вівдею", "Вівді", "Вівдю"},
			{"", "Христя", "Христі", "Христі", "Христю", "Христею", "Христі", "Христю"},
			{"", "Любця", "Любці", "Любці", "Любцю", "Любцею", "Любці", "Любцю"},
			{"", "Зоя", "Зої", "Зої", "Зою", "Зоєю", "Зої", "Зоє"},
			{"", "Майя", "Майї", "Майї", "Майю", "Майєю", "Майї", "Майє"},
			{"", "Наталья", "Натальї", "Натальї", "Наталью", "Натальєю", "Натальї", "Натальє"},
		}

		g.Describe("Inflect names with dash", func() {
			cases := []string{
				"", "Анна-Марія", "Анни-Марії", "Анні-Марії", "Анну-Марію", "Анною-Марією", "Анні-Марії", "Анно-Маріє",
			}

			for caseNum := 2; caseNum < 8; caseNum++ {
				name := cases[caseNum]

				g.It(fmt.Sprintf("[%d] %-16s => %s", caseNum, cases[C_NOMINATIVE], name), func() {
					lowNameParts := strings.Split(strings.ToLower(cases[C_NOMINATIVE]), "-")
					result := femNounPartsToCaseN(lowNameParts, uint8(caseNum))
					g.Assert(result).Equal(name)
				})
			}
		})

		g.Describe("2. Nominative to Genitive", func() {
			for _, cases := range tdtItems {
				g.It(getTitleOf(cases, C_GENITIVE), func() {
					res := femNounToCaseN(strings.ToLower(cases[C_NOMINATIVE]), C_GENITIVE)
					g.Assert(res).Equal(cases[C_GENITIVE])
				})
			}
		})

		g.Describe("3. Nominative to Dative", func() {
			for _, cases := range tdtItems {
				g.It(getTitleOf(cases, C_DATIVE), func() {
					res := femNounToCaseN(strings.ToLower(cases[C_NOMINATIVE]), C_DATIVE)
					g.Assert(res).Equal(cases[C_DATIVE])
				})
			}
		})

		g.Describe("4. Nominative to Accusative", func() {
			for _, cases := range tdtItems {
				g.It(getTitleOf(cases, C_ACCUSATIVE), func() {
					res := femNounToCaseN(strings.ToLower(cases[C_NOMINATIVE]), C_ACCUSATIVE)
					g.Assert(res).Equal(cases[C_ACCUSATIVE])
				})
			}
		})

		g.Describe("5. Nominative to Instrumental", func() {
			for _, cases := range tdtItems {
				g.It(getTitleOf(cases, C_INSTRUMENTAL), func() {
					res := femNounToCaseN(strings.ToLower(cases[C_NOMINATIVE]), C_INSTRUMENTAL)
					g.Assert(res).Equal(cases[C_INSTRUMENTAL])
				})
			}
		})

		g.Describe("6. Nominative to Locative", func() {
			for _, cases := range tdtItems {
				g.It(getTitleOf(cases, C_LOCATIVE), func() {
					res := femNounToCaseN(strings.ToLower(cases[C_NOMINATIVE]), C_LOCATIVE)
					g.Assert(res).Equal(cases[C_LOCATIVE])
				})
			}
		})

		g.Describe("7. Nominative to Vocative", func() {
			for _, cases := range tdtItems {
				g.It(getTitleOf(cases, C_VOCATIVE), func() {
					res := femNounToCaseN(strings.ToLower(cases[C_NOMINATIVE]), C_VOCATIVE)
					g.Assert(res).Equal(cases[C_VOCATIVE])
				})
			}
		})
	})
}
