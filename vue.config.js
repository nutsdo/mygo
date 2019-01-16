module.exports = {
  outputDir: 'public',
  pluginOptions: {
    'style-resources-loader': {
      preProcessor: 'scss',
      patterns: []
    }
  },
  css: {
      loaderOptions: {
          css: {
              localIdentName: '[name]-[hash]',
              camelCase: 'only'
          }
      }
  }
}
