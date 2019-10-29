/* eslint-disable */

const path = require('path');

module.exports = {
  entry: './bridge.js',
  mode: 'production',
  output: {
    path: path.resolve(__dirname, '..', 'assets'),
    filename: 'bridge.js',
    libraryTarget: 'this'
  },
  module: {
    rules: [
      {
        test: /\.m?js$/,
        exclude: /(node_modules|bower_components)/,
        use: {
          loader: 'babel-loader',
          options: {
            plugins: [],
            presets: [
              [
                '@babel/preset-env',
                {
                  'useBuiltIns': 'usage',
                  'corejs': 3,
                  'targets': { "browsers": ["last 2 versions", "ie >= 9"] }
                }
              ]
            ]
          }
        }
      }
    ]
  }
};
