const express = require("express");

function useURLEncoder() {
  console.log("In url encode");
  return express.urlencoded({ extended: false });
}

module.exports = {
  useURLEncoder,
};
