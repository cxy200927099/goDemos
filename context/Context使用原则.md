1、不要把Context放在结构体中，要以参数的方式传递

2、以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。

3、给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO

4、Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递

5、Context是线程安全的，可以放心的在多个goroutine中传递