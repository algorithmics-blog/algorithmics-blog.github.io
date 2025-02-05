class DisjointSet {
    readonly nodes: number[];

    constructor(size: number) {
        this.nodes = new Array(size);

        for (let i = 0; i < size; i++) {
            this.nodes[i] = i;
        }
    }

    find(node: number): number {
        return this.nodes[node];
    }

    union(x: number, y: number): void {
        const rootX = this.find(x);
        const rootY = this.find(y);

        if (rootX !== rootY) {
            for (let i = 0; i < this.nodes.length; i++) {
                if (this.nodes[i] === rootY) {
                    this.nodes[i] = rootX;
                }
            }
        }
    }
}

export const findCircleNum = (isConnected: number[][]): number => {
    const n = isConnected.length;
    const disjointSet = new DisjointSet(n);

    let numberOfProvinces = n;

    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            if (isConnected[i][j] == 1 && disjointSet.find(i) != disjointSet.find(j)) {
                numberOfProvinces -= 1;
                disjointSet.union(i, j);
            }
        }
    }

    return numberOfProvinces
}
