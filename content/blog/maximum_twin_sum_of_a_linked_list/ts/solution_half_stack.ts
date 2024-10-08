import { ListNode } from './ListNode';

export function pairSumStack(head: ListNode): number {
    let stack: number[] = [];
    let currentListElement: ListNode | null = head;
    let oddElementList: ListNode | null = head;

    // Заполняем стек первой половиной списка
    while (oddElementList !== null && oddElementList.next !== null) {
        stack.push(currentListElement!.val);
        currentListElement = currentListElement!.next;
        oddElementList = oddElementList.next.next;
    }

    let res = 0;

    // Сравниваем элементы с конца первой половины и из второй половины
    while (currentListElement !== null) {
        res = Math.max(res, currentListElement.val + stack.pop()!);
        currentListElement = currentListElement.next;
    }

    return res;
}