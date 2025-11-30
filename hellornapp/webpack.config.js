import path from 'path';
import { StyleXPlugin } from 'stylex-webpack';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import MiniCssExtractPlugin from 'mini-css-extract-plugin';

export default {
  mode: 'production',
  entry: './index_web.js',
  output: {
    filename: 'bundle.js',
    path: path.resolve('dist'),
  },
  devServer: {
    historyApiFallback: true,
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: [
          {
            loader: 'babel-loader',
            options: {
              targets: 'defaults',
              comments: false,
              plugins: [
                'transform-class-properties',
                [
                  'module-resolver',
                  {
                    alias: {
                      '@tabler/icons-react-native': '@tabler/icons-react',
                      '^react-native$': 'react-native-web',
                      '@expo-google-fonts/inter': './empty.js',
                      'expo-font': './empty.js',
                    },
                  },
                ],
              ],
              presets: [
                '@babel/preset-react',
                [
                  '@babel/preset-env',
                  {
                    targets: {
                      browsers: 'last 2 versions',
                    },
                    modules: false,
                    loose: false,
                  },
                ],
              ],
            },
          },
        ],
      },
      {
        test: /\.css$/,
        use: [MiniCssExtractPlugin.loader, 'css-loader', 'postcss-loader'],
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: './index.html',
    }),

    new StyleXPlugin({
      // stylex-webpack options goes here, see the following section for more details
    }),
    new MiniCssExtractPlugin(),
  ],
};
