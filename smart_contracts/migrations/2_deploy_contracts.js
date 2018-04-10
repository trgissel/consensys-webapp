var SimpleStorage = artifacts.require("Car");

module.exports = function(deployer) {
  deployer.deploy(SimpleStorage, "BMW");
};