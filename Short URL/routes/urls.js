const express = require("express");
const router = express.Router();
const {
  getShortURL,
  createShortURL,
  geturlvisitercount,
  testhtml,
} = require("../controllers/urls");


router.route("/test").get(testhtml);

// http://localhost:8000/url/Jz2Y7GTL4
router.route("/:shortid").get(getShortURL);

// http://localhost:8000/url/
// JSON data -> {"url": "https://www.youtube.com"}

router.route("/").post(createShortURL);

// http://localhost:8000/url/analytics/Jz2Y7GTL4
router.route("/analytics/:shortid").get(geturlvisitercount);



module.exports = router;

