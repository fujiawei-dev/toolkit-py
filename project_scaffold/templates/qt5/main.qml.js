{{SLASH_COMMENTS}}

function httpRequestExample(params) {
    let router = "/debug"
    let url = "http://" + "localhost:8080" + router

    console.log("POST " + url)

    let request = new XMLHttpRequest()
    let responseBody

    request.onreadystatechange = function () {
        if (request.readyState === request.DONE) {
            console.log("response=" + request.responseText.toString())
            responseBody = JSON.parse(request.responseText.toString())
        }
    }

    let requestBody = JSON.stringify({
        "q": "JavaScript",
        "params": params,
    })

    console.log("request=" + requestBody)

    request.open("POST", url)
    request.setRequestHeader("Content-Type", "application/json")
    request.send(requestBody)
}
