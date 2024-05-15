<div align="center">
    <img src="https://www.hive-net.cn/Assets/SiteGlobal/Hive_blank.png" width="200" alt="HeroPower"/>
    <h1>OnlineCourseAPI大学生网课刷课题库</h1>



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

#### API接口：

```http
https://www.hive-net.cn/backend/wangke/search?token=卡密&question=问题
```


#### 请求参数：
  - `token`：填写你的口令用于验证身份（`free`可以每日免费获取5000次答案）
  - `question`：你想搜索的问题

#### 返回参数
  - `code`：是否搜索到答案（0有，-1无）
  - `type`：题目类型（0单选，1多选）
  - `reason`：答案

#### 使用：

**链接地址不变，将`token=`后的`free`改为您的`token`，再将`question=`后的`在什么情况下N95口罩需要更换?`改为您要搜索的问题（获取token点击：[获取token](#获取token)）**

#### 示例：

请求：
```http
https://www.hive-net.cn/backend/wangke/search?token=free&question=我国的国体是
```
响应：
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
import json
import requests

def get_reason(question):
    url = "https://www.hive-net.cn/backend/wangke/search?token=free&question=" + question
    try:
        response = requests.get(url)
        response.raise_for_status()
        result = json.loads(response.text)
        if result.get("code") == 0:
            print(result.get("data"))
        else:
            print("无答案")
    except:
        print("Connection error")

get_reason("我国的国体是")
```


&emsp;&emsp;


### 学习交流

------

**QQ群：103172845**

**微信公众号：“夜寒信息”，所有的教程，疑难问答，及更新信息都在这里**


&emsp;&emsp;


### 获取token

------

**自助购买：**[点击购买](https://mall.sxjf8789.com/links/51B05514)

**售后群：[103172845](https://qm.qq.com/cgi-bin/qm/qr?k=sJLLnl1RdSdA5nhd7IXbhCxd-k3KaoBl&authKey=ssD9NFl2r5rHhGL4SvyIF56kSJi33zxFu2LqZ0XvUUGIZN3CyhCanNyji7cNXAwo&noverify=0&group_code=103172845#)**

#### **制作不易，点个Star再走，谢谢（将会获得更多免费次数）**
