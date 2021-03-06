/*
当我们在编程过程中，往往需要对线性表进行查找操作。
在顺序表中查找时，需要从表头开始，依次遍历比较a[i]与key的值是否相等，直到相等才返回索引i；
在有序表中查找时，我们经常使用的是二分查找，通过比较key与a[i]的大小来折半查找，直到相等时才返回索引i。最终通过索引找到我们要找的元素。
但是，这两种方法的效率都依赖于查找中比较的次数。我们有一种想法，能不能不经过比较，而是直接通过关键字key一次得到所要的结果呢？这时，就有了散列表查找（哈希表）。 */
package hash

