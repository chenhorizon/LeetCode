// leetcode 211

type WordDictionary struct {
    Val     string
    Left    *WordDictionary
    Right   *WordDictionary
}


func Constructor() WordDictionary {
    return WordDictionary{"", nil, nil}
}


func (this *WordDictionary) AddWord(word string)  {
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
                p.Left = &WordDictionary{word, nil, nil}
                return
            }
            p = p.Left
        }
        if p.Right == nil {
            p.Right = &WordDictionary{word, nil, nil}
            return
        }

        for word > p.Val {
            if p.Right == nil {
                p.Right = &WordDictionary{word, nil, nil}
                return
            }
            p = p.Right
        }
        if p.Left == nil {
            p.Left = &WordDictionary{word, nil, nil}
            return
        }
    }

}


func (this *WordDictionary) Search(word string) bool {
    if wordeq(this.Val, word) {
        return true
    }

    rHas := false
    if this.Right != nil {
        rHas = this.Right.Search(word)
    }

    lHas := false
    if this.Left != nil {
        lHas = this.Left.Search(word)
    }

    return lHas || rHas
}

func wordeq(v string, word string) bool {
    if len(v) != len(word) {
        return false
    }

    for i:=0; i<len(v); i++ {
        if v[i] == word[i] || word[i] == '.' {
            continue
        } else {
            return false
        }
    }

    return true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
