merge(List left, List right) {
  List res = [];

  int i = 0;
  int j = 0;
  while (i < left.length && j < right.length) {
    if (left[i] < right[j]) {
      res.add(left[i]);
      i++;
    } else {
      res.add(right[j]);
      j++;
    }
  }
  res.addAll(left.sublist(i));
  res.addAll(right.sublist(j));
  return res;
}

mergeSort(List arr) {
  if (arr.length < 2) {
    return arr;
  }
  int mid = arr.length ~/ 2;
  List left = arr.sublist(0, mid);
  List right = arr.sublist(mid);
  return merge(mergeSort(left), mergeSort(right));
}

main(List<String> args) {
  var arr = [2, 6, 5, 4, 8, 0, 9, 1, 2, 2];
}
