const fs = require('fs')

// Synchronous writing
fs.writeFileSync('./testdata.txt', 'Testing data');

// Asynchronous writing
// fs.writeFile('./testdata.txt', 'Testing data', (err) => {if(err)console.log(err);});

const readdata = fs.readFileSync('./testdata.txt', "utf-8")
console.log("Sync", readdata)

// fs.readFile('./testdata.txt',"utf-8", (err, succ) => {
//     if(err){
//         console.log(err);
//     }
//     if(succ){
//         console.log("Async", succ)
//     }
// })


// fs.appendFileSync('./testdata.txt', "\nSync append")

// fs.appendFile('./testdata.txt', "\nAsync append","utf-8",(err,succ) => {
//     if(err){
//         console.log(err);
//     }
//     if(succ){
//         console.log("Async", succ)
//     }
// })

// fs.cpSync('./testdata.txt', './copydata.txt')

// fs.unlinkSync('./copydata.txt')

// console.log(fs.statSync('./testdata.txt'))