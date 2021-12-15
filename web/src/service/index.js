import axios from "axios";
import { ElLoading } from "element-plus";

let loading = null;

const baseURL = "http://localhost:8000/api/v1/";

const Axios = axios.create({
  baseURL,
  timeout: 5000,
});

//  添加请求拦截
Axios.interceptors.request.use(
  function (config) {
    loading = ElLoading.service();
    return config;
  },
  function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
Axios.interceptors.response.use(
  function (response) {
    if (loading) {
      setTimeout(function () {
        loading.close();
      }, 200);
    }
    if (response.status !== 200) {
      return;
    }
    return response.data;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    loading.close();
    return Promise.reject(error);
  }
);

export function Get(url) {
  return Axios.get(url);
}

export function Post(url, data) {
  return Axios.post(url, data);
}
