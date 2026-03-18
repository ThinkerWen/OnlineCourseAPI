async function searchAnswer(question, token = "free") {
    const response = await fetch("https://www.hive-net.cn/backend/course/search", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ token, question }),
    });

    if (!response.ok) {
        throw new Error(`Request failed with status ${response.status}`);
    }

    const result = await response.json();
    if (result.code !== 0) {
        return [];
    }

    return result.data?.list ?? [];
}

(async () => {
    try {
        const answers = await searchAnswer("我国的国体是");
        console.log(`Found ${answers.length} answers`);
        for (const item of answers) {
            console.log("Question:", item.question);
            console.log("Answer:", item.answer);
            console.log();
        }
    } catch (error) {
        console.error("Request failed:", error.message);
    }
})();
