const fs = require("fs");
const express = require("express");
const path = require('path');
const status = require('express-status-monitor');
const zlib = require('zlib');
const app = express();
app.use(status());

app.set('view engine', 'ejs');
app.set('views', path.resolve('./views'));

app.get('/', (req, res) => {
    // Goto http://localhost:8000/status to check mem usage
    const stream = fs.createReadStream('./uploads/50MiB.txt', 'utf-8');
    stream.on('data', (chunk) => res.write(chunk));
    stream.on('end', () => res.end());

    // create zip using stream instead of loading all data into varaiable at once
    fs.createReadStream('./uploads/50MiB.txt').pipe(zlib.createGzip().pipe(fs.createWriteStream('./uploads/sample.zip')));
    // fs.readFile('./uploads/50MiB.txt', (err, data) => {
    //     res.end(data);
    // })
})


app.listen(8000, (req, res) => {
    console.log("Express server started at port 8000");
});
