### 作业解析
1. 有三处需阻塞：errgroup.Wait()，等待信号，server的监听

### 解决办法
1. errgroup启动三个goroutine，main函数结尾通过Wait方法保证3个goroutine退出
2. 1个goroutine用来注册和接收信号，接收到信号后进行cancel
3. 2个goroutine用来启动http server。每个server启动一个野生goroutine处理cancel情况，保证优雅退出
