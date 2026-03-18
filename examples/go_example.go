package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

type APIResponse struct {
    Code int `json:"code"`
    Data struct {
        List []map[string]interface{} `json:"list"`
    } `json:"data"`
}

func searchAnswer(question, token string) ([]map[string]interface{}, error) {
    payload := map[string]string{
        "token":    token,
        "question": question,
    }

    body, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    client := &http.Client{Timeout: 10 * time.Second}
    req, err := http.NewRequest(http.MethodPost, "https://www.hive-net.cn/backend/course/search", bytes.NewBuffer(body))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var result APIResponse
    if err := json.Unmarshal(responseBody, &result); err != nil {
        return nil, err
    }

    if result.Code != 0 {
        return []map[string]interface{}{}, nil
    }

    return result.Data.List, nil
}

func main() {
    answers, err := searchAnswer("我国的国体是", "free")
    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }

    fmt.Printf("Found %d answers\n", len(answers))
    for _, item := range answers {
        fmt.Println("Question:", item["question"])
        fmt.Println("Answer:", item["answer"])
        fmt.Println()
    }
}
