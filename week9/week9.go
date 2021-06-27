package week9

// 1、总结几种 socket 粘包的解包方式:
// fix length/delimiter based/length field based frame decoder。尝试举例其应用
//  - fix length
// 	发送端将每个数据包封装为固定长度（不够的可以通过补 0 来填充），接收端每次从接收缓冲区读
// 取数据的时候，都读取固定长的的数据，从而把每个数据包拆分开来。
// 	但是指定的长度过大，消息过短会造成资源浪费。
//  - delimiter based
// 	使用特殊符号来作为数据包结束的分隔符，接收端就可以通过这个边界来进行解包。
//  - length field based
// 	通用解码器，一般协议头中带有长度字段，通过传入特定的参数来解决粘包问题。

// 2、实现一个从 socket connection 中解码出 goim 协议的解码器。
// 	未实现
