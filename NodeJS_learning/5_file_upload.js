const fs = require("fs");
const express = require("express");
const path = require('path');
const multer = require("multer");
const middlewares = require('./middleware/index');


const app = express();
const {MongoDBConnect} = require('./connections');
MongoDBConnect("mongodb://127.0.0.1:27017/mytestdb")
app.use(middlewares.useURLEncoder());
app.set('view engine','ejs');
app.set('views', path.resolve('./views'));

app.get('/', (req, res) => {
    return res.render("homepage")
})

const storage = multer.diskStorage({
    destination: function (req, file, cb) {
        return cb(null, "./uploads");
    },
    filename: function (req, file, cb) {
        return cb(null, `${Date.now()}-${file.originalname}`);
    }
})

const upload = multer({storage});

app.post('/upload', upload.single('profileimage'), (req, res) => {
    console.log(req.body);
    console.log(req.file);
    
    return res.redirect("/");
})
app.listen(8000, (req, res) => {
  console.log("Express server started");
});
