/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution
{
public:
    ListNode *rotateRight(ListNode *head, int k)
    {
        // edge case: empty list or single node
        if (!head || !head->next || k == 0)
            return head;

        int n = 1; // length of list
        ListNode *tail = head;

        // find length and last node
        while (tail->next)
        {
            tail = tail->next;
            n++;
        }

        // make it circular
        tail->next = head;

        // reduce k
        k = k % n;

        // find new tail (n - k - 1 steps)
        int steps = n - k - 1;
        ListNode *newTail = head;

        while (steps--)
        {
            newTail = newTail->next;
        }

        // new head is next of newTail
        ListNode *newHead = newTail->next;

        // break the circle
        newTail->next = nullptr;

        return newHead;
    }
};