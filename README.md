<div align="center">
    <img src="https://www.hive-net.cn/Assets/SiteGlobal/Hive_blank.png" width="200" alt="HeroPower"/>
    <h1>OnlineCourseAPI大学生网课刷课题库</h1>



[![PackageVersion](https://img.shields.io/badge/java-11-orange)](https://www.oracle.com/java/technologies/downloads/#java11)
[![PackageVersion](https://img.shields.io/badge/suggestion-issue-blue)](https://github.com/Raptor-wxw/OnlineCourseAPI/issues)
</div>



### 题库内容

------

**支持各大网课的题目答案查询，超星，智慧树，知到，Welearn，四史，马原，毛概，大学mooc，等各种网课**

**有网页端直接搜索，也有api接口供开发者调用**

#### 支持模糊搜索！！（输入部分题目即可搜索到答案）

&emsp;&emsp;

### **网页端搜索网址**

------

```http
https://www.hive-net.cn/backend/wangke/index
```

**输入题目后回车即可**

&emsp;&emsp;

### 答案API

------

**题库答案获取api如下：**

**JSON接口(更适合微信小程序)**

```http
https://www.hive-net.cn/backend/wangke/search?token=free&question=我国的国体是
```

**链接地址不变，只需要将question=后的“在什么情况下N95口罩需要更换?”更换为你像搜索的问题即可返回答案**

**token=后填写你的口令用于验证身份，token=free可以每日免费获取10000次答案，口令获取方式在最下方**

**返回值为json格式数据，其中"code"为是否搜索到答案（0有，-1无），"question"为问题，"reason"为答案，type为题目类型（0单选，1多选），示例：**

```json
{
    "code": 0,
    "data": {
        "total": 1,
        "reasonList": [
            {
                "id": 409583,
                "question": "我国的国体是",
                "reason": "人民民主专政",
                "type": 1,
                "options": "A:人民代表大会制度,B:人民民主专政,C:共产党领导的多党合作和政治协商制度,D:民族区域自治制度",
                "explanation": "无",
                "course": "中国大学MOOC慕课,中国大学MOOC慕课未分类",
                "tag": "毛概,多党合作,制度"
            }
        ],
        "tokenRemainTimes": 4995,
        "tokenExpireTime": "2030-01-01"
    }
}
```

&emsp;&emsp;


### **获取代码示例**

------

```python
import requests

def get_reason(question):
    url = "https://www.hive-net.cn/backend/wangke/search?token=free&question=" + question
    try:  
        r = requests.get(url)  
        r.raise_for_status()  
        r.encoding = r.apparent_encoding
        print(r.json()['reason'])
    except:  
        print("Fail")

get_reason("我国的国体是")
```


&emsp;&emsp;


### 学习交流

------

**QQ群：103172845**

**微信公众号：“夜寒信息”，所有的教程，疑难问答，及更新信息都在这里**

**哔哩哔哩：”暮至夜寒“，视频教程则会发布到这里**


&emsp;&emsp;


### 获取token

------

**加QQ：296854007**

**加微信：Rem_wife**



#### **制作不易，点个Star再走，谢谢（将会获得更多免费次数）**
