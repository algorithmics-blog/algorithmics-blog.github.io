export const suggestedProducts = (products: string[], searchWord: string): string[][] => {
    products.sort();
    let res: string[][] = [];
    let items = products;

    for (let idx = 0; idx < searchWord.length; idx++) {
        const prefix = searchWord.substring(0, idx + 1);
        items = suggestions(items, prefix);

        if (items.length === 0) {
            break;
        }

        const maxLen = Math.min(3, items.length);
        res.push(items.slice(0, maxLen));
    }

    res = gd(res, searchWord.length - res.length);

    return res;
}

export const suggestions = (products: string[], prefix: string): string[] => {
    const startIdx = findFirstProductWithPrefix(products, prefix);

    if (startIdx === -1) {
        return [];
    }

    const res: string[] = [];

    for (let i = startIdx; i < products.length; i++) {
        if (!products[i].startsWith(prefix)) {
            break;
        }
        res.push(products[i]);
    }

    return res;
}

export const findFirstProductWithPrefix = (products: string[], prefix: string): number => {
    let left = 0;
    let right = products.length - 1;
    let res = -1;

    while (left <= right) {
        const mid = Math.floor((left + right) / 2);

        if (products[mid].startsWith(prefix)) {
            res = mid;
            right = mid - 1; // Ищем самое раннее вхождение
        } else if (products[mid] < prefix) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }

    return res;
}

export const gd = (res: string[][], tailLength: number): string[][] => {
    for (let i = 0; i < tailLength; i++) {
        res.push([]);
    }
    return res;
}
