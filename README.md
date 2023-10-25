
![go-trie-logo](https://github.com/BowlOfSoup/go-trie/blob/main/go-trie-logo.png?raw=true)

### Introduction
A prefix or digital trie (pronounce as 'try') implementation for Go. A trie can be used to quickly construct associative arrays where the keys represent the characters of strings.

Every character (or unicode code point) for a string will its own key; if you input multiple key => values in
the trie, the key will be split by character, and the value for that key will be stored for every character in the key.

**Now why would I use this?** 
* Say you want to match a _big dataset_ on similar strings. 
  * You have for example products that have a product name that ends on a certain revision (like `GolangPlushRev1` and 
  `GolangPlushBlue`). 
  * There is no other way to match these products (like on type or another common property).
  * The products have properties that you need (e.g. color codes) and you want to get a list of that
* You want to build your own dictionary or spelling check
* Substring searches without processing the string every time (you do it once)
* Fast string lookups in general
* You don't want to use it but just stumbled on this repository. Thanks for taking a look.

How does the trie look like?
```
root
 ↓
 G → o → l → a → n → g → P → l → u → s → h
                               ↓
                               R → e → v → 1 [blue, yellow, pink]
                               ↓
                               B → l → u → e [blue]
                               ↓
                               N → o → s → e [black]
                                      ↓
                                      Y → e → l → l → o → w [yellow, pink, blue, green]
```
Important! For every character, all outcomes (values, like `yellow`) is saved. So for the first characters `G` and `o` 
all values (`blue, yellow, pink, blue, black, yellow, pink, blue, green`) are saved.

### Usage
#### Go get
```bash
go get github.com/BowlOfSoup/go-trie
```

#### Examples
```go
package MyPackage

func main() {
	// Create new trie
	trie := NewTrie()
	
	// Insert key => value in trie 
	trie.Insert(key, value)
	
	// Lookup in trie 
	values := trie.Lookup("some key")
	
	// Unique lookup 
	values := trie.Unique("some key")
}
```