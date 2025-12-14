package base

var testBaseTemplate = `// SPDX-License-Identifier: MIT
pragma solidity [version];

import {Test, console} from "forge-std/Test.sol";
import {[name]} from "../src/[name].sol";

contract [name]Test is Test {
    [name] [name_variable];

    function setUp() external {
        [name_variable] = new [name]();
    }
}
`

var Bases = map[string]string{
    "test": testBaseTemplate,
}
