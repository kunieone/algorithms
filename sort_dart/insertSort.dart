import 'dart:math';

List insertSort(List arr) {
  int len = arr.length; //获得目前的长度
  for (int i = 1; i < len; i++) {
    // 从索引一开始
    int key = arr[i];
    // 获得该索引的元素，该元素是要往前滑动比较的元素
    int j = i - 1;
    // 拿到前面的一个的索引
    while ((j >= 0) && (key < arr[j])) {
      // 在索引大于0的范围内，不断往前走，如果遇到更大的数，就停下来。
      arr[j + 1] = arr[j]; //如果是更小的数，则会把这个值往右移动一个位置。
      j--;
    }
    arr[j + 1] = key;
  }
  return arr;
}

List genArrRan(int length) {
  Random r = new Random();
  List list = [];
  for (var i = 0; i < length; i++) {
    list.add(r.nextInt(length));
  }
  return list;
}

void main(List<String> args) {
  print(insertSort([2, 9, 1, 8, 0, 7, 6, 4, 5, 2]));
}
