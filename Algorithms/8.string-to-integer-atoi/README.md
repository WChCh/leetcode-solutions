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