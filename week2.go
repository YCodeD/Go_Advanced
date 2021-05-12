package main

/*
作业问题：

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


答：
应该将 sql.ErrNoRows 往上抛，因为在与数据库操作的时候
会调用到第三方库或者是标准库，这个时候需要考虑使用errors.Wrap 或 errors.Wrapf
保存堆栈信息，然后抛给上层调用方。
*/

/*
不应该把底层的错误信息wrap往上抛，应该进行一层包装之后再往上抛
*/
