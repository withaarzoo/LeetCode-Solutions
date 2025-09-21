import java.util.*;

class MovieRentingSystem {
    private static long key(int shop, int movie) {
        return (((long)shop) << 32) | (movie & 0xffffffffL);
    }

    // price map and status
    private Map<Long, Integer> priceMap = new HashMap<>();
    private Map<Long, Boolean> isRented = new HashMap<>();

    // available: movie -> TreeSet of Pair(price, shop)
    private Map<Integer, TreeSet<Pair>> avail = new HashMap<>();

    // rented: global TreeSet of Triple(price, shop, movie)
    private TreeSet<Triple> rented = new TreeSet<>();

    // Pair for available entries
    private static class Pair implements Comparable<Pair> {
        int price, shop;
        Pair(int p, int s){ price = p; shop = s; }
        public int compareTo(Pair o){
            if (this.price != o.price) return Integer.compare(this.price, o.price);
            return Integer.compare(this.shop, o.shop);
        }
        public boolean equals(Object o){
            if (!(o instanceof Pair)) return false;
            Pair p = (Pair)o;
            return this.price == p.price && this.shop == p.shop;
        }
        public int hashCode(){ return Objects.hash(price, shop); }
    }

    // Triple for rented entries
    private static class Triple implements Comparable<Triple> {
        int price, shop, movie;
        Triple(int p, int s, int m){ price = p; shop = s; movie = m; }
        public int compareTo(Triple o){
            if (this.price != o.price) return Integer.compare(this.price, o.price);
            if (this.shop != o.shop) return Integer.compare(this.shop, o.shop);
            return Integer.compare(this.movie, o.movie);
        }
        public boolean equals(Object o){
            if (!(o instanceof Triple)) return false;
            Triple t = (Triple)o;
            return this.price == t.price && this.shop == t.shop && this.movie == t.movie;
        }
        public int hashCode(){ return Objects.hash(price, shop, movie); }
    }

    public MovieRentingSystem(int n, int[][] entries) {
        for (int[] e : entries) {
            int shop = e[0], movie = e[1], p = e[2];
            long k = key(shop, movie);
            priceMap.put(k, p);
            isRented.put(k, false);

            avail.computeIfAbsent(movie, x -> new TreeSet<>()).add(new Pair(p, shop));
        }
    }

    public List<Integer> search(int movie) {
        List<Integer> res = new ArrayList<>();
        TreeSet<Pair> set = avail.get(movie);
        if (set == null) return res;
        Iterator<Pair> it = set.iterator();
        int cnt = 0;
        while (it.hasNext() && cnt < 5) {
            res.add(it.next().shop);
            cnt++;
        }
        return res;
    }

    public void rent(int shop, int movie) {
        long k = key(shop, movie);
        int p = priceMap.get(k);
        TreeSet<Pair> set = avail.get(movie);
        if (set != null) set.remove(new Pair(p, shop));
        isRented.put(k, true);
        rented.add(new Triple(p, shop, movie));
    }

    public void drop(int shop, int movie) {
        long k = key(shop, movie);
        int p = priceMap.get(k);
        rented.remove(new Triple(p, shop, movie));
        isRented.put(k, false);
        avail.computeIfAbsent(movie, x -> new TreeSet<>()).add(new Pair(p, shop));
    }

    public List<List<Integer>> report() {
        List<List<Integer>> res = new ArrayList<>();
        Iterator<Triple> it = rented.iterator();
        int cnt = 0;
        while (it.hasNext() && cnt < 5) {
            Triple t = it.next();
            res.add(Arrays.asList(t.shop, t.movie));
            cnt++;
        }
        return res;
    }
}

/**
 * Usage:
 * MovieRentingSystem obj = new MovieRentingSystem(n, entries);
 * List<Integer> p1 = obj.search(movie);
 * obj.rent(shop,movie);
 * obj.drop(shop,movie);
 * List<List<Integer>> p4 = obj.report();
 */
