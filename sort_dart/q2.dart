quickSort2(List arr) {
  if (arr.length < 2) {
    return arr;
  }
  var pivot = arr[0];
  print(pivot);
  List left = [];
  List right = [];
  for (var i = 1; i < arr.length; i++) {
    if (arr[i] > pivot) {
      right.add(arr[i]);
    } else {
      left.add(arr[i]);
    }
  }

  return [...quickSort2(left), pivot, ...quickSort2(right)];
}

void main(List<String> args) {
  print(quickSort2([2, 1, 9, 7, 6, 5]));
}
