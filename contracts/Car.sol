pragma solidity ^0.4.0;
contract Car {

    string manufacture;
    // probably want a structs for these things
    string emissionscertification;
    string[] owners;
    string[] inspections;
    string[] accidents;
    bool scrapped;

    /// Create a new car with manufacture
    function Car(string _manufature) public {
        manufacture = _manufature;
    }
}
