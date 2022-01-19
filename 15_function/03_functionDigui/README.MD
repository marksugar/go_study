<!--
 * @Author: Allen_Jol
 * @LastEditors: Allen_Jol
 * @Date: 2022-01-19 10:32:20
 * @LastEditTime: 2022-01-19 11:05:08
 * @version: v1.0
-->

### Instructions
---
递归调用流程分析

![](https://raw.githubusercontent.com/allenjol/images/main/imgs/20220119103157.png?token=AG5BNI3PMWDD2Z23NOCWWW3B454E2)

---
![](https://raw.githubusercontent.com/allenjol/images/main/imgs/20220119110236.png?token=AG5BNI3TIIW25NWDWGHJF5LB457XW)


### 总结
---

函数递归调用：一个函数在函数体内又调用了本身，就称之为递归调用

函数递归调用要遵守的原则：
>- 执行一个函数时就创建一个新的受保护的独立空间(新函数栈)
>- 函数的局部变量是独立的，不会互相影响
>- 递归必须向退出递归的条件逼近，否则就是无限递归
>- 当一个函数执行完毕，或者遇到了return，就会返回。遵守谁调用，就将结果返回给谁的原则