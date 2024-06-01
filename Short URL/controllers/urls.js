const URL = require("../models/urls");
const nanoid = require("shortid");

async function getShortURL(req, res) {
  console.log("getShortURL req obj");
  console.log(req);
  const shortID = req.params.shortid;
  console.log(shortID);
  const resp = await URL.findOneAndUpdate(
    { shortID },
    { $inc: { totalReq: 1 } }
  );
  console.log(`In get short url ${resp}`);
  res.redirect(resp.redirectURL);
}

async function createShortURL(req, res) {
  console.log("Here creating short url");
  const body = req.body;
  const shortID = nanoid(8);
  console.log(`generated shortid ${shortID}`);
  if (!body || !body.url) {
    return res.status(400).json({ error: "url is required" });
  }
  console.log(req);
  const result = await URL.create({
    shortID: shortID,
    redirectURL: body.url,
    createdBy: req.user._id
  });
  return res.render("home", {"id": shortID});
}

async function geturlvisitercount(req, res) {
  shortid = req.params.shortID;
  console.log(req.params.shortID);
  const resp = await URL.findOne({ shortid });
  console.log(`In get short url ${resp}`);
  return res.status(201).json({ count: resp.totalReq });
}

async function testhtml(req, res) {
  console.log("here in test function")
  const allurls = await URL.find({});
  console.log(allurls);
  return res.render("listurls", {
    urls: allurls
  });
}

module.exports = {
  getShortURL,
  createShortURL,
  geturlvisitercount,
  testhtml,
};
