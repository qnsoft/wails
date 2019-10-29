/* eslint-disable */

module.exports = function (api) {
  api.cache(true);

  const presets = [
    [
      "@babel/preset-env",
      {
        "useBuiltIns": "usage",
        "corejs": {
          "version": 3,
          "proposals": true
        }
      }
    ]
  ];

  const plugins = ["@babel/plugin-transform-object-assign", "transform-es2015-shorthand-properties"]

  return {
    presets,
    plugins
  };
}