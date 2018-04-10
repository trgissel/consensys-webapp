const Car = artifacts.require("./contracts/Car.sol");
const web3 = require('web3-utils');

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
    let inEmmisionsId = "123abc"
    let inEmissionsIdb = web3.fromAscii(inEmmisionsId)
    let tx = await instance.setEmmisionInspectionsId(inEmissionsIdb)
    assert.ok(tx)
    let outEmissionsIdB = await instance.getEmmisionInspectionsId()
    let outEmissionsId = web3.toAscii(outEmissionsIdB)
    assert.ok(outEmissionsId.startsWith(inEmmisionsId))
  })


  it("should set setSafetyInspectionsId", async () => {
    let instance = await Car.deployed();
    let inSafteyId = "456def"
    let inSafteyIdHex = web3.fromAscii(inSafteyId)
    let tx = await instance.setSafetyInspectionsId(inSafteyIdHex)
    assert.ok(tx)
    let outSafteyIdHex = await instance.getSafetyInspectionsId()
    let outSafteyId = web3.toAscii(outSafteyIdHex)
    assert.ok(outSafteyId.startsWith(inSafteyId))
  })

  it("should set setScrapped", async () => {
    let instance = await Car.deployed();
    let isScrapped = await instance.getScrapped()
    assert.ok(!isScrapped)
    let tx = await instance.setScrapped(true)
    assert.ok(tx)
    isScrapped = await instance.getScrapped()
    assert.ok(isScrapped)
  })
})

