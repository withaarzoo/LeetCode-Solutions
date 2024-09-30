class CustomStack {
    private int maxSize;
    private List<Integer> stack;
    private List<Integer> inc;

    public CustomStack(int maxSize) {
        this.maxSize = maxSize;
        this.stack = new ArrayList<>();
        this.inc = new ArrayList<>();
    }

    public void push(int x) {
        if (stack.size() < maxSize) {
            stack.add(x);
            inc.add(0); // Initialize increment for this element
        }
    }

    public int pop() {
        if (stack.isEmpty()) {
            return -1;
        }
        int idx = stack.size() - 1;
        int result = stack.get(idx) + inc.get(idx); // Apply any pending increments
        if (idx > 0) {
            inc.set(idx - 1, inc.get(idx - 1) + inc.get(idx)); // Propagate increment to the next element
        }
        stack.remove(idx);
        inc.remove(idx);
        return result;
    }

    public void increment(int k, int val) {
        int limit = Math.min(k, stack.size()) - 1;
        if (limit >= 0) {
            inc.set(limit, inc.get(limit) + val); // Add increment to the bottom k-th element
        }
    }
}