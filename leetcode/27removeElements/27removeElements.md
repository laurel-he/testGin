移除元素
===
题目来源：[27 移除元素](https://leetcode.cn/problems/remove-element/description/?envType=study-plan-v2&envId=top-interview-150)
# 分析
## 注意事项
（1）“原地”删除
（2）交换后确认删除的数据都在末尾，可以直接删除
## 解题思路
在题目要求中加黑提醒：需要原地移除元素，即不能引入额外的内存空间
移除元素后，所有元素都要往前移动
双指针：
一个指向第一个元素一个指向最后一个元素
当匹配到数据时，交换两个指针指向数据，并将左指针向右，右指针向左移动
记录交换的次数，移除