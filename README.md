### LeetCode(Golang版)

#### 已完成列表

##### Easy

- [x] 1.两数之和（数组）

  解法：暴力破解，双for循环

  执行用时：32 ms

  内存消耗：3.6 MB

- [x] 7.整数反转（整数）

  解法：for循环，取余

  执行用时：0 ms

  内存消耗：2.1 MB

- [x] 9.回文数（整数）

  解法：先进行整数反转，再比对是否相等

  执行用时：32 ms

  内存消耗：4.9 MB

- [x] 262.翻转二叉树（数组）

  解法：递归，对左右子节点进行交换

  执行用时：0 ms

  内存消耗：2.1 MB

##### Medium

- [x] 2.两数相加（链表）

  解法：对两条链表进行遍历相加

  执行用时：12 ms

  内存消耗：4.4 MB

- [x] 114.二叉树展开为链表（二叉树）

  解法：递归，将子树拉平，再对右树进行拼接

  执行用时：0 ms

  内存消耗：2.8 MB

- [x] 116.填充每个节点的下一个右侧节点指针（二叉树->N叉树）

  解法：递归，通过双节点的方式进行节点连接的建立，也解决了跨节点的问题

  执行用时：8 ms

  内存消耗：6.3 MB

- [x] 652.寻找重复子树

  解法：递归，中序遍历，将遍历出的结果转换为字符串放入map中进行出现次数累加

  执行用时：16 ms

  内存消耗：12.6 MB

- [x] 654.最大二叉树（二叉树）

  解法：递归，先找出数组中的最值，再将最有两边的数组进行子树的递归构建

  执行用时：16 ms

  内存消耗：7 MB

- [x] 797.所有可能的路径（有向无环图）

  解法：DFS，从0开始进行深度探索，并拷贝经过的路径

  执行用时：8 ms

  内存消耗：6.7 MB

