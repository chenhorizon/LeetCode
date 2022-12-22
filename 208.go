// leetcode 208

import "strings"

type Trie struct {
    Val     string
    Left    *Trie
    Right   *Trie
}


func Constructor() Trie {
    return Trie{"", nil, nil}
}

func (this *Trie) Insert(word string)  {
    if this.Val == "" && this.Left == nil && this.Right == nil {
        this.Val = word
        return
    }

    p := this
    for p != nil {
        if word == p.Val {
            return
        }

        for word < p.Val {
            if p.Left == nil {
                p.Left = &Trie{word, nil, nil}
                return
            }
            p = p.Left
        }
        if p.Right == nil && word > p.Val {
            p.Right = &Trie{word, nil, nil}
            return
        }

        for word > p.Val {
            if p.Right == nil {
                p.Right = &Trie{word, nil, nil}
                return
            }
            p = p.Right
        }
        if p.Left == nil && word < p.Val {
            p.Left = &Trie{word, nil, nil}
            return
        }
    }

}


func (this *Trie) Search(word string) bool {
    if this.Val == word {
        return true
    }
    if this.Val > word {
        if this.Left == nil {
            return false
        }
        return this.Left.Search(word)
    } else if this.Val < word {
        if this.Right == nil {
            return false
        }
        return this.Right.Search(word)
    }

    return false
}


func (this *Trie) StartsWith(prefix string) bool {
    word := prefix
    if strings.HasPrefix(this.Val, word) {
        return true
    }

    if this.Val > word {
        if this.Left == nil {
            return false
        }
        return this.Left.StartsWith(word)
    } else if this.Val < word {
        if this.Right == nil {
            return false
        }
        return this.Right.StartsWith(word)
    }

    return false
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
