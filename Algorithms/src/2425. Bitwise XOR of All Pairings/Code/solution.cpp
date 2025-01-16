class Solution
{
public:
    int xorAllNums(vector<int> &nums1, vector<int> &nums2)
    {
        int xor1 = 0, xor2 = 0;

        // XOR all elements in nums1
        for (int num : nums1)
        {
            xor1 ^= num;
        }

        // XOR all elements in nums2
        for (int num : nums2)
        {
            xor2 ^= num;
        }

        // If nums1 has odd length, include xor2
        // If nums2 has odd length, include xor1
        return (nums1.size() % 2 ? xor2 : 0) ^ (nums2.size() % 2 ? xor1 : 0);
    }
};
