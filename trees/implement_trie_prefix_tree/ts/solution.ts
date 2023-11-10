export class Trie {
    children: Map<string, Trie>
    isFullWord: boolean

    constructor(children?: Map<string, Trie>, isFullWord?: boolean) {
        this.children = children ?? new Map()
        this.isFullWord = isFullWord ?? false
    }

    insert(word: string): void {
        debugger
        if (word.length === 0) {
            this.isFullWord = true
            return
        }

        let child: Trie

        if (!this.children.has(word[0])) {
            child = new Trie()
            this.children.set(word[0], child)
        } else {
            child = this.children.get(word[0])
        }

        child.insert(word.slice(1))
    }

    search(word: string): boolean {
        const node = this.traverse(word)
        if (!node) {
            return false
        }

        return node.isFullWord
    }

    startsWith(prefix: string): boolean {
        const node = this.traverse(prefix)

        return node !== null
    }

    private traverse(prefix: string): Trie | null {
        if (prefix.length === 0) {
            return this
        }

        if (!this.children.has(prefix[0])) {
            return null
        }

        const child = this.children.get(prefix[0])

        return child.traverse(prefix.slice(1))
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * var obj = new Trie()
 * obj.insert(word)
 * var param_2 = obj.search(word)
 * var param_3 = obj.startsWith(prefix)
 */
