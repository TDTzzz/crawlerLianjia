# 单机并发爬虫（武汉链家房源信息）



#### 项目简介

---

每日爬取武汉链家二手房的数据，可视化分析

后端：Golang 1.14

存储：ElasticSearch 7.6.1

可视化：Kibana+Vue



#### TODO

---



- [x] 用go协程实现并发爬取
- [x] 用elastic储存数据
- [x] 将es的数据用kibana可视化
- [x] 按照构想的数据维度，用gin将数据封装成api
- [x] 用vue搭建一个spa图看板页面
- [ ] 编写docker-compose，自动化部署
- [ ] 结合docker搭建分布式并发架构



> ps: 效果图见ex_imgs目录里
