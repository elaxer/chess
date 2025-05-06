const path = require("path");
const autoprefixer = require('autoprefixer')

module.exports = {
    watch: true,
    mode: "development", // "production" для финальной сборки
    entry: "./src/ts/app.ts",
    output: {
        filename: "bundle.js",
        path: path.resolve(__dirname, "./public/dist"),
    },
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: "ts-loader",
                exclude: /node_modules/,
            },
            {
                test: /\.(scss)$/,
                use: [
                  {
                    // Adds CSS to the DOM by injecting a `<style>` tag
                    loader: 'style-loader'
                  },
                  {
                    // Interprets `@import` and `url()` like `import/require()` and will resolve them
                    loader: 'css-loader'
                  },
                  {
                    // Loader for webpack to process CSS with PostCSS
                    loader: 'postcss-loader',
                    options: {
                      postcssOptions: {
                        plugins: [
                          autoprefixer
                        ]
                      }
                    }
                  },
                  {
                    // Loads a SASS/SCSS file and compiles it to CSS
                    loader: 'sass-loader'
                  }
                ]
              }
        ],
    },
    resolve: {
        extensions: [".ts", ".js"], // Позволяет импортировать файлы без расширения
    },
    devServer: {
        static: "./public/dist",
        hot: true,
        open: true,
    },
};