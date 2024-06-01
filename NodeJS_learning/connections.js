const mongoose = require("mongoose");

mongoose.set("debug", true);

async function MongoDBConnect(url) {
  return mongoose
    .connect(url, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    })
    .then(() => {
      console.log("Connected to MongoDB");
    })
    .catch((err) => {
      console.error("Error connecting to MongoDB", err);
    });
}

module.exports = {
    MongoDBConnect,
}