const express = require("express");
const router = express.Router();
const User = require("../models/user");
const { GetAllUsersApi, createUserAPI, handlePutUserByID, handleGetUserByID, handleDeleteUserByID } = require("../controllers/user");

router.get("/", async (req, res) => {
  const allbdusers = await User.find({});
  const htmldata = `
      <ul>
      ${allbdusers.map((user) => `<li>${user.fname}</li>`)}
      `;
  return res.send(htmldata);
});

router.route("/api")
        .get(GetAllUsersApi)
        .post(createUserAPI);

router
  .route("/api/:id")
  .get(handleGetUserByID)
  .patch(handlePutUserByID)
  .delete(handleDeleteUserByID);

module.exports = router;
