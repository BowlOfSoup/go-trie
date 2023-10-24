package trie

type Node struct {
	children map[rune]*Node
	values   []string
	isEnd    bool
}

type Trie struct {
	root *Node
}

func NewNode() *Node {
	return &Node{children: make(map[rune]*Node)}
}

func NewTrie() *Trie {
	return &Trie{root: NewNode()}
}

func (t *Trie) Insert(key string, value string) {
	if key == "" {
		return
	}

	currentNode := t.root
	for _, char := range key {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = &Node{
				children: make(map[rune]*Node),
			}
		}
		currentNode = currentNode.children[char]
		currentNode.values = append(currentNode.values, value)
	}
}

func (t *Trie) Lookup(prefix string) []string {
	currentNode := t.root
	for _, char := range prefix {
		if currentNode.children[char] == nil {
			return []string{} // not found
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.values
}

func (t *Trie) LookupUnique(prefix string) []string {
	lookedUpValues := t.Lookup(prefix)

	existing := make(map[string]bool)
	var result []string
	for _, v := range lookedUpValues {
		if !existing[v] {
			existing[v] = true
			result = append(result, v)
		}
	}
	return result
}
