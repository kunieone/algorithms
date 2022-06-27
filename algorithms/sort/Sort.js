function countingSort(arr) {
  // let min = Math.min(...arr);
  let min = 0;
  let bucket = [];
  let res = [];
  for (let ii = 0; ii < arr.length; ii++) {
    bucket[arr[ii] - min] === undefined
      ? (bucket[arr[ii] - min] = 1)
      : (bucket[arr[ii] - min] += 1);
  }
  for (let i = 0; i < bucket.length; i++) {
    if (bucket[i] !== undefined) while (bucket[i]--) res.push(i + min);
  }
  return res;
}

function genNArr(len) {
  let arr = [];
  for (let i = 0; i < len; i++) arr.push(i);
  return arr.sort((a, b) => (Math.random() >= 0.5 ? a : b));
}

function bubbleSort(arr) {
  for (let i = 0; i < arr.length; i++) {
    for (let j = i; j < arr.length; j++) {
      if (arr[i] > arr[j]) {
        [arr[i], arr[j]] = [arr[j], arr[i]];
      }
    }
  }
  return arr;
}

const merge = (left, right) => {
  let [res, i, j] = [[], 0, 0];
  while (i < left.length && i < right.length) {
    res.push(left[i] < right[j] ? left : right);
    left[i] < right[j] ? i++ : j++;
  }
  res.concat(left.slice(i));
  res.concat(right.slice(j));
};
function mergeSort(arr) {
  let mid = parseInt(arr.length / 2);
  return merge(arr.slice(0, mid), arr.slice(mid));
}

function gcd(a, b) {
  if (b == 0) {
    return a;
  } else {
    return gcd(b, a % b);
  }
}

let arr = genNArr(30_000);
console.time("mergeSort排序时长");
mergeSort(arr);
console.timeEnd("mergeSort排序时长");

console.time("bubbleSort排序时长");
bubbleSort(arr);
console.timeEnd("bubbleSort排序时长");

console.time("countingSort排序时长");
countingSort(arr);
console.timeEnd("countingSort排序时长");

console.time("gcd");
gcd(20000, 300);
console.timeEnd("gcd");
