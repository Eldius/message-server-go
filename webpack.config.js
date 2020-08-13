
//const webpack = require("webpack");
const path = require("path")

if (process.env.WEBPACK_MODE === 'production') {
  
}

module.exports = () => ({
  entry: "./static/js/main.js",
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "static/bundle"),
    publicPath: "static/bundle"
  },
  module: {
    rules: [{
      test: /\.css$/,
      loader: 'style-loader!css-loader'
    },
    { 
        test: /\.(woff|woff2)(\?v=\d+\.\d+\.\d+)?$/, 
        loader: 'url-loader?limit=10000&mimetype=application/font-woff' 
    },
    { 
        test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/, 
        loader: 'url-loader?limit=10000&mimetype=application/octet-stream'
    },
    { 
        test: /\.eot(\?v=\d+\.\d+\.\d+)?$/, 
        loader: 'file-loader' 
    },
    { 
        test: /\.svg(\?v=\d+\.\d+\.\d+)?$/, 
        loader: 'url-loader?limit=10000&mimetype=image/svg+xml' 
    },
    {
        test: /\.(png|svg|jpg|gif)$/,
        use: [
          'file-loader',
        ],
    }]
  }
});
