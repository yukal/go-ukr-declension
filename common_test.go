package main

import (
	"fmt"
	"testing"

	"github.com/franela/goblin"
)

// The possibility of conducting tests in a specific group:
// For more details see testing_helper.go
//
// go test -v -count=1 ./...
// go test -v -count=1 -run Common ./...
// go test -v -count=1 -run Basic ./...
// go test -v -count=1 -run General ./...

func TestBasicGeneralCommon(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Stemming", func() {
		type TDTAct = func(word []rune) []rune

		type TDT struct {
			Action TDTAct
			Data   []rune
			Expect []rune
		}

		tdtItems := []TDT{
			{
				Action: GetStem,
				Data:   []rune("Ярослава"),
				Expect: []rune("Ярослав"),
			},
			{
				Action: GetStem,
				Data:   []rune("Вернидуб"),
				Expect: []rune("Вернидуб"),
			},
			{
				Action: GetStem,
				Data:   []rune("Любов"),
				Expect: []rune("Любов"),
			},
			{
				Action: GetStem,
				Data:   []rune("Олег"),
				Expect: []rune("Олег"),
			},
			{
				Action: GetStem,
				Data:   []rune("Теґ"),
				Expect: []rune("Теґ"),
			},
			{
				Action: GetStem,
				Data:   []rune("Свирид"),
				Expect: []rune("Свирид"),
			},
			{
				Action: GetStem,
				Data:   []rune("Гонгадзе"),
				Expect: []rune("Гонгадз"),
			},
			{
				Action: GetStem,
				Data:   []rune("Фойє"),
				Expect: []rune("Фой"),
			},
			{
				Action: GetStem,
				Data:   []rune("Багаж"),
				Expect: []rune("Багаж"),
			},
			{
				Action: GetStem,
				Data:   []rune("Гарбуз"),
				Expect: []rune("Гарбуз"),
			},
			{
				Action: GetStem,
				Data:   []rune("Мальдіви"),
				Expect: []rune("Мальдів"),
			},
			{
				Action: GetStem,
				Data:   []rune("Гіві"),
				Expect: []rune("Гів"),
			},
			{
				Action: GetStem,
				Data:   []rune("Дубаї"),
				Expect: []rune("Дуб"),
			},
			{
				Action: GetStem,
				Data:   []rune("Дарій"),
				Expect: []rune("Дарій"),
			},
			{
				Action: GetStem,
				Data:   []rune("Щек"),
				Expect: []rune("Щек"),
			},
			{
				Action: GetStem,
				Data:   []rune("Рафаїл"),
				Expect: []rune("Рафаїл"),
			},
			{
				Action: GetStem,
				Data:   []rune("Адам"),
				Expect: []rune("Адам"),
			},
			{
				Action: GetStem,
				Data:   []rune("Бажан"),
				Expect: []rune("Бажан"),
			},
			{
				Action: GetStem,
				Data:   []rune("Ніно"),
				Expect: []rune("Нін"),
			},
			{
				Action: GetStem,
				Data:   []rune("Остап"),
				Expect: []rune("Остап"),
			},
			{
				Action: GetStem,
				Data:   []rune("Володимир"),
				Expect: []rune("Володимир"),
			},
			{
				Action: GetStem,
				Data:   []rune("Тарас"),
				Expect: []rune("Тарас"),
			},
			{
				Action: GetStem,
				Data:   []rune("Орест"),
				Expect: []rune("Орест"),
			},
			{
				Action: GetStem,
				Data:   []rune("Табу"),
				Expect: []rune("Таб"),
			},
			{
				Action: GetStem,
				Data:   []rune("Джозеф"),
				Expect: []rune("Джозеф"),
			},
			{
				Action: GetStem,
				Data:   []rune("Тімох"),
				Expect: []rune("Тімох"),
			},
			{
				Action: GetStem,
				Data:   []rune("Палац"),
				Expect: []rune("Палац"),
			},
			{
				Action: GetStem,
				Data:   []rune("Ткач"),
				Expect: []rune("Ткач"),
			},
			{
				Action: GetStem,
				Data:   []rune("Ярош"),
				Expect: []rune("Ярош"),
			},
			{
				Action: GetStem,
				Data:   []rune("Борщ"),
				Expect: []rune("Борщ"),
			},
			{
				Action: GetStem,
				Data:   []rune("Василь"),
				Expect: []rune("Васил"),
			},
			{
				Action: GetStem,
				Data:   []rune("Ендрю"),
				Expect: []rune("Ендр"),
			},
			{
				Action: GetStem,
				Data:   []rune("Юлія"),
				Expect: []rune("Юл"),
			},
		}

		for _, tdt := range tdtItems {
			title := fmt.Sprintf("%s: %-10s => %s",
				getLastLetter(string(tdt.Data)),
				string(tdt.Data),
				string(tdt.Expect),
			)

			g.It(title, func() {
				hint := fmt.Sprintf("Expect( %s ) Got( %s )", string(tdt.Expect), string(tdt.Data))
				g.Assert(tdt.Action(tdt.Data)).Equal(tdt.Expect, hint)
			})
		}
	})
}
