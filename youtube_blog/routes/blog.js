const { Router } = require('express');
const Blog = require('../models/blog');
const multer = require('multer');
const router = Router();
const path = require('path');
const Comment = require('../models/comment');

const storage = multer.diskStorage({
    destination: function (req, file, cb) {
        cb(null, path.resolve(`./public/uploads/`));
    },
    filename: function (req, file, cb) {
        const fileName = `${Date.now()}-${file.originalname}`;
        cb(null, fileName);
    },
});

const upload = multer({ storage: storage });

router.get("/add-new", (req, res) => {
    return res.render("addblog", {
        user: req.user,
    });
});

router.get("/:id", async (req, res) => {
    const blog = await Blog.findById(req.params.id).populate('createdBy');
    const comments = await Comment.find({ blogid: req.params.id }).populate(
        "createdBy"
      );
    console.log("blog logs", blog);
    console.log("comments logs", comments);
    return res.render("blog", {
        user: req.user,
        blog: blog,
        comments: comments
    });
});

router.post("/", upload.single("coverImage"), async (req, res) => {
    console.log(req.body);
    const { title, body } = req.body;
    const blog = await Blog.create({
        body,
        title,
        createdBy: req.user._id,
        coverImageURL: `/uploads/${req.file.filename}`,
    });
    return res.redirect(`/blog/${blog._id}`);
});

router.post("/comment/:blogid", async (req, res) => {
    console.log(req.body);
    const blogid = req.params.blogid
    const { content } = req.body;
    const blog = await Comment.create({
        content,
        blogid,
        createdBy: req.user._id,
    });
    return res.redirect(`/blog/${blogid}`);
});


module.exports = router;