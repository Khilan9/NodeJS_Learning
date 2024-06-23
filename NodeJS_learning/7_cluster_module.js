
const express = require("express");
const path = require('path');
const app = express();
const os = require('os');
const cluster = require('cluster');
const process = require('node:process');

app.set('view engine', 'ejs');
app.set('views', path.resolve('./views'));


const totalcpu = os.cpus().length;
console.log(os.availableParallelism());
console.log(totalcpu);

if (cluster.isPrimary) {
    console.log(`Primary ${process.pid} is running`);

    // Fork workers.
    for (let i = 0; i < totalcpu; i++) {
        cluster.fork();
    }

    cluster.on('exit', (worker, code, signal) => {
        console.log(`worker ${worker.process.pid} died`);
    });
}
else{
    app.get('/', (req, res) => {
        res.send(`App is running with process id ${process.pid}`)
    })
    
    app.listen(8000, (req, res) => {
        console.log("Express server started at port 8000");
    });
}




