const path = require('path');

module.exports = {
  mode: 'development',
  devtool: 'source-map',
  entry: './src/index.js',
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'main.js',
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/, // .js와 .jsx 파일을 처리
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
          options: {
            presets: ['@babel/preset-env', '@babel/preset-react'], // Babel 프리셋 설정
          },
        },
      },
      {
        test: /\.css$/, // CSS 파일 처리
        use: ['style-loader', 'css-loader', 'postcss-loader'],
      },
    ],
  },
  resolve: {
    extensions: ['.js', '.jsx'], // .js와 .jsx 확장자 처리
  },
  devServer: {
    static: {
      directory: path.join(__dirname, 'public'),
    },
    port: 8082,
    hot: true,
    client: {
      overlay: false,
    },
  },
};