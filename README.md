# 东北大学校园二手市场

## 数据库设计
* 用户
  学号 姓名 性别 年级 院系 专业 校区 班级 头像 
  用户状态		 用法：0 / 1      禁止 / 允许
  手机号 微信号 QQ 用法：联系方式 
  最后登录时间     用法：xx时间前在线
  创建时间        用法：已加入 NeuMark xx 天

* 商品类别
  类别id
  父类别id
  类别
  
* 商品
  商品id
  名称 
  分类 
  描述 文字描述 图片描述
  价格
  浏览量
  状态
  发布人学号
  发布时间
* 商品评论
  评论id
  商品id
  发布人id
  评论内容

## 页面
* 首页
* 分类
  * 商品详情页
* 发布
* 消息
* 我的

  * 我发布的
  * 我收藏的
  * 我买到的
  * 我浏览的

  * 关于NeuMark
  * 意见反馈
  * 联系程序员

## UI库
* ColorUI https://github.com/weilanwl/ColorUI
* WeUI https://github.com/Tencent/weui-wxss
* iView https://github.com/TalkingData/iview-weapp
* ZanUI https://github.com/youzan/zanui-weapp 
* MinUI https://github.com/meili/minui


## 商品类别
电子产品： 手机 耳机 u盘 电脑 相机
书籍： 专业书 课外书
生活用品： 护肤品 衣服 鞋
## 未来规划
店铺 和学校的文创、二手书店合作

## 功能拓展
搜索
图片
高并发
前端美化

## 开发工具
数据库可视化： navicat premium
IDE： GoLand
前端： 微信小程序开发者工具
后端接口测试： postman
