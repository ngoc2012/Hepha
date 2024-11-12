"use strict";

var _fs = _interopRequireDefault(require("fs"));
var _path = _interopRequireDefault(require("path"));
var _react = _interopRequireDefault(require("react"));
var _server = _interopRequireDefault(require("react-dom/server"));
var _App = _interopRequireDefault(require("./App"));
function _interopRequireDefault(e) { return e && e.__esModule ? e : { "default": e }; }
// server.ts

var htmlTemplate = _fs["default"].readFileSync(_path["default"].resolve(__dirname, '../src/index.html'), 'utf8');
var appHtml = _server["default"].renderToString(/*#__PURE__*/_react["default"].createElement(_App["default"]));
var finalHtml = htmlTemplate.replace('<!-- APP -->', appHtml);
_fs["default"].writeFileSync(_path["default"].resolve(__dirname, 'dist/index.html'), finalHtml);
console.log('Static HTML generated successfully!');