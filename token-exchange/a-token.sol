// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract AToken is ERC20 {
    constructor() ERC20("AToken", "AT") {}

    function mint(address account, uint256 amount) public {
        _mint(account, amount);
    }
}
