import express from 'express';
import fs from 'fs';
import path from 'path';
import { renderToString } from 'react-dom/server';
import React from 'react';
import { StaticRouter } from 'react-router-dom/server';
import App from './src/App.tsx';

const app = express();

app.use(express.static(path.resolve(__dirname, 'dist')));

app.get('*', (req, res) => {
  const context = {};

  // Render the app to a string
  const appHtml = renderToString(
    <StaticRouter location={req.url} context={context}>
      <App />
    </StaticRouter>
  );

  // Read the index.html file
  const indexFile = path.resolve(__dirname, 'dist', 'index.html');
  fs.readFile(indexFile, 'utf8', (err, htmlData) => {
    if (err) {
      console.error('Error reading index.html', err);
      return res.status(500).send('Internal Server Error');
    }

    // Replace the placeholder with the rendered app
    const html = htmlData.replace('<div id="root"></div>', `<div id="root">${appHtml}</div>`);

    res.send(html);
  });
});

app.listen(3000, () => {
  console.log('Server is running on http://localhost:3000');
});
