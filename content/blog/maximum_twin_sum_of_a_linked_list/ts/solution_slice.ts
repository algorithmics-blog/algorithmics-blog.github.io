import { ListNode } from './ListNode';

export function pairSumSlice(head: ListNode): number {
    let list: number[] = [head.val];

    // Проходим по списку и заполняем массив значениями
    for (let cur = head.next; cur !== null; cur = cur.next) {
        list.push(cur.val);
    }

    let res = 0;

    // Находим максимальную пару значений
    for (let i = 0; i < Math.floor(list.length / 2); i++) {
        let j = list.length - i - 1;
        res = Math.max(res, list[i] + list[j]);
    }

    return res;
}