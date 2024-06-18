const fs = require("fs");
const express = require("express");

const app = express();
const userRouter = require('./routes/user');
const middlewares = require('./middleware/index');
const { MongoDBConnect } = require('./connections')

// Connect mongodb using below URL
MongoDBConnect("mongodb://127.0.0.1:27017/mytestdb")

//Use Middlewares
app.use(middlewares.useURLEncoder());
app.use(middlewares.Middleware1);
app.use(middlewares.Middleware2);

// User router
app.use("/user", userRouter);


app.listen(8000, (req, res) => {
  console.log("Express server started");
});
