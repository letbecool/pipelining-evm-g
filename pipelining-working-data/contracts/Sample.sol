pragma solidity >=0.4.23;

contract Simple {
    uint256 public val1;
    uint256 public val2;
    
    constructor() {
        val2 = 3;
    }
    function set(uint256 _param) external {
        val1 = _param;
    }
}