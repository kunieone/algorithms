// 鸡尾酒排序是一个双向冒泡的算法.

List cocktailSort(List arr) {
  var len = arr.length;

  bool isSorted = false;
  for (int i = 0; i < len / 2; i++) {
    isSorted = false;

    for (int j = i; j < len - i - 1; j++) {
      if (arr[j] > arr[j + 1]) {
        var tmp = arr[j];
        arr[j] = arr[j + 1];
        arr[j + 1] = tmp;

        isSorted = true;
      }
    }

    for (int j = len - i - 1; j > i; j--) {
      if (arr[j] < arr[j - 1]) {
        var tmp = arr[j];
        arr[j] = arr[j - 1];
        arr[j - 1] = tmp;

        isSorted = true;
      }
      isSorted = true;
    }

    if (isSorted) break;
  }

  return arr;
}

void main(List<String> args) {
  print(cocktailSort([2, 4, 6, 1, 4, 3, 5, 9, 7]));
}
