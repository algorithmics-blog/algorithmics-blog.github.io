import { ListNode } from './ListNode';

export function pairSum(head: ListNode): number {
    let firstHalfList: ListNode | null = head;
    let oddElementList: ListNode | null = head;

    // Ищем середину списка
    while (oddElementList !== null && oddElementList.next !== null) {
        firstHalfList = firstHalfList!.next;
        oddElementList = oddElementList.next.next;
    }

    // Разворачиваем вторую половину списка
    let revers: ListNode | null = reverseList(firstHalfList);

    let maxSum = 0;
    // Сравниваем элементы
    while (revers !== null) {
        maxSum = Math.max(maxSum, head.val + revers.val);
        head = head.next!;
        revers = revers.next;
    }

    return maxSum;
}

function reverseList(head: ListNode | null): ListNode | null {
    let newHead: ListNode | null = null;

    while (head !== null) {
        let next: ListNode | null = head.next;
        head.next = newHead;
        newHead = head;
        head = next;
    }

    return newHead;
}