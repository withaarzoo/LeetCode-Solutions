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
class Solution {
public:
    ListNode* deleteMiddle(ListNode* head) {
        // If there is only one node, deleting the middle leaves an empty list
        if (head->next == nullptr) {
            return nullptr;
        }

        // Slow finds middle, fast moves twice as fast
        ListNode* slow = head;
        ListNode* fast = head;

        // Keeps track of node before slow
        ListNode* prev = nullptr;

        while (fast != nullptr && fast->next != nullptr) {
            prev = slow;           // Store previous node
            slow = slow->next;     // Move slow by 1 step
            fast = fast->next->next; // Move fast by 2 steps
        }

        // Skip the middle node
        prev->next = slow->next;

        return head;
    }
};