package search_suggestions_system

import (
	"sort"
)

func suggestedProductsTrie(products []string, searchWord string) [][]string {
	root := &trie{
		children: make(map[rune]*trie),
	}

	sort.Strings(products)
	for _, p := range products {
		root.Insert(p)
	}

	runeSearch := []rune(searchWord)
	worldStart, exist := root.children[runeSearch[0]]
	if !exist {
		return root.gd(runeSearch)
	}

	suggestions := worldStart.Closest(runeSearch[1:], 3)
	res := make([][]string, len(suggestions))
	prev := []string{}
	i := len(suggestions) - 1
	if i > 0 {
		prev = suggestions[i][:]
		res[i] = prev
		i--
	}

	for ; i >= 0; i-- {
		newItem := suggestions[i][:]
		if len(prev) == 0 {
			res[i] = newItem
			prev = newItem[:]
			continue
		} else if len(newItem) == 0 {
			res[i] = prev[:]
			continue
		}

		newItem = append(newItem, prev...)
		sort.Strings(newItem)
		if len(newItem) > 3 {
			newItem = newItem[:3]
		}

		res[i] = newItem
		prev = newItem[:]
	}

	return res
}

type trie struct {
	world    string
	sort     []rune
	children map[rune]*trie
}

func (t *trie) Insert(word string) {
	t.insert(word, []rune(word))
}

func (t *trie) insert(origin string, word []rune) {
	if len(word) == 0 {
		t.world = origin
		return
	}

	child, exist := t.children[word[0]]
	if !exist {
		child = &trie{
			children: make(map[rune]*trie),
			sort:     make([]rune, 0, 5),
		}

		t.children[word[0]] = child
		t.sort = append(t.sort, word[0])
	}

	child.insert(origin, word[1:])
}

func (t *trie) Closest(prefix []rune, count int) [][]string {
	res := make([][]string, 0, len(prefix))
	var excludeRune *rune
	if len(prefix) > 0 {
		excludeRune = &prefix[0]
	}

	res = append(res, t.exactWorlds(excludeRune, count))

	if len(prefix) == 0 {
		return res
	}

	child, exist := t.children[prefix[0]]
	if !exist {
		res = [][]string{t.exactWorlds(nil, count)}
		return append(res, t.gd(prefix)...)
	}

	return append(res, child.Closest(prefix[1:], count)...)
}

func (t *trie) exactWorlds(excludeRune *rune, n int) []string {
	res := make([]string, 0, n)
	if len(t.world) > 0 {
		res = append(res, t.world)
		n--
	}

	if n == 0 {
		return res
	}

	for _, childRune := range t.sort {
		if excludeRune != nil && childRune == *excludeRune {
			continue
		}
		childWorlds := t.children[childRune].exactWorlds(nil, n)
		if len(childWorlds) >= n {
			return append(res, childWorlds[:n]...)
		}

		res = append(res, childWorlds...)
		n = n - len(childWorlds)
		if n == 0 {
			break
		}
	}

	return res
}

func (t *trie) gd(prefix []rune) [][]string {
	res := make([][]string, 0, len(prefix))
	for range prefix {
		res = append(res, []string{})
	}

	return res
}
