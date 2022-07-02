List quickSort(List arr) {
  // 快速排序，使用的其实就是分而治之的比较
  if (arr.length < 2) {
    return arr;
  }
  var pivot = arr[0];
  List left = [];
  List right = [];
  for (var i = 1; i < arr.length; i++) {
    if (pivot > arr[i]) {
      left.add(arr[i]);
    } else {
      right.add(arr[i]);
    }
  }

  return [...quickSort(left), pivot, ...quickSort(right)];
}

void main(List<String> args) {
  print(quickSort([2, 1, 9, 7, 6, 5]));
}
