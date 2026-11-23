class ProductOfNumbers {
public:
    vector<int> pre;
    ProductOfNumbers() {
        pre = {1};
    }
    
    void add(int num) {
        if (num == 0) {
            pre = {1};
        } else {
            pre.push_back(pre.back() * num);
        }
    }
    
    int getProduct(int k) {
        if (k >= pre.size()) return 0;
        return pre.back() / pre[pre.size() - 1 - k];
    }
};