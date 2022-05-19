import { httpHeader, AxiosPromise } from "./api"

export let getGoodslist = (page: string, page_size: string, author_id: number, warehouse_id: number, start_date: string, end_date: string, name: string): AxiosPromise => {
    let url = `/goods/get-list?page=${page}&page_size=${page_size}`
    if (author_id != 0) {
        url = url + `&author_id=${author_id}`
    }
    if (warehouse_id != 0) {
        url = url + `&warehouse_id=${warehouse_id}`
    }
    if (start_date != "") {
        url = url + `&start_date=${start_date}&end_date=${end_date}`
    }
    if (name != "") {
        url = url + `&name=${name}`
    }
    return httpHeader("GET", url, {}, {})
}

// 定义保存商品参数类型
export interface SaveGoodsParams {
    warehouse_id: number,
    cost_price: number,
    supply_price: number,
    sale_number: number,
    return_number: number,
    link_return_number: number,
    name: string,
    date: string,
    unit: string
    id: number,
}

// 保存
export let saveGoods = (params: SaveGoodsParams): AxiosPromise => {
    return httpHeader("POST", "/goods/save", params, {})
}

// 定义修改商品参数类型
export interface ModifyGoodsParams {
    id: number,
    warehouse_id: number,
    cost_price: number,
    supply_price: number,
    sale_number: number,
    return_number: number,
    link_return_number: number,
    name: string,
    date: string,
    unit: string
}

// 修改
export let modifyGoods = (params: SaveGoodsParams): AxiosPromise => {
    return httpHeader("POST", "/goods/modify", params, {})
}

// 删除
export let deleteGoods = (id: number): AxiosPromise => {
    let params = {
        "id": id
    }
    return httpHeader("POST", "/goods/delete", params, {})
}