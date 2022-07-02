List bubbleSort(List arr) {
  for (var i = 0; i < arr.length; i++) {
    for (var j = i; j < arr.length; j++) {
      if (arr[i] > arr[j]) {
        dynamic tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
      }
    }
  }
  return arr;
}

void main(List<String> args) {
  print(args);

  print(bubbleSort([2.1, 2.13, 2.12, 8, 6, 5, 4]));
}
