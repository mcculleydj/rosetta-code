module.exports = {
  configureWebpack: {
    module: {
      rules: [
        {
          test: /\.txt$/,
          loader: 'raw-loader',
        },
      ],
    },
  },
}
