export function radixSort(arr: number[], base: number) {
  function radixSort(arr: number[], baseNow: number) {
    let maxBase: number = 0;
    arr.forEach((e) => {
      let thiBase = 10 ** (e.toString().length - 1);
      maxBase = maxBase < thiBase ? thiBase : maxBase;
    });
    let bucket: number[][];
    arr.forEach((e) => {
      Math.floor((e % baseNow) / Math.floor(baseNow / base));
    });
  }
}

radixSort([2, 1, 1000, 3], 10);
//第一位:  8  十进制:  8% (10**1) /(10**1/10)=8
//第二位   48 十进制: 48% (10**2) /(10**2/10)=4
// 8  -> 8进制的话  --> 0     8%8 =0

function getDigit(num: number, dig: number, base: number = 10): number {
  return Math.floor(
    (num % base ** (dig + 1)) / Math.floor(10 ** (dig + 1) / base),
  );
}
console.log(getDigit(1236, 2, 10));
