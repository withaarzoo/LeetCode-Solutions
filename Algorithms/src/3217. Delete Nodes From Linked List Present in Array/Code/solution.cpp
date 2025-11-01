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
#include <unordered_set>
#include <vector>
using namespace std;

class Solution
{
public:
    ListNode *modifiedList(vector<int> &nums, ListNode *head)
    {
        // Put nums into a hash set for O(1) lookups
        unordered_set<int> toDelete(nums.begin(), nums.end());

        // Dummy node so deletion of head is easy
        ListNode dummy(0, head);
        ListNode *prev = &dummy;
        ListNode *curr = head;

        while (curr)
        {
            if (toDelete.find(curr->val) != toDelete.end())
            {
                // Remove curr by linking prev->next to curr->next
                prev->next = curr->next;
                // don't advance prev
            }
            else
            {
                // keep curr, advance prev
                prev = curr;
            }
            curr = curr->next;
        }
        return dummy.next;
    }
};
