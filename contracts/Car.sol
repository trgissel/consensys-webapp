pragma solidity ^0.4.0;
//Style Guide: http://solidity.readthedocs.io/en/v0.4.21/style-guide.html?highlight=style
contract Car {

    bytes32 _manufacture;

    struct Owners {
        bytes32 primaryName;
        bytes32 secondaryName;
        address primaryAddress;
    }
    
    struct Inspections {
        bytes32 safetyId;
        bytes32 emmissionsId;
    }
    
    Inspections _inspections;
    Owners _owners;
    bool _scrapped;
    uint64 _miles;
    
    /// Create a new car with manufacture
    function Car(bytes32 inManufature) public {
        _manufacture = inManufature;
        _owners.primaryAddress = msg.sender;
    }
    
    function changeOwners(bytes32 inPrimaryName, bytes32 inSecondaryName, address inOwnerAddress) public {
        if (msg.sender != _owners.primaryAddress) {
             //only owner can change owners
            return;
         }
        _owners.primaryName = inPrimaryName;
        _owners.secondaryName = inSecondaryName;
        _owners.primaryAddress = inOwnerAddress;
    }
    
    function setEmmisionInspectionsId(bytes32 id) public {
        if (msg.sender != _owners.primaryAddress) {
             //only owner can change owners
            return;
         }
         
        _inspections.emmissionsId = id;
    }
    
    function setSafetyInspectionsId(bytes32 id) public {
        if (msg.sender != _owners.primaryAddress) {
             //only owner can change owners
            return;
         }
         
        _inspections.safetyId = id;
    }

    function setScrapped(bool isScrapped) public {
        if (msg.sender != _owners.primaryAddress) {
             //only owner can change owners
            return;
         }
         
        _scrapped = isScrapped;
    }

    function addMiles(uint32 milesToAdd) public {
        _miles = _miles + milesToAdd;
    }
    
    function getMiles() public view
            returns (uint64 miles){
        miles = _miles;
    }
}
