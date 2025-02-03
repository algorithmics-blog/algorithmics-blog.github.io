import { DisjointSet } from "./solution";

describe('DisjointSet', () => {
    const set = new DisjointSet(8)

    set.union(0, 2)
    set.union(0, 1)
    set.union(2, 3)
    set.union(1, 3)

    set.union(4, 5)
    set.union(4, 6)
    set.union(5, 6)

    test("Create disjoint set", () => {
        expect(set.nodes).toEqual([0, 0, 0, 0, 4, 4, 4, 7]);
    })

    test("Find direct connection", () => {
        expect(set.find(3)).toEqual(0);
    })

    test("Find transitive connection", () => {
        expect(set.find(5)).toEqual(4);
    })

    test("Find in set with one element", () => {
        expect(set.find(7)).toEqual(7);
    })

    test("Check direct connection", () => {
        expect(set.connected(4, 5)).toEqual(true);
    })

    test("Check transitive connection", () => {
        expect(set.connected(0, 3)).toEqual(true);
    })

    test("Check root connection", () => {
        expect(set.connected(7, 7)).toEqual(true);
    })

    test("Check not connected elements", () => {
        expect(set.connected(0, 7)).toEqual(false);
    })

    test("Check union", () => {
        set.union(3, 4)
        expect(set.nodes).toEqual([0, 0, 0, 0, 0, 0, 0, 7]);
        expect(set.connected(0, 6)).toEqual(true);
    })
})
