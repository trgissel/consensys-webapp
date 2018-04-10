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

  it("should set setEmmisionInspectionsId", async () => {
    let instance = await Car.deployed();
    let tx = await instance.setEmmisionInspectionsId("123abc");
    assert.ok(tx)
  })


  it("should set setSafetyInspectionsId", async () => {
    let instance = await Car.deployed();
    let tx = await instance.setSafetyInspectionsId("456def");
    assert.ok(tx)
  })

  it("should set setScrapped", async () => {
    let instance = await Car.deployed();
    let tx = await instance.setScrapped(true)
    assert.ok(tx)
  })
})

