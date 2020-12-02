### Week02 作业：
Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A: `sql.ErrNoRows` 是由 **database/sql** 抛出的错误, **dao** 层接收到 `sql.ErrNoRows` 错误应该 `Wrap` 错误并返回给上层应用