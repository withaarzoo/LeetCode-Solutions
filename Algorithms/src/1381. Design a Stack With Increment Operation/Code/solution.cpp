class CustomStack
{
public:
    int maxSize;
    vector<int> stack;
    vector<int> inc;

    CustomStack(int maxSize)
    {
        this->maxSize = maxSize;
    }

    void push(int x)
    {
        if (stack.size() < maxSize)
        {
            stack.push_back(x);
            inc.push_back(0); // Initialize increment for this element
        }
    }

    int pop()
    {
        if (stack.empty())
        {
            return -1;
        }
        int idx = stack.size() - 1;
        int result = stack[idx] + inc[idx]; // Apply any pending increments
        if (idx > 0)
        {
            inc[idx - 1] += inc[idx]; // Propagate increment to the next element
        }
        stack.pop_back();
        inc.pop_back();
        return result;
    }

    void increment(int k, int val)
    {
        int limit = min(k, (int)stack.size()) - 1;
        if (limit >= 0)
        {
            inc[limit] += val; // Add increment to the bottom k-th element
        }
    }
};