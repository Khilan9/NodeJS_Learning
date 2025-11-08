const jwt = require('jsonwebtoken')
const secret = 'test'

function setUser(user){
    return jwt.sign({_id: user._id, email: user.email, role:user.role}, secret);
}

function getUser(token){
    // return sessionIDtousermap.get(id);
    try{
        return jwt.verify(token, secret)
    }
    catch(error){
        return null
    }
}

module.exports = {
    setUser,
    getUser
}