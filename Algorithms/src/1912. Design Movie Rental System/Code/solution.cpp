#include <bits/stdc++.h>
using namespace std;

class MovieRentingSystem {
private:
    // helper to pack shop & movie into a 64-bit key
    inline long long key(int shop, int movie) const {
        return ((long long)shop << 32) | (unsigned long long)movie;
    }

    unordered_map<long long,int> priceMap; // (shop,movie) -> price
    unordered_map<long long,bool> isRented; // (shop,movie) -> rented status

    // available: movie -> set of (price, shop) sorted by price then shop
    unordered_map<int, set<pair<int,int>>> avail;

    // rented: global set of (price, shop, movie) sorted by price, shop, movie
    set<tuple<int,int,int>> rented; 

public:
    MovieRentingSystem(int n, vector<vector<int>>& entries) {
        for (auto &e : entries) {
            int shop = e[0], movie = e[1], p = e[2];
            long long k = key(shop, movie);
            priceMap[k] = p;
            isRented[k] = false;
            avail[movie].insert({p, shop});
        }
    }

    vector<int> search(int movie) {
        vector<int> res;
        auto it_map = avail.find(movie);
        if (it_map == avail.end()) return res;
        auto &s = it_map->second;
        auto it = s.begin();
        for (int i = 0; i < 5 && it != s.end(); ++i, ++it) {
            res.push_back(it->second); // shop
        }
        return res;
    }

    void rent(int shop, int movie) {
        long long k = key(shop, movie);
        int p = priceMap[k];
        // remove from available
        auto it_map = avail.find(movie);
        if (it_map != avail.end()) {
            it_map->second.erase({p, shop});
        }
        // mark rented and add to rented set
        isRented[k] = true;
        rented.insert({p, shop, movie});
    }

    void drop(int shop, int movie) {
        long long k = key(shop, movie);
        int p = priceMap[k];
        // remove from rented set
        rented.erase({p, shop, movie});
        // mark available and insert back to avail set
        isRented[k] = false;
        avail[movie].insert({p, shop});
    }

    vector<vector<int>> report() {
        vector<vector<int>> res;
        auto it = rented.begin();
        for (int i = 0; i < 5 && it != rented.end(); ++i, ++it) {
            int p = get<0>(*it), shop = get<1>(*it), movie = get<2>(*it);
            res.push_back({shop, movie});
        }
        return res;
    }
};

/**
 * Usage:
 * MovieRentingSystem* obj = new MovieRentingSystem(n, entries);
 * vector<int> param_1 = obj->search(movie);
 * obj->rent(shop,movie);
 * obj->drop(shop,movie);
 * vector<vector<int>> param_4 = obj->report();
 */
