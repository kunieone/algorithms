class SegTree {
  key: string;
  value: number;
  left?: SegTree;
  right?: SegTree;
  constructor(start: number, arr: number[]) {
    const end = start + arr.length - 1;
    this.key = `${start}-${end}`;
    const [left, right] = seg(arr);
    this.value = this.sub.reduce((p: number, v: number) => p + v, 0);
    this.left = new SegTree(start, left);
    this.right = new SegTree();
  }
  // 6: 0,1,2,3,4,5 3
}
function ArrToSegTree(arr: number[]): SegTree {
  return new SegTree(0, arr);
}

console.log(JSON.stringify(ArrToSegTree([1, 3, 5, 7, 9])));

// 6 -> 0-5 5/2 =2,2+1 =3
// 0123
const seg = (arr: number[]) => [
  arr.slice(0, Math.floor((arr.length + 1) / 2)),
  arr.slice(Math.floor((arr.length + 1) / 2), arr.length),
];
console.log(seg([2, 12, 1, 6, 8, 0, 0xff]));

//
//      [1,2,3,4,5] len=6
//
//  [1,2,3]  [4,5]
//
//  [1,2] [3]    [4] [5]
