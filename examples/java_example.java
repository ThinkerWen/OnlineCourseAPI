import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;

public class java_example {
    public static void main(String[] args) {
        String token = "free";
        String question = "我国的国体是";
        String requestBody = "{\"token\":\"" + token + "\",\"question\":\"" + question + "\"}";

        HttpClient client = HttpClient.newBuilder()
                .connectTimeout(Duration.ofSeconds(10))
                .build();

        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("https://www.hive-net.cn/backend/course/search"))
                .timeout(Duration.ofSeconds(10))
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(requestBody))
                .build();

        try {
            HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
            System.out.println("Status: " + response.statusCode());
            System.out.println("Response:");
            System.out.println(response.body());
        } catch (IOException | InterruptedException e) {
            System.err.println("Request failed: " + e.getMessage());
            Thread.currentThread().interrupt();
        }
    }
}
