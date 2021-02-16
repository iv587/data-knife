// if (process.env.NODE_ENV !== 'production') {
if (false) {
  console.log('mock mounting')
  const Mock = require('mockjs2')
  require('./service/redis')
  Mock.setup({
    timeout: 800 // setter delay time
  })
  console.log('mock mounted')
}
