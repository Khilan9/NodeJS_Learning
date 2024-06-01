const middlewares = require("./middleware/index");
const path = require('path');
const express = require("express");
const app = express();
const port = 8000;
const urlRouter = require("./routes/urls");
const staticRoute = require('./routes/staticRouter');
const userRoute = require('./routes/user');
const cookieParser = require('cookie-parser');
const { MongoDBConnect } = require("./connections");
const { restrictToLoggedinUserOnly, checkAuth } = require("./middlewares/auth");

// Connect mongodb using below URL
MongoDBConnect("mongodb://127.0.0.1:27017/myurldb");

app.set('view engine','ejs');
app.set('views', path.resolve('./views'));

app.use(middlewares.useURLEncoder());
app.use(cookieParser());
app.use(express.json());

// User router
app.use('/',checkAuth,staticRoute);
app.use('/user',userRoute);
app.use("/url", restrictToLoggedinUserOnly, urlRouter);

app.listen(port, () =>
  console.log(`Shortning URL Service started on port ${port}!`)
);
