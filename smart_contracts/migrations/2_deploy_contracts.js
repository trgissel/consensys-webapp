var Car = artifacts.require("Car");

module.exports = function(deployer) {
  deployer.deploy(Car, "BMW");
};