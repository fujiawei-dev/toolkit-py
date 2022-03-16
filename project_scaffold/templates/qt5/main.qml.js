{{SLASH_COMMENTS}}

Qt.include("http_client.js")

function httpGetExample() {
    HttpClient.get(
        "http://httpbin.org/get",
        function (responseText) {
            console.log(responseText);
        },
        {"q": "typescript"}
    )
}

function httpPostExample() {
    HttpClient.post(
        "http://httpbin.org/post",
        function (responseText) {
            console.log(responseText);
        },
        {"q": "typescript"},
        {"image": "base64", "debug": debugMode},
        {"json": "This is a json object."}
    )
}
