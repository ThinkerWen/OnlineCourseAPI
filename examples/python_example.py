import json
import urllib.error
import urllib.request


def search_answer(question: str, token: str = "free") -> list:
    url = "https://www.hive-net.cn/backend/course/search"
    payload = {
        "token": token,
        "question": question,
    }

    request = urllib.request.Request(
        url,
        data=json.dumps(payload).encode("utf-8"),
        headers={"Content-Type": "application/json"},
        method="POST",
    )

    try:
        with urllib.request.urlopen(request, timeout=10) as response:
            result = json.loads(response.read().decode("utf-8"))
    except urllib.error.URLError as err:
        print(f"Request failed: {err}")
        return []

    if result.get("code") != 0:
        print("No result found")
        return []

    return result.get("data", {}).get("list", [])


if __name__ == "__main__":
    answers = search_answer("我国的国体是")
    print(f"Found {len(answers)} answers")
    for item in answers:
        print(f"Question: {item.get('question', '')}")
        print(f"Answer: {item.get('answer', '')}")
        print()
