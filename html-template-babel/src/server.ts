// server.ts
import fs from 'fs';
import path from 'path';
import React from 'react';
import ReactDOMServer from 'react-dom/server';
import App from './App';

const htmlTemplate = fs.readFileSync(path.resolve(__dirname, '../src/index.html'), 'utf8');

const appHtml = ReactDOMServer.renderToString(React.createElement(App));

const finalHtml = htmlTemplate.replace('<!-- APP -->', appHtml);

fs.writeFileSync(path.resolve(__dirname, 'dist/index.html'), finalHtml);

console.log('Static HTML generated successfully!');
