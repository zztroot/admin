import { httpHeader, AxiosPromise } from "./api"

// 创建用户
let create = (user_name: string, password: string, real_name: string, phone: string, age: number, sex: number | undefined, role_id: number, again_password: string): AxiosPromise => {
    let params = {
        "user_name": user_name,
        "user_password": password,
        "user_real_name": real_name,
        "user_phone": phone,
        "user_age": age,
        "user_sex": sex,
        "user_role_id": role_id,
        "again_password": again_password
    }
    return httpHeader("POST", "/user/register", params, {})
}

// 修改用户
export let modify = (user_name: string, password: string, real_name: string, phone: string, age: number, sex: number | undefined, role_id: number, new_password: string, id: number, again_password: string): AxiosPromise => {
    let params = {
        "user_name": user_name,
        "user_password": password,
        "user_real_name": real_name,
        "user_phone": phone,
        "user_age": age,
        "user_sex": sex,
        "user_role_id": role_id,
        "new_user_password": new_password,
        "id": id,
        "again_password": again_password
    }
    return httpHeader("POST", "/user/modify", params, {})
}

// 获取所有用户分页
export let getlist = (page: string, page_size: string, user_name: string, sex: number): AxiosPromise => {
    let url = `/user/get-list?page=${page}&page_size=${page_size}`
    if (user_name != "") {
        url = url + `&user_name=${user_name}`
    }
    if (sex != 0) {
        url = url + `&sex=${sex}`
    }
    return httpHeader("GET", url, {}, {})
}

// 根据ID查询单个用户
export let getUserOne = (id: number): AxiosPromise => {
    let params = {
        "id": id
    }
    return httpHeader("POST", "/user/get-one", params, {})
}

// 删除一个用户
export let deleteUserOne = (id: number): AxiosPromise => {
    let params = {
        "id": id
    }
    return httpHeader("POST", "/user/delete-one", params, {})
}

// 获取所有用户不分页
export let getUserAll = (): AxiosPromise => {
    return httpHeader("GET", "/user/get-all", {}, {})
}

export {
    create,
}