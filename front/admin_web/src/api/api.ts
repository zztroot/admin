import axios, { AxiosPromise, Method } from "axios";
import router from "@/router";
import { ElMessage } from "element-plus"
import { webConfig } from "../../config/config"

interface httpRequests {
    params: any,
    headers: any
}

export {
    http,
    httpHeader,
    httpRequests,
    AxiosPromise
}

let http = (method: Method, api: string, request: any): AxiosPromise => {
    if (method === "post") {
        return axios({
            method: method,
            url: webConfig.url + api,
            data: request
        })
    } else if (method === "get") {
        return axios({
            method: method,
            url: webConfig.url + api + request,
        })
    }
    console.log("参数错误")
    let a: any
    return a
}

interface resp {
    resp: any
}

let httpHeader = (method: Method, api: string, request: any, header: any): AxiosPromise => {
    header["X-TOKEN"] = JSON.parse(String(localStorage.getItem('user'))).token
    let temp = axios({
        url: webConfig.url + api,
        data: request,
        headers: header,
        method: method,
    })
    temp.then(response => {
        if (response.data.code !== 0) {
            if (response.data.code === 3007) {
                // token已经失效，前端删除缓存，并跳转到登录页
                localStorage.clear()
                ElMessage.error(response.data.message)
                setTimeout(() => {
                    router.push({
                        path: "/"
                    })
                }, 1000)
                return null
            } else {
                ElMessage.error(response.data.message)
            }
        }
    })
    return temp
}

