## 方法一  

使用正则表达式：  

```
^：匹配字符串开头
[\+\-]：代表一个+字符或-字符
?：前面一个字符可有可无
\d：一个数字
+：前面一个字符的一个或多个
\D：一个非数字字符
*：前面一个字符的0个或多个

max(min(数字, 2**31 - 1), -2**31) 用来防止结果越界
```

## 方法二  

使用状态机  

## 方法三  

单指针遍历串  

```

从头遍历串，将i指针定到 字符串中数字 的 最高位：

    遇到 +或-，则其下一位作为 int32 的 最高位，终止遍历；
    遇到数字，则其作为 int32 的 最高位，终止遍历；
    遇到 空格，跳过考虑下一位；
    其它情况，直接返回0。

从 int32 的最高位索引开始遍历串：当遇到非数字，终止遍历，将j指针定到 字符串中数字 的 最低位 的右边。
此时我们已经知道int32的 最高位 和 最低位 的索引。只需要累加起来就得到结果，注意：当遇到overflow时，返回边界值。
```


