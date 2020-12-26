capture_easy
============

一个golang实现的验证码识别程序，不需要依赖其他语言项目
原项目地址:https://github.com/vannnnish/capture

我做了一些改动和调试运行，分享出来
1，解决验证码被固定6位数的bug
2，调试项目，可以直接运行


项目说明：
crm2 文件夹为程序主体部分
Training.dat 为字典库

操作步骤
训练ocr
1，下载100-200个图片文件到本地，例如项目的img文件夹
2，main主函数，调用trainloadsave方法，开始训练，并产生数据字典，注意调整这个
  const (
	  Threshhole = 34000
	  N          = 4   《----验证码长度
  )
3，运行代码，根据命令行输出的二进制阵图，输入你看到的数字
4，输入完成后，系统自动保存数据字典

识别验证码
1，主要查看recognize方法
c := crm2.NewCaptcha(Threshhole, N)   //创建对象
stdModule, err := c.LoadStdModule(`Std.dat`)   //加在上面生成的数据字典
img, err := crm2.ReadImg(`img/` + strconv.Itoa(i) + ".png")
fmt.Println(c.Recognize(img))   //输出识别到的字符串

欢迎共同讨论，微信：bestlive889
