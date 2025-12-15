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

var scriptBaseTemplate = `// SPDX-License-Identifier: MIT
pragma solidity [version];

import {Script} from "forge-std/Script.sol";
import {[name]} from "../src/[name].sol";

contract Deploy[name] is Script {
    function run() external {
        vm.startBroadcast();
        vm.stopBroadcast();
    }
}`

var Bases = map[string]string{
    "test": testBaseTemplate,
    "script": scriptBaseTemplate,
}
