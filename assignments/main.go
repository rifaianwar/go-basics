package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hello world")

	arraySign([]int{2, 1})                    // 1
	arraySign([]int{-2, 1})                   // -1
	arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1

	isAnagram("anak", "kana")       // true
	isAnagram("anak", "mana")       // false
	isAnagram("anagram", "managra") // true

	findTheDifference("abcd", "abcde") // 'e'
	findTheDifference("abcd", "abced") // 'e'
	findTheDifference("", "y")         // 'y'

	canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	canMakeArithmeticProgression([]int{5, 1, 9})    // true; 9, 5, 1 adalah baris aritmatik -4
	canMakeArithmeticProgression([]int{1, 2, 4, 8}) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here
	var result int
	sum := 1
	for i := 0; i < len(nums); i++ {
		sum *= nums[i]
		if sum > 0 {
			result = 1
		}
		if sum < 0 {
			result = -1
		}
		if sum == 0 {
			result = 0
		}
	}
	fmt.Println("SUM:", sum)
	fmt.Println("Result:", result)
	return result
	//return 1 // if positive
	// return -1 // if negative
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here
	var result bool
	runes1 := []rune(s)
	runes2 := []rune(t)
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	if len(s) != len(t) {
		result = false
	}
	if (string(runes1)) == string(runes2) {
		result = true
	}
	fmt.Println(result)
	return result
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	// write code here
	var result byte
	runes2 := []rune(t)

	for i := 0; i < len(t); i++ {
		if strings.Contains(s, string(runes2[i])) {

		} else {
			result = byte(runes2[i])
		}
	}
	fmt.Println(string(result))
	return result
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// write code here
	var temp int
	var result bool
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	swap := arr[1] - arr[0]
	for i := 1; i < len(arr); i++ {
		temp = arr[i] - arr[i-1]

		if temp == swap {
			swap = temp
			if (arr[i-1] + temp) == arr[i] {
				result = true
			} else {
				result = false
			}
		} else {
			result = false
		}
	}
	fmt.Println("Result: ", result)
	return result
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
	d.cards = make([]Card, 0, 52)

	for symbol := 0; symbol < 4; symbol++ {
		for number := 1; number <= 13; number++ {
			d.cards = append(d.cards, Card{symbol: symbol, number: number})
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	if n > len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	if n < len(d.cards) {
		n = len(d.cards)
	}
	return d.cards[:n]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// write code here
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(5)
	fmt.Println("PeekTop 5")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println("PeekCardAtIndex index array 10 - 15")
	fmt.Println(deck.PeekCardAtIndex(10).ToString()) // Jack Spade
	fmt.Println(deck.PeekCardAtIndex(11).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // 2 Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 3 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(5)
	fmt.Println("Deck Shuffle")
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	fmt.Println("Deck Cut")
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
