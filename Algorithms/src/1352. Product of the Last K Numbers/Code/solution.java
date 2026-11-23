class ProductOfNumbers {
    List<Integer> pre;
    public ProductOfNumbers() {
        pre = new ArrayList<>();
        pre.add(1);
    }
    
    public void add(int num) {
        if (num == 0) {
            pre = new ArrayList<>();
            pre.add(1);
        } else {
            pre.add(pre.get(pre.size() - 1) * num);
        }
    }
    
    public int getProduct(int k) {
        if (k >= pre.size()) return 0;
        return pre.get(pre.size() - 1) / pre.get(pre.size() - 1 - k);
    }
}