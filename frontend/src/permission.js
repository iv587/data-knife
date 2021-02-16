import router from './router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { getToken } from '@/utils/auth'


NProgress.configure({ showSpinner: false })
router.beforeEach(async(to, from, next) => {
  NProgress.start()
  if (to.path !== '/login' && !getToken()) {
    next({ path: '/login' })
    NProgress.done()
  }else  {
    if (to.path === '/login' && getToken()) {
      next({ path: '/' })
      NProgress.done()
    }
    next()
  }
})
router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
