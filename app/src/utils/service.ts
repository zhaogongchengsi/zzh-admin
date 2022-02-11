import axios from 'axios';
import type { AxiosInstance } from 'axios'
import { state } from "@/types/response"

const baseURL = "http://localhost:8000/api/v1/"
const http:AxiosInstance = axios.create({
  baseURL,
  timeout: 3000,
  // headers: {'X-Custom-Header': 'foobar'}
})


// 添加请求拦截器
http.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
http.interceptors.response.use(function (response) {
    // 对响应数据做点什么
    return response;
}, function (error) {
    // 对响应错误做点什么
    return Promise.reject(error);
});


export function Get (url: string) {
  return http.get(url)
}

export function Post (url: string, data: unknown) {
  return http.post(url, data)
}

export function confirmStatus (state: state) :boolean {
  if (state.code === 200) {
    return true
  }
  return false
}

export default http;