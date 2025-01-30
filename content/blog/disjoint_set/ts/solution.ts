export class DisjointSet {
    readonly nodes: number[];

    constructor(size: number) {
        // Создаем массив заданного размера по количеству элементов
        this.nodes = new Array(size);

        for (let i = 0; i < size; i++) {
            // При инициализации каждый элемент является рутовым сам для себя
            this.nodes[i] = i;
        }
    }

    find(node: number): number {
        return this.nodes[node];
    }

    connected(x: number, y: number): boolean {
        return this.find(x) === this.find(y);
    }

    union(x: number, y: number): void {
        const rootX = this.find(x);
        const rootY = this.find(y);

        // Объединяем множества, если они не совпадают
        if (rootX !== rootY) {
            for (let i = 0; i < this.nodes.length; i++) {
                if (this.nodes[i] === rootY) {
                    this.nodes[i] = rootX;
                }
            }
        }
    }
}
