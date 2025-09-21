/**
 * @param {number} n
 * @param {number[][]} entries
 */
var MovieRentingSystem = function (_n, entries) {
  this.maxNumSearchResults = 5;

  let sorted = [...entries].sort(([shop1, _1, price1], [shop2, _2, price2]) => {
    let priceDiff = price1 - price2;

    return priceDiff ? priceDiff : shop1 - shop2;
  });

  //{ movie: [ [shop1, priceLowest] ... [shopN, priceHighest] ] }
  this.movies = sorted.reduce((movies, [shop, movie]) => {
    let shops = movies[movie];

    if (shops == undefined) shops = movies[movie] = [];

    shops.push(shop);

    return movies;
  }, {});

  /*
        {
            shop: { movie: { price } }
        }
    */
  this.shops = sorted.reduce((shops, [shop, movie, price]) => {
    let data = shops[shop];

    if (data == undefined) data = shops[shop] = {};

    data[movie] = { price };

    return shops;
  }, {});

  //i = [shop, movie]
  this.rented = [];
};

/**
 * @param {number} movie
 * @return {number[]}
 */
MovieRentingSystem.prototype.search = function (movie) {
  let results = [],
    shops = this.movies[movie];

  if (shops) {
    for (
      let i = 0, l = shops.length;
      results.length < this.maxNumSearchResults && i < l;
      i++
    ) {
      let shop = shops[i];

      if (!this.shops[shop][movie].rented) results.push(shop);
    }
  }

  return results;
};

/**
 * @param {number} shop
 * @param {number} movie
 * @return {void}
 */
MovieRentingSystem.prototype.rent = function (shop, movie) {
  if (this.shops[shop][movie].rented == undefined) {
    let rentData = [shop, movie];

    this.shops[shop][movie].rented = rentData;
    this.rented.push(rentData);
  }
};

/**
 * @param {number} shop
 * @param {number} movie
 * @return {void}
 */
MovieRentingSystem.prototype.drop = function (shop, movie) {
  let movieData = this.shops[shop]?.[movie].rented;

  if (movieData) {
    this.rented.splice(this.rented.indexOf(movieData), 1);
    delete this.shops[shop][movie].rented;
  }
};

/**
 * @return {number[][]}
 */
MovieRentingSystem.prototype.report = function () {
  return this.rented
    .sort(([shop1, movie1], [shop2, movie2]) => {
      let priceDiff =
        this.shops[shop1][movie1].price - this.shops[shop2][movie2].price;

      if (!priceDiff) return shop1 == shop2 ? movie1 - movie2 : shop1 - shop2;

      return priceDiff;
    })
    .slice(0, this.maxNumSearchResults);
};
