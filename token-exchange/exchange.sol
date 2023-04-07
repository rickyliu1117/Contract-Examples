// SPDX-License-Identifier: MIT

pragma solidity ^0.8.2;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Exchange is Ownable {
    address token0;
    address token1;

    uint112 private reserve0;
    uint112 private reserve1;

    uint private token0Num;
    uint private token1Num;

    bytes4 private constant SELECTOR =
        bytes4(keccak256(bytes("transfer(address,uint256)")));

    event Swap(
        address indexed sender,
        uint amount0In,
        uint amount1In,
        uint amount0Out,
        uint amount1Out,
        address indexed to
    );

    event Sync(uint112 reserve0, uint112 reserve1);

    constructor(address _token0, address _token1) {
        token0 = _token0;
        token1 = _token1;
    }

    function setExchangeRate(
        uint256 _token0Num,
        uint256 _token1Num
    ) public onlyOwner {
        require(
            _token0Num > 0 && _token1Num > 0,
            "Exchange: token0Num and token1Num must be greater than 0"
        );
        token0Num = _token0Num;
        token1Num = _token1Num;
    }

    function getReserves()
        public
        view
        returns (uint112 _reserve0, uint112 _reserve1)
    {
        _reserve0 = reserve0;
        _reserve1 = reserve1;
    }

    function queryExchangeRate() public view returns (uint256, uint256) {
        return (token0Num, token1Num);
    }

    function _safeTransfer(address token, address to, uint value) private {
        (bool success, bytes memory data) = token.call(
            abi.encodeWithSelector(SELECTOR, to, value)
        );
        require(
            success && (data.length == 0 || abi.decode(data, (bool))),
            "Exchange: transfer failed"
        );
    }

    function _update(uint balance0, uint balance1) private {
        reserve0 = uint112(balance0);
        reserve1 = uint112(balance1);
        emit Sync(reserve0, reserve1);
    }

    function updateExchangePool() public onlyOwner {
        uint256 balance0 = IERC20(token0).balanceOf(address(this));
        uint256 balance1 = IERC20(token1).balanceOf(address(this));
        _update(balance0, balance1);
    }

    function swap(address to) external {
        (uint112 _reserve0, uint112 _reserve1) = getReserves();
        uint balance0;
        uint balance1;
        uint256 amount0Out;
        uint256 amount1Out;

        balance0 = IERC20(token0).balanceOf(address(this));
        balance1 = IERC20(token1).balanceOf(address(this));
        uint amount0In = balance0 - _reserve0;
        uint amount1In = balance1 - _reserve1;

        require(
            amount0In > 0 || amount1In > 0,
            "Exchange: insufficient input amount"
        );
        if (amount0In > 0) {
            amount1Out = (amount0In * token1Num) / token0Num;
            require(
                amount1Out <= balance1,
                "Exchange: all token1 of the contract is not enough to provide exchange"
            );
            _safeTransfer(token1, to, amount1Out);
        }
        if (amount1In > 0) {
            amount0Out = (amount1In * token0Num) / token1Num;
            require(
                amount0Out <= balance0,
                "Exchange: all token0 of the contract is not enough to provide exchange"
            );
            _safeTransfer(token0, to, amount0Out);
        }
        balance0 = IERC20(token0).balanceOf(address(this));
        balance1 = IERC20(token1).balanceOf(address(this));
        _update(balance0, balance1);
        emit Swap(msg.sender, amount0In, amount1In, amount0Out, amount1Out, to);
    }
}
