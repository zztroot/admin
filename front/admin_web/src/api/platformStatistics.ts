import { httpHeader, AxiosPromise } from "./api"

export interface platformStatisticsSetup {
    id?: number,
    platform_id: number,
    price: number,
    goods: string,
    comments?: string
}

// 获取数据分页
export let getPlatformStatisticsList = (page: string, page_size: string, author_id: number, platform_id: number, goods_name: string): AxiosPromise => {
    let url = `/platform-statistics/get-list?page=${page}&page_size=${page_size}`
    if (author_id != 0) {
        url = url + `&author_id=${author_id}`
    }
    if (platform_id != 0) {
        url = url + `&platform_id=${platform_id}`
    }
    if (goods_name != "") {
        url = url + `&goods_name=${goods_name}`
    }
    return httpHeader("GET", url, {}, {})
}

// 保存数据
export let savePlatformStatistics = (params: platformStatisticsSetup): AxiosPromise => {
    return httpHeader("POST", "/platform-statistics/save", params, {})
}

// 修改数据
export let modifyPlatformStatistics = (params: platformStatisticsSetup): AxiosPromise => {
    return httpHeader("POST", "/platform-statistics/modify", params, {})
}

// 删除数据
export let deletePlatformStatistics = (id: number): AxiosPromise => {
    let params = {
        id: id
    }
    return httpHeader("POST", "/platform-statistics/delete", params, {})
}