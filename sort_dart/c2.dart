List countSort(List arr) {
  int min = arr[0];
  int max = arr[0];
  List result = [];
  for (var i = 0; i < arr.length; i++) {
    if (arr[i] < min) {
      min = arr[i];
    }

    if (arr[i] > max) {
      max = arr[i];
    }
  }
  print("目前的最小值：$min最大值：$max");

  // 开始真正的计数排序：

  List bucket = List.generate(max - min + 1, (i) => 0);

  for (var i = 0; i < arr.length; i++) {
    int index = arr[i] - min;
    bucket[index]++;
    print(bucket);
  }

// bucket
  for (var i = 0; i < bucket.length; i++) {
    int index = bucket[i];
    while (index > 0) {
      result.add(i + min);
      index--;
    }
  }
  return result;
}

// [1, 1, 0, 0, 2]   -> [1,2,5,5]

void main(List<String> args) {
  print(countSort([100, 1]));
  // min =1
}
