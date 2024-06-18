const http = require('http')
const fs = require('fs')
const url = require('url');

const myserver = http.createServer((req, res) => {
    // console.log("Req recieved");
    const myUrl = url.parse(req.url, true)
    console.log(myUrl)
    // console.log(myUrl.pathname)
    const logadder = `${Date.now()}: ${req.url} Req Recieved\n`;
    fs.appendFile('./reqfile.log', logadder, "utf-8", (err, suc) => {
        if (err) {
            console.log(err)
        }
        else {
            switch (myUrl.pathname) {
                case "/":
                    // localhost:8000/
                    console.log(myUrl.query)
                    res.end("Homepage")
                    break;
                case "/about":
                    console.log(req.method)
                    console.log(myUrl.query.q)
                    console.log(myUrl.query.r)
                    // Use localhost:8000/about?q=abc&r=xyz                
                    if (req.method == "GET") {
                        res.end("About page get")
                    }
                    else if (req.method == "POST") {
                        res.end("About page post")
                    }
                    break;
                default:
                    res.end("404 not found")
                    break;
            }
        }
    })
});


myserver.listen(8000, () => { console.log("Server started") })