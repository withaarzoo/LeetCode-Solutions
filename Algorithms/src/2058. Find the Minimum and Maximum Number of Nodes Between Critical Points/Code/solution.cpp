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
    vector<int> nodesBetweenCriticalPoints(ListNode *head)
    {
        // The first critical point position; -1 means we have not found one yet.
        int first = -1;

        // The most recently found critical point position.
        int last = -1;

        // The minimum distance between two consecutive critical points.
        int minDistance = INT_MAX;

        // Position of the current node in the linked list.
        int position = 1;

        // Start from the second node because a critical point needs a previous node.
        ListNode *prev = head;

        // The current node starts from the second node.
        ListNode *curr = head->next;

        // We need curr->next to exist, so the last node is not checked.
        while (curr != nullptr && curr->next != nullptr)
        {
            // Check whether curr is a local maximum or a local minimum.
            bool isCritical =
                (curr->val > prev->val && curr->val > curr->next->val) ||
                (curr->val < prev->val && curr->val < curr->next->val);

            // Process the node only if it is a critical point.
            if (isCritical)
            {
                // This is the first critical point we have found.
                if (first == -1)
                {
                    first = position;
                }
                else
                {
                    // Compare this critical point with the previous critical point.
                    minDistance = min(minDistance, position - last);
                }

                // Store the current critical point as the latest one.
                last = position;
            }

            // Move the previous pointer forward for the next iteration.
            prev = curr;

            // Move the current pointer forward for the next iteration.
            curr = curr->next;

            // Move to the next position in the linked list.
            position++;
        }

        // Fewer than two critical points means no valid distance exists.
        if (first == -1 || first == last)
        {
            return {-1, -1};
        }

        // The distance between the first and last critical points is the maximum.
        int maxDistance = last - first;

        // Return the minimum and maximum distances.
        return {minDistance, maxDistance};
    }
};