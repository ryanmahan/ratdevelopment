var webpack = require('webpack');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
    entry: {main: "./src/index.tsx"},
    output: {
        filename: "[name].[chunkhash].js",
        path: __dirname + "/dist",
        // Location where webpack will put its compiled assets (JS, CSS etc.)
        publicPath: "/"
    },
    plugins: [
        new webpack.ProvidePlugin({
            "React": "react",
            "ReactDOM" : "react-dom",
        }),
        new MiniCssExtractPlugin({
            // Options similar to the same options in webpackOptions.output
            // both options are optional
            filename: 'style.[contenthash].css'
        }),
        new HtmlWebpackPlugin({
            inject: false,
            hash: true,
            template: './src/index.html',
            filename: 'index.html'
        })
    ],
    // Enable sourcemaps for debugging webpack's output.
    devtool: "source-map",

    resolve: {
        // Add '.ts' and '.tsx' as resolvable extensions.
        extensions: [".ts", ".tsx", ".js", ".json"]
    },

    module: {
        rules: [
            // All files with a '.ts' or '.tsx' extension will be handled by 'awesome-typescript-loader'.
            { test: /\.tsx?$/, loader: "awesome-typescript-loader" },

            // All output '.js' files will have any sourcemaps re-processed by 'source-map-loader'.
            { enforce: "pre", test: /\.js$/, loader: "source-map-loader" },

            // { test: /\.json$/, loader: "json-loader"},

            {
                test: /\.(sa|sc|c)ss$/,
                use: [
                    'style-loader',
                    MiniCssExtractPlugin.loader,
                    'css-loader',
                    "postcss-loader",
                    'sass-loader',
                ],
            }
        ]
    },



    // When importing a module whose path matches one of the following, just
    // assume a corresponding global variable exists and use that instead.
    // This is important because it allows us to avoid bundling all of our
    // dependencies, which allows browsers to cache those libraries between builds.
    // externals: {
    //     "react": "React",
    //     "react-dom": "ReactDOM"
    // },

    devServer: {
        contentBase: './dist',
        compress: true,
        host: '0.0.0.0',
        port: 8080,
        // React routing will work relative to index.html, so we set "historyApiFallback" 
        // to true, so we use index.html when finding views on the front-end.
        historyApiFallback: true
    }
};
