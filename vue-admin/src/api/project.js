import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/project/list',
    method: 'get',
    params: query
  })
}

export function fetchProject(data) {
  return request({
    url: '/project/detail/' + data,
    method: 'get'
  })
}


export function createProject(data) {
  return request({
    url: '/project/create',
    method: 'post',
    data
  })
}

export function updateProject(data) {
  return request({
    url: '/project/update',
    method: 'put',
    data
  })
}

export function deleteProject(data) {
  return request({
    url: '/project/delete/' + data,
    method: 'delete',
    data
  })
}
