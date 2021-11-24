# OnlineCourseAPI大学生网课刷课题库



### 题库内容

------

**支持各大网课的题目答案查询，超星，智慧树，知到，Welearn，四史，马原，毛概，大学mooc，等各种网课**

**有网页端直接搜索，也有api接口供开发者调用**

#### 支持模糊搜索！！（输入部分题目即可搜索到答案）



### **网页端搜索网址**

------

```http
http://wangke.hive-net.cn/wangke
```

**输入题目后回车即可**

### 答案API

------

**题库答案获取api如下：**

**JSON接口(更适合微信小程序)**

```http
https://wangke.hive-net.cn/wechat/search/?token=free&question=在什么情况下N95口罩需要更换?
```

**链接地址不变，只需要将question=后的“在什么情况下N95口罩需要更换?”更换为你像搜索的问题即可返回答案**

**token=后填写你的口令用于验证身份，token=free可以每日免费获取10000次答案，口令获取方式在下方**

**返回值为json格式数据，其中"has_reason"为是否搜索到答案，"question"为问题，"reason"为答案，汉字为Unicode编码，转化格式即可显示汉字，示例：**

```json
{"has_reason": 1, "question": "\u5728\u4ec0\u4e48\u60c5\u51b5\u4e0bN95\u53e3\u7f69\u9700\u8981\u66f4\u6362?", "reason": "\n1\uff1a\u53e3\u7f69\u6709\u7834\u635f.\u635f\u574f\u6216\u4e0e\u9762\u90e8\u65e0\u6cd5\u5bc6\u5408\u65f6\n2\uff1a\u53e3\u7f69\u53d7\u6c61\u67d3(\u5982\u67d3\u6709\u8840\u6e0d\u6216\u98de\u6cab\u7b49\u5f02\u7269\u65f6)\n3\uff1a\u547c\u5438\u963b\u6297\u660e\u663e\u589e\u52a0\u65f6\n4\uff1a\u66fe\u4f7f\u7528\u4e8e\u4e2a\u4f8b\u75c5\u623f\u6216\u75c5\u60a3\u63a5\u89e6(\u56e0\u4e3a\u8be5\u53e3\u7f69\u5df2\u88ab\u6c61\u67d3)"}
```
&emsp;&emsp;
**HTML接口(直接显示文本)**

```http
https://wangke.hive-net.cn/wechat/searchHtml/?token=free&question=在什么情况下N95口罩需要更换?
```

**链接地址不变，只需要将question=后的“在什么情况下N95口罩需要更换?”更换为你像搜索的问题即可返回答案**

**token=后填写你的口令用于验证身份，token=free可以每日免费获取10000次答案，口令获取方式在下方**

**返回值为文本数据，其中"has_reason"为是否搜索到答案，"question"为问题，"reason"为答案，示例：**
```json
{"has_reason":1, "question":在什么情况下N95口罩需要更换, "reason": 1：口罩有破损.损坏或与面部无法密合时 2：口罩受污染(如染有血渍或飞沫等异物时) 3：呼吸阻抗明显增加时 4：曾使用于个例病房或病患接触(因为该口罩已被污染), "remaining_times":868}
```


### **获取代码示例**

------

```python
import requests

def get_reason(question):
    url = "https://www.hive-net.cn:8443/wechat/search/?token=free&question=" + question
    try:  
        r = requests.get(url)  
        r.raise_for_status()  
        r.encoding = r.apparent_encoding
        print(r.json()['reason'])
    except:  
        print("Fail")

get_reason("在什么情况下N95口罩需要更换?")
```



### 学习交流

------

**QQ群：103172845**

**微信公众号：“夜寒信息”，所有的教程，疑难问答，及更新信息都在这里**

**哔哩哔哩：”暮至夜寒“，视频教程则会发布到这里**



### 获取token

------

**加QQ：296854007**

**加微信：Rem_wife**



#### **制作不易，点个Star再走，谢谢（将会获得更多免费次数）**
