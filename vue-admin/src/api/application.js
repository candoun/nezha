import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/application/list',
    method: 'get',
    params: query
  })
}

export function fetchApplication(id) {
  return request({
    url: '/application/detail/' + id,
    method: 'get'
  })
}


export function createApplication(data) {
  return request({
    url: '/application',
    method: 'post',
    data
  })
}

export function updateApplication(data) {
  return request({
    url: '/application/update',
    method: 'post',
    data
  })
}
