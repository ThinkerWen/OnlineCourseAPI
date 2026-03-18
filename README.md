# 🎓 在线网课题库API文档

<div align="center">
    <img src="https://www.hive-net.cn/assets/site/Hive_blank.png" width="120"alt="HeroPower">
    <h2>OnlineCourseAPI - 大学生网课答案查询服务</h2>
    <p><strong>支持全国各大网课平台的题目答案查询</strong></p>
    
[![Issues](https://img.shields.io/badge/Feedback-Welcome-success)](https://github.com/ThinkerWen/OnlineCourseAPI/issues)

</div>

---

## 📋 功能特性

- ✅ **支持多平台**：超星、智慧树、知到、Welearn、四史、马原、毛概、大学MOOC等
- ✅ **模糊搜索**：输入部分题目关键词即可匹配答案
- ✅ **双端支持**：网页端和API接口双管齐下
- ✅ **实时更新**：题库持续更新维护

---

## 🌐 快速开始

### 方式一：网页端搜索（推荐新手）

**访问地址：** https://www.hive-net.cn/backend/course/index

只需输入题目关键词，回车即可查看答案。

---

## 🔌 API 接口文档

### 接口地址

```http
GET https://www.hive-net.cn/backend/course/search?token={token}&question={question}
```

```http
POST https://www.hive-net.cn/backend/course/search
Content-Type: application/json
```

### 请求参数

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| `token` | string | ✅ | 身份认证令牌（`free` 表示免费账户，每日10000次） |
| `question` | string | ✅ | URL编码后的问题关键词 |

> POST 请求时，这两个参数放在 JSON 请求体中。

POST JSON 请求体示例：

```json
{
    "token": "free",
    "question": "我国的国体是"
}
```

### 响应参数

| 字段 | 类型 | 说明 |
|------|------|------|
| `code` | int | 状态码：0=成功，-1=无结果 |
| `data.list` | array | 答案列表 |
| `data.remain_times` | int | 剩余查询次数 |
| `data.expire_time` | string | Token过期时间 |

#### 答案对象字段

| 字段 | 说明 |
|------|------|
| `id` | 答案ID |
| `question` | 完整问题 |
| `reason` | 答案内容 |
| `type` | 题型（0=单选，1=多选） |
| `options` | 选项列表 |
| `explanation` | 解析说明 |
| `course` | 所属课程 |
| `tag` | 标签分类 |

---

## 💡 使用示例

### 请求示例

```http
POST https://www.hive-net.cn/backend/course/search
Content-Type: application/json

{
    "token": "free",
    "question": "我国的国体是"
}
```

### 响应示例

```json
{
    "code": 0,
    "data": {
        "list": [
            {
                "question": "我国的国体是",
                "answer": "人民民主专政",
                "answer_type": 0,
                "answer_options": "A:人民代表大会制度,B:人民民主专政,C:共产党领导的多党合作和政治协商制度,D:民族区域自治制度",
                "answer_explanation": "解析见答案",
                "course": "中国大学MOOC慕课",
                "tag": "毛概,多党合作,制度"
            }
        ],
        "remain_times": 9995,
        "expire_time": "2030-01-01"
    }
}
```

---

## 📦 代码示例

### Python

```python
import requests

def search_answer(question: str, token: str = "free") -> dict:
    """查询网课答案"""
    url = "https://www.hive-net.cn/backend/course/search"
    payload = {
        "token": token,
        "question": question
    }
    
    try:
        response = requests.post(url, json=payload, timeout=10)
        response.raise_for_status()
        result = response.json()
        
        if result.get("code") == 0:
            print(f"✅ 找到 {len(result['data']['list'])} 条答案")
            return result["data"]["list"]
        else:
            print("❌ 未找到相关答案")
            return []
            
    except requests.RequestException as e:
        print(f"❌ 请求失败: {e}")
        return []

# 使用示例
if __name__ == "__main__":
    answers = search_answer("我国的国体是")
    for item in answers:
        print(f"问题: {item['question']}")
        print(f"答案: {item['answer']}")
        print(f"题型: {item['answer_type']}\n")
```

### JavaScript / Node.js

```javascript
const axios = require('axios');

async function searchAnswer(question, token = 'free') {
    try {
        const response = await axios.post('https://www.hive-net.cn/backend/course/search', {
            token,
            question
        });
        
        const { code, data } = response.data;
        if (code === 0) {
            console.log(`✅ 找到 ${data.list.length} 条答案`);
            return data.list;
        } else {
            console.log('❌ 未找到相关答案');
            return [];
        }
    } catch (error) {
        console.error('❌ 请求失败:', error.message);
        return [];
    }
}

// 使用示例
searchAnswer('我国的国体是').then(answers => {
    answers.forEach(item => {
        console.log(`问题: ${item.question}`);
        console.log(`答案: ${item.answer}\n`);
    });
});
```

---

## 🔑 获取 Token

### 方式一：免费试用

- Token: `free`
- 限制：每日10000次查询
- 有效期：永久

### 方式二：购买正式Token

👉 [自助购买 Token](https://c.fakamiao.top/shopDetail/yyN1OW?classifyId=1976)

---

## 📞 技术支持

| 渠道 | 联系方式 |
|------|--------|
| **QQ群** | [103172845](https://qm.qq.com/cgi-bin/qm/qr?k=sJLLnl1RdSdA5nhd7IXbhCxd-k3KaoBl&authKey=ssD9NFl2r5rHhGL4SvyIF56kSJi33zxFu2LqZ0XvUUGIZN3CyhCanNyji7cNXAwo&noverify=0&group_code=103172845#) |
| **微信公众号** | 极客范式 |

> 💡 **说明**：公众号提供教程、疑难问答、更新公告等完整支持

---

## ⚠️ 常见问题

**Q: 为什么提示"无答案"？**
A: 题目可能不在数据库中，或者搜索关键词不够准确。建议尝试更换关键词或使用模糊搜索。

**Q: 免费Token有次数限制吗？**
A: 有，每日10000次。超过限制需要购买正式Token或等待次日重置。

**Q: 支持哪些网课平台？**
A: 超星、智慧树、知到、Welearn、四史、马原、毛概、大学MOOC等主流平台。

---

## 📝 更新日志

- **v1.0.0** (2021-04-17)
  - ✨ 正式上线，支持多平台答案查询
  - ✨ 实现模糊搜索功能
  - ✨ 提供网页端和API接口

---

## 🙏 致谢

> 感谢您的使用和支持！如果觉得有帮助，请给项目一个⭐Star

---

<div align="center">
    <p><strong>Made with ❤️ by Hive Team</strong></p>
    <p>© 2024 All Rights Reserved</p>
</div>

