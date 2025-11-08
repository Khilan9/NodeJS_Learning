const { Schema, model } = require('mongoose')
const { randomBytes, createHmac } = require('node:crypto');
const { createToken } = require('../services/authentication');

const userSchema = new Schema({
    fullName: {
        type: String,
        required: true
    },
    email: {
        type: String,
        required: true,
        unique: true
    },
    salt: {
        type: String
    },
    password: {
        type: String,
        required: true
    },
    profileImageUrl: {
        type: String,
        default: "/images/avatar.jpg"
    },
    role: {
        type: String,
        enum: ["USER", "ADMIN"],
        default: "USER"
    }
}, { timestamps: true });

userSchema.pre("save", function (next) {
    const user = this
    console.log("Here model");

    if (!user.isModified("password")) return;

    // const salt = "randomstr";
    const salt = randomBytes(16).toString();

    const hasedpass = createHmac("sha256", salt).update(user.password).digest("hex")
    this.salt = salt;
    this.password = hasedpass;
    next();
})

userSchema.static('matchPasswordandgenerateToken', async function (email, password) {
    console.log("Here called");
    console.log(email);
    console.log(password);
    const user = await this.findOne({ email });
    if (!user) return false;
    const salt = user.salt;
    // const salt = "randomstr";
    const hashedPasswd = user.password;
    const userProvidedhash = createHmac("sha256", salt).update(password).digest("hex");
    if (hashedPasswd !== userProvidedhash) throw new Error("Incorrect pass");
    const token = createToken(user);
    return token;
})

const User = model('user', userSchema);


module.exports = User;