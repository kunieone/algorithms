import 'dart:math';

List baseSort(List arr, int base) {
  int maxDigit = findMaxDigit(arr, base);
  int digit = 0;
  while (digit <= maxDigit) {
    List<List> bucket = List.generate(base, (i) => []);
    for (var i = 0; i < arr.length; i++) {
      int number = select(arr[i], digit, base);
      bucket[number].add(arr[i]);
    }
    arr = putOut(bucket);
    print(digit);
    digit++;
  }
  return arr;
}

List putOut(List<List> bucket) {
  List result = [];
  for (var i = 0; i < bucket.length; i++) {
    result.addAll(bucket[i]);
  }
  return result;
}

int select(int number, int digit, int base) {
  return number % (pow(base, digit + 1)) ~/ pow(base, digit);
}

int findMaxDigit(List arr, int base) {
  int max = 0;
  for (var i = 0; i < arr.length; i++) {
    if (max < arr[i].toRadixString(base).length) {
      max = arr[i].toRadixString(base).length;
    }
  }
  return max;
}

void main(List<String> args) {
  var now = new DateTime.now();
  print(baseSort([1000, 200, 3288, 24, 3, 4, 11, 7], 16));
  var now2 = new DateTime.now();
  var difference = now2.difference(now);
  print(difference.toString());
}
