const flatten = (arr) => {
  let res = [];
  for (const i of arr) {
    if (Array.isArray(i)) {
      res = res.concat(flatten(i));
    } else {
      res.push(i);
    }
  }
  return res;
};

const array = [1, 2, [3, 4, [5, 6, [7, 8, [9, 10]]]]];
console.log(flatten(array));
