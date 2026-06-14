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
    int pairSum(ListNode *head)
    {
        // Find the middle of the linked list
        ListNode *slow = head;
        ListNode *fast = head;

        while (fast && fast->next)
        {
            slow = slow->next;
            fast = fast->next->next;
        }

        // Reverse the second half
        ListNode *prev = nullptr;
        while (slow)
        {
            ListNode *nextNode = slow->next;
            slow->next = prev;
            prev = slow;
            slow = nextNode;
        }

        // Compare first half and reversed second half
        int ans = 0;
        ListNode *first = head;
        ListNode *second = prev;

        while (second)
        {
            ans = max(ans, first->val + second->val);

            first = first->next;
            second = second->next;
        }

        return ans;
    }
};