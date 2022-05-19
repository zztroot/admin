import { httpRequests, httpHeader, AxiosPromise } from "./api"

// 创建角色
let createRole = (name: string): AxiosPromise => {
    let params = {
        "name": name
    }
    return httpHeader("POST", "/role/save", params, {})
}

// 修改角色
export let modifyRole = (params: httpRequests): AxiosPromise => {
    return httpHeader("POST", "/role/modify", params.params, params.headers)
}

// 删除角色
let deleteRole = (id: number): AxiosPromise => {
    let params = {
        "id": id
    }
    return httpHeader("POST", "/role/del", params, {})
}

// 获取全部数据
let getRoleList = (params: httpRequests): AxiosPromise => {
    return httpHeader("GET", "/role/get-list", params.params, params.headers)
}

// 获取全部菜单
export let getMenuList = (params: httpRequests): AxiosPromise => {
    return httpHeader("GET", "/menu", params.params, params.headers)
}

// 根据名称查询
export let getOneByName = (params: httpRequests): AxiosPromise => {
    return httpHeader("POST", "/roleget-one", params.params, params.headers)
}

export {
    createRole,
    deleteRole,
    getRoleList,
}
