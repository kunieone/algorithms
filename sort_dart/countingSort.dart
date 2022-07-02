countingSort(List arr) {
  var max = arr[0];
  var min = arr[0];

  for (var i = 0; i < arr.length; i++) {
    if (arr[i] < min) {
      min = arr[i];
    }
    if (arr[i] > max) {
      max = arr[i];
    }
  }
  var result = [];
  var buck = List.generate(max - min + 1, (i) => 0);
  for (var i = 0; i < arr.length; i++) {
    buck[arr[i] - min] += 1;
  }

  for (var i = 0; i < buck.length; i++) {
    int times = buck[i];
    while (times > 0) {
      result.add(i + min);
      times--;
    }
  }
  return result;
}

// [10,12,11]
// min=10   max =12

//  10
// [0,0,0,0,00,0,0,]
// [1,5,3,4,4,10000]
//  0 1 2 3
// [1,3,4,4,5]

// [0,1,0,1,1,1,0,0,0]
// 1,3,4,4,5
// 1,3,4,5

bool testSort(Function cb) {
  var testArr = [1, 4, 7, 9, 0, 2, 9, 3, 7, 8];
  if ([0, 1, 2, 3, 4, 7, 7, 8, 9, 9].toString() == cb(testArr).toString()) {
    print(cb(testArr).toString());
    print(cb.toString() + "通过！");
    return true;
  } else {
    print(cb(testArr).toString());
    print(cb.toString() + "不通过！");
    return false;
  }
}

void main() {
  testSort(countingSort);
}
