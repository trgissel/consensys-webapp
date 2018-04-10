const Car = artifacts.require("./contracts/Car.sol");

contract('Car test', async (accounts) => {

  it("should get miles", async () => {
    let instance = await Car.deployed();
    let miles = await instance.getMiles();
    assert.equal(miles, 0);
  })

  it("should add miles", async () => {
    let instance = await Car.deployed();
    let beforeMiles = await instance.getMiles();
    assert.ok(beforeMiles > -1);
    let tr = await instance.addMiles(5);
    let afterMiles = await instance.getMiles();
    assert.ok(afterMiles > beforeMiles);
  })

})

