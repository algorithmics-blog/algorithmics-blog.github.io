import { isSubsequence } from "./solution";

type Suit = {
    name: string
    s: string
    t: string
    out: boolean
}

const suits: Suit[] = [
    {
        name: "s = \"abc\", t = \"ahbgdc\"",
        s:    "abc",
        t:    "ahbgdc",
        out:  true,
    },
    {
        name: "s = \"axc\", t = \"ahbgdc\"",
        s:    "axc",
        t:    "ahbgdc",
        out:  false,
    },
    {
        name: "s = \"\", t = \"ahbgdc\"",
        s:    "",
        t:    "ahbgdc",
        out:  true,
    },
    {
        name: "s = \"n\", t = \"c\"",
        s:    "b",
        t:    "c",
        out:  false,
    },
    {
        name: "s = \"b\", t = \"abc\"",
        s:    "b",
        t:    "abc",
        out:  true,
    },
    {
        name: "s = \"abc\", t = \"b\"",
        s:    "abc",
        t:    "b",
        out:  false,
    },
]

describe('decodeString', () => {
    suits.forEach(suit => {
        test(suit.name, () => {
            const res = isSubsequence(suit.s, suit.t)
            expect(res).toEqual(suit.out);
        })
    })
})
