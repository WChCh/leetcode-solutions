## 题解1 - 链表长度(TLE) <i class="fa fa-thumbs-o-down"></i>

直观角度来考虑，如果把链表视为数组来处理，那么我们要做的就是依次将下标之和为`n`的两个节点链接到一块儿，使用两个索引即可解决问题，一个索引指向`i`, 另一个索引则指向其之后的第`n - 2*i`个节点(对于链表来说实际上需要获取的是其前一个节点), 直至第一个索引大于第二个索引为止即处理完毕。

既然依赖链表长度信息，那么要做的第一件事就是遍历当前链表获得其长度喽。获得长度后即对链表进行遍历，小心处理链表节点的断开及链接。用这种方法会提示 TLE，也就是说还存在较大的优化空间！  

### 源码分析

1. 异常处理，对于节点数目在两个以内的无需处理。
2. 遍历求得链表长度。
3. 遍历链表，第一个索引处的节点使用`last`表示，第二个索引处的节点的前一个节点使用`beforeTail`表示。
4. 处理链表的链接与断开，迭代处理下一个`last`。

### 复杂度分析

1. 遍历整个链表获得其长度，时间复杂度为 `$$O(n)$$`.
2. 双重`for`循环的时间复杂度为 `$$(n-2) + (n-4) + ... + 2 = O(\frac{1}{2} \cdot n^2)$$`.
3. 总的时间复杂度可近似认为是 `$$O(n^2)$$`, 空间复杂度为常数。

> **Warning** 使用这种方法务必注意`i`和`j`的终止条件，若取`i < length + 1 - i`, 则在处理最后两个节点时会出现环，且尾节点会被删掉。在对节点进行遍历时务必注意保留头节点的信息！

## 题解2 - 反转链表后归并

既然题解1存在较大的优化空间，那我们该从哪一点出发进行优化呢？擒贼先擒王，题解1中时间复杂度最高的地方在于双重`for`循环，在对第二个索引进行遍历时，`j`每次都从`i`处开始遍历，要是`j`能从链表尾部往前遍历该有多好啊！这样就能大大降低时间复杂度了，可惜本题的链表只是单向链表... 有什么特技可以在单向链表中进行反向遍历吗？还真有——反转链表！一语惊醒梦中人。  

### 源码分析

相对于题解1，题解2更多地利用了链表的常用操作如反转、找中点、合并。

1. 找中点：我在九章算法模板的基础上增加了对`head->next`的异常检测，增强了鲁棒性。
2. 反转：非常精炼的模板，记牢！
3. 合并：也可使用九章提供的模板，思想是一样的，需要注意`left`, `right`和`dummy`三者的赋值顺序，不能更改任何一步。

### 复杂度分析

找中点一次，时间复杂度近似为 `$$O(n)$$`. 反转链表一次，时间复杂度近似为 `$$O(n/2)$$`. 合并左右链表一次，时间复杂度近似为 `$$O(n/2)$$`. 故总的时间复杂度为 `$$O(n)$$`.

## Reference

- [Reorder List | 九章算法](http://www.jiuzhang.com/solutions/reorder-list/)

