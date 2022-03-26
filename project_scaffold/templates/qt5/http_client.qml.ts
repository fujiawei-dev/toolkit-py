{{SLASH_COMMENTS}}

class HttpClient {
    static get(
        url: string,
        callback: Function = responseText => {
            console.log(JSON.stringify(responseText))
        },
        params: object = {},
    ) {
        HttpClient.request(url, "GET", callback, params);
    }

    static post(
        url: string,
        callback: Function = responseText => {
            console.log(JSON.stringify(responseText))
        },
        params: object = {},
        data: object = {},
        json: object = {},
    ) {
        HttpClient.request(url, "POST", callback, params, data, json);
    }

    // I can't understand. Since all parameters must be passed, what is the meaning of default parameters?
    static request(
        url: string,
        method: string = "GET",
        callback: Function = responseText => {
            console.log(JSON.stringify(responseText));
        },
        params: object = {},
        data: object = {},
        json: object = {},
    ) {
        method = method.toUpperCase();

        let query: (string)[] = [];

        for (let key in params) {
            if (params.hasOwnProperty(key)) {
                query.push(key + "=" + params[key]);
            }
        }

        for (let key in data) {
            if (data.hasOwnProperty(key)) {
                query.push(key + "=" + data[key]);
            }
        }

        if (query.length > 0) {
            url += "?" + query.join("&");
        }

        if (url.substring(0, 4) !== "http") {
            url = "http://" + url;
        }

        console.log(method + " " + url);

        let client = new XMLHttpRequest();

        client.onreadystatechange = () => {
            if (client.readyState === client.DONE) {
                let responseText = client.responseText.toString();
                console.log("response=" + responseText);
                callback(responseText);
            }
        }

        client.open(method, url, true);

        if (method !== "GET") {
            let requestBody = JSON.stringify(json);
            if (requestBody !== "{}") {
                client.setRequestHeader("Content-Type", "application/json");
                console.log("request=" + requestBody);
                client.send(requestBody);
            } else {
                client.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
                client.send();
            }
        } else {
            client.send()
        }
    }
}
