const User = require("../models/user");

async function GetAllUsersApi(req, res) {
  // X stats that this is the custom header
  const allusers = await User.find({});
  res.setHeader("X-myvar", "testvalue");
  return res.json(allusers);
}

async function createUserAPI(req, res) {
  const body = req.body;
  if (!body || !body.fname) {
    console.log("Here checking body");
    return res.status(400).json({ msg: "All fields are required" });
  }
  // users.push({ ...body, id: users.length + 1 });
  // fs.writeFile("./MOCK_DATA.json", JSON.stringify(users), (err, data) => {
  //   console.log(err);
  //   console.log(data);
  // });
  console.log(body);
  const result = await User.create({
    fname: body.fname,
  });
  console.log(result);
  return res.status(201).json({ status: "success", id: result._id });
}


async function handleGetUserByID(req, res) {
  const userbyid = await User.findById(req.params.id);
  // const user = users.find((user) => user.id == id);
  return res.send(userbyid);
}

async function handlePutUserByID(req, res) {
  await User.findByIdAndUpdate(req.params.id, { fname: "changing static" });
  return res.json({ msg: "patch done" });
}

async function handleDeleteUserByID(req, res) {
  await User.findByIdAndDelete(req.params.id, { fname: "changing static" });
  return res.json({ msg: "delete done" });
}

module.exports = {
  GetAllUsersApi,
  createUserAPI,
  handleGetUserByID,
  handlePutUserByID,
  handleDeleteUserByID
};
