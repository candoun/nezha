import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/403',
    component: () => import('@/views/403'),
    hidden: true
  },

  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [{
      path: 'dashboard',
      name: '首页',
      component: () => import('@/views/dashboard/index'),
      meta: { title: '首页', icon: 'dashboard' }
    }]
  },

  {
    path: '/application',
    component: Layout,
    redirect: '/application/list',
    name: '应用部署',
    meta: { title: '应用部署', icon: 'example' },
    children: [
      {
        path: 'list',
        name: 'Application',
        component: () => import('@/views/application/list'),
        meta: { title: '应用列表', icon: 'table' }
      },
      {
        path: 'create',
        hidden: true,
        component: () => import('@/views/application/create'),
        name: 'CreateApplicaiton',
        meta: { title: '新建Application', icon: 'create' }
      },
      {
        path: 'update/:id',
        hidden: true,
        component: () => import('@/views/application/update'),
        name: 'EditApplicaiton',
        meta: { title: '编辑Application', icon: 'edit' }
      },
      {
        path: 'detail/:id',
        hidden: true,
        component: () => import('@/views/application/detail'),
        name: 'DetailApplicaiton',
        meta: { title: '查看Application', icon: 'detail' }
      },
      {
        path: 'applog',
        name: 'AppLog',
        component: () => import('@/views/deploy/applog/index'),
        meta: { title: '部署日志', icon: 'tree' }
      }
    ]
  },
  {
    path: '/k8s',
    component: Layout,
    children: [
      {
        path: 'index',
        name: 'K8s集群',
        component: () => import('@/views/user/index'),
        meta: {
          title: 'K8s集群',
          icon: 'form'
        }
      }
    ]
  },
]

// 异步挂载的路由
// 动态需要根据权限加载的路由表
export const asyncRoutes = [
  {
    path: '/user',
    component: Layout,
    children: [
      {
        path: 'index',
        name: '用户管理',
        component: () => import('@/views/user/index'),
        meta: {
          title: '用户管理',
          icon: 'form',
          roles: ['admin']
        }
      }
    ]
  },
  {
    path: '/form',
    component: Layout,
    redirect: '/form/list',
    name: '文章管理',
    meta: {
      title: '文章管理',
      icon: 'form',
      roles: ['admin']
    },
    children: [
      {
        path: 'create',
        component: () => import('@/views/form/create'),
        name: 'CreateArticle',
        meta: { title: '新建文章', icon: 'edit', roles: ['admin'] }
      },
      {
        path: 'edit/:id(\\d+)',
        component: () => import('@/views/form/edit'),
        name: 'EditArticle',
        meta: { title: '编辑文章', noCache: true, activeMenu: '/form/list', roles: ['admin'] },
        hidden: true
      },
      {
        path: 'list',
        component: () => import('@/views/form/list'),
        name: 'ArticleList',
        meta: { title: '文章列表', icon: 'list', roles: ['admin'] }
      }
    ]
  },
  {
    path: '/project',
    component: Layout,
    redirect: '/project/list',
    name: '项目管理',
    meta: {
      title: '项目管理',
      icon: 'form',
      roles: ['admin']
    },
    children: [
      {
        path: 'list',
        name: 'Project',
        component: () => import('@/views/project/list'),
        meta: { title: '项目列表', icon: 'table' }
      },
      {
        path: 'create',
        hidden: true,
        component: () => import('@/views/project/create'),
        name: 'CreateProject',
        meta: { title: '新建Project', icon: 'create' }
      },
      {
        path: 'update/:id',
        hidden: true,
        component: () => import('@/views/project/update'),
        name: 'EditProject',
        meta: { title: '编辑Project', icon: 'edit' }
      }
    ]
  },
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
